package main

import (
	"fmt"
)

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}

	slice1 := array1[1:3]

	slice2 := slice1[:2]

	fmt.Println(slice1)
	fmt.Println(slice2)
}
