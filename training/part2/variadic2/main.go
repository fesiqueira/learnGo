package main

func foo(numbers ...int) {

}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3}
	foo(aSlice...)
	foo()
}
