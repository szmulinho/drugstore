package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"io/ioutil"
	"net/http"
)

func UpdateDrug(w http.ResponseWriter, r *http.Request) {
	DrugID := mux.Vars(r)["id"]
	var updatedDrug model.Drug
	drugs := []model.Drug{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the task title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedDrug)

	for i, singleDrug := range drugs {
		if singleDrug.DrugID == DrugID {
			singleDrug.Name = updatedDrug.Name
			singleDrug.Type = updatedDrug.Type
			singleDrug.Price = updatedDrug.Price
			drugs = append(drugs[:i], singleDrug)
			json.NewEncoder(w).Encode(singleDrug)
		}
	}
}