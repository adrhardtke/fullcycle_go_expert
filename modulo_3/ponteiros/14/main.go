package main

func main() {
	// Memoria -> Endereco -> Valor
	a := 10
	println(a)  // 10
	println(&a) // endereco de onde está o "a" armazenado na memoria

	var ponteiro *int = &a
	println(ponteiro) // mesmo enderecamento de "a"
	*ponteiro = 20
	println(a) // 20
	b := &a
	println(b)  // mesmo endereco de "a"
	println(*b) // 20
}
