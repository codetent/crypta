package cmd

import (
	"github.com/codetent/crypta/pkg/daemon"
	"github.com/codetent/crypta/pkg/store"
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
	store := store.NewLocalSecretStore(
		store.WithEnvPrefix("CRYPTA_SECRET_"),
		store.WithLocalPath("/var/run/secrets/crypta"),
	)

	server := daemon.NewDaemonServer(store)
	return server.ListenAndServe(c.global.ip, c.global.port)
}
