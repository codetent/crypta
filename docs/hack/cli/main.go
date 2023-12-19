package main

import (
	"path/filepath"
	"strings"

	"github.com/codetent/crypta/cmd"
	"github.com/spf13/cobra/doc"
)

func filePrepender(s string) string {
	base := strings.TrimSuffix(filepath.Base(s), filepath.Ext(s))
	cmdline := strings.Join(strings.Split(base, "_"), " ")
	pre := []string{
		"---",
		"hide_title: true",
		"sidebar_label: " + cmdline,
		"---",
	}
	return strings.Join(pre, "\n") + "\n"
}

func linkHandler(s string) string {
	return s
}

func main() {
	cmd := cmd.NewRootCmd()
	err := doc.GenMarkdownTreeCustom(cmd, "docs/pages/cli", filePrepender, linkHandler)
	if err != nil {
		panic(err)
	}
}
