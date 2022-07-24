class tempconv:
    def __init__(self):
        # AbsoluteZeroC representa a temperatura zero absoluto em Celsius
        self.AbsoluteZeroC Celsius = -273.15

        # FreezingC representa a temperatura de congelamento da água em Celsius
        self.FreezingC Celsius = 0

        # BoilingC representa a temperatura de ebulição da água em Celsius
        self.BoilingC Celsius = 100


    # CToS imprime uma temperatura 'n' em Celsius no formato n°C
    def CToS(celsius):
        return str(celsius) + "°C"
    # FToS imprime uma temperatura 'n' em Fahrenheit no formato n°F
    def FToS(fahrenheit):
        return str(fahrenheit) + "°F"
    # KToS imprime uma temperatura 'n' em Kelvin no formato nK
    def KToS(kelvin):
        return str(kelvin) + "K"
    
    # CToF converte uma temperatura em Celsius para Fahrenheit
    def CToF(celsius):
        return celsius*9/5 + 32
    # CToK converte uma temperatura em Celsius para Kelvin
    def CToK(celsius):
        return celsius+273.15
    
    # FToC converte uma temperatura em Fahrenheit para Celsius
    def FToC(fahrenheit):
        return (fahrenheit-32) * 5/9
    # FToK converte uma temperatura em Fahrenheit para Kelvin
    def FtoK(fahrenheit):
        return (fahrenheit-32) * 5/9 + 273.15
    
    # KToC converte uma temperatura em Kelvin para Celsius
    def KToC(kelvin):
        return kelvin-273.15
    # KToF converte uma temperatura em Kelvin para Fahrenheit
    def KToF(kelvin):
        return (kelvin-273.15)*9/5 + 32