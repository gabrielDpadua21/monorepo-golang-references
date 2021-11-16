package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array - tamanho fixo
	arr := [3]int{1, 2, 3}

	// slice - tamanho variavel
	slic := []int{1, 2, 3}

	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(slic))

	// SLICE NAO E UM ARRAY ...

	gatos := [3]string{"frajola", "thor", "lucifer"}

	gatinhos := gatos[0:2]

	gatoes := gatos[1:3]

	fmt.Println(gatos)

	fmt.Println(gatinhos)

	fmt.Println(gatoes)
}
