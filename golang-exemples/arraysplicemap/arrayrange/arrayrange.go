package main

import "fmt"

func main() {

	// array de tamanho contato pelo compilador ...
	numeros := [...]int{10, 20, 30, 40, 50}

	// range com indice ...
	for i, item := range numeros {
		fmt.Println(i)
		fmt.Println(item)
	}

	// range sem indice ...
	for _, item := range numeros {
		fmt.Println(item)
	}
}
