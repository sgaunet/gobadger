package badge

import (
	"fmt"
	"io"
	"os"
)

// OSFileWriter implements FileWriter using the OS filesystem.
type OSFileWriter struct{}

// NewOSFileWriter creates a new OS file writer.
func NewOSFileWriter() *OSFileWriter {
	return &OSFileWriter{}
}

// Create creates a new file.
func (o *OSFileWriter) Create(name string) (io.WriteCloser, error) {
	// #nosec G304 - Path is sanitized by the caller (badge.Generate)
	f, err := os.Create(name)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	return f, nil
}