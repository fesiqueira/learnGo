package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	"golang.org/x/crypto/ssh"
)

var sshConfig = &ssh.ClientConfig{
	User: os.Getenv("USER"),
	Auth: []ssh.AuthMethod{
		ssh.Password(os.Getenv("SSH_PASS")),
	},
	HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error { return nil },
}

var termModes = ssh.TerminalModes{
	ssh.ECHO:          0,
	ssh.ECHOCTL:       0,
	ssh.TTY_OP_ISPEED: 14400,
	ssh.TTY_OP_OSPEED: 14400,
}

func main() {
	conn, err := ssh.Dial("tcp", "localhost:22", sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// if err = session.RequestPty(os.Getenv("TERM"), 80, 40, termModes); err != nil {
	// 	session.Close()
	// 	log.Fatal(err)
	// }

	stdin, _ := session.StdinPipe()
	go io.Copy(stdin, os.Stdin)
	stdout, _ := session.StdoutPipe()
	go io.Copy(os.Stdout, stdout)
	stderr, _ := session.StderrPipe()
	go io.Copy(os.Stderr, stderr)

	if err = session.Shell(); err != nil {
		session.Close()
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			session.Signal(
				ssh.Signal([]byte{3}),
			)
			fmt.Println("sent", ssh.Signal([]byte{3}))
		}
	}()

	if err = session.Wait(); err != nil {
		session.Close()
		log.Fatal(err)
	}
}
