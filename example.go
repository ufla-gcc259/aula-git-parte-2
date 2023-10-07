package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		uso()
	} else{
		if value, err := strconv.ParseFloat(argsWithProg[1], 64); err != nil { 
			uso()
		} else {
			tempCelsius := tempconv.Celsius(value)
			tempFahrenheit := tempconv.Fahrenheit(value)
			
			if (tempFahrenheit < tempconv.CToF(tempconv.AbsoluteZeroC)){
				fmt.Printf("O valor entrado (%f) está abaixo do Zero Absoluto na escala Fahrenheit (%v) e na escala Celsius (%v)\n", value, tempconv.CToF(tempconv.AbsoluteZeroC), tempconv.AbsoluteZeroC)
			} else if (tempCelsius < tempconv.AbsoluteZeroC){
				fmt.Printf("O valor entrado (%f) está abaixo do Zero Absoluto na escala Celsius (%v) \n", value, tempconv.AbsoluteZeroC)
				fmt.Printf("%v = %v\n", tempFahrenheit, tempconv.FToC(tempFahrenheit))
			} else {
				fmt.Printf("%v = %v\n", tempCelsius, tempconv.CToF(tempCelsius))
				fmt.Printf("%v = %v\n", tempFahrenheit, tempconv.FToC(tempFahrenheit))
			}
		}
	}
}

func uso(){
	fmt.Println("Para executar corretamente o exemplo, passe apenas um valor numérico como parametro:\ngo run ./example.go <valor numérico>")
}