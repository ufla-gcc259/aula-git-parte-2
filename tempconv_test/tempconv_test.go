package tempconv_test

import (
	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
	"testing"
)

func TestCToFAccept(t *testing.T) {
	testes := []struct {
		Valores  tempconv.Celsius
		Esperado tempconv.Fahrenheit
	}{
		{Valores: 0, Esperado: 32},
		{Valores: 100, Esperado: 212},
		{Valores: -5, Esperado: 23},
	}
	for _, teste := range testes {
		resultado := tempconv.CToF(teste.Valores)
		if resultado != teste.Esperado {
			t.Fatalf("Valor esperado: %g - Valor retornado: %g", teste.Esperado, resultado)
		}
	}
}

func TestFToFAccept(t *testing.T) {
	testes := []struct {
		Valores  tempconv.Fahrenheit
		Esperado tempconv.Celsius
	}{
		{Valores: 32, Esperado: 0},
		{Valores: 212, Esperado: 100},
	}

	for _, teste := range testes {
		resultado := tempconv.FToC(teste.Valores)
		if resultado != teste.Esperado {
			t.Fatalf("Valor esperado: %g - Valor retornado: %g", teste.Esperado, resultado)
		}
	}
}
