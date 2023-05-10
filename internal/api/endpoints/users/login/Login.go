package login

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/drugstore/database"
	"github.com/szmulinho/drugstore/internal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := database.DB
	var userFromDB model.User
	result := db.Where("login = ?", user.Login).First(&userFromDB)
	if result.Error != nil {
		http.Error(w, "Invalid username", http.StatusUnauthorized)
		return
	}
	hashPasswordFromDatabase := userFromDB.Password

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(hashPasswordFromDatabase))

	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": user.Login,
		"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, err := token.SignedString(model.JwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString})
}
