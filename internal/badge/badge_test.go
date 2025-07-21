package badge

import (
	"errors"
	"io"
	"path/filepath"
	"strings"
	"testing"
)

// mockWriteCloser implements io.WriteCloser for testing
type mockWriteCloser struct {
	data       []byte
	writeErr   error
	closeErr   error
	closeCalled bool
}

func (m *mockWriteCloser) Write(p []byte) (n int, err error) {
	if m.writeErr != nil {
		return 0, m.writeErr
	}
	m.data = append(m.data, p...)
	return len(p), nil
}

func (m *mockWriteCloser) Close() error {
	m.closeCalled = true
	return m.closeErr
}

// mockFileWriter implements FileWriter for testing
type mockFileWriter struct {
	createFunc func(name string) (io.WriteCloser, error)
}

func (m *mockFileWriter) Create(name string) (io.WriteCloser, error) {
	if m.createFunc != nil {
		return m.createFunc(name)
	}
	return &mockWriteCloser{}, nil
}

func TestGenerator_Generate(t *testing.T) {
	tests := []struct {
		name       string
		outputFile string
		title      string
		value      string
		color      string
		setupMock  func() FileWriter
		wantErr    bool
		errContains string
	}{
		{
			name:       "successful badge generation",
			outputFile: "test.svg",
			title:      "Coverage",
			value:      "95%",
			color:      "#5272B4",
			setupMock: func() FileWriter {
				return &mockFileWriter{
					createFunc: func(name string) (io.WriteCloser, error) {
						return &mockWriteCloser{}, nil
					},
				}
			},
			wantErr: false,
		},
		{
			name:       "file creation error",
			outputFile: "test.svg",
			title:      "Coverage",
			value:      "95%",
			color:      "#5272B4",
			setupMock: func() FileWriter {
				return &mockFileWriter{
					createFunc: func(name string) (io.WriteCloser, error) {
						return nil, errors.New("permission denied")
					},
				}
			},
			wantErr:     true,
			errContains: "failed to create file",
		},
		{
			name:       "file write error",
			outputFile: "test.svg",
			title:      "Coverage",
			value:      "95%",
			color:      "#5272B4",
			setupMock: func() FileWriter {
				return &mockFileWriter{
					createFunc: func(name string) (io.WriteCloser, error) {
						return &mockWriteCloser{
							writeErr: errors.New("disk full"),
						}, nil
					},
				}
			},
			wantErr:     true,
			errContains: "failed to write badge",
		},
		{
			name:       "file close error",
			outputFile: "test.svg",
			title:      "Coverage",
			value:      "95%",
			color:      "#5272B4",
			setupMock: func() FileWriter {
				return &mockFileWriter{
					createFunc: func(name string) (io.WriteCloser, error) {
						return &mockWriteCloser{
							closeErr: errors.New("close failed"),
						}, nil
					},
				}
			},
			wantErr:     true,
			errContains: "failed to close file",
		},
		{
			name:       "path traversal attempt",
			outputFile: "../../../etc/passwd",
			title:      "Coverage",
			value:      "95%",
			color:      "#5272B4",
			setupMock: func() FileWriter {
				return &mockFileWriter{
					createFunc: func(name string) (io.WriteCloser, error) {
						// Verify that path was cleaned (should resolve to etc/passwd)
						expectedPath := filepath.Clean("../../../etc/passwd")
						if name != expectedPath {
							t.Errorf("Path not properly cleaned: got %s, want %s", name, expectedPath)
						}
						return &mockWriteCloser{}, nil
					},
				}
			},
			wantErr: false,
		},
		{
			name:       "empty color uses default",
			outputFile: "test.svg",
			title:      "Coverage",
			value:      "95%",
			color:      "",
			setupMock: func() FileWriter {
				return &mockFileWriter{
					createFunc: func(name string) (io.WriteCloser, error) {
						return &mockWriteCloser{}, nil
					},
				}
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator(tt.setupMock())
			err := g.Generate(tt.outputFile, tt.title, tt.value, tt.color)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
			
			if err != nil && tt.errContains != "" {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Generate() error = %v, want error containing %v", err, tt.errContains)
				}
			}
		})
	}
}

func TestGenerator_GenerateVerifyOutput(t *testing.T) {
	mockWriter := &mockWriteCloser{}
	fw := &mockFileWriter{
		createFunc: func(name string) (io.WriteCloser, error) {
			return mockWriter, nil
		},
	}
	
	g := NewGenerator(fw)
	err := g.Generate("test.svg", "Build", "Passing", "#4c1")
	
	if err != nil {
		t.Fatalf("Generate() unexpected error: %v", err)
	}
	
	// Verify SVG output
	output := string(mockWriter.data)
	if !strings.Contains(output, "<svg") {
		t.Error("Output does not contain SVG tag")
	}
	if !strings.Contains(output, "Build") {
		t.Error("Output does not contain title")
	}
	if !strings.Contains(output, "Passing") {
		t.Error("Output does not contain value")
	}
	
	// Verify file was closed
	if !mockWriter.closeCalled {
		t.Error("File was not closed")
	}
}

func TestNewGenerator(t *testing.T) {
	fw := &mockFileWriter{}
	g := NewGenerator(fw)
	
	if g == nil {
		t.Fatal("NewGenerator() returned nil")
	}
	
	if g.fileWriter != fw {
		t.Error("NewGenerator() did not set fileWriter correctly")
	}
}