package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type DateLogin struct {
	Nume   string `json:"nume"`
	Parola string `json:"parola"`
}

type DateInregistrare struct {
	Email string `json:"email"`
	DateLogin
}

// POST /utilizatori
func Inregistrare(w http.ResponseWriter, r *http.Request, db *sqlx.DB) (interface{}, error) {
	var date DateInregistrare
	err := DecodeJSON(r, &date)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	hash := sha256.New()
	hash.Write([]byte(date.Parola))
	date.Parola = base64.URLEncoding.EncodeToString(hash.Sum(nil))
	_, err = db.NamedExec("INSERT INTO utilizatori (email, parola, nume) VALUES (:email, :parola, :nume)", date)
	switch {
	case err == nil:
	case VerifCod(err, UniqueViolation):
		return http.StatusBadRequest, err
	default:
		return http.StatusInternalServerError, err
	}
	var utilizator Utilizator
	db.Get(&utilizator, `SELECT * FROM utilizatori WHERE nume = $1`, date.Nume)

	jwt, err := utilizator.GetJWT()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return []byte("{\"token\": \"" + jwt + "\"}"), nil
}

// PUT /utilizatori
func Login(w http.ResponseWriter, r *http.Request, db *sqlx.DB) (interface{}, error) {
	var date DateLogin
	err := DecodeJSON(r, &date)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	var utilizator Utilizator
	err = db.Get(&utilizator, `SELECT * FROM utilizatori WHERE nume = $1`, date.Nume)
	switch {
	case err == nil:
	case err == sql.ErrNoRows:
		return http.StatusNotFound, err
	default:
		return http.StatusInternalServerError, err
	}

	hash := sha256.New()
	hash.Write([]byte(date.Parola))
	if base64.URLEncoding.EncodeToString(hash.Sum(nil)) != utilizator.Parola {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}
	jwt, err := utilizator.GetJWT()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return []byte("{\"token\": \"" + jwt + "\"}"), nil
}

// GET /utilizatori?nume=<>
func Info(w http.ResponseWriter, r *http.Request, db *sqlx.DB) (interface{}, error) {
	nume := r.URL.Query().Get("nume")
	if nume == "" {
		return http.StatusBadRequest, errors.New("numele nu poate fi gol")
	}

	var utilizator Utilizator
	err := db.Get(&utilizator, `SELECT * FROM utilizatori WHERE nume = $1`, nume)
	switch {
	case err == nil:
	case err == sql.ErrNoRows:
		return http.StatusNotFound, err
	default:
		return http.StatusInternalServerError, err
	}

	toReturn, _ := json.Marshal(utilizator)
	return toReturn, nil
}
