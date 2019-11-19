package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Bug struct {
	Issue string `json:"Issue"`
	Desc  string `json:"Desc"`
}

type Bugs []Bug

func allBugs(w http.ResponseWriter, r *http.Request) {
	Bugs := Bugs{{Issue: "Main bug", Desc: "Check it"}}
	fmt.Fprint(w, "Endpoint all BUGS")
	json.NewEncoder(w).Encode(Bugs)
}

func postBugs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post")
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/bugs", allBugs).Methods("GET")
	myRouter.HandleFunc("/bugs", postBugs).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
func main() {
	handleRequests()
}
