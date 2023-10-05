public class TempConv {
    public static class Celsius {
        private double value;

        public Celsius(double value) {
            this.value = value;
        }

        public double getValue() {
            return value;
        }

        @Override
        public String toString() {
            return String.format("%g°C", value);
        }
    }

    public static class Fahrenheit {
        private double value;

        public Fahrenheit(double value) {
            this.value = value;
        }

        public double getValue() {
            return value;
        }

        @Override
        public String toString() {
            return String.format("%g°F", value);
        }
    }

    public static final Celsius AbsoluteZeroC = new Celsius(-273.15);
    public static final Celsius FreezingC = new Celsius(0);
    public static final Celsius BoilingC = new Celsius(100);

    public static Fahrenheit CToF(Celsius celsius) {
        double value = celsius.getValue() * 9 / 5 + 32;
        return new Fahrenheit(value);
    }

    public static Celsius FToC(Fahrenheit fahrenheit) {
        double value = (fahrenheit.getValue() - 32) * 5 / 9;
        return new Celsius(value);
    }

    public static void main(String[] args) {
        Celsius celsiusValue = new Celsius(25.0);
        Fahrenheit fahrenheitValue = CToF(celsiusValue);

        System.out.println(celsiusValue);  // Imprime 25°C
        System.out.println(fahrenheitValue); // Imprime 77°F
    }
}
