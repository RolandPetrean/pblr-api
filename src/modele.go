package main

import (
	"time"
)

type Utilizator struct {
	Email       string    `json:"email" db:"email"`
	Parola      string    `json:"parola" db:"parola"`
	Nume        string    `json:"nume" db:"nume"`
	Rol         string    `json:"rol" db:"rol"`
	DataCreatie time.Time `json:"data_creatie" db:"data_creatie"`
}

type Problema struct {
	ID            int64  `json:"id" db:"id"`
	Nume          string `json:"nume" db:"nume"`
	Descriere     string `json:"descriere" db:"descriere"`
	Dificultate   int    `json:"dificultate" db:"dificultate"`
	LimitaTimp    int    `json:"limita_timp" db:"limita_timp"`
	LimitaMemorie int    `json:"limita_memorie" db:"limita_memorie"`
	Evaluator     string `json:"evaluator" db:"evaluator"`
	Autor         string `json:"autor" db:"autor"`
}

type Solutie struct {
	ID         int64  `json:"id" db:"id"`
	Sursa      string `json:"sursa" db:"sursa"`
	Problema   int    `json:"problema" db:"problema"`
	Utilizator string `json:"utilizator" db:"utilizator"`
}

type Test struct {
	ID       int64  `json:"id" db:"id"`
	Intrare  string `json:"intrare" db:"intrare"`
	Iesire   string `json:"iesire" db:"iesire"`
	Problema string `json:"problema" db:"problema"`
}

type Rezultat struct {
	ID      int64 `json:"id" db:"id"`
	Timp    int   `json:"timp" db:"timp"`
	Memorie int   `json:"memorie" db:"memorie"`
	Solutie int   `json:"solutie" db:"solutie"`
	Test    int   `json:"test" db:"test"`
}
