class tempConv:
    
    def __init__(self):
        # absoluteZeroC representa a temperatura zero absoluto em Celsius
        self.absoluteZeroC = -273.15
        # freezingC representa a temperatura de congelamento da água em Celsius
        self.freezingC  = 0
        # boilingC representa a temperatura de ebulição da água em Celsius
        self.boilingC = 100

    
    # imprime uma temperatura 'n' em Celsius no formato n°C
    def printCelsius (c):
        return (str(c) + '°C')

    # imprime uma temperatura 'n' em Fahrenheit no formato n°F
    def printFahrenheit (f):
        return (str(f) + '°F')
    
    # CToF converte uma temperatura em Celsius para Fahrenheit
    def CToF (c):
        return c*9/5 + 32

    # FToC converte uma temperatura em Fahrenheit para Celsius
    def FToC (f):
        return (f - 32) * 5 / 9



