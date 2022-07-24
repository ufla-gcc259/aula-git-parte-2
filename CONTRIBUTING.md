# Contribuindo com o tempconv

Toda contribuição é bem-vinda para ajudar e melhorar o **tempconv**. A seguir, os tópicos com instruções importantes
sobre como contribuir com este projeto.

  - [Mensagens de Commit](#commit)

## <a name="commit"></a> Mensagens de Commit

Para regularizar as mensagens de commit, siga as instruções e padrões apresentados aqui.

### Estrutura do commit

O commit é composto por: cabeçalho (header), corpo (body) e rodapé (footer), separados por espaços em brancos entre si
(blank line), sendo o cabeçalho obrigatório, já o corpo e rodapé quando for necessário.

```
<header>
<blank_line>
<body>
<blank_line>
<footer>
```

Exemplo:
```
Adiciona condicional para determinar conversão de temperatura escolhida

Adiciona trecho de condicional para determinar quais temperaturas escolhidas para a conversão.

<comando_para_issue>
```

O **cabeçalho** deve ser simples e direto, limitado a 72 caracteres, também é o único campo obrigatório. Deve descrever
a ação feita.

Porém, se o cabeçalho não for o suficiente para descrição e você desejar explicar melhor o que foi feito, ou porquê,
você pode utilizar o **corpo** para isso. Adiciona uma linha branca entre os campos e adiciona seu texto. Cada linha é
limitada a 120 caracteres.

Por fim, caso aplique uma ação sobre a sua issue de contribuição ao projeto, você pode adicionar uma ação no campo
**rodapé**. Basta adicionar uma linha em branco entre o esse campo e o anterior, seja o **cabeçalho** ou o **corpo**
do seu commit.
