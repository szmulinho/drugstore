package model

import (
	"html/template"
	"os"
	"time"
)

var Templates *template.Template

type Drug struct {
	DrugID string `json:"drug-id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  string `json:"price"`
}

var Prescs []CreatePrescInput

type CreatePrescInput struct {
	PreId      string    `json:"pre-id"`
	Drugs      []string  `json:"drugs"`
	Expiration time.Time `json:"expiration"`
}

var PreId string

var Drugz []string

var Expiration time.Time

var Prescription CreatePrescInput

var drugs []Drug

type Port struct {
	Port string
}

var Drugs []Drug

type jwtUser struct {
	Jwtuser     string "jwt-user"
	Jwtpassword string "jwtpassword"
}

var Juser jwtUser

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

type Response struct {
	Data string `json:"data"`
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
