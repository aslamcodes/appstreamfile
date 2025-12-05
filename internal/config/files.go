package config

import (
	"bufio"
	"io"
	"os"
)

type File struct {
	Path    string `yaml:"path"`
	Content string `yaml:"content"`
}

func (f *File) Deploy(w io.Writer) error {
	file, err := os.Create(f.Path)

	if err != nil {
		return err
	}

	defer file.Close()

	bw := bufio.NewWriter(file)

	if _, err := bw.WriteString(f.Content); err != nil {
		return err
	}

	return bw.Flush()
}
