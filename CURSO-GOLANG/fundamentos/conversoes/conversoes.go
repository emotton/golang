package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado
	fmt.Println("Teste " + string(123)) // caracter 123 unicode

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 100)

	b, _ := strconv.ParseBool("True") // Apenas true ou "1"
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}
}
