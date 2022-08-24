package main

import "fmt"

func main() {
	x := 3.1495
	// fmt.Println("O valor de x é " + x)
	// não é permitido concactenar tipos diferentes
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
