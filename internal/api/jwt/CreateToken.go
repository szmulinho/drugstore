package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
	"time"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&model.Juser)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": model.Juser.Jwtuser,
		"password": model.Juser.Jwtpassword,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, error := token.SignedString(model.JwtKey)
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString})
}
