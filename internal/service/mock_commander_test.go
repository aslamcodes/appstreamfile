package service_test

import (
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

type FakeCmd struct {
	Command string
	Args    []string
}

func (c *FakeCmd) CombinedOutput() ([]byte, error) {
	return []byte{}, nil
}

func (c *FakeCmd) String() string {
	return ""
}

type FakeCommander struct {
	LastCommand string
	LastArgs    []string

	LookPathErr error
}

func (fc *FakeCommander) LookPath(file string) (string, error) {
	return file, fc.LookPathErr
}

func (fc *FakeCommander) Command(name string, args ...string) execx.Cmd {
	fc.LastCommand = name
	fc.LastArgs = args
	return &FakeCmd{
		Command: name,
		Args:    args,
	}
}
