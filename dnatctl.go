package main

import (
	"./pkg/cmd"
	"os"
)


func main() {

	command := cmd.NewDnatctlCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}