package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/codetent/crypta/pkg/cli"
	"github.com/codetent/crypta/pkg/daemon"
	"github.com/spf13/cobra"
)

type getCmd struct {
	global *globalFlags

	hidden bool
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

	cc.Flags().BoolVar(&c.hidden, "hidden", false, "Hide user input")

	return cc
}

func (c *getCmd) Run(args []string) error {
	name := args[0]

	client := daemon.NewDaemonClient(c.global.ip, c.global.port)
	value, err := client.GetSecret(context.Background(), name)

	if errors.Is(err, daemon.ErrSecretNotExists) {
		if c.global.interactive {
			prompt := fmt.Sprintf("Value for %s", name)

			if c.hidden {
				value, err = cli.AskPassword(os.Stdin, os.Stderr, prompt)
			} else {
				value, err = cli.AskInput(os.Stdin, os.Stderr, prompt)
			}

			if err != nil {
				return err
			}

			client.SetSecret(context.Background(), name, value)
		} else {
			return err
		}
	}

	fmt.Println(value)
	return nil
}
