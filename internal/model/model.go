package model

import (
	"os"
	"time"
)

type Drug struct {
	DrugID string `json:"drug-id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Price  string `json:"price"`
}

type Presc struct {
	PreId      string    `json:"pre-id"`
	DrugId     string    `json:"drug-id"`
	Expiration time.Time `json:"expiration"`
}

type server struct {
	prescs []Presc
}

type Port struct {
	Port string
}

type User struct {
	UserID   string `json:"user-id"`
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

var Users = []User{
	User{
		UserID:   "user1",
		Password: "password1",
	},
	User{
		UserID:   "user2",
		Password: "password2",
	},
}
