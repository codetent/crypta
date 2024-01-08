package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type globalFlags struct {
	ip          string
	port        string
	interactive bool
	verbose     bool
}

func NewRootCmd() *cobra.Command {
	gf := &globalFlags{}

	cc := &cobra.Command{
		Use:   "crypta",
		Short: "",
		Long:  ``,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !gf.verbose {
				log.SetOutput(io.Discard)
			}
		},
	}

	cc.PersistentFlags().StringVar(&gf.ip, "ip", "127.0.0.1", "IP to bind daemon")
	cc.PersistentFlags().StringVar(&gf.port, "port", "35997", "Port to bind daemon")
	cc.PersistentFlags().BoolVar(&gf.interactive, "interactive", true, "Allow user input")
	cc.PersistentFlags().BoolVarP(&gf.verbose, "verbose", "v", false, "Enable verbose logging")

	cc.AddCommand(NewDaemonCmd(gf))
	cc.AddCommand(NewGetCmd(gf))
	cc.AddCommand(NewSetCmd(gf))

	return cc
}

func Execute() {
	cmd := NewRootCmd()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
