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
### Estrutura Geral da Mensagem:

- **Tipo:**
  - Use um dos seguintes tipos: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`.
  - Exemplo: `feat:` para uma nova feature.

- **Âmbito (opcional):**
  - Indique o âmbito da alteração dentro do projeto.

- **Breve Descrição (linha de assunto):**
  - Comece com uma letra maiúscula.
  - Seja conciso e objetivo (máximo de 50 caracteres).
  - Use o tempo presente.

- **Corpo da Mensagem:**
  - Forneça uma descrição detalhada da alteração.
  - Mantenha linhas com cerca de 72 caracteres.

- **Mensagem de Rodapé (opcional):**
  - Adicione informações sobre quebras de compatibilidade, problemas fechados, etc.
  - Use palavras-chave como `Closes`, `Fixes`, `Resolves`, seguidas pelo número do problema.

### Exemplo de Mensagem de Commit:

```bash
git commit -m "docs(commit-guidelines): Adiciona diretrizes para mensagens de commit

- Adiciona seção sobre a estrutura geral das mensagens de commit.
- Inclui exemplos e explicações detalhadas para cada parte.
- Fornecer um guia claro para manter um histórico de commits organizado.

Closes #número_da_issue"
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
