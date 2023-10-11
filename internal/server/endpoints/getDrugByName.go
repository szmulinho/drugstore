package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func (h *handlers) GetDrugByName(w http.ResponseWriter, r *http.Request) {
	drugName := mux.Vars(r)["name"]
	var foundDrugs []model.Drug

	for _, singleDrug := range model.Drugs {
		if singleDrug.Name == drugName {
			foundDrugs = append(foundDrugs, singleDrug)
		}
	}

	if len(foundDrugs) == 0 {
		http.Error(w, "Drug not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(foundDrugs)
}
