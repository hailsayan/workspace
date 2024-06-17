package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	addItem(&numbers)

	fmt.Printf("%v \n", numbers)

}

func changeNumbers(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 3
	}
}

func addItem(numbers *[]int){
	*numbers = append(*numbers, 6)
}