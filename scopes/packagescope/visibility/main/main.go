package main

import (
	"fmt"

	"github.com/fsiqp/learnGo/scopes/packagescope/visibility/visi"
)

func main() {
	fmt.Println(visi.MyName)
	visi.PrintVar()

}
