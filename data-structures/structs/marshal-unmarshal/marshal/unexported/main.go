package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"Felipe", "Siqueira", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)

	// The JSON is empty because the Marshal only converts the Exported fields (capitalized fields)
	fmt.Println(string(bs))
}
