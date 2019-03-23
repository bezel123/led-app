package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/CheckConnection", CheckConnection).Methods("GET")
	router.HandleFunc("/v1/GetDevices", GetDevices).Methods("GET")
	router.HandleFunc("/v1/CreateDevice", CreateDevice).Methods("POST")
	router.HandleFunc("/v1/DeleteDevice", DeleteDevice).Methods("DELETE")
	router.HandleFunc("/v1/UpdateDevice", UpdateDevice).Methods("PUT")
	log.Fatal(http.ListenAndServe(":80", router))
}
