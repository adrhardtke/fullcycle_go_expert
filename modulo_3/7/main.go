package main

import "fmt"

func main() {
	salarios := map[string]int{
		"Wesley": 1000,
		"Joao":   2000,
		"Maria":  3000,
	}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley") // remove a chave Wesley
	salarios["Wes"] = 5000
	fmt.Println(salarios["Wes"])

	// outra forma de criar map
	sal := make(map[string]int)
	sal["Wesley"] = 2000

	// outra forma de criar map
	sal1 := map[string]int{}
	sal1["Wesley"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é de R$ %d,00\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é de R$ %d,00\n", salario)
	}
}
