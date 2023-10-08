# _tempconv_

Biblioteca escrita em C para conversões de temperatura. Desenvolvido como atividade para a disciplina de Software Livre.

## Como usar?

Faça o download da biblioteca:

`go get https://github.com/ufla-gcc259/aula-git-parte-2@v1.0.1`

Pronto, agora é só usar:

```c
#include "tempconv/tempconv.c"

void main()
{
    Celsius c = -10.0;
    Fahrenheit f = 32.0;

    printf("Celsius: ");
    printCelsius(c);

    printf("Fahrenheit: ");
    printFahrenheit(f);

    printf("Celsius to Fahrenheit: ");
    printFahrenheit(celsiusToFahrenheit(c));

    printf("Fahrenheit to Celsius: ");
    printCelsius(fahrenheitToCelsius(f));
}
```

## Outras linguagens?

Versões da biblioteca _tempconv_ para outras linguagens:
