package execx

import "context"

type Commander interface {
	LookPath(file string) (string, error)
	Command(name string, arg ...string) Cmd
	CommandContext(ctx context.Context, name string, arg ...string) Cmd
}

type Cmd interface {
	CombinedOutput() ([]byte, error)
	String() string
}
