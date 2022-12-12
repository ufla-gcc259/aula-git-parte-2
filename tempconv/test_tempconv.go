package tempconv

import (
	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
	"testing"
)

func CelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius  tempconv.Celsius
		fahrenheit tempconv.Fahrenheit
	} {
		{ celsius: 0, fahrenheit: 32 },
		{ celsius: 50, fahrenheit: 122 },
		{ celsius: -50, fahrenheit: -58 },
	}
	for _, test := range tests {
		result := tempconv.CToF(test.celsius)
		if result != test.fahrenheit {
			t.Fatalf("Era esperado o valor %g em fahrenheit, mas foi recebido %g", test.fahrenheit, result)
		}
	}
}

func FahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		fahrenheit  tempconv.Fahrenheit
		celsius tempconv.Celsius
	} {
		{ fahrenheit: 32, celsius: 0 },
		{ fahrenheit: 122, celsius: 50 },
		{ fahrenheit: -58, celsius: -50 },
	}
	for _, test := range tests {
		result := tempconv.FToC(test.fahrenheit)
		if result != test.celsius {
			t.Fatalf("Era esperado o valor %g em celsius, mas foi recebido %g", test.celsius, result)
		}
	}
}