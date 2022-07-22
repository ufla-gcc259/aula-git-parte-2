*tempconv*
=====
Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan. 

Como usar?
----
Faça o download da biblioteca:

`go get https://github.com/ufla-gcc259/aula-git-parte-2@v1.0.1`

Pronto, agora é só usar:
```go
package main

import (
	"fmt"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	fmt.Printf("Que frio! %v\n", tempconv.AbsoluteZeroC) // Que frio! -273.15°C
	fmt.Printf("Fervendo! %v\n", tempconv.CToF(tempconv.BoilingC)) // Fervendo! 212°F
}
```

Exemplo de Utilização:
----
```go
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
```


Outras linguagens?
----
Versões da biblioteca *tempconv* para outras linguagens:

> *Todo*


Licença
-----

> *Todo*


Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
