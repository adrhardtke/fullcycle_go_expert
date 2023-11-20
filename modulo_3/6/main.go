package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=9 cap=9 [10 20 30 50 60 70 80 90 100]

	/*
		s[:0] --> a partir do INICIO do meu array (:), coloque "0" coisas, ou seja, retorna o array vazio -> []
		s[2:] --> ignorou os dois primeiros itens, e pegou o restante (30 50 60 70 80 90 100)
	*/
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // len=0 cap=9 []

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // len=4 cap=9 [10 20 30 50]

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // len=7 cap=7 [30 50 60 70 80 90 100]

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// len=10 cap=18 [10 20 30 50 60 70 80 90 100 110], no go, quando der um append e não tiver mais capacidade, ele duplica o tamanho do array, como tinha capacidade = 9, agora tem capacidade = 18

}
