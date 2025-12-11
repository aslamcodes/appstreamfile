package execx

import "os/exec"

type ExecCommander struct{}

func (ex *ExecCommander) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

func (ex *ExecCommander) Command(name string, arg ...string) Cmd {
	return &ExecCmd{exec.Command(name, arg...)}
}
