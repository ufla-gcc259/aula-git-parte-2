from tempconv_python.tempconv import tempConv 

def main():

    print("Fervendo! ", tempConv.CToF(tempConv().boilingC))
    print("Que frio! ", tempConv().absoluteZeroC)
    

if __name__ == "__main__":
	main()