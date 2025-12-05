package config

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

type File struct {
	Path    string `yaml:"path"`
	Content string `yaml:"content"`
}

func (f *File) Deploy(w io.Writer) error {
	if err := os.MkdirAll(filepath.Dir(f.Path), 0770); err != nil {
		return err
	}

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
