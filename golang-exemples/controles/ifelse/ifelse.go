package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	a1 := 10.0
	a2 := 5.0

	imprimirResultado(a1)
	imprimirResultado(a2)
}
