package main

import "fmt"

func fibo(j int) []int {
	var ns []int
	a, n := 0, 1
	for i := 1; i < j; i++ {
		a, n = n, a+n
		ns = append(ns, n)
	}
	return ns
}

func sumEven(x int) int {
	var sum int
	ns := fibo(x)
	for _, n := range ns {
		if n%2 == 0 && sum < 4000000 {
			sum += n
		}
	}
	return sum
}

func main() {
	fmt.Println(fibo(10))
	fmt.Println(sumEven(10))
}
