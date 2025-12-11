package execx

type Commander interface {
	LookPath(file string) (string, error)
	Command(name string, arg ...string) Cmd
}

type Cmd interface {
	CombinedOutput() ([]byte, error)
}
