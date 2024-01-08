package cmd

import (
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
			// PersistentPreRun seems to be run after the FlagError functions have been called.
			// The usage shall be printed if there is an error in the usage (e.g., flags are missing)
			// But shall not printed if the RunE functions return errors (i.e. application errors)
			cmd.SilenceUsage = true

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
		os.Exit(1)
	}
}
