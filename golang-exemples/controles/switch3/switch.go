package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"

	case float64:
		return "real"

	case string:
		return "String"

	case func():
		return "Funcao"

	default:
		return "Não sei!"

	}
}

func main() {
	fmt.Println(tipo(2))
	fmt.Println(tipo(2.0))
	fmt.Println(tipo("Frajola"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
