package api

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetAllPrescriptions(w http.ResponseWriter, r *http.Request) {
	prescs := []model.Presc{}
	json.NewEncoder(w).Encode(prescs)

}
