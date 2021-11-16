package main

import (
	"fmt"
)

// funcao sem parametro de entra ou saida
func funcao1() {
	fmt.Println("Iniciando Funcao 1...")
}

// funcao com 2 parametros de entrada
func funcao2(p1 string, p2 string) {
	fmt.Printf("parametro 1: %s, parametro 2: %s \n", p1, p2)
}

// Funcao com parametro de saida - return
func funcao3() string {
	return "retorno 3"
}

// funcao com parametro de entrada e saida
func funcao4(p1, p2 string) string {
	return fmt.Sprintf("Funcao 4: %s %s", p1, p2)
}

func funcao5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	funcao1()
	funcao2("teste", "teste2")
	fmt.Println(funcao3())
	fmt.Println(funcao4("teste", "teste2"))

	r51, r52 := funcao5()
	fmt.Println("R1:", r51)
	fmt.Println("R2:", r52)
}
