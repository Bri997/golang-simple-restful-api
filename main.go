package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/bugs", allBugs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequests()
}
