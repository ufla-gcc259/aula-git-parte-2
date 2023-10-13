package tempconv

import "testing"
import "math"

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}


func TestTempConv(t *testing.T) {
    testes := []struct {
        tempF   Fahrenheit
    }{
    	{tempF:   23.120},
    	{tempF:   86.920},
    	{tempF: - 11.310},
    	{tempF:  103.670},
    	{tempF:   56.210},
    	{tempF:  110.300},
    }
	
    for _, teste := range testes {
    
		testeTempC := FToC(teste.tempF)
		testeTempF := CToF(testeTempC)

    	if toFixed( float64(testeTempF), 2) != float64(teste.tempF) {
        	t.Fatalf("Erro --> Valor esperado: %g - Valor retornado: %g", teste.tempF, testeTempF)
    	}
    }
}