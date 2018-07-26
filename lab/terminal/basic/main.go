package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	cmd := exec.Command("bash")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(cmd.Stdin)
	w := bufio.NewWriter(cmd.Stdout)

	rw := bufio.NewReadWriter(r, w)

	term := terminal.NewTerminal(rw, ">")

	n, err := term.Write([]byte("ls"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
