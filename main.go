package main

import (
	"github.com/Kuniwak/name/cmd"
	"os"
)

func main() {
	exitStatus := cmd.MainCmd(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	os.Exit(int(exitStatus))
}
