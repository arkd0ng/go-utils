package fileutil

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// Test file reading operations
// 파일 읽기 작업 테스트
func TestFileReading(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("ReadFile", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "read-test.txt")
		testData := []byte("Hello, World!")
		if err := WriteFile(testFile, testData); err != nil {
			t.Fatal(err)
		}

		data, err := ReadFile(testFile)
		if err != nil {
			t.Fatalf("ReadFile failed: %v", err)
		}

		if !bytes.Equal(data, testData) {
			t.Errorf("Expected %s, got %s", testData, data)
		}
	})

	t.Run("ReadString", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "string.txt")
		testString := "Hello, String!"
		if err := WriteString(testFile, testString); err != nil {
			t.Fatal(err)
		}

		content, err := ReadString(testFile)
		if err != nil {
			t.Fatalf("ReadString failed: %v", err)
		}

		if content != testString {
			t.Errorf("Expected %s, got %s", testString, content)
		}
	})

	t.Run("ReadLines", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "lines.txt")
		testLines := []string{"Line 1", "Line 2", "Line 3"}
		if err := WriteLines(testFile, testLines); err != nil {
			t.Fatal(err)
		}

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

	t.Run("ReadJSON", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "test.json")
		testData := map[string]interface{}{
			"name":  "test",
			"value": 123,
			"flag":  true,
		}

		if err := WriteJSON(testFile, testData); err != nil {
			t.Fatal(err)
		}

		var result map[string]interface{}
		if err := ReadJSON(testFile, &result); err != nil {
			t.Fatalf("ReadJSON failed: %v", err)
		}

		if result["name"] != "test" {
			t.Errorf("Expected name=test, got %v", result["name"])
		}
	})

	t.Run("ReadYAML", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "test.yaml")
		testData := map[string]interface{}{
			"name":  "test",
			"value": 123,
		}

		if err := WriteYAML(testFile, testData); err != nil {
			t.Fatal(err)
		}

		var result map[string]interface{}
		if err := ReadYAML(testFile, &result); err != nil {
			t.Fatalf("ReadYAML failed: %v", err)
		}

		if result["name"] != "test" {
			t.Errorf("Expected name=test, got %v", result["name"])
		}
	})

	t.Run("ReadCSV", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "test.csv")
		testData := [][]string{
			{"Name", "Age", "City"},
			{"Alice", "30", "Seoul"},
			{"Bob", "25", "Busan"},
		}

		if err := WriteCSV(testFile, testData); err != nil {
			t.Fatal(err)
		}

		records, err := ReadCSV(testFile)
		if err != nil {
			t.Fatalf("ReadCSV failed: %v", err)
		}

		if len(records) != len(testData) {
			t.Errorf("Expected %d records, got %d", len(testData), len(records))
		}
	})

	t.Run("ReadBytes", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "bytes.txt")
		testData := []byte("0123456789")
		if err := WriteFile(testFile, testData); err != nil {
			t.Fatal(err)
		}

		// Read bytes 2-5 (offset=2, length=3)
		data, err := ReadBytes(testFile, 2, 3)
		if err != nil {
			t.Fatalf("ReadBytes failed: %v", err)
		}

		expected := "234"
		if string(data) != expected {
			t.Errorf("Expected %s, got %s", expected, string(data))
		}
	})

	t.Run("ReadChunk", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "chunks.txt")
		testData := strings.Repeat("A", 1000)
		if err := WriteString(testFile, testData); err != nil {
			t.Fatal(err)
		}

		var chunks int
		err := ReadChunk(testFile, 100, func(chunk []byte) error {
			chunks++
			return nil
		})
		if err != nil {
			t.Fatalf("ReadChunk failed: %v", err)
		}

		if chunks != 10 {
			t.Errorf("Expected 10 chunks, got %d", chunks)
		}
	})
}

// Test file writing operations
// 파일 쓰기 작업 테스트
func TestFileWriting(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("WriteFile", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "deep", "nested", "write.txt")
		testData := []byte("test data")

		if err := WriteFile(testFile, testData); err != nil {
			t.Fatalf("WriteFile failed: %v", err)
		}

		if !Exists(testFile) {
			t.Error("File should exist")
		}
	})

	t.Run("WriteString", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "string-write.txt")
		testString := "Hello, World!"

		if err := WriteString(testFile, testString); err != nil {
			t.Fatalf("WriteString failed: %v", err)
		}

		content, _ := ReadString(testFile)
		if content != testString {
			t.Errorf("Expected %s, got %s", testString, content)
		}
	})

	t.Run("WriteLines", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "lines-write.txt")
		testLines := []string{"Line 1", "Line 2", "Line 3"}

		if err := WriteLines(testFile, testLines); err != nil {
			t.Fatalf("WriteLines failed: %v", err)
		}

		lines, _ := ReadLines(testFile)
		if len(lines) != len(testLines) {
			t.Errorf("Expected %d lines, got %d", len(testLines), len(lines))
		}
	})

	t.Run("WriteJSON", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "write.json")
		testData := map[string]string{"key": "value"}

		if err := WriteJSON(testFile, testData); err != nil {
			t.Fatalf("WriteJSON failed: %v", err)
		}

		var result map[string]string
		ReadJSON(testFile, &result)
		if result["key"] != "value" {
			t.Errorf("Expected value, got %v", result["key"])
		}
	})

	t.Run("WriteYAML", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "write.yaml")
		testData := map[string]string{"key": "value"}

		if err := WriteYAML(testFile, testData); err != nil {
			t.Fatalf("WriteYAML failed: %v", err)
		}

		var result map[string]string
		ReadYAML(testFile, &result)
		if result["key"] != "value" {
			t.Errorf("Expected value, got %v", result["key"])
		}
	})

	t.Run("WriteCSV", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "write.csv")
		testData := [][]string{
			{"Name", "Age"},
			{"Alice", "30"},
		}

		if err := WriteCSV(testFile, testData); err != nil {
			t.Fatalf("WriteCSV failed: %v", err)
		}

		records, _ := ReadCSV(testFile)
		if len(records) != 2 {
			t.Errorf("Expected 2 records, got %d", len(records))
		}
	})

	t.Run("WriteAtomic", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "atomic.txt")
		testData := []byte("atomic write")

		if err := WriteAtomic(testFile, testData); err != nil {
			t.Fatalf("WriteAtomic failed: %v", err)
		}

		content, _ := ReadFile(testFile)
		if !bytes.Equal(content, testData) {
			t.Errorf("Expected %s, got %s", testData, content)
		}
	})

	t.Run("AppendFile", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "append.txt")
		WriteString(testFile, "Line 1\n")

		if err := AppendString(testFile, "Line 2\n"); err != nil {
			t.Fatalf("AppendString failed: %v", err)
		}

		content, _ := ReadString(testFile)
		expected := "Line 1\nLine 2\n"
		if content != expected {
			t.Errorf("Expected %s, got %s", expected, content)
		}
	})

	t.Run("AppendLines", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "append-lines.txt")
		WriteLines(testFile, []string{"Line 1"})

		if err := AppendLines(testFile, []string{"Line 2", "Line 3"}); err != nil {
			t.Fatalf("AppendLines failed: %v", err)
		}

		lines, _ := ReadLines(testFile)
		if len(lines) != 3 {
			t.Errorf("Expected 3 lines, got %d", len(lines))
		}
	})
}

// Test path operations
// 경로 작업 테스트
func TestPathOperations(t *testing.T) {
	t.Run("Join", func(t *testing.T) {
		path := Join("home", "user", "file.txt")
		expected := filepath.Join("home", "user", "file.txt")
		if path != expected {
			t.Errorf("Expected %s, got %s", expected, path)
		}
	})

	t.Run("Split", func(t *testing.T) {
		_, file := Split("path/to/file.txt")
		if file != "file.txt" {
			t.Errorf("Expected file.txt, got %s", file)
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

	t.Run("Ext", func(t *testing.T) {
		ext := Ext("file.txt")
		if ext != ".txt" {
			t.Errorf("Expected .txt, got %s", ext)
		}

		ext = Ext("file")
		if ext != "" {
			t.Errorf("Expected empty string, got %s", ext)
		}
	})

	t.Run("Abs", func(t *testing.T) {
		abs, err := Abs(".")
		if err != nil {
			t.Fatalf("Abs failed: %v", err)
		}

		if abs == "" {
			t.Error("Absolute path should not be empty")
		}
	})

	t.Run("CleanPath", func(t *testing.T) {
		cleaned := CleanPath("path//to///file.txt")
		expected := filepath.Clean("path//to///file.txt")
		if cleaned != expected {
			t.Errorf("Expected %s, got %s", expected, cleaned)
		}
	})

	t.Run("Normalize", func(t *testing.T) {
		normalized, err := Normalize("./path/../file.txt")
		if err != nil {
			t.Fatalf("Normalize failed: %v", err)
		}

		if normalized == "" {
			t.Error("Normalized path should not be empty")
		}
	})

	t.Run("ToSlash", func(t *testing.T) {
		// ToSlash converts OS-specific path to forward slashes
		// On Unix systems, it may already be forward slashes
		path := ToSlash("path/to/file.txt")
		if !strings.Contains(path, "/") {
			t.Error("ToSlash should produce forward slashes")
		}
	})

	t.Run("FromSlash", func(t *testing.T) {
		path := FromSlash("path/to/file.txt")
		expected := filepath.FromSlash("path/to/file.txt")
		if path != expected {
			t.Errorf("Expected %s, got %s", expected, path)
		}
	})

	t.Run("IsAbs", func(t *testing.T) {
		if IsAbs("relative/path") {
			t.Error("Should not be absolute")
		}

		abs, _ := filepath.Abs(".")
		if !IsAbs(abs) {
			t.Error("Should be absolute")
		}
	})

	t.Run("IsValid", func(t *testing.T) {
		if !IsValid("path/to/file.txt") {
			t.Error("Should be valid")
		}

		if IsValid("") {
			t.Error("Empty path should be invalid")
		}
	})

	t.Run("IsSafe", func(t *testing.T) {
		root := "/home/user"
		if !IsSafe("/home/user/file.txt", root) {
			t.Error("Should be safe")
		}

		if IsSafe("/etc/passwd", root) {
			t.Error("Should not be safe (outside root)")
		}
	})

	t.Run("Match", func(t *testing.T) {
		matched, err := Match("*.txt", "file.txt")
		if err != nil {
			t.Fatalf("Match failed: %v", err)
		}

		if !matched {
			t.Error("Should match pattern")
		}
	})

	t.Run("WithoutExt", func(t *testing.T) {
		path := WithoutExt("file.txt")
		if path != "file" {
			t.Errorf("Expected file, got %s", path)
		}
	})

	t.Run("ChangeExt", func(t *testing.T) {
		path := ChangeExt("file.txt", ".md")
		if path != "file.md" {
			t.Errorf("Expected file.md, got %s", path)
		}
	})

	t.Run("HasExt", func(t *testing.T) {
		if !HasExt("file.txt", ".txt") {
			t.Error("Should have .txt extension")
		}

		if HasExt("file.txt", ".md") {
			t.Error("Should not have .md extension")
		}

		if !HasExt("file.txt", ".txt", ".md", ".go") {
			t.Error("Should match at least one extension")
		}
	})
}

// Test file information operations
// 파일 정보 작업 테스트
func TestFileInformation(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "info-test.txt")
	WriteString(testFile, "test content with some data")

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

	t.Run("IsSymlink", func(t *testing.T) {
		// Most temp files are not symlinks
		if IsSymlink(testFile) {
			t.Error("Regular file should not be a symlink")
		}
	})

	t.Run("Size", func(t *testing.T) {
		size, err := Size(testFile)
		if err != nil {
			t.Fatalf("Size failed: %v", err)
		}

		if size == 0 {
			t.Error("Size should not be zero")
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

		if !strings.Contains(sizeStr, "B") {
			t.Error("SizeHuman should contain 'B' for bytes")
		}
	})

	t.Run("Chmod", func(t *testing.T) {
		if err := Chmod(testFile, 0644); err != nil {
			t.Fatalf("Chmod failed: %v", err)
		}
	})

	t.Run("IsReadable", func(t *testing.T) {
		if !IsReadable(testFile) {
			t.Error("File should be readable")
		}
	})

	t.Run("IsWritable", func(t *testing.T) {
		if !IsWritable(testFile) {
			t.Error("File should be writable")
		}
	})

	t.Run("ModTime", func(t *testing.T) {
		modTime, err := ModTime(testFile)
		if err != nil {
			t.Fatalf("ModTime failed: %v", err)
		}

		if modTime.IsZero() {
			t.Error("ModTime should not be zero")
		}
	})

	t.Run("Touch", func(t *testing.T) {
		oldModTime, _ := ModTime(testFile)
		time.Sleep(10 * time.Millisecond)

		if err := Touch(testFile); err != nil {
			t.Fatalf("Touch failed: %v", err)
		}

		newModTime, _ := ModTime(testFile)
		if !newModTime.After(oldModTime) {
			t.Error("ModTime should be updated after Touch")
		}
	})
}

// Test copy operations
// 복사 작업 테스트
func TestCopyOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("CopyFile", func(t *testing.T) {
		srcFile := filepath.Join(tempDir, "source.txt")
		dstFile := filepath.Join(tempDir, "destination.txt")
		testContent := "test content for copy"

		WriteString(srcFile, testContent)

		if err := CopyFile(srcFile, dstFile); err != nil {
			t.Fatalf("CopyFile failed: %v", err)
		}

		content, _ := ReadString(dstFile)
		if content != testContent {
			t.Errorf("Expected %s, got %s", testContent, content)
		}
	})

	t.Run("CopyFile_WithOverwrite", func(t *testing.T) {
		srcFile := filepath.Join(tempDir, "src-overwrite.txt")
		dstFile := filepath.Join(tempDir, "dst-overwrite.txt")

		WriteString(srcFile, "new content")
		WriteString(dstFile, "old content")

		if err := CopyFile(srcFile, dstFile, WithOverwrite(true)); err != nil {
			t.Fatalf("CopyFile with overwrite failed: %v", err)
		}

		content, _ := ReadString(dstFile)
		if content != "new content" {
			t.Errorf("Expected new content, got %s", content)
		}
	})

	t.Run("CopyFile_WithProgress", func(t *testing.T) {
		srcFile := filepath.Join(tempDir, "src-progress.txt")
		dstFile := filepath.Join(tempDir, "dst-progress.txt")

		WriteString(srcFile, strings.Repeat("A", 1000))

		var progressCalled bool
		err := CopyFile(srcFile, dstFile, WithProgress(func(written, total int64) {
			progressCalled = true
		}))

		if err != nil {
			t.Fatalf("CopyFile with progress failed: %v", err)
		}

		if !progressCalled {
			t.Error("Progress callback should be called")
		}
	})

	t.Run("CopyDir", func(t *testing.T) {
		srcDir := filepath.Join(tempDir, "src-dir")
		dstDir := filepath.Join(tempDir, "dst-dir")

		MkdirAll(srcDir)
		WriteString(filepath.Join(srcDir, "file1.txt"), "content1")
		WriteString(filepath.Join(srcDir, "file2.txt"), "content2")

		if err := CopyDir(srcDir, dstDir); err != nil {
			t.Fatalf("CopyDir failed: %v", err)
		}

		if !Exists(filepath.Join(dstDir, "file1.txt")) {
			t.Error("file1.txt should exist in destination")
		}

		if !Exists(filepath.Join(dstDir, "file2.txt")) {
			t.Error("file2.txt should exist in destination")
		}
	})
}

// Test move operations
// 이동 작업 테스트
func TestMoveOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("MoveFile", func(t *testing.T) {
		srcFile := filepath.Join(tempDir, "move-src.txt")
		dstFile := filepath.Join(tempDir, "move-dst.txt")
		testContent := "move test"

		WriteString(srcFile, testContent)

		if err := MoveFile(srcFile, dstFile); err != nil {
			t.Fatalf("MoveFile failed: %v", err)
		}

		if Exists(srcFile) {
			t.Error("Source file should not exist after move")
		}

		if !Exists(dstFile) {
			t.Error("Destination file should exist")
		}

		content, _ := ReadString(dstFile)
		if content != testContent {
			t.Errorf("Expected %s, got %s", testContent, content)
		}
	})

	t.Run("Rename", func(t *testing.T) {
		oldPath := filepath.Join(tempDir, "old-name.txt")
		newPath := filepath.Join(tempDir, "new-name.txt")

		WriteString(oldPath, "rename test")

		if err := Rename(oldPath, newPath); err != nil {
			t.Fatalf("Rename failed: %v", err)
		}

		if Exists(oldPath) {
			t.Error("Old path should not exist")
		}

		if !Exists(newPath) {
			t.Error("New path should exist")
		}
	})

	t.Run("RenameExt", func(t *testing.T) {
		oldPath := filepath.Join(tempDir, "file.txt")
		WriteString(oldPath, "extension test")

		err := RenameExt(oldPath, ".md")
		if err != nil {
			t.Fatalf("RenameExt failed: %v", err)
		}

		newPath := WithoutExt(oldPath) + ".md"

		if Exists(oldPath) {
			t.Error("Old file should not exist")
		}

		if !Exists(newPath) {
			t.Error("New file should exist")
		}

		if Ext(newPath) != ".md" {
			t.Errorf("Expected .md extension, got %s", Ext(newPath))
		}
	})
}

// Test delete operations
// 삭제 작업 테스트
func TestDeleteOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("DeleteFile", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "delete-me.txt")
		WriteString(testFile, "delete me")

		if err := DeleteFile(testFile); err != nil {
			t.Fatalf("DeleteFile failed: %v", err)
		}

		if Exists(testFile) {
			t.Error("File should not exist after deletion")
		}
	})

	t.Run("DeleteDir", func(t *testing.T) {
		testDir := filepath.Join(tempDir, "empty-dir")
		MkdirAll(testDir)

		if err := DeleteDir(testDir); err != nil {
			t.Fatalf("DeleteDir failed: %v", err)
		}

		if Exists(testDir) {
			t.Error("Directory should not exist after deletion")
		}
	})

	t.Run("DeleteRecursive", func(t *testing.T) {
		testDir := filepath.Join(tempDir, "recursive-delete")
		MkdirAll(filepath.Join(testDir, "sub"))
		WriteString(filepath.Join(testDir, "file.txt"), "test")
		WriteString(filepath.Join(testDir, "sub", "file2.txt"), "test2")

		if err := DeleteRecursive(testDir); err != nil {
			t.Fatalf("DeleteRecursive failed: %v", err)
		}

		if Exists(testDir) {
			t.Error("Directory should not exist after recursive deletion")
		}
	})

	t.Run("DeletePattern", func(t *testing.T) {
		testDir := filepath.Join(tempDir, "pattern-delete")
		MkdirAll(testDir)
		WriteString(filepath.Join(testDir, "file1.txt"), "test")
		WriteString(filepath.Join(testDir, "file2.txt"), "test")
		WriteString(filepath.Join(testDir, "file.md"), "test")

		pattern := filepath.Join(testDir, "*.txt")
		if err := DeletePattern(pattern); err != nil {
			t.Fatalf("DeletePattern failed: %v", err)
		}

		if Exists(filepath.Join(testDir, "file1.txt")) {
			t.Error("file1.txt should be deleted")
		}

		if !Exists(filepath.Join(testDir, "file.md")) {
			t.Error("file.md should still exist")
		}
	})

	t.Run("DeleteFiles", func(t *testing.T) {
		file1 := filepath.Join(tempDir, "multi1.txt")
		file2 := filepath.Join(tempDir, "multi2.txt")
		WriteString(file1, "test")
		WriteString(file2, "test")

		if err := DeleteFiles(file1, file2); err != nil {
			t.Fatalf("DeleteFiles failed: %v", err)
		}

		if Exists(file1) || Exists(file2) {
			t.Error("Files should not exist after deletion")
		}
	})

	t.Run("Clean", func(t *testing.T) {
		testDir := filepath.Join(tempDir, "clean-dir")
		MkdirAll(testDir)
		WriteString(filepath.Join(testDir, "file.txt"), "test")

		if err := Clean(testDir); err != nil {
			t.Fatalf("Clean failed: %v", err)
		}

		if !Exists(testDir) {
			t.Error("Directory itself should still exist")
		}

		isEmpty, _ := IsEmpty(testDir)
		if !isEmpty {
			t.Error("Directory should be empty after Clean")
		}
	})

	t.Run("RemoveEmpty", func(t *testing.T) {
		testDir := filepath.Join(tempDir, "remove-empty")
		emptySubDir := filepath.Join(testDir, "empty")
		MkdirAll(emptySubDir)

		if err := RemoveEmpty(testDir); err != nil {
			t.Fatalf("RemoveEmpty failed: %v", err)
		}

		if Exists(emptySubDir) {
			t.Error("Empty subdirectory should be removed")
		}
	})
}

// Test directory operations
// 디렉토리 작업 테스트
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
			t.Error("Nested directory should exist")
		}
	})

	t.Run("CreateTemp", func(t *testing.T) {
		tempFilePath, err := CreateTemp(tempDir, "temp-*.txt")
		if err != nil {
			t.Fatalf("CreateTemp failed: %v", err)
		}

		if !Exists(tempFilePath) {
			t.Error("Temp file should exist")
		}
	})

	t.Run("CreateTempDir", func(t *testing.T) {
		tmpDir, err := CreateTempDir(tempDir, "temp-dir-*")
		if err != nil {
			t.Fatalf("CreateTempDir failed: %v", err)
		}

		if !IsDir(tmpDir) {
			t.Error("Temp directory should exist")
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		emptyDir := filepath.Join(tempDir, "empty")
		MkdirAll(emptyDir)

		empty, err := IsEmpty(emptyDir)
		if err != nil {
			t.Fatalf("IsEmpty failed: %v", err)
		}

		if !empty {
			t.Error("Directory should be empty")
		}

		WriteString(filepath.Join(emptyDir, "file.txt"), "test")
		empty, _ = IsEmpty(emptyDir)
		if empty {
			t.Error("Directory should not be empty")
		}
	})

	t.Run("DirSize", func(t *testing.T) {
		sizeDir := filepath.Join(tempDir, "size-test")
		MkdirAll(sizeDir)
		WriteString(filepath.Join(sizeDir, "file1.txt"), "test1")
		WriteString(filepath.Join(sizeDir, "file2.txt"), "test2")

		size, err := DirSize(sizeDir)
		if err != nil {
			t.Fatalf("DirSize failed: %v", err)
		}

		if size == 0 {
			t.Error("Directory size should not be zero")
		}
	})

	t.Run("ListFiles", func(t *testing.T) {
		listDir := filepath.Join(tempDir, "list-test")
		MkdirAll(listDir)
		WriteString(filepath.Join(listDir, "file1.txt"), "test")
		WriteString(filepath.Join(listDir, "file2.txt"), "test")
		MkdirAll(filepath.Join(listDir, "subdir"))

		files, err := ListFiles(listDir)
		if err != nil {
			t.Fatalf("ListFiles failed: %v", err)
		}

		if len(files) != 2 {
			t.Errorf("Expected 2 files, got %d", len(files))
		}
	})

	t.Run("ListDirs", func(t *testing.T) {
		listDir := filepath.Join(tempDir, "list-dirs")
		MkdirAll(filepath.Join(listDir, "dir1"))
		MkdirAll(filepath.Join(listDir, "dir2"))
		WriteString(filepath.Join(listDir, "file.txt"), "test")

		dirs, err := ListDirs(listDir)
		if err != nil {
			t.Fatalf("ListDirs failed: %v", err)
		}

		if len(dirs) != 2 {
			t.Errorf("Expected 2 directories, got %d", len(dirs))
		}
	})

	t.Run("ListAll", func(t *testing.T) {
		listDir := filepath.Join(tempDir, "list-all")
		MkdirAll(filepath.Join(listDir, "dir1"))
		WriteString(filepath.Join(listDir, "file1.txt"), "test")

		all, err := ListAll(listDir)
		if err != nil {
			t.Fatalf("ListAll failed: %v", err)
		}

		if len(all) != 2 {
			t.Errorf("Expected 2 entries, got %d", len(all))
		}
	})

	t.Run("FindFiles", func(t *testing.T) {
		findDir := filepath.Join(tempDir, "find-test")
		MkdirAll(filepath.Join(findDir, "sub"))
		WriteString(filepath.Join(findDir, "file.txt"), "test")
		WriteString(filepath.Join(findDir, "file.md"), "test")
		WriteString(filepath.Join(findDir, "sub", "file2.txt"), "test")

		txtFiles, err := FindFiles(findDir, func(path string, info os.FileInfo) bool {
			return Ext(path) == ".txt"
		})

		if err != nil {
			t.Fatalf("FindFiles failed: %v", err)
		}

		if len(txtFiles) != 2 {
			t.Errorf("Expected 2 .txt files, got %d", len(txtFiles))
		}
	})
}

// Test hash operations
// 해시 작업 테스트
func TestHashOperations(t *testing.T) {
	tempDir, err := CreateTempDir("", "fileutil-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "hash-test.txt")
	testContent := "test content for hashing"
	WriteString(testFile, testContent)

	t.Run("MD5", func(t *testing.T) {
		hash, err := MD5(testFile)
		if err != nil {
			t.Fatalf("MD5 failed: %v", err)
		}

		if len(hash) != 32 {
			t.Errorf("Expected 32 characters, got %d", len(hash))
		}
	})

	t.Run("SHA1", func(t *testing.T) {
		hash, err := SHA1(testFile)
		if err != nil {
			t.Fatalf("SHA1 failed: %v", err)
		}

		if len(hash) != 40 {
			t.Errorf("Expected 40 characters, got %d", len(hash))
		}
	})

	t.Run("SHA256", func(t *testing.T) {
		hash, err := SHA256(testFile)
		if err != nil {
			t.Fatalf("SHA256 failed: %v", err)
		}

		if len(hash) != 64 {
			t.Errorf("Expected 64 characters, got %d", len(hash))
		}
	})

	t.Run("SHA512", func(t *testing.T) {
		hash, err := SHA512(testFile)
		if err != nil {
			t.Fatalf("SHA512 failed: %v", err)
		}

		if len(hash) != 128 {
			t.Errorf("Expected 128 characters, got %d", len(hash))
		}
	})

	t.Run("Hash", func(t *testing.T) {
		hash, err := Hash(testFile, "md5")
		if err != nil {
			t.Fatalf("Hash failed: %v", err)
		}

		if len(hash) != 32 {
			t.Errorf("Expected 32 characters for MD5, got %d", len(hash))
		}
	})

	t.Run("HashBytes", func(t *testing.T) {
		// HashBytes는 파일 경로를 받아서 바이트를 반환합니다
		testFile2 := filepath.Join(tempDir, "hashbytes-test.txt")
		WriteString(testFile2, "test data for bytes")

		hashData, err := HashBytes(testFile2)
		if err != nil {
			t.Fatalf("HashBytes failed: %v", err)
		}

		if len(hashData) == 0 {
			t.Error("HashBytes should return non-empty data")
		}
	})

	t.Run("CompareFiles", func(t *testing.T) {
		file1 := filepath.Join(tempDir, "compare1.txt")
		file2 := filepath.Join(tempDir, "compare2.txt")
		WriteString(file1, testContent)
		WriteString(file2, testContent)

		same, err := CompareFiles(file1, file2)
		if err != nil {
			t.Fatalf("CompareFiles failed: %v", err)
		}

		if !same {
			t.Error("Files with same content should be equal")
		}

		WriteString(file2, "different content")
		same, _ = CompareFiles(file1, file2)
		if same {
			t.Error("Files with different content should not be equal")
		}
	})

	t.Run("CompareHash", func(t *testing.T) {
		file1 := filepath.Join(tempDir, "hash1.txt")
		file2 := filepath.Join(tempDir, "hash2.txt")
		WriteString(file1, testContent)
		WriteString(file2, testContent)

		same, err := CompareHash(file1, file2)
		if err != nil {
			t.Fatalf("CompareHash failed: %v", err)
		}

		if !same {
			t.Error("Files with same hash should be equal")
		}
	})

	t.Run("Checksum", func(t *testing.T) {
		checksum, err := Checksum(testFile)
		if err != nil {
			t.Fatalf("Checksum failed: %v", err)
		}

		if checksum == "" {
			t.Error("Checksum should not be empty")
		}
	})

	t.Run("VerifyChecksum", func(t *testing.T) {
		checksum, _ := Checksum(testFile)

		valid, err := VerifyChecksum(testFile, checksum)
		if err != nil {
			t.Fatalf("VerifyChecksum failed: %v", err)
		}

		if !valid {
			t.Error("Checksum should be valid")
		}

		valid, _ = VerifyChecksum(testFile, "invalid-checksum")
		if valid {
			t.Error("Invalid checksum should not be valid")
		}
	})
}

// Benchmarks
// 벤치마크
func BenchmarkWriteFile(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	data := []byte("benchmark test data")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testFile := filepath.Join(tempDir, "bench.txt")
		WriteFile(testFile, data)
	}
}

func BenchmarkReadFile(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "bench.txt")
	WriteString(testFile, "benchmark test data")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReadFile(testFile)
	}
}

func BenchmarkWriteString(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	testString := "benchmark test string"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testFile := filepath.Join(tempDir, "bench-string.txt")
		WriteString(testFile, testString)
	}
}

func BenchmarkReadString(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "bench-string.txt")
	WriteString(testFile, "benchmark test string")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReadString(testFile)
	}
}

func BenchmarkCopyFile(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	srcFile := filepath.Join(tempDir, "src.txt")
	WriteString(srcFile, strings.Repeat("A", 10000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dstFile := filepath.Join(tempDir, "dst.txt")
		CopyFile(srcFile, dstFile)
		DeleteFile(dstFile)
	}
}

func BenchmarkSHA256(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "hash.txt")
	WriteString(testFile, strings.Repeat("A", 10000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SHA256(testFile)
	}
}

func BenchmarkMD5(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "hash.txt")
	WriteString(testFile, strings.Repeat("A", 10000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MD5(testFile)
	}
}

func BenchmarkJSON(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "bench.json")
	testData := map[string]interface{}{
		"name":  "test",
		"value": 123,
		"items": []string{"a", "b", "c"},
	}

	b.Run("WriteJSON", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			WriteJSON(testFile, testData)
		}
	})

	b.Run("ReadJSON", func(b *testing.B) {
		WriteJSON(testFile, testData)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			ReadJSON(testFile, &result)
		}
	})
}

func BenchmarkYAML(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "bench.yaml")
	testData := map[string]interface{}{
		"name":  "test",
		"value": 123,
	}

	b.Run("WriteYAML", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			WriteYAML(testFile, testData)
		}
	})

	b.Run("ReadYAML", func(b *testing.B) {
		WriteYAML(testFile, testData)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var result map[string]interface{}
			ReadYAML(testFile, &result)
		}
	})
}

func BenchmarkListFiles(b *testing.B) {
	tempDir, _ := CreateTempDir("", "fileutil-bench-*")
	defer os.RemoveAll(tempDir)

	// Create 100 test files
	for i := 0; i < 100; i++ {
		WriteString(filepath.Join(tempDir, "file"+string(rune('0'+i))+".txt"), "test")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ListFiles(tempDir)
	}
}
