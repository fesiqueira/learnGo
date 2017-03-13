package main

import "fmt"

func findGreatest(numbers ...float64) (greatest float64) {
	for _, n := range numbers {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}

func main() {
	fmt.Println(findGreatest(1, 2, 3, 4, 5, 11, 6, 7.1))
}
