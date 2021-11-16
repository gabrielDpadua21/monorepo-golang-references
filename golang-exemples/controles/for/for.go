package main

import (
	"fmt"
	"time"
)

func main() {
	// Formas de usar o laco for
	// Em go nao e possivel usar o while

	i := 1

	for i <= 10 {
		fmt.Printf("%d", i)
		i++
	}

	fmt.Printf("\n")

	for j := 0; j < 10; j++ {
		fmt.Printf("%d", j)
	}

	// Misturando estruturas
	fmt.Printf("\nMisturando\n")

	for y := 0; y < 10; y++ {
		if y%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// loop infinito
	for {
		fmt.Println("For ever...")
		time.Sleep(time.Second)
	}

}
