package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get the Moby Dick book
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// Create a slice to hold counts
	buckets := make([]int, 200)

	// Iterate over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])
	// fmt.Println("*****************")
	// for i := 28; i <= 126; i++ {
	// 	fmt.Printf("%v - %c - %v\n", i, i, buckets[i])
	// }
}

func HashBucket(word string) int {
	return int(word[0])
}
