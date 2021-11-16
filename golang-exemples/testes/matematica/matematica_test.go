package matematica

import "testing"

const erroPadrao = "Valor esperado %v mais foi encontrado %v"

func TestMedia(t *testing.T) {
	valorEsperado := 8.0
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
