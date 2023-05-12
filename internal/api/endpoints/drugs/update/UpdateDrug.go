package update

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szmulinho/drugstore/internal/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UpdateDrug(w http.ResponseWriter, r *http.Request) {
	drugIDStr := mux.Vars(r)["id"]
	DrugID, err := strconv.ParseInt(drugIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid drug ID", http.StatusBadRequest)
		return
	}
	var updatedDrug model.Drug
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the task title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedDrug)

	for i, singleDrug := range model.Drugs {
		if singleDrug.DrugID == DrugID {
			singleDrug.Name = updatedDrug.Name
			singleDrug.Price = updatedDrug.Price
			model.Drugs = append(model.Drugs[:i], singleDrug)
			json.NewEncoder(w).Encode(singleDrug)
		}
	}
}
