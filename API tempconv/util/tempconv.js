// AbsoluteZeroC representa a temperatura zero absoluto em Celsius
const AbsoluteZeroC = -273.15;
// FreezingC representa a temperatura de congelamento da água em Celsius
const FreezingC = 0;
// BoilingC representa a temperatura de ebulição da água em Celsius
const BoilingC = 100;

class tempconv {
  // String imprime uma temperatura 'n' em Celsius no formato n°C
  static celsius(temp) {
    return `${temp}°C`;
  }
  // String imprime uma temperatura 'n' em Fahrenheit no formato n°F
  static fahrenheit(temp) {
    return `${temp}°F`;
  }
  // CToF converte uma temperatura em Celsius para Fahrenheit
  static CToF(temp){
    return this.fahrenheit(temp*9/5 + 32);
  }
  // FToC converte uma temperatura em Fahrenheit para Celsius
  static FToC(temp){
    return this.celsius(parseFloat((temp - 32) * 5 / 9).toFixed(2));
  }
}

module.exports = tempconv;