package service_test

import (
	"github.com/aslamcodes/appstreamfile/internal/execx"
)

type FakeCmd struct {
}

func (c *FakeCmd) CombinedOutput() ([]byte, error) {
	return []byte{}, nil
}

type FakeCommander struct {
	LastCommand string
	LastArgs    []string

	LookPathErr error
}

func (fc *FakeCommander) LookPath(file string) (string, error) {
	return file, fc.LookPathErr
}

func (fc *FakeCommander) Command(name string, arg ...string) execx.Cmd {
	fc.LastCommand = name
	fc.LastArgs = arg
	return &FakeCmd{}
}
