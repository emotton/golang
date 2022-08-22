package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Teste")
	uuid := uuid.New()
	fmt.Println(uuid)
}
