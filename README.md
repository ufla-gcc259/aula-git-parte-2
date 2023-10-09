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

## Diretrizes para Mensagens de Commit

Este documento descreve as diretrizes para mensagens de commit, inspiradas nas diretrizes de mensagens de commit do projeto Angular Material - versão seção Commit Message Guidelines.

### Estrutura da Mensagem de Commit

Cada mensagem de commit deve seguir a seguinte estrutura:

> tipo(<âmbito>): mensagem


- `<tipo>`: Indica a natureza do commit. Deve ser um dos seguintes:
  - `feat`: Novo recurso (feature).
  - `fix`: Correção de bug.
  - `docs`: Atualização na documentação.
  - `style`: Mudanças que não afetam o comportamento do código (formatação, espaçamento, etc.).
  - `refactor`: Refatoração de código.
  - `test`: Adição ou modificação de testes.
  - `chore`: Tarefas de manutenção geral.

- `<âmbito>` (opcional): Indica a área ou módulo do projeto afetado pelo commit.

- `<mensagem>`: Uma breve descrição do que o commit faz, em até 72 caracteres.

### Exemplos de Mensagens de Commit

Aqui estão alguns exemplos de mensagens de commit seguindo estas diretrizes:

> feat(button): Adiciona novo botão de envio no formulário de contato
> 
> fix(navbar): Corrige problema de sobreposição no menu de navegação


### Boas Práticas

- Mantenha as mensagens de commit curtas, claras e descritivas.
- Use o tempo presente ("Adiciona" em vez de "Adicionou") na mensagem.
- Evite incluir detalhes técnicos na mensagem; eles podem estar na descrição do commit, se necessário.
- Use letras minúsculas para `<tipo>` e `<âmbito>`, separando palavras com hífen.

Lembrando que aderir a diretrizes de mensagens de commit consistentes ajuda a manter o histórico de versionamento do projeto organizado e fácil de entender.






Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
