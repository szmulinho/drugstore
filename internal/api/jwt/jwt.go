package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
	"strings"
	"time"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserID,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, error := token.SignedString(model.JwtKey)
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString})
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return model.JwtKey, nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(model.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					json.NewEncoder(w).Encode(model.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(model.Exception{Message: "An authorization header is required"})
		}
	})
}
