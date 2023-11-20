package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// automaticamente se uma struct possui esse metodo Desativar, ela já implementa a interface Pessoa
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

// adiciona o metodo a struct Cliente
func (client Cliente) Desativar() {
	client.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", client.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	usuario := Cliente{
		Nome:  "Adriano",
		Idade: 27,
		Ativo: true,
	}
	usuario.Cidade = "Pelotas"
	usuario.Desativar()

	// pode passar o usuario, pois o mesmo é uma Pessoa, já que implementa o metodo Desativar, que está na interface Pessoa.
	Desativacao(usuario)
}
