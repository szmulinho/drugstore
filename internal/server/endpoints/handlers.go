package endpoints

import (
	"gorm.io/gorm"
	"net/http"
)

type Handlers interface {
	AddDrug(w http.ResponseWriter, r *http.Request)
	CreateToken(w http.ResponseWriter, r *http.Request, userID int64, isCustomer bool) (string, error)
	DeleteDrug(w http.ResponseWriter, r *http.Request)
	GetDrugByName(w http.ResponseWriter, r *http.Request)
	GetOneDrug(w http.ResponseWriter, r *http.Request)
	UpdateDrug(w http.ResponseWriter, r *http.Request)
	GetAllDrugs(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) Handlers {
	return &handlers{
		db: db,
	}
}
