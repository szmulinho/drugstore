package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *handlers) UpdateDrug(w http.ResponseWriter, r *http.Request) {
	drugIDStr := mux.Vars(r)["id"]
	DrugID, err := strconv.ParseInt(drugIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}

	var updatedDrug model.Drug
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Kindly enter data with the drug name and price only in order to update", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &updatedDrug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var existingDrug model.Drug
	result := h.db.First(&existingDrug, DrugID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	existingDrug.Name = updatedDrug.Name
	existingDrug.Price = updatedDrug.Price

	result = h.db.Save(&existingDrug)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(existingDrug)
}
