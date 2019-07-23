package main

import (
	"fmt"
	"github.com/koind/go-envdir/exec"
	flag "github.com/spf13/pflag"
)

var envDir, command string

func init() {
	flag.StringVarP(&envDir, "env", "e", "", "path to environment variables")
	flag.StringVarP(&command, "command", "c", "", "unix utility name")
}

func main() {
	flag.Parse()

	if envDir == "" || command == "" {
		fmt.Println("Specify parameters")
		return
	}

	exec.Command(command, envDir)
}
