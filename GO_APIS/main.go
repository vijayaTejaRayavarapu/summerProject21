package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type BasicResponse struct {
	Status  string
	Message string
}

type AdvancedResponseGet struct {
	Status     string
	Message    string
	EmployeeId string
	SkillId    string
	HobbyId    string
	Country    string
	State      string
	Phone      string
	City       string
}

type BasicRequest struct {
	InputString string `json:"inputString"`
}

type Address struct {
	Line1   string `json:"line1"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type BasicDetails struct {
	Id           int32    `json:"id"`
	Name         string   `json:"name"`
	PhoneNumbers []string `json:"phoneNumbers"`
	Address      Address  `json:"address"`
}

type Skill struct {
	SkillId int32  `json:"skillId"`
	Skill   string `json:"skill"`
}

type Hobby struct {
	HobbyId int32  `json:"hobbyId"`
	Hobby   string `json:"hobby"`
}

type AdditionalDetails struct {
	Skills  []Skill `json:"skills"`
	Hobbies []Hobby `json:"hobbies"`
}

type EmployeeDetails struct {
	Id                int32             `json:"id"`
	BasicDetails      BasicDetails      `json:"basicDetails"`
	AdditionalDetails AdditionalDetails `json:"additionalDetails"`
}

type AdvancedRequest struct {
	InputString     string            `json:"inputString"`
	EmployeeDetails []EmployeeDetails `json:"employeeDetails"`
}

func advanced(w http.ResponseWriter, r *http.Request) {
	var method = r.Method
	if method == "GET" {
		var stat = "success"
		var mess = r.FormValue("inputString")
		var emp = r.FormValue("employeeId")
		var sid = r.FormValue("skillId")
		var hid = r.FormValue("hobbyId")
		var phone = r.FormValue("phone")
		var city = r.FormValue("city")
		var state = r.FormValue("state")
		var country = r.FormValue("country")
		var resp = AdvancedResponseGet{Status: stat, Message: mess, EmployeeId: emp, SkillId: sid, HobbyId: hid, Phone: phone, City: city, State: state, Country: country}
		json.NewEncoder(w).Encode(resp)
	} else if method == "POST" {
		var req AdvancedRequest
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
		var resp = req
		json.NewEncoder(w).Encode(resp)
	}

}
func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Home")
}
func basic(w http.ResponseWriter, r *http.Request) {
	var method = r.Method
	if method == "GET" {
		var stat = "success"
		var mess = r.FormValue("inputString")
		var resp = BasicResponse{Status: stat, Message: mess}
		json.NewEncoder(w).Encode(resp)
	} else if method == "POST" {
		var stat = "success"
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
	http.HandleFunc("/", home)
	http.HandleFunc("/basic", basic)
	http.HandleFunc("/advanced", advanced)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
