package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type BasicResponse struct {
	Status  string
	Message string
}
type BasicRequest struct {
	InputString string `json:"inputString"`
}

func advanced(w http.ResponseWriter, r *http.Request) {}

func basic(w http.ResponseWriter, r *http.Request) {
	var method = r.Method
	if method == "GET" {
		var stat = "success"
		var mess = r.FormValue("inputString")
		var resp = BasicResponse{Status: stat, Message: mess}
		json.NewEncoder(w).Encode(resp)
		fmt.Println("Endpoint Hit: basic get")
	} else if method == "POST" {
		fmt.Println("Endpoint Hit: basic post")
		var stat = "success"
		r.ParseForm()
		var req BasicRequest
		var unmarshalErr *json.UnmarshalTypeError
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&req)
		if err != nil {
			if errors.As(err, &unmarshalErr) {
				errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
			} else {
				errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
			}
			return
		}
		var resp = BasicResponse{Status: stat, Message: req.InputString}
		json.NewEncoder(w).Encode(resp)
		fmt.Println("Endpoint Hit: basic get")
	}

}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["Message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
func handleRequests() {
	http.HandleFunc("/basic", basic)
	http.HandleFunc("/advanced", advanced)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
