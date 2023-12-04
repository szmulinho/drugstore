package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/common/model"
	"net/http"
)

func (h *handlers) GetAllDrugs(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Find(&model.Drugs{}).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Drugs{})
}
