package main

import (
	"os"

	"github.com/codetent/crypta/cmd"
)

var (
	version = "dev"
)

func main() {
	root := cmd.NewRootCmd()
	root.Version = version

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
