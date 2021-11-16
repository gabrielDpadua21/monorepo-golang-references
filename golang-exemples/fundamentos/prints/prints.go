package main

import (
	"fmt"
)

func main() {
	// Printa os valores em uma mesma linha
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	// printa valores com quebra de linha
	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 3.141592

	// nao e possivel usar a funcao print para concatenar valores
	// fmt.Print("X: ", x) erro

	// Essa funcao formata o campo para poder printar o valor no console
	xs := fmt.Sprint(x)

	fmt.Print("X: ", xs)
	fmt.Println("X: ", x)
	fmt.Printf("X: %.4f", x)

	// Funcão do print formatado

	a := 1
	b := 1.9999
	c := true
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
