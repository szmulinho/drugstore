package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/database"
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

	var existingDrug model.Drug
	result := database.DB.First(&existingDrug, DrugID)
	if result.Error != nil {
		http.Error(w, "Drug not found", http.StatusNotFound)
		return
	}

	result = database.DB.Delete(&existingDrug)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "The drug with ID %v has been deleted successfully", DrugID)
}
