package main

import "fmt"

func main() {
	const pi float64 = 3.14
	const (
		name   = "Psyon"
		age = 20
		city   = "teheran"
	)

	fmt.Printf("%f  \n", pi)

	fmt.Printf("name: %s age: %d city: %s pi: %f \n", name, age, city, pi)
}