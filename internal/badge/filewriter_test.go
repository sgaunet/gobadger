package badge

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOSFileWriter_Create(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	
	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "create file successfully",
			filename: filepath.Join(tmpDir, "test.svg"),
			wantErr:  false,
		},
		{
			name:     "create file in non-existent directory",
			filename: filepath.Join(tmpDir, "nonexistent", "dir", "test.svg"),
			wantErr:  true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fw := NewOSFileWriter()
			f, err := fw.Create(tt.filename)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			
			if err == nil {
				// Clean up
				f.Close()
				os.Remove(tt.filename)
			}
		})
	}
}

func TestNewOSFileWriter(t *testing.T) {
	fw := NewOSFileWriter()
	if fw == nil {
		t.Fatal("NewOSFileWriter() returned nil")
	}
}