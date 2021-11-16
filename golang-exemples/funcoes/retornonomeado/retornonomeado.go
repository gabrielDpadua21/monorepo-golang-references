package main

import "fmt"

func trocar(p1, p2 int) (seg int, pri int) {
	seg = p2
	pri = p1

	return // retorno limpo
}

// Retorno os valores de em ordem de declaracao
func teste() (a, b string) {
	a = "retorno 1"
	b = "retorno 2"

	return
}

func main() {
	x, y := trocar(2, 3)

	fmt.Println(x, y)

	seg, pri := trocar(1, 7)

	fmt.Println(seg, pri)

	fmt.Println(teste())
}
