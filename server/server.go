package server

import (
	"log"
	"net/http"
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
	w.Write([]byte("Hello, world!"))
}

func appendDose(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func deleteDose(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.PathValue("id")))
}
