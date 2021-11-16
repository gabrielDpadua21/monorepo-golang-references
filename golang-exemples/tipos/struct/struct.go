package main

import "fmt"

type produto struct {
	nome string
	preco float64
	desconto float64
}

// Método: funcao com receiver (receptor)
func (p produto) precoComDesconto() float64{
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome: "lapis",
		preco: 2.0,
		desconto: 0.10,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
}

