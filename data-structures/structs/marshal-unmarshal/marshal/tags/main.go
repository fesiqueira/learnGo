package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`            // the "-" is throwing the Last field away
	Age   int    `json:"wisdom score"` // wisdom score is renaming (tagging) the Age field when the struct is converted to JSON
}

func main() {
	p1 := Person{"Felipe", "Siqueira", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
