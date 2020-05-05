package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Items is a struct to hold the json body from the post request
type Items struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Place string `json:"place"`
}

var data []string

var peopleData []Items

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/teststring", teststring).Methods("GET")
	router.HandleFunc("/testjson", testjson).Methods("GET")
	router.HandleFunc("/addstring/{datavalue}", addString).Methods("POST")
	router.HandleFunc("/addjson", addItems).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func teststring(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application is Running"))
}

func testjson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status string
	}{"Application is Running"})
}

func addString(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mainData := mux.Vars(r)["datavalue"]
	data = append(data, mainData)
	json.NewEncoder(w).Encode(data)
}

func addItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newItem Items
	json.NewDecoder(r.Body).Decode(&newItem)
	peopleData = append(peopleData, newItem)
	json.NewEncoder(w).Encode(peopleData)
}
