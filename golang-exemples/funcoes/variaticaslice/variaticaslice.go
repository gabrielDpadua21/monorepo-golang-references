package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("lista de aprovados...")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Frajola", "Thor", "Lucifer", "Sheldon"}
	imprimirAprovados(aprovados...)
}
