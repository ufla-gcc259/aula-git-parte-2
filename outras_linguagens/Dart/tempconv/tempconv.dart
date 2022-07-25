// Tempconv realiza conversões de Celsius e Fahrenheit
class Tempconv {
  // absoluteZeroC representa a temperatura zero absoluto em Celsius
  static const double absoluteZeroC = -273.15;
  // freezingC representa a temperatura de congelamento da água em Celsius
  static const double freezingC = 0;
  // boilingC representa a temperatura de ebulição da água em Celsius
  static const double boilingC = 100;

  double CtoF(double celsius) {
    return celsius * 9 / 5 + 32;
  }

  double FtoC(double fahrenheit) {
    return (fahrenheit - 32) * 5 / 9;
  }
}
