package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Substração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10 (2)
	fmt.Println("OU =>", a|b)  // 11 | 10 = 11 (3)
	fmt.Println("XOR =>", a^b) // 11 ^ 10 = 01 (1)

	c := 3.0
	d := 2.0

	// Outras operações usando math
	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
