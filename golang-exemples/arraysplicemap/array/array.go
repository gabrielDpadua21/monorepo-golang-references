package main

import "fmt"

func main() {
	// arrays sao estruturas homogeneas e imutaveis de tamanho

	var gatos [3]string

	fmt.Println(gatos)

	gatos[0], gatos[1], gatos[2] = "frajola", "thor", "lucifer"

	fmt.Println(gatos)

	msg := "Tenho 3 gatos eles sao: "

	for i := 0; i < len(gatos); i++ {
		msg += gatos[i] + ", "
	}

	fmt.Println(msg)
}
