package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetOneDrug(w http.ResponseWriter, r *http.Request) {
	DrugID := mux.Vars(r)["id"]
	drugs := []model.Drug{}
	for _, singleDrug := range drugs {
		if singleDrug.DrugID == DrugID {
			json.NewEncoder(w).Encode(singleDrug)
		}
	}
}
