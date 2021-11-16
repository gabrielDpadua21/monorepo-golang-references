package main

import "fmt"

type carro struct {
	nome string
	velocidadeAtual int
}

type ferrari struct {
	carro // campo anonimo -> composicao
	turboLigado bool	
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 100
	f.turboLigado = true

	fmt.Printf("A ferrari %s esta com turbo ligado? %v", f.nome, f.turboLigado)
	fmt.Println(f)
}
