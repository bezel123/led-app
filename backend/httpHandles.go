package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Devices retrieves and sends all connected Devices
func GetDevices(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("GetDevices Hit!")
}

func CheckConnection(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("CheckConnection Hit!")
}

func rgb(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("lmaoo")
}
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Create Hit!")
}
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Delete Hit!")
}

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
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
