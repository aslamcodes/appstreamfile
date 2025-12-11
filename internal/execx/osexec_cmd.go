package execx

import "os/exec"

type ExecCmd struct {
	*exec.Cmd
}

func (c *ExecCmd) CombinedOutput() ([]byte, error) {
	return c.Cmd.CombinedOutput()
}
