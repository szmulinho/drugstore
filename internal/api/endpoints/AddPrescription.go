package api

import (
	"encoding/json"
	"fmt"
	"github.com/szmulinho/drugstore/internal/model"
	"io/ioutil"
	"net/http"
)

func AddPrescription(w http.ResponseWriter, r *http.Request) {
	var newPresc model.Presc
	prescs := []model.Presc{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(reqBody, &newPresc)
	if err != nil {
		panic(err)
	}

	for _, singlePresc := range prescs {
		fmt.Println(singlePresc)
		if singlePresc.PreId == newPresc.PreId {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Prescription %s already exist", newPresc.PreId)})
			return
		}

	}

	prescs = append(prescs, newPresc)

	fmt.Printf("created new task %+v\n", newPresc)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPresc)
}
