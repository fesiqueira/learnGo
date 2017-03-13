package main

import "fmt"

func fibo(until int) (numbers []int) {
	prev := 1
	var number, prev2 int
	for i := 0; i < until; i++ {
		number = prev + prev2
		prev2 = prev
		prev = number
		numbers = append(numbers, number)
	}
	return numbers
}

func sumEven(numbers []int) (sum int) {
	for _, n := range numbers {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}

func main() {
	fmt.Println("Sum of the even numbers:", sumEven(fibo(10)))
}
