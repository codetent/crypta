package main

import (
	"github.com/codetent/crypta/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := cmd.NewRootCmd()
	err := doc.GenMarkdownTree(cmd, "docs/pages/cli")
	if err != nil {
		panic(err)
	}
}
