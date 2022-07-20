# Diretrizes para criação de branch

Antes de criar sua branch siga o padrão Conventional em seguido utilize o numero da issue:

features/NUMERO-ISSUE <br />
fix/NUMERO-ISSUE <br />
docs/NUMERO-ISSUE <br />
style/NUMERO-ISSUE <br />
refactor/NUMERO-ISSUE <br />
ci/NUMERO-ISSUE <br />
revert/NUMERO-ISSUE <br />
test/NUMERO-ISSUE <br />
build/NUMERO-ISSUE <br />
chore/NUMERO-ISSUE <br />

## Por exemplo:

```
features/7
```


# Diretrizes para mensagens de commit

Antes de enviar seu commit por favor siga as instruções abaixo para podermos estabelecer um padrão.

## Formato da mensagem do commit
Nesse caso iremos seguir o padrão Convencional commits.
Cada commit é necessário seguir um formato de cabeçalho para termos um padrão.  
Para isso, é de suma importância entender bem o que será como é feito o commit. 
<br/>
Siga o padrão: 

```
<type>:
<LINHA EM BRANCO>
<body>
<LINHA EM BRANCO>
```
<br/>
Os tipos podem ser:

**feat**: referente a uma nova feature implementada no projeto.

**fix**: referente a resolver algum bug no projeto.

**docs**: referente a mudança de documentação do projeto.

**style**: mudança referente ao estilo.

**refactor**: melhoria de código solicitada.

**ci**: referente arquivos de configuração do projeto.

**revert**: referente a reverter um commit.

**test**: referente a implementação ou correção de testes.

**build**: referente a implementações ou modificações que alteram a instalação de dependências do projeto

**chore**: quando não se aplica nenhum dos outros tipos do convencional commits.

Para ter um melhor entendimento após o type: colocar uma palavra na terceira pessoa do presente referente a implementação por exemplo:

**adiciona**

**modifica**

**melhora**


## Mensagem de exemplo de commit: 

```
feat: adiciona conversão em Kelvin para Celsius
```


