package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Scan the page
	// NewScanner takes a reader and res.Body implements the reader interface
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// Iterate over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
