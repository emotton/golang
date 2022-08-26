package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Eduardo": 20000.0,
			"Célia":   15000.0,
		},
		"H": {
			"Bheatriz": 10000.0,
		},
	}

	fmt.Println(funcsPorLetra)
}
