package cmd

import (
	"context"
	"fmt"

	"github.com/codetent/crypta/pkg/daemon"
	"github.com/spf13/cobra"
)

type getCmd struct {
	global *globalFlags
}

func NewGetCmd(global *globalFlags) *cobra.Command {
	c := &getCmd{global: global}
	cc := &cobra.Command{
		Use:   "get",
		Short: "Get secret value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Run(args)
		},
	}

	return cc
}

func (c *getCmd) Run(args []string) error {
	client := daemon.NewDaemonClient(c.global.ip, c.global.port)
	val, err := client.GetSecret(context.Background(), args[0])
	if err != nil {
		return err
	}

	fmt.Println(val)
	return nil
}
