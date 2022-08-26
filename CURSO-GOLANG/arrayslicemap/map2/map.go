package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Eduardo":  20000.0,
		"Célia":    15000.0,
		"Henrique": 10000.0,
	}

	funcsESalarios["Bheatriz"] = 5000.0

	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "Inexistente") // Não gera problema
}
