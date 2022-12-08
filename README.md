# _tempconv_

Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan.

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

> _Todo_

## Licença

> _Todo_

## Como contribuir?

Escolha uma _issue_ dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um _Pull Request_, quando terminar.

## Como deve ser os commits

Cada mensagem de commit deve conter um cabeçalho, um corpo e um rodapé
confirmação consiste em um cabeçalho, um corpo e um rodapé.

```
<cabecalho>
<Pular linha>
<corpo>
<Pular linha>
<rodapé>
```

O cabeçalho é obrigatório. Qualquer linha da mensagem de confirmação não pode ter mais de 100 caracteres! Isso permite que a mensagem seja mais fácil de ler no GitHub, bem como em várias ferramentas git.

Exemplo:

```
Covertendo a temperatura de celsius para fahrenheit

Acrescentando a melhoria de converter temperatura de celcius para fahrenheit.
Para isso foi feito uma funcao que recebe o valor de celsius e retorna esse
valor em fahrenheit.

Melhorias nº 1234
```

### Revert

Se o commit reverter um commit anterior, ele deve começar com revert: , seguido pelo cabeçalho do commit revertido. No corpo deve dizer: Isto reverte o commit (número do commit).

Exemplo:

```
Revert: Covertendo a temperatura de celsius para fahrenheit

Isso reverte o commit Melhorias nº1234
Acrescentando a melhoria de converter temperatura de celcius para fahrenheit.
Para isso foi feito uma funcao que recebe o valor de celsius e retorna esse
valor em fahrenheit.

Revert nº 12
```

### Pacote

A mensagem no corpo de commit deve especificar qual pacote é afetado pela mudança. Por exemplo: A alteração impacta no pacote cadatro_temp.

### Corpo

Use o imperativo, tempo presente. O corpo deve incluir a motivação para a mudança e compará-la com o comportamento anterior.

### Rodape

O rodapé deve conter todas as informações sobre alterações recentes ou reprovações e também é o local para fazer referência aos problemas do GitHub que este commit fecha.
