package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetOnePrescription(w http.ResponseWriter, r *http.Request) {
	prescPreId := mux.Vars(r)["id"]
	for _, singlePresc := range model.Prescs {
		if singlePresc.PreId == prescPreId {
			json.NewEncoder(w).Encode(singlePresc)
		}
	}
}
