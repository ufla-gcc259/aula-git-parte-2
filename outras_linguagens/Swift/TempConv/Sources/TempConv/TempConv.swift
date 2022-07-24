//
//  TempConv.swift
//  
//
//  Created by Paulo Vieira on 22/07/22.
//

import Foundation

public struct TempConv {
    public static func celsiusToFahrenheit(_ temperature: Temperature) -> Temperature {
        return Temperature(
            value: temperature.value * 9 / 5 + 32,
            scale: .fahrenheit
        )
    }
    
    public static func fahrenheitToCelsius(_ temperature: Temperature) -> Temperature {
        return Temperature(
            value: (temperature.value - 32) * 5 / 9,
            scale: .celsius
        )
    }
}
