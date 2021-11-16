package main

import (
	"fmt"
)

func notaConceito(n float64) string {
	nota := int(n)

	switch {
	case nota <= 10 && nota >= 9:
		return "A"

	case nota >= 7 && nota < 9:
		return "B"

	case nota >= 5 && nota < 7:
		return "C"

	case nota >= 3 && nota < 5:
		return "D"

	case nota >= 0 && nota < 3:
		return "E"

	default:
		return "Nota inválida!"
	}
}

func main() {
	fmt.Println(notaConceito(10.0))
	fmt.Println(notaConceito(2.0))
	fmt.Println(notaConceito(-1.0))
	fmt.Println(notaConceito(5.0))
}
