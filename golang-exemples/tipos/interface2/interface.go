package main

import "fmt"

type esportivo interface {
	modoTurbo()
}

type ferrari struct {
	modelo string
	turboLigado bool
	velocidadeAtual int
}

func (f *ferrari) modoTurbo() {
	f.turboLigado = true
}


func main() {
	carro1 := ferrari{"f40", false, 0}
	carro1.modoTurbo()

	var carro2 esportivo = &ferrari{"f50", false, 0}
	carro2.modoTurbo()

	fmt.Println(carro1)
	fmt.Println(carro2)
}
