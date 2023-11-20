package main

import (
	"fmt"
)

func main() {
	fmt.Println(somador(1, 2, 3, 4))
}

// funcao recebendo um numero indefinido de parametros
func somador(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		fmt.Println(numero)
		total += numero
	}
	return total
}
