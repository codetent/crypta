package main

import (
	"os"

	"github.com/codetent/crypta/cmd"
)

var (
	version = "0.0.0"
)

func main() {
	root := cmd.NewRootCmd()
	root.Version = version

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
