package main

import (
	"errors"
	"fmt"
)

type AnimalDato struct {
	Perro     float64
	Gato      float64
	Hamster   float64
	Tarántula float64
}

func TestAnimal(animal string) (func(float64) float64, error) {
	dato := AnimalDato{
		Perro:     10.0,
		Gato:      5.0,
		Hamster:   0.25,
		Tarántula: 0.15,
	}

	switch animal {
	case "perro":
		return func(cantidad float64) float64 {
			return dato.Perro * cantidad
		}, nil
	case "gato":
		return func(cantidad float64) float64 {
			return dato.Gato * cantidad
		}, nil
	case "hamster":
		return func(cantidad float64) float64 {
			return dato.Hamster * cantidad
		}, nil
	case "tarántula":
		return func(cantidad float64) float64 {
			return dato.Tarántula * cantidad
		}, nil
	default:
		return nil, errors.New("Animal desconocido")
	}
}

func main() {
	animalTipo := "gato"
	cantidad := 3.0

	animalFunc, err := TestAnimal(animalTipo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalFood := animalFunc(cantidad)
	fmt.Println("Para", cantidad, animalTipo+"s,", "se necesita", totalFood, "kg de alimento.")
}

//comentario
