package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
)

func GetAllDrugs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "C:\\Program Files\\Go\\src\\github.com\\szmulinho\\drugstore\\cmd\\server\\www\\getAllDrugs\\index.html")
	json.NewEncoder(w).Encode(model.Drugs)

}
