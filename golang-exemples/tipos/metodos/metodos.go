package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome string
	sobrenome string
}

// leitura
func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// gravacao
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
} 

func main() {
	p1 := pessoa{"Frajola", "Gatinho"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Thor Chato")
	fmt.Println(p1.getNomeCompleto())
}



