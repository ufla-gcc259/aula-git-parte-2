# include <stdio.h>

typedef double Celsius;
typedef double Fahrenheit;

const Celsius AbsoluteZeroC = -273.15;
const Celsius FreezingC = 0;
const Celsius BoilingC = 100;

void printCelsius(Celsius c)
{
    printf("%.2f°C\n", c);
}

void printFahrenheit(Fahrenheit f)
{
    printf("%.2f°F\n", f);
}

Fahrenheit celsiusToFahrenheit(Celsius c)
{
    return c * 9 / 5 + 32;
}

Celsius fahrenheitToCelsius(Fahrenheit f)
{
    return (f - 32) * 5 / 9;
}

int main()
{
    Celsius c = -10.0;
    Fahrenheit f = 32.0;

    printf("Celsius: ");
    printCelsius(c);

    printf("Fahrenheit: ");
    printFahrenheit(f);

    printf("Celsius to Fahrenheit: ");
    printFahrenheit(celsiusToFahrenheit(c));

    printf("Fahrenheit to Celsius: ");
    printCelsius(fahrenheitToCelsius(f));

    return 0;
}
