package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/codetent/crypta/pkg/cli"
	"github.com/codetent/crypta/pkg/daemon"
	"github.com/spf13/cobra"
)

type setCmd struct {
	global *globalFlags

	hidden bool
}

func NewSetCmd(global *globalFlags) *cobra.Command {
	c := &setCmd{global: global}
	cc := &cobra.Command{
		Use:   "set",
		Short: "Set secret value",
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Run(args)
		},
	}

	cc.Flags().BoolVar(&c.hidden, "hidden", false, "Hide user input")

	return cc
}

func (c *setCmd) Run(args []string) error {
	name := args[0]
	var value string

	if len(args) < 2 {
		if c.global.interactive {
			prompt := fmt.Sprintf("Value for %s", name)

			var err error
			if c.hidden {
				value, err = cli.AskPassword(os.Stdin, os.Stderr, prompt)
			} else {
				value, err = cli.AskInput(os.Stdin, os.Stderr, prompt)
			}
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("value not provided")
		}
	} else {
		value = args[1]
	}

	client := daemon.NewDaemonClient(c.global.ip, c.global.port)
	err := client.SetSecret(context.Background(), name, value)
	if err != nil {
		return err
	}

	return nil
}
