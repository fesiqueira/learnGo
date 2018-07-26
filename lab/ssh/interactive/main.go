// TODO: the remote exit status does not reset when pressing ENTER
// Clean this shitty code
// Somehow send os.Exit (or something similar) but honor the deferred functions
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"unsafe"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var sshConfig = &ssh.ClientConfig{
	User: "fesiqp",
	Auth: []ssh.AuthMethod{
		ssh.Password("1123"),
	},
	HostKeyCallback: func(string, net.Addr, ssh.PublicKey) error { return nil },
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func main() {
	conn, err := ssh.Dial("tcp", "192.168.100.150:22", sshConfig)
	if err != nil {
		log.Fatal("Dial failed:", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("New session failed: ", err)
	}
	defer session.Close()

	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = terminal.Restore(int(os.Stdin.Fd()), oldState)
	}()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:  0,
		ssh.IGNCR: 0,
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatal("request pty failed:", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			fmt.Println("resizing")
			if err := changeSize(os.Stdin, session); err != nil {
				log.Printf("error resizing: %s", err)
			}
		}
	}()

	if err = session.Shell(); err != nil {
		log.Fatal("shell failed:", err)
	}

	ch <- syscall.SIGWINCH
	if err = session.Wait(); err != nil {
		if exitStatus, ok := err.(*ssh.ExitError); ok {
			os.Exit(exitStatus.ExitStatus())
		}
	}
}

func changeSize(r io.Reader, session *ssh.Session) error {
	ws := &winsize{}
	retCode, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if int(retCode) == -1 {
		return errors.New("something is fucked" + err.Error())
	}
	return session.WindowChange(int(ws.Row), int(ws.Col))
}
