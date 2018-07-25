package main

import "fmt"

var numbers = []int{1, 1, 1, 1, 3, 3, 5, 5, 7, 7, 8, 8, 9, 9}

func main() {
	pairs := make([][]int, 0)
	end := len(numbers) - 1

	for index := 0; index < end; index++ {
		if numbers[index] == numbers[index+1] {
			pairs = append(pairs, numbers[index:index+2])
			index += 1
		}
	}

	fmt.Println(pairs)
}
