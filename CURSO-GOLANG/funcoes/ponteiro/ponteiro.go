package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// * usando para desreferenciar, ter acesso ao valor ao qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	inc2(&n) // por referência
	fmt.Println(n)
}
