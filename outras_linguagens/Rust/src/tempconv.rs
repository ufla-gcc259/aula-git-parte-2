#[derive(Debug)]
pub enum Temperatures {
    Celsius(f64),
    Fahrenheit(f64),
    Kelvin(f64),
}

impl Temperatures {
    pub fn boiling(&self) -> Self {
        match self {
            Self::Celsius(_) => Self::Celsius(100.00),
            Self::Fahrenheit(_) => Self::Fahrenheit(212.00),
            Self::Kelvin(_) => Self::Kelvin(373.15),
        }
    }

    pub fn absolute_zero(&self) -> Self {
        match self {
            Self::Celsius(_) => Self::Celsius(-273.15),
            Self::Fahrenheit(_) => Self::Fahrenheit(-459.67),
            Self::Kelvin(_) => Self::Kelvin(0.0),
        }
    }

    pub fn freezing(&self) -> Self {
        match self {
            Self::Celsius(_) => Self::Celsius(0.0),
            Self::Fahrenheit(_) => Self::Fahrenheit(32.0),
            Self::Kelvin(_) => Self::Kelvin(273.15),
        }
    }

    pub fn to_celsius(&self) -> Self {
        match self {
            Self::Celsius(temperature) => Self::Celsius(*temperature),
            Self::Fahrenheit(temperature) => Self::Celsius(5.0 / 9.0 * temperature - 32.0),
            Self::Kelvin(temperature) => Self::Celsius(temperature - 273.15),
        }
    }

    pub fn to_fahrenheit(&self) -> Self {
        match self {
            Self::Celsius(temperature) => Self::Fahrenheit(9.0 / 5.0 * temperature + 32.0),
            Self::Fahrenheit(temperature) => Self::Fahrenheit(*temperature),
            Self::Kelvin(temperature) => Self::Fahrenheit(9.0 / 5.0 * temperature - 459.67),
        }
    }

    pub fn to_kelvin(&self) -> Self {
        match self {
            Self::Celsius(temperature) => Self::Kelvin(temperature + 273.15),
            Self::Fahrenheit(temperature) => Self::Kelvin(5.0 / 9.0 * temperature + 459.67),
            Self::Kelvin(temperature) => Self::Kelvin(*temperature),
        }
    }
}
