package main

import (
	"fmt"
)

func main() {
	funcsESalarios := map[string]float64{
		"Frajola": 10000.0,
		"Thor":    15000.20,
		"Lucifer": 30000.0,
	}

	funcsESalarios["Tom"] = 12000.0

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "Tom")

	for nome, salario := range funcsESalarios {
		fmt.Printf("Nome: %s, Salario: %.2f\n", nome, salario)
	}
}
