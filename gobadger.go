package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/narqo/go-badge"
)

func main() {
	var outputFile, color, title, value string
	flag.StringVar(&outputFile, "o", "badge.svg", "output file name")
	flag.StringVar(&color, "c", "#5272B4", "color of badge")
	flag.StringVar(&title, "t", "", "title")
	flag.StringVar(&value, "v", "", "Value for the title")
	flag.Parse()

	if title == "" || value == "" {
		fmt.Fprintln(os.Stderr, "title and value are required")
		flag.Usage()
		os.Exit(1)
	}

	err := generateBadge(outputFile, title, value, color)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func generateBadge(outputFile, title, value, color string) error {
	// Clean the output file path to prevent path traversal
	cleanPath := filepath.Clean(outputFile)
	
	f, err := os.Create(cleanPath)
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
