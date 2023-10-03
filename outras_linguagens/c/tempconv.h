#define AbsoluteZeroC -273.15
#define FreezingC 0.0
#define BoilingC 100.0


char* formatCelsius (float celsius){
    static char buffer[100];
    sprintf(buffer, "%.2f °C", celsius);
    return buffer;
}

char* formatFahrenheit (float fahrenheit){
    static char buffer[100];
    sprintf(buffer, "%.2f °F", fahrenheit);
    return buffer;
}

float CToF (float celsius){
    return celsius * 9.0 / 5.0 + 32.0;
}

float FToC (float fahrenheit){
    return (fahrenheit - 32.0) * 5.0 / 9.0;
}