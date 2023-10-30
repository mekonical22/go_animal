package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func main() {
	Index := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			i := rand.Intn(len(animalsData)-1)
			data := GetAnimalData(animalsData[i])

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(*data)
			return
		}

		w.WriteHeader(http.StatusBadRequest) // reject all methods except GET
		return
	}
	http.HandleFunc("/", Index)

	http.ListenAndServe(":3000", nil)
}

var dog = &Animal{
	Name: "Dog",
	MaxAge: 13,
	ColorPatterns: [][2]string{
		{"Black", "White"},
		{"Black", "Black"},
		{"Orange", "Black"},
		{"Brown", "White"},
	},
}

var cat = &Animal{
	Name: "Cat",
	MaxAge: 18,
	ColorPatterns: [][2]string{
		{"Orange", "Orange"},
		{"Orange", "White"},
		{"Gray", "Black"},
		{"White", "Black"},
	},
}

var animalsData = [...]Animal{*dog, *cat}
