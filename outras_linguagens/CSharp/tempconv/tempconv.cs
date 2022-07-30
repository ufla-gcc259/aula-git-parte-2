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

        public static double CToK(double cels){
            return cels - Constants.absZeroCelsius;
        }

        public static double FToK(double fahr){
            return CToK(FToC(fahr));
        }

        public static double KToC(double kelv){
            return kelv + Constants.absZeroCelsius;
        }

        public static double KToF(double kelv){
            return CToF(KToC(kelv));
        }
    }
}
