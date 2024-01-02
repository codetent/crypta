package cmd

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
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
	if err = cmd.Start(); err != nil {
		return err
	}

	// check if the daemon is started up and responsive
	client := daemon.NewDaemonClient(c.global.ip, c.global.port)

	for timeout := time.After(2 * time.Second); ; {
		select {
		case <-timeout:
			return errors.New("checking whether the daemon has been started timed out")
		default:
			// query the interface of the daemon to determine if it is available
			if _, unresponsive := client.GetProcessId(context.Background()); unresponsive == nil {
				return nil
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (c *daemonCmd) stop() error {
	client := daemon.NewDaemonClient(c.global.ip, c.global.port)

	// get the process id of the running daemon
	pid, err := client.GetProcessId(context.Background())

	if err != nil {
		return err
	}

	log.Println("Returned PID:", pid)

	// try to stop the running daemon
	p, err := process.NewProcess(pid)
	if err != nil {
		if err == process.ErrorProcessNotRunning {
			return nil
		}

		return err
	}

	log.Println("Running processes:")
	processes, _ := process.Processes()
	for _, p := range processes {
		name, _ := p.Name()
		log.Println(name, ":", p.Pid)
	}

	if err = p.Terminate(); err != nil {
		return err
	}

	log.Println("Before checking if daemon has stopped")

	// check if the daemon has been stopped
	for timeout := time.After(2 * time.Second); ; {
		select {
		case <-timeout:
			log.Println("Timeout elapsed")
			return errors.New("checking whether the daemon has stopped timed out")
		default:
			exists, err := process.PidExistsWithContext(context.Background(), pid)

			log.Println("PidExistsWithContext Exists:", exists, "err:", err)

			if err != nil {
				return err
			}

			if !exists {
				return nil
			}

			err = p.SendSignal(syscall.Signal(0))
			if err == nil {
				log.Println("Process is still found")
			}
			if errors.Is(err, os.ErrProcessDone) {
				log.Println("Process is not found: ErrProcessDone")
			}
			var errno syscall.Errno
			if !errors.As(err, &errno) {
				log.Println("Process is not found: errors.As(err, &errno)")
			}
			switch errno {
			case syscall.ESRCH:
				log.Println("Process is not found: ESRCH")
			case syscall.EPERM:
				log.Println("Process is still found: EPERM")
			default:
			}

			log.Println("Running processes:")
			processes, _ := process.Processes()
			for _, p := range processes {
				name, _ := p.Name()
				log.Println(name, ":", p.Pid)
			}

			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (c *daemonCmd) run() error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()

	server := daemon.NewDaemonServer(c.global.ip, c.global.port)
	return server.ListenAndServe()
}
