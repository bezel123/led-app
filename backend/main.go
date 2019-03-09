package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Devices retrieves and sends all connected Devices
func Devices(w http.ResponseWriter, r *http.Request) {

}

func test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ayy")
}

func rgb(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("lmaoo")
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/CheckConnection", test).Methods("GET")
	router.HandleFunc("/v1/GetDevices", Devices).Methods("GET")
	router.HandleFunc("/v1/SendValues", rgb).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}
