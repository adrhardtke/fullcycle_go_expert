package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	usuario := Cliente{
		Nome:  "Adriano",
		Idade: 27,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", usuario.Nome, usuario.Idade, usuario.Ativo)
}
