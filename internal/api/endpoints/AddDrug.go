package api

import (
	"encoding/json"
	"fmt"
	"github.com/szmulinho/drugstore/internal/model"
	"io/ioutil"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func AddDrug(w http.ResponseWriter, r *http.Request) {
	var newDrug model.Drug
	drugs := []model.Drug{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &newDrug)
	if err != nil {
		panic(err)
	}

	for _, singleDrug := range drugs {
		fmt.Println(singleDrug)
		if singleDrug.DrugID == newDrug.DrugID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Drug %s already exist", newDrug.DrugID)})
			return
		}

	}

	drugs = append(drugs, newDrug)

	fmt.Printf("added new drug %+v\n", newDrug)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newDrug)
}