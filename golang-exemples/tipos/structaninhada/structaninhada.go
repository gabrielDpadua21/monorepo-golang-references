package main

import "fmt"

type item struct {
	produtoID int
	qtd int
	preco float64
}

type pedido struct {
	userID int
	itens []item
}

// Receiver..
func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}

	return total
}

func main() {
	pedido := pedido {
		userID: 1,
		itens: []item {
			item{1, 2, 2.0},
			item{2, 4, 2.4},
			item{3, 5, 12.0},
		},

	}

	fmt.Printf("O valor total do pedido e R$ %.2f", pedido.valorTotal())
}
