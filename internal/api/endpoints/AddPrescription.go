package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/szmulinho/drugstore/internal/model"
	"io/ioutil"
	"net/http"
)

func AddPrescription(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &model.Prescription)
	if err != nil {
		panic(err)
	}

	for _, singlePresc := range model.Prescs {
		fmt.Println(singlePresc)
		if singlePresc.PreId == model.Prescription.PreId {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Prescription %s already exist", model.Prescription.PreId)})
			return
		}
	}
	model.Prescs = append(model.Prescs, model.Prescription)

	fmt.Printf("created new presciption %+v\n", model.Prescription)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Prescription)
}
