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


Diretrizes para mensagens de commit
---

Antes de enviar seu commit é importante atentar às regras que serão descritas a seguir. Elas servem para manter mensagens mais legíveis e fáceis de seguir ao examinar o histórico do projeto. O commit deverá ter o seguinte formato:

```
<tipo>(<escopo>): <descrição>
<linha em branco>
[corpo opcional]
<linha em branco>
[rodapé opcional]

```

Onde: 

* **Tipo**: especifica a atividade que foi realizada.

* **Escopo**: parte do código que sofreu a alteração.

* **Descrição**: o que foi feito no commit.

* **Corpo**: detalhes sobre o que foi feito no commit.

* **Rodapé**: informa issue, id ou task utilizada na alteração do código commitado.


Os tipos aceitos são:

* *docs*: quando for alterado algo na documentação;

* *feat*: quando for adicionada uma nova funcionalidade;

* *fix*: quando um bug for corrigido;

* *perf*: quando a alteração melhorou o desempenho;

* *refactor*: quando for realizada uma mudança no código que não afetou funcionalidades;

* *style*: quando forem realizadas mudanças na formatação do próprio código.

Observações:

* A mensagem deve conter menos de 150 caracteres no total; deve-se utilizar sempre o modo imperativo;

* A primeira linha não pode conter mais do que 70 caracteres; a segunda linha será sempre vazia;

* As demais linhas devem conter até 80 caracteres;

* O tipo e o escopo devem ser sempre em caixa baixa;

* O escopo pode ser vazio em casos que os parênteses são omitidos.

Exemplo:

```
fix(button): impossível desmarcar seleção do botão do tipo radio

Corrige um bug no componente `button` onde os botões do tipo `radio` não podem ser desabilitados devido a um problema na classe.

Fixes #12

```