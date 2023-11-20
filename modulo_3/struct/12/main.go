package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	// end Endereco // também pode ser assim
}

func main() {
	usuario := Cliente{
		Nome:  "Adriano",
		Idade: 27,
		Ativo: true,
	}
	usuario.Cidade = "Pelotas"
	// ou
	// usuario.Endereco.Cidade = "Pelotas"
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", usuario.Nome, usuario.Idade, usuario.Ativo)
}
