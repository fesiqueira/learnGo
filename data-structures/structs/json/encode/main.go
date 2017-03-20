package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}

	// NewEncoder writes the encoding in some place, in this case it's writing at the StandardOut
	// And the encoding result is assigned to the p1
	json.NewEncoder(os.Stdout).Encode(p1)
}
