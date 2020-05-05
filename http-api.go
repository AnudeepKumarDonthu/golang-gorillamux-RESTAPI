package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UserAddress is a strut which will hold the details of user Address details
type UserAddress struct {
	Area    string `json:"area"`
	City    string `json:"city"`
	PinCode string `json:"pincode"`
}

// Items is a struct to hold the json body from the post request
type User struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Age         int         `json:"age"`
	Useraddress UserAddress `json:"useraddress"`
}

var newID int

var data []string

var peopleData []User

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/teststring", teststring).Methods("GET")
	router.HandleFunc("/testjson", testjson).Methods("GET")

	router.HandleFunc("/addstring/{datavalue}", addString).Methods("POST")

	router.HandleFunc("/adduser", addUser).Methods("POST")
	router.HandleFunc("/getallusers", getAllUsers).Methods("GET")

	http.ListenAndServe(":8080", router)
}

/*
	All End point handler Function Definitions
*/

//teststring
func teststring(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application is Running"))
}

//testjson
func testjson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status string
	}{"Application is Running"})
}

//addstring/{datavalue}
func addString(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mainData := mux.Vars(r)["datavalue"]
	data = append(data, mainData)
	json.NewEncoder(w).Encode(data)
}

//adduser
func addUser(w http.ResponseWriter, r *http.Request) {
	newID = newID + 1
	var newItem User
	newItem.ID = newID
	json.NewDecoder(r.Body).Decode(&newItem)
	peopleData = append(peopleData, newItem)
	w.Write([]byte("User Added Successfully with id : " + strconv.Itoa(newID)))
}

//getallusers
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peopleData)
}
