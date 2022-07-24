# Diretrizes para um bom commit

Os commits realizados devem ser semânticos (conventional commit), ou seja, devem ser simples e claros quanto ao seu objetivo e também devem ser atômicos (tratar de apenas um objetivo / issue)

# Estrutura

    - tipo[escopo opcional]: <descrição> 
    - corpo opcional
    - rodapé opcional

Quanto mais completo um commit, maior será o tempo de processamento. Portanto, recomenda-se seu uso simplificado, não utilizando suas partes opcionais. Há casos específicos que necessitam um maior descritivo do processo executado.

## Tipos
A primeira e principal descrição de um commit semântico, refere-se a seu tipo, os quais possuem a finalidade de comunicar a intenção de processamento que o utilizador teve em sua execução.

> Abaixo será enumerado os principais types descritos na documentação do
> _Angular Commit Message Guidelines_:

1. **build:**  Alterações que afetam o sistema de construção ou dependências externas (escopos de exemplo: gulp, broccoli, npm),
2.  **ci:** Mudanças nos arquivos de configuração e scripts de integração contínua (CI).
3.  **docs:**  referem-se a inclusão ou alteração somente de arquivos de documentação;
4.  **feat:**  Tratam adições de novas funcionalidades ou de quaisquer outras novas implantações ao código;
5.  **fix:**  Essencialmente definem o tratamento de correções de bugs;
6.  **perf:**  Uma alteração de código que melhora o desempenho;
7.  **refactor:**  Tipo utilizado em quaisquer mudanças que sejam executados no código, porém não alterem a funcionalidade final da tarefa impactada;
8.  **style:**  Alterações referentes a formatações na apresentação do código que não afetam o significado do código, como por exemplo: espaço em branco, formatação, ponto e vírgula ausente etc.);
9.  **test:** Adicionando testes ausentes ou corrigindo testes existentes nos processos de testes automatizados (TDD);
10.  **chore:**  Atualização de tarefas que não ocasionam alteração no código de produção, mas mudanças de ferramentas, mudanças de configuração e bibliotecas que realmente não entram em produção;
11.  **env:**  basicamente utilizado na descrição de modificações ou adições em arquivos de configuração em processos e métodos de integração contínua (CI), como parâmetros em arquivos de configuração de containers.

> Também, o Guidelines, recomenda o tipo improvement para commits que melhoram uma implementação atual sem adicionar um novo recurso ou consertar um bug.

## Descrições

As descrições se trata de um dos parâmetros obrigatórios da sintaxe, e devem ser preferencialmente apresentadas com letras minúsculas, e obrigatoriamente iniciando com letra minúscula, suportando letras maiúsculas em descrição de arquivos ou classes específicas.

**Também devem ser suficientemente claras, utilizando seu espaço até o máximo permitido da linha.**

## Corpo e Rodapé
O corpo, por sua vez, caracteriza-se por ser opcional e raramente recomendado sua utilização. Resumindo-se em casos que necessitem de uma apresentação mais completa do conteúdo implantado no commit, o qual a descrição anterior não foi o suficiente para elucidar e esclarecer.

O rodapé por sua vez, também não possui uso obrigatório. Restringindo-se às alterações de estado via smart  _commit_, como resoluções de problemas (issues), através de chamados de atendimentos, ou sprints de projetos de implantação os quais podem ser descritos no rodapé.

Pode ser fornecido um ou mais rodapés, o primeiro sempre iniciando uma linha em branco após o corpo. Cada rodapé deve consistir em um token de palavra, seguido pelo símbolo  **“:”**  (dois pontos) e posteriormente um espaço em branco e o símbolo  **“#”**  (sustenido) como separador da string descritiva do rodapé (conceito inspirado na convenção do Git Trailer).

O token de um rodapé DEVE usar o símbolo  **“-”** (hífen) no lugar de caracteres de espaço em branco, por exemplo, **Reviewed-by**, permitindo uma diferenciação de um rodapé em relação a um corpo com vários parágrafos.

```
fix: correct minor typos in code
see the issue for details
on typos fixed.
Reviewed-by: Elisandro Mello
Refs #133
```

