package cli

import (
	"io"
	"os"
)

type ProcInout struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func DefaultProcInout() ProcInout {
	return ProcInout{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}
