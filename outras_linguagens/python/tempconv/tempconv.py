class tempconv:

    def __init__(self):
        # AbsoluteZeroC representa a temperatura zero absoluto em Celsius
        self.ABSOLUTE_ZERO_C = -273.15

        # FreezingC representa a temperatura de congelamento da água em Celsius
        self.FREEZING_C = 0

        # BoilingC representa a temperatura de ebulição da água em Celsius
        self.BOILING_C = 100

    # String imprime uma temperatura 'n' em Celsius no formato n°C
    def formatCelsius(self, c):
        return ( str(c) + "°C") 

    # String imprime uma temperatura 'n' em Fahrenheit no formato n°F
    def formatFahrenheit(self, f):
        return ( str(f) + "°F") 

    # CToF converte uma temperatura em Celsius para Fahrenheit
    def CToF(self, c):
        return c*9/5 + 32

    # FToC converte uma temperatura em Fahrenheit para Celsius
    def FToC(self, f):
        return (f - 32) * 5 / 9
