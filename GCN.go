package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	x := [9999]int{}
	var y []int

	// Creamos slice con combinacion
	for i := range x {
		if i > 1000 {
			y = append(y, i)
		}
	}

	// Gestionamos login shell
	cont := 1
	for _, a := range y {
		z := rand.Intn(len(y))
		fmt.Println("combinacion ordenada  : " + strconv.Itoa(a))
		fmt.Println("combinacion aleatoria : " + strconv.Itoa(y[z]))
		cont++
	}
	fmt.Println("contador total: " + strconv.Itoa(cont))
}
