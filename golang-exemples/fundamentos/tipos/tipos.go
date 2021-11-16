package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Tipos de dados em go

	// Inteiro literal
	fmt.Println("Tipo inteiro literal: ", reflect.TypeOf(1000))

	// tipos inteiros positivos uint8, uint16, uint32, uint64

	// tipo byte e alias para uint8
	var b byte = 255
	fmt.Println("Tipo byte ", reflect.TypeOf(b))

	// Numeros inteiros com sinal (tanto positivos quando negativos)
	// int8, int16, int32, int64

	i1 := math.MaxInt64
	fmt.Println("Valor maximo de int", i1)
	fmt.Println("Tipo de i1:", reflect.TypeOf(i1))

	// representa um mapeamento da tabela Unicode (int32)
	var i2 rune = 'a'

	fmt.Println("Rune e", reflect.TypeOf(i2))
	fmt.Println("Valor de i2", i2)

	// Numeros reais float32 float64
	var f float32 = 2.3
	fmt.Println("Tipo de f", reflect.TypeOf(f))
	fmt.Println("Tipo literal", reflect.TypeOf(2.3))

	// Tipo boolean - pode ser trabalhado com operadores lógicos
	var bo bool = true
	fmt.Println("Tipo de bo", reflect.TypeOf(bo))

	// String - somente e delimitado por aspas duplas
	s1 := "lorem lore lorem"
	fmt.Println("Tipo de s1", reflect.TypeOf(s1))
	// A funcao len tras o tamanho de um string
	fmt.Println("Tamamho de s1", len(s1))

	// String de multiplas linha
	s2 := `lorem
	  lorem
	`

	fmt.Println("Tamanho de s2", len(s2))

}
