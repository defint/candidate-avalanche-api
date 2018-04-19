package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe" })
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetPeople).Methods("GET")
	router.HandleFunc("/person", CreatePerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}