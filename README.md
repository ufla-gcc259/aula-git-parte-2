*tempconv*
=====
Biblioteca escrita em Go para conversões simples de temperatura.

Como usar?
----
Faça o download da biblioteca:

`go get https://github.com/ufla-gcc259/aula-git-parte-2`

Pronto, agora é só usar:
```go
package main

import (
	"fmt"

	"github.com/paulojunior-ufla/aula-git-parte-2/tempconv"
)

func main() {
	fmt.Printf("Que frio! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Fervendo! %v\n", tempconv.CToF(tempconv.BoilingC))
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
