package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"golang.org/x/crypto/ssh"
)

var sshConfig = &ssh.ClientConfig{
	User: "fesiqp",
	Auth: []ssh.AuthMethod{
		ssh.Password("1123"),
	},
	HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error { return nil },
}

func main() {
	conn, err := ssh.Dial("tcp", "localhost:22", sshConfig)
	if err != nil {
		log.Fatal("Dial failed:", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("New session failed: ", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		// ssh.ECHO:  0,
		// ssh.IGNCR: 0,
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatal("request pty failed:", err)
	}

	if err = session.Shell(); err != nil {
		log.Fatal("shell failed:", err)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for {
			<-c
			log.Println("sigint")
		}
	}()

	if err = session.Wait(); err != nil {
		log.Fatal("wait:", err)
	}
}
