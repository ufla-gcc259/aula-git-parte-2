// AbsoluteZeroC representa a temperatura zero absoluto em Celsius
const AbsoluteZeroC = -273.15

// FreezingC representa a temperatura de congelamento da água em Celsius
const FreezingC = 0

// BoilingC representa a temperatura de ebulição da água em Celsius
const BoilingC = 100

// imprime uma temperatura 'n' em Celsius no formato n°C
function imprimeCelsius(n) {
    return n + "°C";
}

//Imprime uma temperatura 'n' em Fahrenheit no formato n°F
function imprimeFahrenheit(n) {
    return n + "°F";
}

// CToF converte uma temperatura em Celsius para Fahrenheit
function CToF (c) {
    return (c*9/5 + 32);
}

// FToC converte uma temperatura em Fahrenheit para Celsius
function FToC (f) {
    return ((f - 32) * 5 / 9);
}


console.log("Que frio! " + imprimeCelsius(AbsoluteZeroC) + "\n");
console.log("Fervendo! " + imprimeFahrenheit(CToF(BoilingC)) + "\n");
