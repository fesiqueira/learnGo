package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person
	// When we have a string, and want to turn that string into a reader, we can use the strings.NewReader
	rdr := strings.NewReader(`{"First":"Felipe", "Last":"Siqueira", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
