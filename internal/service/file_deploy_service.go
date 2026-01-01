package service

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aslamcodes/appstreamfile/internal/config"
)

type FileDeploySvc struct {
}

func (s *FileDeploySvc) DeployFile(ctx context.Context, f *config.File) error {

	if err := ctx.Err(); err != nil {
		return err
	}

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
