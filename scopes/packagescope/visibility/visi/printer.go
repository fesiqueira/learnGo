package visi

// Imports have a file scope
// If we needed to use fmt package in another package visi file
// We should import it again
import "fmt"

// MyName is exported, so its accessible from any main file that imports visi package
var MyName = "Felipe"

// myname is unexported, so it's just accessible from the files that is part of the visi package
var myname = "Siqueira"

func PrintVar() {
	fmt.Println(MyName)

	// myname is accessible from here, because it's declared at the package scope
	fmt.Println(myname)
}
