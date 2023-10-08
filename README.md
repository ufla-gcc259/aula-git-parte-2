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

## Padrão de commits

Para realizar commits no projeto, utilize o "Conventional Commits" em sua versão 1.0.0

Essa convenção traz alguns padrões de mensagens de commit, fornecendo um conjunto fácil de regras para criar um histórico explícito de commits.

A mensagem deve ser estruturada da seguinte forma:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

<strong>Elementos estruturais do commit:</strong>

1. <strong>fix:</strong> um commit do tipo fix corrige um bug em sua base de código;
2. <strong>feat:</strong> um commit do tipo feat introduz uma nova funcionalidade na base de código;
3. <strong>BREAKING CHANGE:</strong> um commit que possui um rodapé BREAKING CHANGE: ou anexa um ! após o tipo/escopo introduz uma mudança na API que quebra a compatibilidade;
4. Outros tipos além de fix: e feat: são permitidos, por exemplo, @commitlint/config-conventional (com base na convenção do Angular) recomenda build:, chore:, ci:, docs:, style:, refactor:, perf:, test:, e outros.
5. Tipos adicionais não são obrigatórios pela especificação Conventional Commits e não têm efeito implícito no Versionamento Semântico (a menos que incluam uma BREAKING CHANGE). Um escopo pode ser fornecido ao tipo de commit para oferecer informações contextuais adicionais e é contido entre parênteses, por exemplo, feat(parser): adicionar capacidade de analisar matrizes.

<strong>Exemplos:</strong>

```
feat: Add tests for the Kelvin scale
```

```
fix: Modify parameter of CToF function to fix calculation
```

Documentação completa do [Conventional Commits.](https://www.conventionalcommits.org/en/v1.0.0/#specification)

## Outras linguagens?

Versões da biblioteca _tempconv_ para outras linguagens:

> _Todo_

## Licença

> _Todo_

## Como contribuir?

Escolha uma _issue_ dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um _Pull Request_, quando terminar.
