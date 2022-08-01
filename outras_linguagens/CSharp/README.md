# Biblioteca para Conversões Simples de Temperatura em C#

## Pré-requisitos
- .NET
- Visual Studio IDE


### **Para fazer uso da biblioteca, basta incluir o arquivo .nupkg no seu projeto C#!** 
## E como fazer isso?

1. Crie um novo projeto em .NET (ou utilize um já feito) pela IDE Visual Studio ou pelo terminal utilizando o seguinte comando:

```
dotnet new console -o NomeDoProjeto
```

2. No Visual Studio, no projeto, selecione a opção *Manage NuGet Packages...*

![imagem1](https://github.com/yanpisani/aula-git-parte-2/blob/main/outras_linguagens/CSharp/src/1.png)

3. No *NuGet Package Manager*, selecione a engrenagem.

![imagem2](https://github.com/yanpisani/aula-git-parte-2/blob/main/outras_linguagens/CSharp/src/2.png)

4. Após aberta a janela, clique no botão **+**.

![imagem3](https://github.com/yanpisani/aula-git-parte-2/blob/main/outras_linguagens/CSharp/src/3.png)

5. Feito isso, dê um nome para o local onde está localizada o arquivo *.nupkg* (biblioteca), e o seu local.

![imagem4](https://github.com/yanpisani/aula-git-parte-2/blob/main/outras_linguagens/CSharp/src/4.png)

6. Agora aperte no botão *Install* para incluir a biblioteca nas dependências do seu projeto.

![imagem5](https://github.com/yanpisani/aula-git-parte-2/blob/main/outras_linguagens/CSharp/src/5.png)

## E pronto! Agora para fazer uso da biblioteca apenas inclua:

```cs
using tempconv
```

- Programa de exemplo:
```cs
using tempconv;
using System;

namespace testing{
    class Program{
        static void Main(string[] args){
            Console.WriteLine("Que frio! " + TemperatureFormatter.FormatCelsius(Constants.absZeroCelsius));
            Console.WriteLine("Fervendo! " + TemperatureFormatter.FormatFahrenheit(TemperatureConverter.CToF(Constants.boilingCelsius)));
        }
    }
}
```
- Resultado:

![imagem6](https://github.com/yanpisani/aula-git-parte-2/blob/main/outras_linguagens/CSharp/src/6.png)
