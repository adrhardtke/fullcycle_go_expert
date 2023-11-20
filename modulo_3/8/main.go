package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(soma(2, 2))

	value, isValidSum := somaPositivos(-1, 0)

	if isValidSum {
		fmt.Printf("Operação valida, o valor é de %d\n", value)
	} else {
		fmt.Printf("Operação invalida, o valor é de %d\n", value)
	}

	value2, hasError := somaPositivosComErro(-1, 0)

	if hasError != nil {
		fmt.Printf("Operação valida, o valor é de %d\n", value2)
	} else {
		fmt.Printf("Operação invalida, o valor é de %d\n", value2)
	}
}

func soma(a int, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

// funcao retornando mais de 2 tipos possiveis

func somaPositivos(a, b int) (int, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}
	return a + b, true
}

func somaPositivosComErro(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("A soma é invalida")
	}
	return a + b, nil // nil é o mesmo que o null do javascript, é um valor em branco, nulo
}
