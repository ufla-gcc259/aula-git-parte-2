import tempconv.TempConv.Celsius;
import tempconv.TempConv.Fahrenheit;

public class Main {
    public static void main(String[] args) {
        System.out.printf("Que frio! %s\n", TempConv.AbsoluteZeroC);
        System.out.printf("Fervendo! %s\n", TempConv.CToF(TempConv.BoilingC));
    }
}