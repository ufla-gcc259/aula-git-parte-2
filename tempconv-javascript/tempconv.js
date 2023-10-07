const ABSOLUTE_ZERO_C = -273.15; 
const FREEZING_C = 0; 
const BOILING_C = 100; 

const tempconv = {
    ABSOLUTE_ZERO_C,
    FREEZING_C,
    BOILING_C,

    // String imprime uma temperatura 'n' em Celsius no formato n°C
    stringCelsius(value) {
        return `${value}°C`;
    },

    // String imprime uma temperatura 'n' em Fahrenheit no formato n°F
    stringFahrenheit(value) {
        return `${value}°F`;
    },
    
    // CToF converte uma temperatura em Celsius para Fahrenheit
    CToF(value) {
        return (value*(9/5) + 32);
    },

    // FToC converte uma temperatura em Fahrenheit para Celsius
    FToC(value) {
        return ((value-32) * (5/9));
    }


}

module.exports = tempconv;
