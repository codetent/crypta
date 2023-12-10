package cmd

import (
	"context"
	"fmt"

	"github.com/codetent/crypta/pkg/daemon"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get secret value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := daemon.NewDaemonClient(flagIP, flagPort)
		val, err := client.GetSecret(context.Background(), args[0])
		if err != nil {
			return err
		}

		fmt.Println(val)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
