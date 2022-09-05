package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

func DecodeJSON(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func (utilizator *Utilizator) GetJWT() (string, error) {
	claims := jwt.MapClaims{
		"iss":  "pblr",
		"data": utilizator.Nume,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("PRIVATE_KEY")))
}

func VerifyAuthorization(r *http.Request) (string, error) {
	var authorization = r.Header.Get("Authorization")
	if authorization == "" {
		return "", errors.New("Unauthorized")
	}

	signedJWT := strings.Split(authorization, " ")[1]
	token, err := jwt.Parse(signedJWT, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("PRIVATE_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	return token.Claims.(jwt.MapClaims)["data"].(string), nil
}

const (
	UniqueViolation = "duplicate key"
)

func VerifCod(err error, cod string) bool {
	return strings.Contains(err.Error(), cod)
}
