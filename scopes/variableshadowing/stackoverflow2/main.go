package main

import "fmt"

func main() {
	x := 3

	// The named return in the function signature is declaring a variable called x within the double() scope
	// So, the x within double() is another x, not the same declared at the line
	// When it's returned, the "other x" is returned, and the "first x" keeping unchanged

	// If the function signature returns only (int) and not (x int), the "new x" wouldn't be
	// re-declared, and the function would double the value, changing the "first x"
	double := func() (x int) {
		x = x * 2
		return x
	}

	double()
	double()

	fmt.Println(x)
}
