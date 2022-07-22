# _tempconv_

Biblioteca escrita em Node com TypeScript para conversões simples de temperatura.

## Como usar?

Baixe o [Node](https://nodejs.org/pt-br/download/) e instale o [Yarn](https://classic.yarnpkg.com/lang/en/docs/install/#windows-stable), podendo utilizar o npm tambem.

na pasta do projeto rode o camando `yarn`, para instalar as dependencias.

depois das dependencia terem sido instaladas, rode o seguinte comando:

`yarn start`

Pronto, agora é só usar:

```ts
import { AbsoluteZeroC, BoilingC } from "./constants";
import { formatCelsius, formatFahrenheit } from "./formatTemp";
import { CToF } from "./tempConv";

console.log(`Que frio! ${formatCelsius(AbsoluteZeroC)}\n`); // Que frio! -273.15°C
console.log(`Fervendo! ${formatFahrenheit(CToF(BoilingC))}\n`); // Fervendo! 212°F
```

## Licença

MIT
