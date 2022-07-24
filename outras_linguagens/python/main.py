from tempconv.tempconv import tempconv as tc

def main():
	print("Que frio! ", tc().formatCelsius(tc().ABSOLUTE_ZERO_C))
	print("Fervendo! ", tc().formatCelsius(tc().CToF(tc().BOILING_C)))


if __name__ == "__main__":
	main()