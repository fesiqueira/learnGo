package main

import (
	"fmt"

	"github.com/fel-siqueira/learnGo/training/part2/ezcli"
)

func main() {
	program := ezcli.NewProgram()
	fmt.Println(program)
	fmt.Println(program.Compiled)
}
