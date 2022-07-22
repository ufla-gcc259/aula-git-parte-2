// FToC converte uma temperatura em Fahrenheit para Celsius
export function FToC(fahrenheit: number) {
  return ((fahrenheit - 32) * 5) / 9;
}

// CToF converte uma temperatura em Celsius para Fahrenheit
export function CToF(celsius: number) {
  return (celsius * 9) / 5 + 32;
}
