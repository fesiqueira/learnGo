package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// The underline represents the blank identifier
// The http.Get and the ioutil.ReadAll have multiple returns,
// But using the blank identifier, we are just keeping the res and page variables,
// and throwing away the err that returns together from each method.

// It means that the code below isn't using error checking, which is bad
// for production!
func main() {
	res, _ := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
