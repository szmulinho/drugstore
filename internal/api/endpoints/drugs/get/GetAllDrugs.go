package get

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/database"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetAllDrugs(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(model.Drugs)

	result := database.DB.Where(&model.Drugs)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

}
