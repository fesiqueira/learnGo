package main

import (
	"fmt"

	"github.com/fel-siqueira/learnGo/training/part2/ezcli"
)

func main() {
	s1 := ezcli.NewUser()
	fmt.Println(s1.OS)
	fmt.Println(s1.Username)
}
