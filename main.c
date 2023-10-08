#include "tempconv/tempconv.c"

void main()
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
}