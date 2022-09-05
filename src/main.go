package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func OpenDB() *sqlx.DB {
	var dsn = "host=" + os.Getenv("PGHOST") + " user=" + os.Getenv("PGUSER") + " port=" + os.Getenv("PGPORT") + " dbname=" + os.Getenv("PGDATABASE") + " sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func NewEndpointWithDB(f func(http.ResponseWriter, *http.Request, *sqlx.DB) (interface{}, error), db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// CORS
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		data, err := f(w, r, db)
		if err != nil {
			w.WriteHeader(data.(int))
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data.([]byte))
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := OpenDB()
	r := mux.NewRouter()
	r.HandleFunc("/utilizatori", NewEndpointWithDB(Inregistrare, db)).Methods("OPTIONS", "POST")
	r.HandleFunc("/utilizatori", NewEndpointWithDB(Login, db)).Methods("OPTIONS", "PUT")
	r.HandleFunc("/utilizatori", NewEndpointWithDB(Info, db)).Methods("OPTIONS", "GET")

	fmt.Println("Serverul a pornit!")
	panic(http.ListenAndServe(":3000", r))
}
