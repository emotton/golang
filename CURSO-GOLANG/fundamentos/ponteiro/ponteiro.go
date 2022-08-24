package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // pegando o endereço da variável

	*p++ // desreferenciando
	i++

	fmt.Println(p, *p, i)

	// p++
	// Go não tem aritmética de ponteiros

}
