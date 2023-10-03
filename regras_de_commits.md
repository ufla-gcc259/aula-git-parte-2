## Git Flow
Usamos o padrão baseado no Git Flow para versionamento de branches e código. Para isso temos 2 branches principais main e develop. A develop, como o próprio nome já diz, é usada durante o processo de desenvolvimento, ela irá possuir o código mais atual e será a branch a qual o desenvolvedor usará de base para criar as branches de novas features. As branches de feature devem seguir o padrão feature/<id>-<feature> onde o <id> é o código da tarefa e <feature> é uma breve descrição, em português, separada por "-" ex: feature/126-adiciona-pipeline-build. Para correções de bugs/problemas o padrão de nome das branches é hotfix/<id>-<fix> que é bem semelhante ao padrão anterior.

<IMG src="Hotfix branches.svg" alt="Fluxo de trabalho Git flow – Ramificações de hotfix"/>

## Versionamento
O versionamento das releases [Semantic Version](https://semver.org/) para facilitar a geração das versões da biblioteca. Em resumo, o padrão é formado por três números onde a versão é dada por `<major>.<minor>.<patch>` e cada feature pode incrementar um número no <minor> ou <major> dependendo da compatibilidade com as versões anteriores. Correções normalmente incrementam o <patch>.

## Padrão de Commit
Dado o padrão de versionamento definido anteriormente, um padrão de commit foi adotado para facilitar e automatizar a geração de versões dentro da biblioteca. O padrão de [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) segue o padrão de versionamento também adotado e facilita para que ferramentas de automatização consigam gerar versões automaticamente dado os commits feitos.

O padrão do conventional commits em sumo é:


```
<tipo>[escopo]: <descrição>

[opcional corpo texto]

[opcional rodapé(s)]
```

- O <tipo> é um dos tipos definidos nos padrões, pode ser `fix`, `feature`, `chore`
- [escopo] é um parâmetro opcional que define algum escopo mais amplo da alteração
- <descrição> é uma descrição de no máximo 100 caracteres, sempre com tudo minúsculo, com uma breve descrição do que foi alterado/corrigido. Algumas diretrizes são recomendadas de padronização, mas nao são obrigatórias, como sempre descrever o que foi alterado de forma imperativa (ex: corrige ao invés de corrigindo) e evitar mensagens muito genéricas como "correção de bug".
- [opcional corpo texto] é utilizado para dar mais informações e contextos quando somente a descrição não é suficiente
- [opcional rodapé(s)] são utilizados para dar mais informações estruturadas ao commit. Devem sepre seguir o padrão <rodapé>: <informação> ou <rodapé> <informação>. Alguns rodapés mais utilizados são `Reviewed-by: Usuário`,  `Refs #133` e `BREAKING CHANGE: <texto>`. Importante falar que o rodapé `BREAKING CHANGE:` causa um update da versão <major>.
