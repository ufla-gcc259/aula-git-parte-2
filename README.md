# _tempconv_

Biblioteca escrita em C++ para conversões de temperatura. Desenvolvido como atividade para a disciplina de Software Livre.

## Como usar?

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

## Outras linguagens?

Versões da biblioteca _tempconv_ para outras linguagens:

[c++](https://github.com/IgorFreiredeMorais/aula-git-parte-2)

## Licença

> _Todo_

## Como contribuir?

Escolha uma _issue_ dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um _Pull Request_, quando terminar.
