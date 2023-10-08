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

## Mensagem de Commit
Cada mensagem de confirmação consiste em um tipo, escopo, descrição e issue. 

```php
<tipo>(<escopo>): <descrição> (<issue>)
```

**Aqui está uma explicação de cada parte:**


> O **tipo** do commit, como "feat" (para novos recursos), "fix" (para correções de bugs), "docs" (para atualizações na documentação), "style" (para mudanças de estilo de código),etc.
> O **escopo** opcional descreve a parte específica do projeto que foi afetada pela alteração. Isso pode ser um nome de módulo, um componente ou outra divisão relevante do projeto.
> Uma breve **descrição** da alteração.
> O número referente à **issue** que está resolvendo.
**Aqui estão alguns exemplos:**

* feat(login): Adiciona autenticação via OAuth2 (#123)
* fix(bug-form): Corrige validação de email no formulário (#124)
* docs(readme): Atualiza instruções de instalação (#125)
* style(css): Refatora estilos de botões (#126)

**Para manter a consistência, vamos definir alguns padrões:**

1. Use verbos no imperativo, como: *("Adiciona", "Corrige", "Atualiza", "Remove" em vez de "Adicionou", "Corrigindo", "Atualizando", "Removendo").*
2. Seja conciso e descritivo: A descrição deve ser clara e concisa, comunicando o que a alteração faz.
3. Se uma alteração precisar de mais detalhes, você pode adicionar uma mensagem mais longa abaixo da issue.
4. Mantenha uma mensagem de commit por linha: Evite mensagens de commit muito longas ou quebradas em várias linhas. Cada mensagem deve ser uma única frase.