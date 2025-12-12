package execx

import "os/exec"

type ExecCmd struct {
	*exec.Cmd
}

func (c *ExecCmd) CombinedOutput() ([]byte, error) {
	return c.Cmd.CombinedOutput()
}

func (c *ExecCmd) String() string{
	return c.Cmd.String()
}

type ExecCommander struct{}

func (ex *ExecCommander) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

func (ex *ExecCommander) Command(name string, arg ...string) Cmd {
	return &ExecCmd{exec.Command(name, arg...)}
}
