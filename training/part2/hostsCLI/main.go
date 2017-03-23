package main

import (
	"fmt"

	"github.com/fel-siqueira/learnGo/training/part2/ezcli"
)

func main() {
	arguments := ezcli.UserArgs(2)
	fmt.Println(arguments)
}
