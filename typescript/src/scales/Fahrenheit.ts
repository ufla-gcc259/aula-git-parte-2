export class Fahrenheit extends Number {
  toString(radix?: number | undefined): string {
    return `${super.toString(radix)}ºC`;
  }
}
