package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio  && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch(n) {
		case n.entre(9.0, 10):
			return "A"

		default:
			return "E"
	}	
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(3.4))
	fmt.Println(notaParaConceito(2.0))
	fmt.Println(notaParaConceito(6.7))
}
