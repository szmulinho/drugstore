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
	router.HandleFunc("/", api.HomeLink)
	router.HandleFunc("/drug", api.AddDrug).Methods("POST")
	router.HandleFunc("/drugs", api.GetAllDrugs).Methods("GET")
	router.HandleFunc("/drugs/{id}", api.GetOneDrug).Methods("GET")
	router.HandleFunc("/drugs/{id}", api.UpdateDrug).Methods("PATCH")
	router.HandleFunc("/drugs/{id}", api.DeleteDrug).Methods("DELETE")
	router.HandleFunc("/presc", api.AddPrescription).Methods("POST")
	router.HandleFunc("/prescs", jwt.ValidateMiddleware(api.GetAllPrescriptions)).Methods("GET")
	router.HandleFunc("/presc/{id}", api.GetOnePrescription).Methods("GET")
	router.HandleFunc("/authenticate", jwt.CreateToken).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8081"), router))
}

func server() {
	log.Println("server with port 8081 is starting")
	Run()
}
