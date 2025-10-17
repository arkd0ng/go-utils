package validation

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid file paths
		{"valid absolute path", "/usr/bin/test", false},
		{"valid relative path", "./test.txt", false},
		{"valid relative path 2", "test.txt", false},
		{"valid path with dir", "dir/file.txt", false},
		{"valid nested path", "a/b/c/file.txt", false},
		{"valid path with spaces", "path with spaces/file.txt", false},

		// Invalid file paths
		{"empty string", "", true},
		{"dot only", ".", true},
		{"non-string value", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "file_path")
			v.FilePath()

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "test_file_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "test_dir_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid existing paths
		{"existing file", tmpFile.Name(), false},
		{"existing directory", tmpDir, false},

		// Invalid paths
		{"non-existing file", "/nonexistent/path/file.txt", true},
		{"non-existing directory", "/nonexistent/directory", true},
		{"empty string", "", true},
		{"non-string value", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "file_path")
			v.FileExists()

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

func TestFileReadable(t *testing.T) {
	// Create a temporary readable file
	tmpFile, err := os.CreateTemp("", "test_readable_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("test content")
	tmpFile.Close()

	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid readable files
		{"readable file", tmpFile.Name(), false},

		// Invalid cases
		{"non-existing file", "/nonexistent/path/file.txt", true},
		{"non-string value", 123, true},
		{"nil value", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "file_path")
			v.FileReadable()

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

func TestFileWritable(t *testing.T) {
	// Create a temporary writable file
	tmpFile, err := os.CreateTemp("", "test_writable_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Create a temporary directory for testing new file creation
	tmpDir, err := os.MkdirTemp("", "test_dir_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a read-only file (if possible on this OS)
	readOnlyFile, err := os.CreateTemp("", "test_readonly_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	readOnlyFilePath := readOnlyFile.Name()
	readOnlyFile.Close()
	defer os.Remove(readOnlyFilePath)
	os.Chmod(readOnlyFilePath, 0444)       // Read-only permissions
	defer os.Chmod(readOnlyFilePath, 0644) // Restore permissions for cleanup

	// Create a read-only directory to test non-writable parent
	readOnlyDir, err := os.MkdirTemp("", "test_readonly_dir_*")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.Chmod(readOnlyDir, 0755) // Restore permissions for cleanup
		os.RemoveAll(readOnlyDir)
	}()
	os.Chmod(readOnlyDir, 0555) // Read-only directory

	// Create a file inside a directory, then make the directory inaccessible
	noAccessDir, err := os.MkdirTemp("", "test_noaccess_dir_*")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.Chmod(noAccessDir, 0755) // Restore permissions for cleanup
		os.RemoveAll(noAccessDir)
	}()
	noAccessFile := filepath.Join(noAccessDir, "file.txt")
	os.WriteFile(noAccessFile, []byte("test"), 0644)
	os.Chmod(noAccessDir, 0000) // No permissions - cannot stat files inside

	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid writable paths
		{"existing writable file", tmpFile.Name(), false},
		{"new file in writable directory", filepath.Join(tmpDir, "newfile.txt"), false},

		// Invalid cases
		{"new file in non-existing directory", "/nonexistent/dir/file.txt", true},
		{"non-string value", 123, true},
		{"nil value", nil, true},
		{"directory path", tmpDir, true},                                                     // Directory instead of file
		{"read-only file", readOnlyFilePath, true},                                           // Read-only file
		{"new file in read-only directory", filepath.Join(readOnlyDir, "newfile.txt"), true}, // Cannot create file in read-only dir
		{"parent is not directory", filepath.Join(tmpFile.Name(), "subfile.txt"), true},      // Parent is a file, not a directory
		{"file in inaccessible directory", noAccessFile, true},                               // Cannot stat file in directory with no permissions
		{"file with null byte", "/tmp/file\x00test.txt", true},                               // Invalid file path with null byte
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "file_path")
			v.FileWritable()

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

func TestFileSize(t *testing.T) {
	// Create a temporary file with known size
	tmpFile, err := os.CreateTemp("", "test_size_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := "Hello, World!" // 13 bytes
	tmpFile.WriteString(content)
	tmpFile.Close()

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "test_dir_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name      string
		value     interface{}
		min       int64
		max       int64
		wantError bool
	}{
		// Valid size ranges
		{"exact range", tmpFile.Name(), 13, 13, false},
		{"within range", tmpFile.Name(), 10, 20, false},
		{"min boundary", tmpFile.Name(), 13, 100, false},
		{"max boundary", tmpFile.Name(), 0, 13, false},

		// Invalid size ranges
		{"too small", tmpFile.Name(), 20, 30, true},
		{"too large", tmpFile.Name(), 1, 10, true},
		{"non-existing file", "/nonexistent/file.txt", 0, 100, true},
		{"directory instead of file", tmpDir, 0, 100, true},
		{"non-string value", 123, 0, 100, true},
		{"nil value", nil, 0, 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "file_path")
			v.FileSize(tt.min, tt.max)

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

func TestFileExtension(t *testing.T) {
	tests := []struct {
		name       string
		value      interface{}
		extensions []string
		wantError  bool
	}{
		// Valid extensions
		{"valid .txt", "file.txt", []string{".txt"}, false},
		{"valid .txt without dot", "file.txt", []string{"txt"}, false},
		{"valid multiple extensions", "file.txt", []string{".md", ".txt", ".pdf"}, false},
		{"valid .log", "app.log", []string{".log", ".txt"}, false},
		{"valid nested path", "dir/subdir/file.txt", []string{".txt"}, false},

		// Invalid extensions
		{"invalid extension", "file.txt", []string{".md", ".pdf"}, true},
		{"no extension", "filename", []string{".txt"}, true},
		{"wrong extension", "file.doc", []string{".txt", ".pdf"}, true},
		{"non-string value", 123, []string{".txt"}, true},
		{"nil value", nil, []string{".txt"}, true},
		{"empty string", "", []string{".txt"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "file_path")
			v.FileExtension(tt.extensions...)

			if tt.wantError {
				if len(v.GetErrors()) == 0 {
					t.Errorf("expected error but got none")
				}
			} else {
				if len(v.GetErrors()) > 0 {
					t.Errorf("expected no error but got %v", v.GetErrors())
				}
			}
		})
	}
}

// Test StopOnError behavior for File validators
func TestFileValidatorsStopOnError(t *testing.T) {
	t.Run("FilePath StopOnError", func(t *testing.T) {
		v := New("", "file_path").StopOnError()
		v.FilePath()
		v.FilePath() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("FileExists StopOnError", func(t *testing.T) {
		v := New("/nonexistent/file.txt", "file_path").StopOnError()
		v.FileExists()
		v.FileExists() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("FileReadable StopOnError", func(t *testing.T) {
		v := New("/nonexistent/file.txt", "file_path").StopOnError()
		v.FileReadable()
		v.FileReadable() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("FileWritable StopOnError", func(t *testing.T) {
		v := New("/nonexistent/dir/file.txt", "file_path").StopOnError()
		v.FileWritable()
		v.FileWritable() // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("FileSize StopOnError", func(t *testing.T) {
		v := New("/nonexistent/file.txt", "file_path").StopOnError()
		v.FileSize(0, 100)
		v.FileSize(0, 100) // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})

	t.Run("FileExtension StopOnError", func(t *testing.T) {
		v := New("filename", "file_path").StopOnError()
		v.FileExtension(".txt")
		v.FileExtension(".txt") // Should not add another error
		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
		}
	})
}
