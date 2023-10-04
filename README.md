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

# Diretrizes para mensagens de Commit

## Formato de Mensagem de Commit
Cada mensagem de commit deve seguir o formato padrão composto por um cabeçalho, um corpo e um rodapé. O cabeçalho é obrigatório e deve conter um tipo, um pacote (opcional), um escopo (opcional) e um assunto. O formato é o seguinte:

<type>(<package>/<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>

Os campos package e scope são obrigatórios para alterações mostradas no changelog (fix, feat, perf, revert). Caso contrário, eles podem ser omitidos se a alteração não afetar um pacote específico e não for exibida no changelog (por exemplo, alterações de compilação ou refatorações).

## Tamanho Máximo da Linha
Nenhuma linha da mensagem de commit deve exceder 100 caracteres. Isso garante que a mensagem seja legível no GitHub e em outras ferramentas Git.

## Exemplo de Mensagem de Commit
fix(material/button): impossibilidade de desativar o botão 

Corrige um bug no componente `button` do Angular Material em que os botões
não podem ser desativados através de uma ligação. Isso ocorre porque a entrada
`disabled` não define a classe `.mat-button-disabled` no elemento hospedeiro.

Corrige #1234



## Revert 
Se o commit reverter um commit anterior, ele deve começar com "revert:" seguido pelo cabeçalho do commit revertido. O corpo do revert deve conter a seguinte frase: "This reverts commit <hash>.", onde <hash> é o SHA do commit que está sendo revertido.

## Tipo 
O campo "type" deve ser um dos seguintes valores:

- feat: Cria um novo recurso.
- fix: Corrige uma falha/bug descoberto anteriormente.
- docs: Mudanças que afetam exclusivamente a documentação.
- refactor: Refatoração sem alteração na funcionalidade ou API (inclui alterações de estilo).
- perf: Melhora o desempenho sem alteração na funcionalidade ou API.
- test: Melhorias ou correções feitas no conjunto de testes do projeto.
- build: Mudanças no sistema de compilação do repositório local e nas ferramentas.
- ci: Mudanças na configuração do CI e nas ferramentas específicas do CI.
- release: Um ponto de lançamento no repositório.

## Assunto 
O campo "subject" deve conter uma descrição sucinta da mudança. Deve seguir as seguintes diretrizes:

- Use o imperativo presente, por exemplo, "mudar" e não "mudou" ou "muda".
- Não coloque a primeira letra em maiúscula.
- Não inclua um ponto (.) no final.

## Corpo 
Assim como no assunto, use o imperativo presente, por exemplo, “mudar” e não “mudou” ou “muda”. O corpo deve incluir a motivação para a mudança e compará-la com o comportamento anterior.

## Rodapé
O rodapé deve conter informações sobre alterações recentes ou descontinuações. É também o local para fazer referência aos problemas do GitHub que este commit fecha. Alterações significativas devem começar com a palavra "BREAKING CHANGE:" seguida de um espaço ou duas novas linhas. O restante da mensagem de commit é usado para descrever a quebra de compatibilidade. Depreciações devem começar com a palavra "DEPRECATED:". O restante da mensagem de commit é usado como conteúdo da nota.


Seguindo essas diretrizes, você poderá criar mensagens de commit consistentes e informativas para o projeto.