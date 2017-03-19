package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	var nounce [24]byte
	fmt.Println(nounce)

	// nounce[:] is the slice of the nounce array
	io.ReadFull(rand.Reader, nounce[:])
	fmt.Println(nounce)
}
