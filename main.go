package main

import (
	"fmt"

	"github.com/paulojunior-ufla/aula-git-parte-2/tempconv"
)

func main() {
	fmt.Printf("Que frio! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Fervendo! %v\n", tempconv.CToF(tempconv.BoilingC))
}
