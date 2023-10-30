package main

import (
	"math/rand"
)

type Animal struct {
	Name          string
	MaxAge        uint
	ColorPatterns [][2]string
}

type AnimalData struct{
	Name          string    `json:"name"`
	Age           int       `json:"age"`
	ColorPatterns [2]string `json:"color_patterns"`
}

func GetAnimalData(animal Animal) *AnimalData {
	var index int
	if i := rand.Intn(len(animal.ColorPatterns)-1); i <= 0 {
		index = 1
	} else {
		index = i
	}

	return &AnimalData{
		Name: animal.Name,
		Age: rand.Intn(int(animal.MaxAge)),
		ColorPatterns: animal.ColorPatterns[index],
	}
}
