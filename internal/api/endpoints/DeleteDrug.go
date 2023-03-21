package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func DeleteDrug(w http.ResponseWriter, r *http.Request) {
	DrugID := mux.Vars(r)["id"]
	drugs := []model.Drug{}

	for i, singleDrug := range drugs {
		if singleDrug.DrugID == DrugID {
			drugs = append(drugs[:i], drugs[i+1:]...)
			fmt.Fprintf(w, "The drug with ID %v has been deleted successfully", DrugID)
		}
	}
}
