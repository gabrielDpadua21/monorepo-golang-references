package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - indices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r e show", "Cod3r", 0},
		{"", "", 0},
		{"Frajola", "Fraj", 1},
		{"Opa", "opa", -1},
	}

	for _, teste := range testes {
		t.Logf("Massa %v", testes)

		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}

}
