package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return somador(1, 2, 3, 4) * 2
	}()

	fmt.Println(total)
}

func somador(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		fmt.Println(numero)
		total += numero
	}
	return total
}
