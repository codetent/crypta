package cmd

import (
	"github.com/codetent/crypta/pkg/daemon"
	"github.com/codetent/crypta/pkg/store"
	"github.com/spf13/cobra"
)

type daemonCmd struct {
	global *globalFlags
	path   string
	envP   string
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

	cc.Flags().StringVar(&c.path, "path", "/var/run/secrets/crypta", "Path to secret file folder to load at startup")
	cc.Flags().StringVar(&c.envP, "env", "CRYPTA_SECRET_", "Prefix for environment variables to load at startup")

	return cc
}

func (c *daemonCmd) Run(args []string) error {
	store := store.NewLocalSecretStore(
		store.WithEnvPrefix(c.envP),
		store.WithLocalPath(c.path),
	)

	server := daemon.NewDaemonServer(store)
	return server.ListenAndServe(c.global.ip, c.global.port)
}
