//
//  Temperature.swift
//  
//
//  Created by Paulo Vieira on 22/07/22.
//

import Foundation

public struct Temperature {
    public enum Scale {
        case celsius
        case fahrenheit
        
        public var unit: String {
            switch self {
            case .celsius: return "ºC"
            case .fahrenheit: return "ºF"
            }
        }
    }
    
    public struct Constants {
        public static let absoluteZeroInCelsius = Temperature(value: -273.15, scale: .celsius)
        public static let waterFreezingPointInCelsius = Temperature(value: 0.0, scale: .celsius)
        public static let waterBoilingPointInCelsius = Temperature(value: 100.0, scale: .celsius)
    }
    
    let value: Double
    let scale: Scale
    
    public init(value: Double, scale: Scale) {
        self.value = value
        self.scale = scale
    }
}

extension Temperature: CustomStringConvertible {
    public var description: String {
        String(format: "%.2f\(scale.unit)", value)
    }
}
