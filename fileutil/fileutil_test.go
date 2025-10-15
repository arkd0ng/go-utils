package fileutil

import (
	"os"
	"path/filepath"
	"testing"
)

// Test file operations / 파일 작업 테스트
func TestFileOperations(t *testing.T) {
	// Create temp directory for tests / 테스트용 임시 디렉토리 생성
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Test WriteFile and ReadFile / WriteFile 및 ReadFile 테스트
	t.Run("WriteFile_ReadFile", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "test", "file.txt")
		testData := []byte("Hello, World!")

		// Write file / 파일 쓰기
		if err := WriteFile(testFile, testData); err != nil {
			t.Fatalf("WriteFile failed: %v", err)
		}

		// Read file / 파일 읽기
		data, err := ReadFile(testFile)
		if err != nil {
			t.Fatalf("ReadFile failed: %v", err)
		}

		if string(data) != string(testData) {
			t.Errorf("Expected %s, got %s", testData, data)
		}
	})

	// Test WriteString and ReadString / WriteString 및 ReadString 테스트
	t.Run("WriteString_ReadString", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "string.txt")
		testString := "Hello, String!"

		// Write string / 문자열 쓰기
		if err := WriteString(testFile, testString); err != nil {
			t.Fatalf("WriteString failed: %v", err)
		}

		// Read string / 문자열 읽기
		content, err := ReadString(testFile)
		if err != nil {
			t.Fatalf("ReadString failed: %v", err)
		}

		if content != testString {
			t.Errorf("Expected %s, got %s", testString, content)
		}
	})

	// Test WriteLines and ReadLines / WriteLines 및 ReadLines 테스트
	t.Run("WriteLines_ReadLines", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "lines.txt")
		testLines := []string{"Line 1", "Line 2", "Line 3"}

		// Write lines / 줄 쓰기
		if err := WriteLines(testFile, testLines); err != nil {
			t.Fatalf("WriteLines failed: %v", err)
		}

		// Read lines / 줄 읽기
		lines, err := ReadLines(testFile)
		if err != nil {
			t.Fatalf("ReadLines failed: %v", err)
		}

		if len(lines) != len(testLines) {
			t.Errorf("Expected %d lines, got %d", len(testLines), len(lines))
		}

		for i, line := range lines {
			if line != testLines[i] {
				t.Errorf("Line %d: expected %s, got %s", i, testLines[i], line)
			}
		}
	})
}

// Test path operations / 경로 작업 테스트
func TestPathOperations(t *testing.T) {
	t.Run("Join", func(t *testing.T) {
		path := Join("home", "user", "file.txt")
		expected := filepath.Join("home", "user", "file.txt")
		if path != expected {
			t.Errorf("Expected %s, got %s", expected, path)
		}
	})

	t.Run("Ext", func(t *testing.T) {
		ext := Ext("file.txt")
		if ext != ".txt" {
			t.Errorf("Expected .txt, got %s", ext)
		}
	})

	t.Run("Base", func(t *testing.T) {
		base := Base("path/to/file.txt")
		if base != "file.txt" {
			t.Errorf("Expected file.txt, got %s", base)
		}
	})

	t.Run("Dir", func(t *testing.T) {
		dir := Dir("path/to/file.txt")
		expected := filepath.Join("path", "to")
		if dir != expected {
			t.Errorf("Expected %s, got %s", expected, dir)
		}
	})
}

// Test file information / 파일 정보 테스트
func TestFileInfo(t *testing.T) {
	// Create temp file / 임시 파일 생성
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "test.txt")
	if err := WriteString(testFile, "test content"); err != nil {
		t.Fatal(err)
	}

	t.Run("Exists", func(t *testing.T) {
		if !Exists(testFile) {
			t.Error("File should exist")
		}

		if Exists(filepath.Join(tempDir, "nonexistent.txt")) {
			t.Error("File should not exist")
		}
	})

	t.Run("IsFile", func(t *testing.T) {
		if !IsFile(testFile) {
			t.Error("Should be a file")
		}

		if IsFile(tempDir) {
			t.Error("Directory should not be a file")
		}
	})

	t.Run("IsDir", func(t *testing.T) {
		if !IsDir(tempDir) {
			t.Error("Should be a directory")
		}

		if IsDir(testFile) {
			t.Error("File should not be a directory")
		}
	})

	t.Run("Size", func(t *testing.T) {
		size, err := Size(testFile)
		if err != nil {
			t.Fatalf("Size failed: %v", err)
		}

		if size != int64(len("test content")) {
			t.Errorf("Expected size %d, got %d", len("test content"), size)
		}
	})

	t.Run("SizeHuman", func(t *testing.T) {
		sizeStr, err := SizeHuman(testFile)
		if err != nil {
			t.Fatalf("SizeHuman failed: %v", err)
		}

		if sizeStr == "" {
			t.Error("SizeHuman should return non-empty string")
		}
	})
}

// Test copy operations / 복사 작업 테스트
func TestCopyOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("CopyFile", func(t *testing.T) {
		srcFile := filepath.Join(tempDir, "source.txt")
		dstFile := filepath.Join(tempDir, "destination.txt")

		// Create source file / 소스 파일 생성
		if err := WriteString(srcFile, "test content"); err != nil {
			t.Fatal(err)
		}

		// Copy file / 파일 복사
		if err := CopyFile(srcFile, dstFile); err != nil {
			t.Fatalf("CopyFile failed: %v", err)
		}

		// Verify copy / 복사 확인
		content, err := ReadString(dstFile)
		if err != nil {
			t.Fatalf("ReadString failed: %v", err)
		}

		if content != "test content" {
			t.Errorf("Expected 'test content', got %s", content)
		}
	})
}

// Test delete operations / 삭제 작업 테스트
func TestDeleteOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("DeleteFile", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "delete-me.txt")

		// Create file / 파일 생성
		if err := WriteString(testFile, "delete me"); err != nil {
			t.Fatal(err)
		}

		// Delete file / 파일 삭제
		if err := DeleteFile(testFile); err != nil {
			t.Fatalf("DeleteFile failed: %v", err)
		}

		// Verify deletion / 삭제 확인
		if Exists(testFile) {
			t.Error("File should not exist after deletion")
		}
	})
}

// Test hash operations / 해시 작업 테스트
func TestHashOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "hash-test.txt")
	if err := WriteString(testFile, "test content"); err != nil {
		t.Fatal(err)
	}

	t.Run("SHA256", func(t *testing.T) {
		hash, err := SHA256(testFile)
		if err != nil {
			t.Fatalf("SHA256 failed: %v", err)
		}

		if hash == "" {
			t.Error("Hash should not be empty")
		}

		if len(hash) != 64 { // SHA256 produces 64 hex characters
			t.Errorf("Expected 64 characters, got %d", len(hash))
		}
	})

	t.Run("MD5", func(t *testing.T) {
		hash, err := MD5(testFile)
		if err != nil {
			t.Fatalf("MD5 failed: %v", err)
		}

		if hash == "" {
			t.Error("Hash should not be empty")
		}

		if len(hash) != 32 { // MD5 produces 32 hex characters
			t.Errorf("Expected 32 characters, got %d", len(hash))
		}
	})
}

// Test directory operations / 디렉토리 작업 테스트
func TestDirectoryOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("MkdirAll", func(t *testing.T) {
		nestedDir := filepath.Join(tempDir, "a", "b", "c")

		if err := MkdirAll(nestedDir); err != nil {
			t.Fatalf("MkdirAll failed: %v", err)
		}

		if !IsDir(nestedDir) {
			t.Error("Directory should exist")
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		emptyDir := filepath.Join(tempDir, "empty")
		if err := MkdirAll(emptyDir); err != nil {
			t.Fatal(err)
		}

		empty, err := IsEmpty(emptyDir)
		if err != nil {
			t.Fatalf("IsEmpty failed: %v", err)
		}

		if !empty {
			t.Error("Directory should be empty")
		}

		// Add a file / 파일 추가
		testFile := filepath.Join(emptyDir, "test.txt")
		if err := WriteString(testFile, "test"); err != nil {
			t.Fatal(err)
		}

		empty, err = IsEmpty(emptyDir)
		if err != nil {
			t.Fatalf("IsEmpty failed: %v", err)
		}

		if empty {
			t.Error("Directory should not be empty")
		}
	})

	t.Run("ListFiles", func(t *testing.T) {
		testDir := filepath.Join(tempDir, "list-test")
		if err := MkdirAll(testDir); err != nil {
			t.Fatal(err)
		}

		// Create test files / 테스트 파일 생성
		for i := 1; i <= 3; i++ {
			file := filepath.Join(testDir, "file"+string(rune('0'+i))+".txt")
			if err := WriteString(file, "test"); err != nil {
				t.Fatal(err)
			}
		}

		files, err := ListFiles(testDir)
		if err != nil {
			t.Fatalf("ListFiles failed: %v", err)
		}

		if len(files) != 3 {
			t.Errorf("Expected 3 files, got %d", len(files))
		}
	})
}

// Benchmark file operations / 파일 작업 벤치마크
func BenchmarkWriteFile(b *testing.B) {
	tempDir, err := CreateTempDir("", "fileutil-bench-*")
	if err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	data := []byte("benchmark test data")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testFile := filepath.Join(tempDir, "bench.txt")
		if err := WriteFile(testFile, data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadFile(b *testing.B) {
	tempDir, err := CreateTempDir("", "fileutil-bench-*")
	if err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "bench.txt")
	if err := WriteString(testFile, "benchmark test data"); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := ReadFile(testFile); err != nil {
			b.Fatal(err)
		}
	}
}
