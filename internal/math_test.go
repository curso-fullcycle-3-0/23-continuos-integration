package internal

import (
	"testing"
)

func TestSoma(t *testing.T) {
	total := Calc(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invsálido: Resultado %d Esperado: %d", total, 30)
	}
}
