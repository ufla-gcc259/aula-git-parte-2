//
//  SampleCLIApp.swift
//  
//
//  Created by Paulo Vieira on 22/07/22.
//

import Foundation
import TempConv

func main() -> Void {
    print("Que frio! 🥶", Temperature.Constants.absoluteZeroInCelsius)
    print("Fervendo! 🥵🔥", TempConv.celsiusToFahrenheit(Temperature.Constants.waterBoilingPointInCelsius))
}

main()

