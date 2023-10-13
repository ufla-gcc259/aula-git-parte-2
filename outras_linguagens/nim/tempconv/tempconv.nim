import strformat

type
  Celsius* = distinct float64
  Fahrenheit* = distinct float64

const
  AbsoluteZeroC* = -273.15.Celsius
  FreezingC* = 0.Celsius
  BoilingC* = 100.Celsius

func `$`*(c: Celsius): string = &"{c.float64}°C"
func `$`*(f: Fahrenheit): string = &"{f.float64}°F"

func toF*(c: Celsius): Fahrenheit = Fahrenheit(c.float64 * 9/5 + 32)
func toC*(f: Fahrenheit): Celsius = Celsius((f.float64 - 32) * 5/9)
