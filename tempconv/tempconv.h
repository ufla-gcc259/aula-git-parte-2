typedef float Celsius;
typedef float Fahrenheit;

const Celsius AbsoluteZeroC = -273.15;
const Celsius FreezingC = 0;
const Celsius BoilingC = 100;

Fahrenheit CToF(Celsius c){
    return Fahrenheit(c*9/5 + 32);
}

Celsius FToC(Fahrenheit f){
    return Celsius((f - 32) * 5/9);
}