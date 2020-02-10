package main

import (
	"github.com/kenchaaan/dnatctl/pkg/cmd"
	"os"
)

func a() {

	command := cmd.NewDeafultDnatctlCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
