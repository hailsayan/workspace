package main

func main() {
	i := 0
	//lst := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	lst1 := []int{12, 15, 14, 13, 16, 17, 18, 19, 20, 21}
	for {
		println("")
		break
	}

	for i < 10 {
		println("num: ", i)
		i++
	}

	for j := 0; j < 10; j++ {
		println("print: ", j)
	}

	for index, item := range lst1 {
		println("index&item ", index, item)
	}
}
