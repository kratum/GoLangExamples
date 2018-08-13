package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Issue struct {
	Url  string `json:"Url"`
	Band string `json:"Band"`
	Jahr string `json:"Jahr"`
}

var issues = []Issue{}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data/mvgn_ausgaben.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// ReadAll of jsonFile
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal the byteValues to persons
	json.Unmarshal([]byte(byteValue), &issues)

	//Mux Router for Routing
	router := mux.NewRouter()

	//Handler for Endpoints
	router.HandleFunc("/issues/", GetEntriesEndpoint).Methods("GET")
	router.HandleFunc("/issues/{id}", GetEntryEndpoint).Methods("GET")
	router.HandleFunc("/issues/{id}", CreateEntryEndpoint).Methods("POST")
	router.HandleFunc("/issues/{id}", DeleteEntryEndpoint).Methods("DELETE")
	router.HandleFunc("/issues/{id}", UpdateEntryEndpoint).Methods("PUT")

	//Run the Server and Log Errors
	log.Fatal(http.ListenAndServe(":8080", router))

}

//GetEntriesEndpoint Handler lists all Entries
func GetEntriesEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Get All Entries:")
	json.NewEncoder(w).Encode(issues)
}

//GetEntryEndpoint Handler lists one Entries
func GetEntryEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
	for item := range issues {
		if issues[item].Jahr == params["id"] {
			json.NewEncoder(w).Encode(issues[item])
			return
		}
	}
	json.NewEncoder(w).Encode(&Issue{})
}

//UpdateEntryEndpoint updates an Entry
func UpdateEntryEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var issue Issue
	fmt.Println(issue)

	for index, _ := range issues {
		if issues[index].Jahr == params["id"] {
			fmt.Println(issues[index])

			issues = append(issues[:index], issues[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&issue)
			issue.Jahr = params["id"]
			issues = append(issues, issue)
			json.NewEncoder(w).Encode(issue)
			break
		}
	}
	json.NewEncoder(w).Encode(issues)

}

//CreateEntryEndpoint Creates an Entry
func CreateEntryEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var issue Issue
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&issue)
	issue.Jahr = params["id"]
	issues = append(issues, issue)
	json.NewEncoder(w).Encode(issues)

}

//DeleteEntryEndpoint delets an Entry
func DeleteEntryEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
	for index, _ := range issues {
		if issues[index].Jahr == params["id"] {
			issues = append(issues[:index], issues[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(issues)

}
