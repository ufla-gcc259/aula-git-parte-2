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

Diretrizes para mensagens de commit
-----

* As mensagens devem ser simples
* As mensagens devem ser sucintas (100 caracteres)
* As mensagens devem dizer se o commit é referente a uma:
  * melhoria 
  * correção de bug 
  * teste 
  * documentação 
  * refatoração 
  * etc...

- As mensagens devem mencionar uma Issue específica
- Caso o commit seja para reverter outro commit, ele deve ser mencionado
- As mensagens devem evitar vocabulário vago
- As mensagens devem mencionar o escopo do commit

Exemplo
> **(Issue 7)** Documentação: este commit adiciona as diretrizes de commit do projeto 
