package main

import (
	"fmt"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	fmt.Printf("Que frio! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Fervendo! %v\n", tempconv.CToF(tempconv.BoilingC))
}
