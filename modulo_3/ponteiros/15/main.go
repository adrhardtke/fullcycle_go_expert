package main

func soma(a, b int) int {
	a = 50
	return a + b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	// está passando uma copia dos valores e nao o endereco de fato
	soma(minhaVar1, minhaVar2)
	println(minhaVar1) // 10, mesmo alterando em soma para 50

	somaPonteiro(&minhaVar1, &minhaVar2)
	println(minhaVar1) // 50
}

func somaPonteiro(a, b *int) int {
	*a = 50
	return *a + *b
}
