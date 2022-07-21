## Diretrizes para mensagens de commit

Mensagens de commits bem redigidas são a melhor maneira de comunicar o contexto de uma alteração para outros desenvolvedores que também estejam trabalhando no projeto. Portanto, antes de enviar sua contribuição, leia as instruções abaixo.

### Formato do commit

Cada commit é formado por um cabeçalho, um corpo e um rodapé. O cabeçalho, além de possuir o tipo, inclui também o número da issue, o nome da branch de trabalho e o assunto.

```
<type>(<issue>/<branch>) : <assunto>
<linha em branco>
<body>
<linha em branco>
<footer>
```

O cabeçalho é **obrigatório**, bem como o número da issue, o nome da branch e o assunto. O corpo e o rodapé são **opcionais**.

O assunto do commit não pode ter mais de *80* caracteres, permitindo que a mensagem seja fácil de ser lida. Deve começar com letra maiúscula e não terminar com pontuação. Use o tom imperativo para descrever o que o commit faz, em vez do que ele fez.

* Nem todos os commits são complexos o suficiente para garantir um corpo, portanto, seu uso é opcional. Use o corpo para explicar o **quê** e o **porquê**, também no tom imperativo. A linha em branco é obrigatória e o comprimento de cada linha não pode passar de 80 caracteres.

* O rodapé é opcional e é usado para fazer referências.

#### Tipos

*fix*: a resolução de um bug

*style*: recurso e atualizações relacionadas à estilização

*refactor*: refatoração de uma seção específica da base de código

*test*: tudo o que for relacionado a testes

*docs*: tudo o que for relacionado à documentação

*chore*: manutenção regular do código

_**Exemplo**_

```
refactor(7/refactor-edicao-colaboradores) : Reescreve edição de colaboradores

Fix #2
```