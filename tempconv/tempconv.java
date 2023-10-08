package tempconv;
public class tempconv {
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
            return value + "°C";
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
            return value + "°F";
        }
    }

    public static final Celsius AbsoluteZeroC = new Celsius(-273.15);
    public static final Celsius FreezingC = new Celsius(0);
    public static final Celsius BoilingC = new Celsius(100);

    public static Fahrenheit celsiusToFahrenheit(Celsius celsius) {
        double value = celsius.getValue() * 9/5 + 32;
        return new Fahrenheit(value);
    }

    public static Celsius fahrenheitToCelsius(Fahrenheit fahrenheit) {
        double value = (fahrenheit.getValue() - 32) * 5/9;
        return new Celsius(value);
    }

  }
