package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 1
	myArray[2] = 1

	fmt.Println(len(myArray))
	fmt.Println(myArray[len(myArray)-1])

	for i, v := range myArray {
		fmt.Printf("O valor do indice %v é %d\n", i, v)
	}

}
