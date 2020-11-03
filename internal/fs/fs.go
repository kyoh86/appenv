package fs

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
)

type FileManager interface {
	Open(filename string) (io.WriteCloser, error)
}

type Dir string

func (d Dir) Open(filename string) (io.WriteCloser, error) {
	return os.OpenFile(filepath.Join(string(d), filename), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
}

type Memory map[string]nopCloser

type nopCloser struct {
	*bytes.Buffer
}

func (nopCloser) Close() error { return nil }

func (b Memory) Open(filename string) (io.WriteCloser, error) {
	b[filename] = nopCloser{Buffer: &bytes.Buffer{}}
	return b[filename], nil
}

func (b Memory) Result(filename string) string {
	buffer, ok := b[filename]
	if !ok {
		return ""
	}
	return buffer.String()
}
