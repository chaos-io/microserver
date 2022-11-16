package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Local is an implementation of the Storage interface which work with
// the local disk on the current machine
type Local struct {
	basePath    string
	maxFileSize int // maximum number of bytes for files
}

// NewLocal creates a new Local filesystem with the given base path
// bashPath is base directory to save files to
// maxSize is the max number of bytes that a file can be
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: p}, nil
}

// Save the contents of the Writer to the given path
// path is relative path, basePath will be appended
func (l *Local) Save(path string, contents io.Reader) error {
	fp := l.fullPath(path)

	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}

	_, err := os.Stat(fp)
	if err == nil {
		if err = os.Remove(fp); err != nil {
			return fmt.Errorf("unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("unable to get file info: %w", err)
	}

	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer f.Close()

	if _, err = io.Copy(f, contents); err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}

	return nil
}

// Get the file at the given path and return a Reader
//
func (l *Local) Get(path string) (*os.File, error) {
	fp := l.fullPath(path)

	f, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	return f, nil
}

// get the absolute path
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
