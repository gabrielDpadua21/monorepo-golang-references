package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"F": {
			"Frajola": 1000.0,
		},
		"T": {
			"Thor": 12000.0,
			"Tom":  13000.0,
		},
		"L": {
			"Lucifer": 30000.0,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "T")

	for letra, funcs := range funcsPorLetra {
		for nome, salario := range funcs {
			fmt.Printf("Letra %s, Nome: %s, Salario %.2f\n", letra, nome, salario)
		}
	}
}
