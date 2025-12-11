package service_test

type FakeCmd struct {
}

func (c *FakeCmd) CombinedOutput() ([]byte, error) {
	return []byte{}, nil
}

type FakeCommander struct {
}

func (FakeCommander) LookPath(file string) (string, error) {
	return "", nil
}

func (FakeCommander) Command(name string, arg ...string) FakeCmd {
	return FakeCmd{}
}
