*tempconv*
=====
Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan. 

Como usar?
----
Faça o download da biblioteca:

`go get https://github.com/ufla-gcc259/aula-git-parte-2@v1.0.1`

ou em Java
'java get https://github.com/luisteixeira13/aula-git-parte-2.git'


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


ou em Java
import tempconv.tempconv;

public class main {
    public static void main(String[] args) {
        tempconv.Celsius absoluteZeroC = new tempconv.Celsius(-273.15);
        tempconv.Celsius boilingC = new tempconv.Celsius(100);

        System.out.printf("Que frio! %s%n", absoluteZeroC);
        System.out.printf("Fervendo! %s%n", tempconv.celsiusToFahrenheit(boilingC));
    }
}


Outras linguagens?
Java
Versões da biblioteca *tempconv* para outras linguagens:

> Java

Licença
-----

> *Todo*


Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
