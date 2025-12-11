package execx

import "os/exec"

type ExecCommander struct{}

func (ExecCommander) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

func (ExecCommander) Command(name string, arg ...string) Cmd {
	return &ExecCmd{exec.Command(name, arg...)}
}
