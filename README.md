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
### Exemplo de Uso
----
Você também pode conferir um exemplo de uso da biblioteca tempconv em um programa que converte temperaturas entre as escalas Celsius e Fahrenheit. O código de exemplo pode ser encontrado no arquivo example.go. Para executar o exemplo, siga estas etapas:
1. Clone este repositório:
```
git clone https://github.com/ufla-gcc259/aula-git-parte-2.git
```

2. Navegue até o diretório do projeto:
```
cd aula-git-parte-2
```

3. Compile e execute o exemplo:
```
go run example.go 100
``` 
Isso imprimirá o valor convertido nas escalas Celsius e Fahrenheit.

4. Experimente diferentes valores numéricos para ver as conversões. Aproveite a biblioteca tempconv para suas conversões de temperatura em Go!

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
