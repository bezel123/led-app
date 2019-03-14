package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Values are the values send from the frontend
type Values struct {
	Name string
	R    string
	G    string
	B    string
}

//Devices retrieves and sends all connected Devices
func Devices(w http.ResponseWriter, r *http.Request) {

}

func test(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ayy")
}

func rgb(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("lmaoo")
}

func SendValues(w http.ResponseWriter, r *http.Request) {
	val := Values{
		Name: r.PostFormValue("name"),
		R:    r.PostFormValue("r"),
		G:    r.PostFormValue("g"),
		B:    r.PostFormValue("b"),
	}

	log.Println(val)

	answer := fmt.Sprintln(
		"Successfully get values with name: ", val.Name,
		" r: ", val.R,
		" g: ", val.G,
		" b: ", val.B)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, answer)
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/CheckConnection", test).Methods("GET")
	router.HandleFunc("/v1/GetDevices", Devices).Methods("GET")
	router.HandleFunc("/v1/SendValues", SendValues).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}
