import tempconv.tempconv;

public class main {
    public static void main(String[] args) {
        tempconv.Celsius absoluteZeroC = new tempconv.Celsius(-273.15);
        tempconv.Celsius boilingC = new tempconv.Celsius(100);

        System.out.printf("Que frio! %s%n", absoluteZeroC);
        System.out.printf("Fervendo! %s%n", tempconv.celsiusToFahrenheit(boilingC));
    }
}
