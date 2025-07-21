package badge

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/narqo/go-badge"
)

// FileWriter interface for file operations.
type FileWriter interface {
	Create(name string) (io.WriteCloser, error)
}

// Generator handles badge generation.
type Generator struct {
	fileWriter FileWriter
}

// NewGenerator creates a new badge generator.
func NewGenerator(fw FileWriter) *Generator {
	return &Generator{
		fileWriter: fw,
	}
}

// Generate creates a badge SVG file.
func (g *Generator) Generate(outputFile, title, value, color string) (err error) {
	// Clean the output file path to prevent path traversal
	cleanPath := filepath.Clean(outputFile)
	
	f, err := g.fileWriter.Create(cleanPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			if err == nil {
				err = fmt.Errorf("failed to close file: %w", cerr)
			}
		}
	}()

	b, err := badge.RenderBytes(title, value, badge.Color(color))
	if err != nil {
		return fmt.Errorf("failed to render badge: %w", err)
	}
	_, err = f.Write(b)
	if err != nil {
		return fmt.Errorf("failed to write badge: %w", err)
	}
	return nil
}