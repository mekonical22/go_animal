package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	Index := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			i := len(animalsData)-1
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

var animalsData = [...]Animal{}
