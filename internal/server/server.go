package server

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/server/endpoints"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Run(ctx context.Context, db *gorm.DB) {
	handler := endpoints.NewHandler(db)
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/drug", handler.AddDrug).Methods("POST")
	router.HandleFunc("/drugs", handler.GetAllDrugs).Methods("GET")
	router.HandleFunc("/drugs/{id}", handler.GetOneDrug).Methods("GET")
	router.HandleFunc("/drugs/{name}", handler.GetDrugByName).Methods("GET")
	router.HandleFunc("/drugs/{id}", handler.UpdateDrug).Methods("PATCH")
	router.HandleFunc("/drugs/{id}", handler.DeleteDrug).Methods("DELETE")
	router.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		userID := uint(1)
		isDoctor := true
		token, err := handler.CreateToken(w, r, int64(userID), isDoctor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
	}).Methods("POST")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	go func() {
		err := http.ListenAndServe(":8081", cors(router))
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

}
