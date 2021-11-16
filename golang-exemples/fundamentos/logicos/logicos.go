package main

import (
	"fmt"
)

func compras(trab1, trab2 bool) (bool, bool, bool) {
	// Operador e
	comprarTv50 := trab1 && trab2

	// Operador ou exclusivo nao existe porem pela diferenca e possivel fazer a validacao
	comprarTv32 := trab1 != trab2

	// Operador ou inclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)

	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
