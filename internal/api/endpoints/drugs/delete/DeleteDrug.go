package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
	"strconv"
)

func DeleteDrug(w http.ResponseWriter, r *http.Request) {
	drugIDStr := mux.Vars(r)["id"]
	DrugID, err := strconv.ParseInt(drugIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}

	for i, singleDrug := range model.Drugs {
		if singleDrug.DrugID == DrugID {
			model.Drugs = append(model.Drugs[:i], model.Drugs[i+1:]...)
			fmt.Fprintf(w, "The drug with ID %v has been deleted successfully", DrugID)
		}
	}
}
