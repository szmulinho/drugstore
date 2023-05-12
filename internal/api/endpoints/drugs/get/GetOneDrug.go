package get

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
	"strconv"
)

func GetOneDrug(w http.ResponseWriter, r *http.Request) {
	drugIDStr := mux.Vars(r)["id"]
	DrugID, err := strconv.ParseInt(drugIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}
	for _, singleDrug := range model.Drugs {
		if singleDrug.DrugID == DrugID {
			json.NewEncoder(w).Encode(singleDrug)
		}
	}
}
