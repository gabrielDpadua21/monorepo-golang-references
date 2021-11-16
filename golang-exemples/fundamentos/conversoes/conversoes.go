package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 1.2
	y := 1

	// Converte um tipo inteiro para float
	fmt.Println(x / float64(y))

	// Converte um topo float para inteiro
	fmt.Println(int(x) / y)

	// Convertendo um tipo inteiro em string
	// Não e feito a conversao literal e sim o numero correspondente a tabela ascii
	fmt.Println("Teste ", string(80))

	// Conversao de inteiro para string literal
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para inteiro
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
	fmt.Println(reflect.TypeOf(num))

	// string para boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
