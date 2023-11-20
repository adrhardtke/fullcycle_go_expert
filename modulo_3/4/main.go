package main

import "fmt"

type ID int

var t1 ID = 1
var t2 int = 2

// %v = tras o valor
// %T = tras o tipo

func main() {
	fmt.Printf("O tipo de t1 é %T\n", t1)
	fmt.Printf("O tipo de t2 é %T\n", t2)
}
