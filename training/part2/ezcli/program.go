package ezcli

import (
	"os"
	"path/filepath"
	"time"
)

type Program struct {
	Name        string
	Version     string
	Description string
	Compiled    time.Time
	Author      Author
}

func NewProgram() *Program {
	return &Program{
		Name:     filepath.Base(os.Args[0]),
		Version:  "0.0.0",
		Compiled: compileTime(),
	}
}

func compileTime() time.Time {
	info, err := os.Stat(os.Args[0])
	if err != nil {
		return time.Now()
	}
	return info.ModTime()
}
