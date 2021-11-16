package main

import "fmt"

func inc1(n int) {
	n++
}

// ponteiro e representado por *
func inc2(n *int) {
	//* e usado para desreferenciar, ou seja
	// fazer o ponteiro ter acesso a valor no qual aponta
	*n++
}

func main() {
    n := 1

    inc1(n) // passagem por valor

    fmt.Println(n)

    // & e usado para obter o endereco da variável
    inc2(&n) // passagem por referencia -> evitar

    fmt.Println(n)
}
