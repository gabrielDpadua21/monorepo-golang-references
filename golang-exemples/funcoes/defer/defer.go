package main

import "fmt"

func obterValor(numero int) int {
       // atrasa e execucao do metodo 	
       defer fmt.Println("fim...")
	if numero > 5000 {
		fmt.Println("Valor Alto...")
		return 5000
	}
	// o else e desnecessario já que o if está retornando um valor da funcao
	fmt.Println("Valor Baixo...")
	return numero
}

func main() {
	fmt.Println(obterValor(6000))
}
