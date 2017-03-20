package main

import "fmt"
import "encoding/json"

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	// In this example, the JSON fields should have the same name of the Struct fields
	// Initializing a slice of bytes with a JSON
	bs := []byte(`{"First":"Felipe", "Last":"Siqueira", "Age":20}`)

	// Unmarshal is assigning the JSON fields to the Struct fields, it receives the slice of bytes containing the JSON and the pointer to the target Struct
	json.Unmarshal(bs, &p1)

	fmt.Println("-----------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T\n", p1)
}
