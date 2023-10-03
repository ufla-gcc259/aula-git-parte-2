*tempconv*
=====
Transcição da biblioteca tempconv escrita em Go para C.

Como usar?
----
Faça o download da biblioteca:

`git clone https://github.com/ufla-gcc259/aula-git-parte-2@v1.0.1`

Adicione o arquivo de cabeçalho *tempconv.h* ao seu projeto e inclua-o:

```c
#include <stdio.h>
#include "tempconv.h"

int main(){

    printf("Que frio!: %.2f\n", AbsoluteZeroC); // Que frio!: -273.15
    printf("Fervendo!: %s\n", formatFahrenheit(CToF((float)BoilingC))); // Fervendo!: 212.00 °F
}
```
Licensa
----
Consultar a licensa presente no repositorio.

Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
