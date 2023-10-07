class Celsius:
    def __init__(self, value):
        self.value = value

    def __str__(self):
        return f"{self.value}°C"

class Fahrenheit:
    def __init__(self, value):
        self.value = value

    def __str__(self):
        return f"{self.value}°F"

def c_to_f(celsius):
    return Fahrenheit(celsius.value * 9/5 + 32)

def f_to_c(fahrenheit):
    return Celsius((fahrenheit.value - 32) * 5/9)

absolute_zero_c = Celsius(-273.15)
freezing_c = Celsius(0)
boiling_c = Celsius(100)

if __name__ == "__main__":
    print(f"Que frio! {absolute_zero_c}")
    print(f"Fervendo! {c_to_f(boiling_c)}")
