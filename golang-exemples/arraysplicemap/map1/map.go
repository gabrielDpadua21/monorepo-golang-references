package main

import "fmt"

func main() {
	// Chave -> valor
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[987654321] = "Pedro"
	aprovados[987623401] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// Acessando um valor de um map
	fmt.Println(aprovados[123456789])

	// excluindo o valor de um map
	delete(aprovados, 123456789)

	fmt.Println(aprovados)
}
