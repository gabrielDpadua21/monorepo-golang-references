package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// Nao e possivel dar um append de uma array dentro de outro
	// so como slice
	//array2 := append(array1, 1, 2, 3)

	slice1 = append(slice1, 1, 2, 3)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	// copia os valores do slice1 para o slice2
	// o slice 2 nao precisa ser nessaria mente do tamanho 1
	copy(slice2, slice1)

	fmt.Println(slice2)
}
