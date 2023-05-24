package getUsers

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/internal/database"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	result := database.DB.Find(&model.Users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.Users)
}
