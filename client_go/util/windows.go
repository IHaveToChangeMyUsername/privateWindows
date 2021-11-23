package util

import (
	"io"
	"os"
	"os/exec"
)

type Windows struct {
	command    *exec.Cmd
	shellStdIn *io.WriteCloser
}

func (w *Windows) CreateSession() error {
	ps, _ := exec.LookPath("powershell.exe")
	cmd := exec.Command(ps, "-NoProfile", "-NonInteractive")

	stdIn, err := cmd.StdinPipe()
	w.shellStdIn = &stdIn
	w.command = cmd
	if err != nil {
		return err
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		return err
	}
	return nil
}

func (w *Windows) DestroySession() {
	(*w.shellStdIn).Close()
}

func (w *Windows) Run(command string) {
	io.WriteString(*w.shellStdIn, command+"\n")
}
