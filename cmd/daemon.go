package cmd

import (
	"context"

	"github.com/codetent/crypta/pkg/daemon"
	"github.com/codetent/crypta/pkg/trace"
	"github.com/spf13/cobra"
)

type daemonCmd struct {
	global *globalFlags
}

func NewDaemonCmd(global *globalFlags) *cobra.Command {
	c := &daemonCmd{global: global}
	cc := &cobra.Command{
		Use:   "daemon",
		Short: "Run the crypta daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Run(args)
		},
	}

	return cc
}

func (c *daemonCmd) Run(args []string) error {
	shutdown, err := trace.SetupTracing()
	if err != nil {
		return err
	}
	defer shutdown(context.Background())

	server := daemon.NewDaemonServer(c.global.ip, c.global.port)
	return server.ListenAndServe()
}
