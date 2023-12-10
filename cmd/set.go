package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/codetent/crypta/pkg/cli"
	"github.com/codetent/crypta/pkg/daemon"
	"github.com/spf13/cobra"
)

var (
	flagInteractive bool
	flagHidden      bool
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set secret value",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		var value string

		if len(args) < 2 {
			if flagInteractive {
				prompt := fmt.Sprintf("Value for %s", name)

				var err error
				if flagHidden {
					value, err = cli.AskPassword(os.Stdin, os.Stdout, prompt)
				} else {
					value, err = cli.AskInput(os.Stdin, os.Stdout, prompt)
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

		client := daemon.NewDaemonClient(flagIP, flagPort)
		err := client.SetSecret(context.Background(), name, value)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	setCmd.Flags().BoolVar(&flagHidden, "hidden", false, "Hide user input")
	setCmd.Flags().BoolVar(&flagInteractive, "interactive", true, "Ask for secret value")
	rootCmd.AddCommand(setCmd)
}
