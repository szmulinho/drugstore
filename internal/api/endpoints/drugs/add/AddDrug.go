package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/drugstore/database"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func AddDrug(w http.ResponseWriter, r *http.Request) {
	var newDrug model.Drug

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &newDrug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := database.DB.Create(&newDrug)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, singleDrug := range model.Drugs {
		fmt.Println(singleDrug)
		if singleDrug.DrugID == newDrug.DrugID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Drug %s already exist", newDrug.DrugID)})
			return
		}
	}

	model.Drugs = append(model.Drugs, newDrug)

	fmt.Printf("added new drug %+v\n", newDrug)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newDrug)
}
