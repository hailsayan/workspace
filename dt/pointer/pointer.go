package main

import "fmt"

func main() {
	i, j := 3, 7

	var iA *int = &i
	var jA *int = &j

	fmt.Println("l'adresse de i: ", &i)
	fmt.Println("l'adresse de i: ", iA)

	fmt.Println("l'adresse de j: ", &j)
	fmt.Println("l'adresse de j: ", jA)
}