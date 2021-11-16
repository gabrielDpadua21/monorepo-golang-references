package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 2

	// operador posfixado
	// mesmo que x += 1 ou x = x + 1
	x++
	fmt.Println(x)

	y--
	fmt.Println(y)

	// nao e possivel usar operador de atribuicao em operacoes
	// x == y++
}
