ZERO_ABSOLUTO_C = -273.15
CONGELAMENTO_C = 0
EBULICAO_C = 100

def imprime_celsius(n):
    print(n, '°C')
    
def imprime_fahrenheit(n):
    print(n, '°F')
    
def celsius_para_fahrenheit(c):
    return (c * 9 / 5 + 32)

def fahrenheit_para_celsius(f):
    return (f - 32) * 5 / 9
    

# imprime_celsius(celsius_para_fahrenheit((CONGELAMENTO_C)))
# imprime_celsius(celsius_para_fahrenheit((EBULICAO_C)))
# imprime_fahrenheit(fahrenheit_para_celsius((32)))
# imprime_fahrenheit(fahrenheit_para_celsius((212)))
