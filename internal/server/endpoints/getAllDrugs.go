package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func (h *handlers) GetAllDrugs(w http.ResponseWriter, r *http.Request) {

	result := h.db.Find(&model.Drugs)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Drugs)
}
