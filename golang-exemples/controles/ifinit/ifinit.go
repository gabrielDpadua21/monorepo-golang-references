package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	// Gera um numero aleatorio a partir de uma data no sistema
	s := rand.NewSource(time.Now().UnixNano())

	r := rand.New(s)

	// Gera um numero aleatorio ate 10
	return r.Intn(10)
}

func main() {
	// Inicia a variavel antes da condicao
	// E aplicavel no switch
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}
}
