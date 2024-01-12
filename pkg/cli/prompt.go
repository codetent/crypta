package cli

import (
	"fmt"
	"io"
	"os"
	"os/signal"

	"golang.org/x/term"
)

func AskInput(r io.Reader, w io.Writer, prompt string) (string, error) {
	fmt.Fprintf(w, "%s: ", prompt)

	var input string
	_, err := fmt.Fscanln(r, &input)
	if err != nil {
		if err.Error() == "unexpected newline" {
			return "", nil
		} else {
			return "", err
		}
	}

	return input, nil
}

func AskPassword(f *os.File, w io.Writer, prompt string) (string, error) {
	fd := int(f.Fd())

	oldState, err := term.GetState(fd)
	if err != nil {
		return "", err
	}
	defer func() { _ = term.Restore(fd, oldState) }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		for range sigCh {
			_ = term.Restore(fd, oldState)
			os.Exit(1)
		}
	}()

	fmt.Fprintf(w, "%s: ", prompt)

	passB, err := term.ReadPassword(fd)
	if err != nil {
		return "", err
	}

	fmt.Fprintln(w)

	return string(passB), nil
}
