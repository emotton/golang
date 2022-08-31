package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	resultado := media(7.7, 8.8, 3.4, 1.1)
	fmt.Println(resultado)
}
