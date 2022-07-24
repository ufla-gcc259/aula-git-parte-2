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

>imagem 1

3. No NuGet Package Manager, selecione a engrenagem.

>imagem 2

4. Após aberta a janela, clique no botão +.

>imagem 3

5. Feito isso, dê um nome para o local onde está localizada o arquivo .nupkg (biblioteca), e o seu local.

>imagem 4

6. Agora aperte no botão *Install* para incluir a biblioteca nas dependências do seu projeto.

>imagem 5

E pronto! Agora para fazer uso da biblioteca apenas inclua:

```cs
using tempconv
```

Programa de exemplo:
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


