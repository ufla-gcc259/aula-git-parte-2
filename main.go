package main

import (
	"fmt"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	var kelvinValue tempconv.Kelvin

	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scanln(&kelvinValue)

	kelvinToCelsius := tempconv.KToC(kelvinValue)
	kelvinToFahrenheit := tempconv.KToF(kelvinValue)

	fmt.Printf("%.2f K é igual a %.2f°C e %.2f°F\n", kelvinValue, kelvinToCelsius, kelvinToFahrenheit)
}
