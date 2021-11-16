package main

// Label para o import
import (
	f "fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415 // constante com tipo e valor
	var raio = 3.4            // Variavel com tipo (float64) inferido pelo compilador

	// forma reduzida de criar variaveis

	area := PI * m.Pow(raio, 2)

	f.Println("Area de circuferencia e:", area)

	// Blocos de construcao

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	f.Println(a, b, c, d)
}
