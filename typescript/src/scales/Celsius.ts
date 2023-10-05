export class Celsius extends Number {
  toString(radix?: number | undefined): string {
    return `${super.toString(radix)}ºC`;
  }
}
