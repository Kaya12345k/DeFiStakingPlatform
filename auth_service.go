package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret = os.Getenv("JWT_SECRET")
)

type User struct {
	Username string
	Password string
}

var usersDb = map[string]string{
	"admin": "$2a$14$N0N07Rpite7Rmv1KPf2t4OvWjv/CQ.FF4vl5h9HdZ8XZ12O8mHdE6",
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func signin(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expectedPassword, ok := usersDb[user.Username]

	if !ok || !checkPasswordHash(user.Password, expectedPassword) {
		http.Error(w, "username or password is invalid", http.StatusUnauthorized)
		return
	}

	tokenString, err := generateJWT(user.Username)
	if err != nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(username string) (string, error) {
	expirationTime := jwt.NewNumericDate(jwt.TokenExpiresAt())
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/signin", signin)

	http.Handle("/protected", authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the protected route!"))
	})))

	log.Fatal(http.ListenAndServe(":8000", nil))
}