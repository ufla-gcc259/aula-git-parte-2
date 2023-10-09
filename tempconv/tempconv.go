package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	// AbsoluteZeroC representa a temperatura zero absoluto em Celsius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC representa a temperatura de congelamento da água em Celsius
	FreezingC Celsius = 0

	// BoilingC representa a temperatura de ebulição da água em Celsius
	BoilingC Celsius = 100
)

// String imprime uma temperatura 'n' em Celsius no formato n°C
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// String imprime uma temperatura 'n' em Fahrenheit no formato n°F
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// CToF converte uma temperatura em Celsius para Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converte uma temperatura em Fahrenheit para Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5/9) }

// FToK converte uma temperatura em Fahrenheit para Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f - 32) * 5/9 + 273.15) }

// CToK converte uma temperatura em Celsius para Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToF converte uma temperatura em Kelvin para Fahrenheit
func CToF(k Kelvin) Fahrenheit { return Fahrenheit((k - 273.15) * 9/5 + 32) }

// KToC converte uma temperatura em Kelvin para Celsius
func KToC(k Kelvin) Celsius { return Celsius(c - 273.15) }