import { AbsoluteZeroC, BoilingC } from "./constants";
import { formatCelsius, formatFahrenheit } from "./formatTemp";
import { CToF } from "./tempConv";

console.log(`Que frio! ${formatCelsius(AbsoluteZeroC)}\n`);
console.log(`Fervendo! ${formatFahrenheit(CToF(BoilingC))}\n`);
