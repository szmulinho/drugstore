package api

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetAllDrugs(w http.ResponseWriter, r *http.Request) {
	drugs := []model.Drug{}
	json.NewEncoder(w).Encode(drugs)

}
