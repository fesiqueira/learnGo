package main

import "fmt"

func zero(x *int) {
	fmt.Printf("%p\n", x)
	*x = 0
}

// Notice that both printed addresses are the same
// It means the both x are the same
// Changing any of them, will change the other too
func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
}
