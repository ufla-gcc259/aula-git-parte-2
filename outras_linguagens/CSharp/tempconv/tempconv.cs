using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace tempconv{
    public class TemperatureConverter{
        public static double CToF(double cels){
            return (cels * 9) / 5 + 32;
        }

        public static double FToC(double fahr){
            return ((fahr - 32) * 5) / 9;
        }
    }
}
