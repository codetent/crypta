package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	flagIP   string
	flagPort string
)

var rootCmd = &cobra.Command{
	Use:   "crypta",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&flagIP, "ip", "127.0.0.1", "IP to bind daemon")
	rootCmd.PersistentFlags().StringVar(&flagPort, "port", "35997", "Port to bind daemon")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
