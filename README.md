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

Diretrizes para Mensagens de Commit
----

1. **Mantenha as mensagens curtas e descritivas**: As mensagens de commit devem ser concisas e transmitir claramente o propósito do commit.

2. **Use um prefixo para categorizar os commits**: Você pode usar prefixos para categorizar seus commits, como "feat" para novos recursos, "fix" para correções de bugs, "docs" para atualizações de documentação, "refactor" para refatorações de código, "test" para testes e assim por diante. Isso torna mais fácil entender rapidamente a natureza do commit.

3. **Comece com um verbo no imperativo**: Inicie a mensagem de commit com um verbo no imperativo, como "Adicionar", "Corrigir", "Atualizar", "Remover", etc. Isso ajuda a indicar o que o commit faz.

4. **Forneça detalhes no corpo da mensagem (opcional)**: Se necessário, forneça informações adicionais no corpo da mensagem para explicar o motivo por trás do commit, incluindo links para problemas ou solicitações de pull relacionados.

5. **Referencie problemas e solicitações de pull**: Se o commit estiver relacionado a um problema específico no GitHub, você pode fazer referência a ele usando "#NúmeroDoIssue" na mensagem de commit, por exemplo, "Corrigir #123: Problema com a conversão de Fahrenheit para Celsius".
