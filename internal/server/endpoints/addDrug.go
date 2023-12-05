package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func (h *handlers) AddDrug(w http.ResponseWriter, r *http.Request) {
	var newDrug model.Drug

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request body: %s", err), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &newDrug)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to unmarshal JSON: %s", err), http.StatusBadRequest)
		return
	}

	result := h.db.Create(&newDrug)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("failed to create new drug in the database: %s", result.Error), http.StatusInternalServerError)
		return
	}

	for _, singleDrug := range model.Drugs {
		fmt.Println(singleDrug)
		if singleDrug.DrugID == newDrug.DrugID {
			w.WriteHeader(http.StatusConflict)
			errorMsg := fmt.Sprintf("Drug %v already exists", newDrug.DrugID)
			if _, err := w.Write([]byte(errorMsg)); err != nil {
				http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
			}
			return
		}
	}

	model.Drugs = append(model.Drugs, newDrug)

	fmt.Printf("added new drug %+v\n", newDrug)

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newDrug); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response JSON: %s", err), http.StatusInternalServerError)
	}
}
