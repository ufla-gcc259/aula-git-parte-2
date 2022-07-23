# _tempconv_

Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan.

---

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

---

## API Pública

#### Está disponível uma api para realizar as conversões

[Acesse aqui](https://api-convtemp.herokuapp.com/) ou cole https://api-convtemp.herokuapp.com/

- Para utilizar basta acrescentar a escala **atual** e o valor que deseja realizar a conversão, por exemplo:

  - `https://api-convtemp.herokuapp.com/celsius/10` [( testar )](https://api-convtemp.herokuapp.com/celsius/10)
  - `https://api-convtemp.herokuapp.com/fahrenheit/10` [( testar )](https://api-convtemp.herokuapp.com/fahrenheit/10)
  - `https://api-convtemp.herokuapp.com/kelvin/10` [( testar )](https://api-convtemp.herokuapp.com/kelvin/10)

- A resposta virá em formato JSON, seguindo o modelo

```json
{
  "celsius": 10,
  "fahrenheit": 50,
  "kelvin": 283.15
}
```

## Contribua com a API clicando [**aqui**](https://github.com/ThomazSIUFLA/aula-git-parte-2)

**Ou acesse https://github.com/ThomazSIUFLA/aula-git-parte-2**

---

## Outras linguagens?

Versões da biblioteca _tempconv_ para outras linguagens:

> _Todo_

---

## Licença

> _Todo_

---

## Como contribuir?

Escolha uma _issue_ dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um _Pull Request_, quando terminar.
