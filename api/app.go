package main

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"

	ctrl "github.com/dymzfp/tes-netmonk/controller"
)

const (
	ListenPort = ":8081"
)

func Handle() {
	r := mux.NewRouter()
	r.HandleFunc("/api/snmp", ctrl.GetSnmp).Methods(http.MethodGet)
	r.HandleFunc("/api/available", ctrl.Available).Methods(http.MethodGet)
	r.HandleFunc("/api/mttr", ctrl.Mttr).Methods(http.MethodGet)

	log.Printf("Start server at http://localhost%v", ListenPort)
	log.Fatal(http.ListenAndServe(ListenPort, r))
}

func main() {
	Handle()
}