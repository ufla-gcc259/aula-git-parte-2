import tempconv.TempConv;

public class Main {
  public static void main(String[] args) {
    TempConv tc = new TempConv();
    System.out.printf("Que frio! %s\n", tc.formatCelsius(tc.getAbsoluteZeroC()));
    System.out.printf("Fervendo! %s\n", tc.formatFahrenheit(tc.CToF(tc.getBoilingC())));
  }
}
