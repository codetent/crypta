package cmd

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"time"

	"github.com/codetent/crypta/pkg/daemon"
	"github.com/shirou/gopsutil/v3/process"
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

	// get the process id of the running daemon
	pid, err := client.GetProcessId(context.Background())

	if err != nil {
		return err
	}

	// try to stop the running daemon
	p, err := process.NewProcess(pid)
	if err != nil {
		if err == process.ErrorProcessNotRunning {
			return nil
		}

		return err
	}

	if err = p.Terminate(); err != nil {
		return err
	}

	// check if the daemon has been stopped
	for timeout := time.After(2 * time.Second); ; {
		select {
		case <-timeout:
			return errors.New("checking whether the daemon has stopped timed out")
		default:
			exists, err := process.PidExistsWithContext(context.Background(), pid)

			if err != nil {
				return err
			}

			if !exists {
				return nil
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (c *daemonCmd) run() error {
	server := daemon.NewDaemonServer(c.global.ip, c.global.port)
	return server.ListenAndServe()
}
