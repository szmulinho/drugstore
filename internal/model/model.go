package model

import (
	"os"
)

type Drug struct {
	DrugID      int64  `json:"drug_id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Image       string `json:"image"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}

var Drugs []Drug

type Exception struct {
	Message string `json:"message"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
