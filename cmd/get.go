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
		Short: "Get cached secret value",
		Long:  "Retrieves the cached secret value which has been either set by using the provided `set` command, or by setting environment variables. The environment variables have the following format: `CRYPTA_SECRET_<KEY>`. The value given for `<KEY>` will be used as the name with which the value of the environment variable will be added to the secret store.",
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

			err = client.SetSecret(context.Background(), name, value)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else if err != nil {
		return fmt.Errorf("the daemon does not seem to be running: %w", err)
	}

	fmt.Println(value)
	return nil
}
