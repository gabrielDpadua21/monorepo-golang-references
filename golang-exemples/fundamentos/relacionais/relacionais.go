package main

import (
	"fmt"
	"time"
)

func main() {
	// Operadores relacionais
	fmt.Println("Strings:", "frajola" == "frajola")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1)
	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Data com Equals", d1.Equal(d2))

	type pessoa struct {
		nome string
	}

	p1 := pessoa{"frajola"}
	p2 := pessoa{"frajola"}

	fmt.Println(p1)
	fmt.Println(p1 == p2)

}
