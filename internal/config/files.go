package config

import (
	"bufio"
	"fmt"
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
		return fmt.Errorf("error creating required directories for %s: %w", f.Path, err)
	}

	file, err := os.Create(f.Path)

	if err != nil {
		return fmt.Errorf("failed to create file on %s: %w", f.Path, err)
	}

	defer file.Close()

	bw := bufio.NewWriter(file)

	if _, err := bw.WriteString(f.Content); err != nil {
		return fmt.Errorf("unable to write to file %s: %w", f.Path, err)
	}

	return bw.Flush()
}
