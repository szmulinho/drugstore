package endpoints

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func DeleteDrug(w http.ResponseWriter, r *http.Request) {
	DrugID := mux.Vars(r)["id"]

	for i, singleDrug := range model.Drugs {
		if singleDrug.DrugID == DrugID {
			model.Drugs = append(model.Drugs[:i], model.Drugs[i+1:]...)
			fmt.Fprintf(w, "The drug with ID %v has been deleted successfully", DrugID)
		}
	}
}
