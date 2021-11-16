package main

import "fmt"

func main() {
	var b byte = 3

	fmt.Println(b)

	// inferencia de tipo
	a := 2
	// mesma coisa que a = a + 1 ou a++
	a += 1
	a -= 1
	a /= 1
	a *= 1
	a %= 1

	fmt.Println(a)

	// atribuindo mais de uma varivel
	x, y := 2, 3
	fmt.Println(x, y)

	y, x = x, y
	fmt.Println(x, y)
}
