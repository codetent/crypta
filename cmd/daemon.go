package cmd

import (
	"context"
	"os"
	"os/exec"

	"github.com/codetent/crypta/pkg/daemon"
	"github.com/spf13/cobra"
)

type daemonCmd struct {
	global *globalFlags

	detached bool
}

func NewDaemonCmd(global *globalFlags) *cobra.Command {
	c := &daemonCmd{global: global}
	cc := &cobra.Command{
		Use:   "daemon",
		Short: "Run the crypta daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			if !c.detached {
				return cmd.Help()
			}

			return c.run()
		},
	}

	start := &cobra.Command{
		Use:   "start",
		Short: "Starts the crypta daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.start()
		},
	}

	stop := &cobra.Command{
		Use:   "stop",
		Short: "Stops the crypta daemon",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.stop()
		},
	}

	cc.AddCommand(start)
	cc.AddCommand(stop)

	cc.Flags().BoolVar(&c.detached, "detached", false, "Runs the daemon detached")
	_ = cc.Flags().MarkHidden("detached")

	return cc
}

func (c *daemonCmd) start() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	ex, err := os.Executable()
	if err != nil {
		return err
	}

	// FIXME: Pass global flags to command, in order to be able to set the endpoint
	cmd := exec.Command(ex, "daemon", "--detached")
	cmd.Dir = cwd
	return cmd.Start()
}

func (c *daemonCmd) stop() error {
	client := daemon.NewDaemonClient(c.global.ip, c.global.port)
	return client.Shutdown(context.Background())
}

func (c *daemonCmd) run() error {
	server := daemon.NewDaemonServer(c.global.ip, c.global.port)
	return server.ListenAndServe()
}
