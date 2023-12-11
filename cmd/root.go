package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type globalFlags struct {
	ip          string
	port        string
	interactive bool
}

func NewRootCmd() *cobra.Command {
	cc := &cobra.Command{
		Use:   "crypta",
		Short: "",
		Long:  ``,
	}

	gf := &globalFlags{}
	cc.PersistentFlags().StringVar(&gf.ip, "ip", "127.0.0.1", "IP to bind daemon")
	cc.PersistentFlags().StringVar(&gf.port, "port", "35997", "Port to bind daemon")
	cc.PersistentFlags().BoolVar(&gf.interactive, "interactive", true, "Allow user input")

	cc.AddCommand(NewDaemonCmd(gf))
	cc.AddCommand(NewGetCmd(gf))
	cc.AddCommand(NewSetCmd(gf))

	return cc
}

func Execute() {
	cmd := NewRootCmd()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
