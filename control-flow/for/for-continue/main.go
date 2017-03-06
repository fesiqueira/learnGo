package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			// Continue makes the program to run from the start of the loop again
			// It doesn't let all for block of code run completelly
			// So after the continue, the next step will be i++ again
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			// Break will make the program run after the outerest loop
			// Or stop, if there's no program left to run
			break
		}
	}
}
