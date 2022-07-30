using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace tempconv{
    public class TemperatureFormatter{
        public static String FormatCelsius(double cels){
            return cels.ToString() + " ºC";
        }

        public static String FormatFahrenheit(double fahr){
            return fahr.ToString() + " ºF";
        }

        public static String FormatKelvin(double kelv){
            return kelv.ToString() + " K";
        }
    }
}
