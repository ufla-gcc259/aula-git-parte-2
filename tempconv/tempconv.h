#include <string.h>

typedef float Celsius;
typedef float Fahrenheit;

const Celsius AbsoluteZeroC = -273.15f;
const Celsius FreezingC = 0.0f;
const Celsius BoilingC = 100.0f;

Fahrenheit CToF(Celsius c) {
   return c * 9/5 + 32;
}

Celsius FToC(Fahrenheit f) {
   return (f - 32) * 5 / 9;
}