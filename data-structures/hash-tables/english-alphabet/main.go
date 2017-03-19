package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bs)
	fmt.Println(str)
}
