package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s : %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	go fale("frajola", "miau", 200)
	fale("thor", "miauuu", 400)

	fmt.Println("fim")
}
