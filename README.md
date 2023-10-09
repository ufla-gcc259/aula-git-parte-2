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


# Diretrizes para Mensagens de Commit

Para manter um histórico de commits organizado e informativo, siga as seguintes diretrizes ao fazer commits neste projeto:

## Tipo de Commit:

- **feat:** Adiciona uma nova funcionalidade ou recurso.
- **fix:** Corrige um erro ou bug existente.
- **docs:** Alterações exclusivamente na documentação.
- **refactor:** Refatorações no código sem alterações de funcionalidade.
- **perf:** Melhorias de desempenho sem alterações de funcionalidade.
- **test:** Alterações relacionadas aos testes do projeto.
- **build:** Mudanças no sistema de construção (build) ou ferramentas.
- **ci:** Mudanças na configuração de integração contínua.
- **release:** Ponto de lançamento no repositório.

## Pacote:

Especifique qual pacote ou componente do projeto está sendo afetado pela alteração. Por exemplo, "material/button".

## Escopo:

Descreva o local ou a parte específica do pacote que está sendo alterada. Por exemplo, "material/datepicker", "cdk/dialog", etc.

## Assunto do Commit:

Use o tempo presente e a voz imperativa. Exemplo: "adiciona função de conversão de temperatura".

- Evite capitalizar a primeira letra.
- Não coloque ponto final no final do assunto.

### Exemplo de mensagem de commit:

```bash
feat(material/button): adiciona estilo hover aos botões

Adiciona um estilo de hover aos botões do componente Material. Agora os botões mudam de cor quando o cursor passa sobre eles.

Fecha #123
Corpo do Commit:
Forneça uma explicação detalhada da alteração, incluindo a motivação por trás dela e como ela se compara ao comportamento anterior.

Rodapé do Commit:
Use o rodapé para listar informações sobre mudanças significativas ou obsoletas (Breaking Changes ou Deprecations).

Para Breaking Changes, inicie com "BREAKING CHANGE:" seguido das alterações relevantes.
Para Deprecations, inicie com "DEPRECATED:" seguido das informações relevantes.
Use o rodapé para fazer referência a problemas do GitHub que foram resolvidos ou afetados por este commit.





