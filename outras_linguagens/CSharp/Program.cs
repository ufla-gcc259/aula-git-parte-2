using tempconv;
using System;

namespace testing{
    class Program{
        static void Main(string[] args){
            Console.WriteLine("Que frio! " + TemperatureFormatter.FormatCelsius(Constants.absZeroCelsius));
            Console.WriteLine("Fervendo! " + TemperatureFormatter.FormatFahrenheit(TemperatureConverter.CToF(Constants.boilingCelsius)));
        }
    }
}
