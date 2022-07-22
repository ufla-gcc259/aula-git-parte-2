package main

import (
	"fmt"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

//App simples que testa algumas funcionalidades do pacote tempconv
func main() {
	//Valor que sera lido pelo terminal para ser verificado
	var valor float64
	fmt.Printf("Digite um valor em °Celsius: ")
	_, erro := fmt.Scanf("%f", &valor)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	//Variavel do tipo Celsius
	celsius := tempconv.Celsius(valor)
	// Comparações utilizando as constantes definidas no pacote
	// Primeira comparação é se o valor é valido
	if celsius < tempconv.AbsoluteZeroC {
		fmt.Println("Temperatura invalida")
		return
		// Segunda comparação é se a temperatura é o zero absoluto
	} else if celsius == tempconv.AbsoluteZeroC {
		fmt.Println("Está zero absoluto ")
		// Terceira é se a temperatura esta congelando em relação a água
	} else if celsius <= tempconv.FreezingC {
		fmt.Println("Está congelando")
		// Quarta é se a temperatura esta entre o congelamento e a ebulição em relação a água
	} else if celsius < tempconv.BoilingC {
		fmt.Println("Temperatura normal")
		// Ultima é se esta no ponto de ebulição da água
	} else {
		fmt.Println("Está evaporando")
	}
	// Por fim é exibido na tela a temperatura convertida em fahrenheit
	fmt.Printf("Temperatura em fahrenheit: %v\n", tempconv.CToF(celsius))
}
