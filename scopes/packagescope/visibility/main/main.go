package main

import (
	"fmt"

	"github.com/fel-siqueira/learnGo/scopes/packagescope/visibility/visi"
)

func main() {
	fmt.Println(visi.MyName)
	visi.PrintVar()

}
