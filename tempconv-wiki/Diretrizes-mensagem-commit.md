# Diretrizes para a mensagem do commit

Estas diretrizes tem como objetivo padronizar os commits para além de facilitar descobrir do que se trata cada commit e rastrear quais issues foram solucionadas a partir daquele commit, ajudar os contribuidores a entender o caminho pelo qual a aplicação seguiu apenas lendo os textos dos commits.

## Compondo a mensagem do commit

Utilizaremos como base as diretrizes de Linus Torvalds que comentou um modelo de uma boa mensagem de commit:

> Cabeçalho: Explique esse commit em uma linha (Use a linguagem imperativa)
> O corpo da mensagem de commit são algumas linhas de texto, explicando em mais detalhes e possivelmente apresentando mais contexto sobre o 
problema sendo tratado.

> O corpo da mensagem de commit pode ser diversos paragráfos e por favor façam corretamente a quebra de linha e mantenham as colunas menos  que 74 caracteres. Assim, o comando "git log" irá mostrar a mensagem de forma agradavel mesma que esteja identada.

> Faça questão de explicar a sua solução e por que voce está fazendo o que está fazendo, ao invés de apenas descrever o que está fazendo de forma superficial. Pense que revisores e o seu eu-futuro irão ler essas mudanças, mas podem não entender por que determinada solução foi implementada.

> Reported-by: quem-reportou
> Signed-off-by: Seu Nome [email@host.com](mailto:email@host.com)

### Estrutura mensagem de commit

```
[Cabeçalho]*

[Corpo]

[Rodapé]
```

Os que estão marcados com * são obrigatórios.

**Cabeçalho** terá a seguinte estrutura:

`<tipo>: <descrição>`

Os tipos possíveis serão:

 - _fix_: corrige um bug no seu código.
 - _feat_: adiciona uma nova funcionalidade no seu código.
 - _chore_: Mudanças de configuração ou de código que não entra em produção;
 - _docs_: Mudanças na documentação;
 - _style_: Alteração apenas no estilo do código, sem mudança de algoritmo;
 - _refactor_: Refatoração de determinado bloco de código;
 - _perf_: Alterações que impactam o desempenho da aplicação;
 - _test_: Mudanças na estrutura ou na forma de testar o projeto.

**Corpo**: Uma descrição do que foi feito no commit

**Rodapé**: Identificará a issue na qual este commit se refere, ou se é uma grande mudança no código.

#### BREAKING CHANGE
um commit que introduz uma grande mudança no código, deverá apresentar um ! depois do seu tipo ou então no rodapé da mensagem como nos exemplos abaixo.
Exemplo 1:
```
chore!: Atualização para a versão go1.18
```

Exemplo 2:
```
chore: Atualização para a versão go1.18

BREAKING CHANGE: a versão minima do go agora é a 1.18.
```

### Exemplo

```
docs: Adicionado diretrizes para a mensagem do commit
    
Incluído no wiki do projeto diretrizes para a comunidade sobre como escrever suas mensagens de commit

Ref.: #7
```