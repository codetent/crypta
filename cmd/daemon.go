package cmd

import (
	"github.com/codetent/crypta/pkg/daemon"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the crypta daemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := daemon.NewDaemonServer(flagIP, flagPort)
		return server.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
