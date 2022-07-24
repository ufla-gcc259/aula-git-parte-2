// tempconv realiza conversões de Celsius e Fahrenheit
package tempconv;

public class TempConv {
  // AbsoluteZeroC representa a temperatura zero absoluto em Celsius
  private double AbsoluteZeroC;

  // FreezingC representa a temperatura de congelamento da água em Celsius
  private double FreezingC;

  // BoilingC representa a temperatura de ebulição da água em Celsius
  private double BoilingC;

  public TempConv() {
    this.AbsoluteZeroC = -273.15;
    this.FreezingC = 0;
    this.BoilingC = 100;
  }

  public double getAbsoluteZeroC() {
    return AbsoluteZeroC;
  }

  public double getFreezingC() {
    return FreezingC;
  }

  public double getBoilingC() {
    return BoilingC;
  }

  // String imprime uma temperatura 'n' em Celsius no formato n°C
  public String formatCelsius(double c) {
    return String.format("%.2f°C", c);
  }

  // String imprime uma temperatura 'n' em Fahrenheit no formato n°F
  public String formatFahrenheit(double f) {
    return String.format("%.2f°F", f);
  }

  // CToF converte uma temperatura em Celsius para Fahrenheit
  public double CToF(double c) {
    return (c * 9 / 5 + 32);
  }

  // FToC converte uma temperatura em Fahrenheit para Celsius
  public double FToC(double f) {
    return ((f - 32) * 5 / 9);
  }

}
