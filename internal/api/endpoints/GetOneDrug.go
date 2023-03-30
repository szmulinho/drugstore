package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetOneDrug(w http.ResponseWriter, r *http.Request) {
	DrugID := mux.Vars(r)["id"]
	for _, singleDrug := range model.Drugs {
		if singleDrug.DrugID == DrugID {
			json.NewEncoder(w).Encode(singleDrug)
		}
	}
}
