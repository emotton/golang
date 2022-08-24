package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(32000))

	// sem sinal (uint8 uint16 uint32 uint64)
	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	// com sinal (int8 int16 int32 int64)
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de um int64 é", i1)

	var i2 rune = 'A' // tabela Unicode (int32)
	fmt.Println(i2)
	fmt.Println(reflect.TypeOf(i2))

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	// string
	s1 := "Olá meu nome é Edu"
	fmt.Println(s1)
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
}
