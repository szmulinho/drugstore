package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/api/endpoints"
	"github.com/szmulinho/drugstore/internal/api/jwt"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", endpoints.HomeLink)
	router.HandleFunc("/drug", endpoints.AddDrug).Methods("POST")
	router.HandleFunc("/drugs", endpoints.GetAllDrugs).Methods("GET")
	router.HandleFunc("/drugs/{id}", endpoints.GetOneDrug).Methods("GET")
	router.HandleFunc("/drugs/{id}", endpoints.UpdateDrug).Methods("PATCH")
	router.HandleFunc("/drugs/{id}", endpoints.DeleteDrug).Methods("DELETE")
	router.HandleFunc("/presc", endpoints.AddPrescription).Methods("POST")
	router.HandleFunc("/prescs", jwt.ValidateMiddleware(endpoints.GetAllPrescriptions)).Methods("GET")
	router.HandleFunc("/presc/{id}", endpoints.GetOnePrescription).Methods("GET")
	router.HandleFunc("/authenticate", jwt.CreateToken).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), router))
}

func server() {
	log.Println("server with port 8081 is starting")
	Run()
}
