package server

import (
	"encoding/json"
	"log"
	"net/http"
	"skrive/data"
)

func Serve(address string) {
	server := &http.Server{
		Addr: address,
	}
	http.HandleFunc("GET /v1/doses", getDoses)
	http.HandleFunc("POST /v1/doses/apend", appendDose)
	http.HandleFunc("DELETE /v1/doses/{id}", deleteDose)
	log.Print("Serving on " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func getDoses(w http.ResponseWriter, req *http.Request) {
	if !handleAuthentication(w, req) {
		return
	}

	doses, err := data.ApplicationStorage.FetchAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var response []byte
	response, err = json.Marshal(doses)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write(response)
}

func appendDose(w http.ResponseWriter, req *http.Request) {
	if !handleAuthentication(w, req) {
		return
	}
	w.Write([]byte("Hello, world!"))
}

func deleteDose(w http.ResponseWriter, req *http.Request) {
	if !handleAuthentication(w, req) {
		return
	}
	w.Write([]byte(req.PathValue("id")))
}

func handleAuthentication(w http.ResponseWriter, req *http.Request) bool {
	// TODO
	return true
}
