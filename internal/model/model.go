package model

import (
	"os"
)

type Drug struct {
	DrugID int64  `json:"drugid" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	Price  string `json:"price"`
}

var Prescs []CreatePrescInput

type CreatePrescInput struct {
	PreId      int64  `json:"preid"`
	Drugs      string `json:"drugs"`
	Expiration string `json:"expiration"`
}

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
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique"`
	Password string
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
