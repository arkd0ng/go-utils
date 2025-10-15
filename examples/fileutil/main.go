package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
)

func main() {
	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(fmt.Sprintf("logs/fileutil-example-%s.log", time.Now().Format("20060102-150405"))),
		logging.WithLevel(logging.INFO),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("=== fileutil Package Examples Started ===")

	// Create temp directory for examples / 예제를 위한 임시 디렉토리 생성
	tempDir, err := fileutil.CreateTempDir("", "fileutil-examples-*")
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	defer fileutil.DeleteRecursive(tempDir)

	logger.Info("Using temp directory", "path", tempDir)

	// Example 1: File Writing and Reading / 파일 쓰기 및 읽기
	example1FileOperations(logger, tempDir)

	// Example 2: Path Operations / 경로 작업
	example2PathOperations(logger)

	// Example 3: File Information / 파일 정보
	example3FileInfo(logger, tempDir)

	// Example 4: File Copying / 파일 복사
	example4FileCopy(logger, tempDir)

	// Example 5: File Hashing / 파일 해싱
	example5FileHash(logger, tempDir)

	// Example 6: Directory Operations / 디렉토리 작업
	example6DirectoryOperations(logger, tempDir)

	// Example 7: File Deletion / 파일 삭제
	example7FileDeletion(logger, tempDir)

	logger.Info("=== All Examples Completed Successfully! ===")
}

// Example 1: File Writing and Reading / 파일 쓰기 및 읽기
func example1FileOperations(logger *logging.Logger, tempDir string) {
	logger.Info("--- Example 1: File Writing and Reading ---")

	// Write string to file / 파일에 문자열 쓰기
	file1 := filepath.Join(tempDir, "example1", "hello.txt")
	if err := fileutil.WriteString(file1, "Hello, World!"); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Written to file", "path", file1)

	// Read string from file / 파일에서 문자열 읽기
	content, err := fileutil.ReadString(file1)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Read from file", "content", content)

	// Write lines to file / 파일에 줄 쓰기
	file2 := filepath.Join(tempDir, "example1", "lines.txt")
	lines := []string{"Line 1", "Line 2", "Line 3"}
	if err := fileutil.WriteLines(file2, lines); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Written lines to file", "count", len(lines), "path", file2)

	// Read lines from file / 파일에서 줄 읽기
	readLines, err := fileutil.ReadLines(file2)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Read lines from file", "count", len(readLines), "path", file2)

	// Append to file / 파일에 추가
	if err := fileutil.AppendString(file1, "\nAppended line"); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Appended to file", "path", file1)
}

// Example 2: Path Operations / 경로 작업
func example2PathOperations(logger *logging.Logger) {
	logger.Info("--- Example 2: Path Operations ---")

	// Join paths / 경로 결합
	path := fileutil.Join("home", "user", "documents", "file.txt")
	logger.Info("✓ Joined path", "value", path)

	// Get base name / 기본 이름 가져오기
	base := fileutil.Base(path)
	logger.Info("✓ Base name", "value", base)

	// Get directory / 디렉토리 가져오기
	dir := fileutil.Dir(path)
	logger.Info("✓ Directory", "value", dir)

	// Get extension / 확장자 가져오기
	ext := fileutil.Ext(path)
	logger.Info("✓ Extension", "value", ext)

	// Change extension / 확장자 변경
	newPath := fileutil.ChangeExt(path, ".md")
	logger.Info("✓ Changed extension", "value", newPath)

	// Check if has extension / 확장자 확인
	hasExt := fileutil.HasExt(path, ".txt", ".md")
	logger.Info("✓ Has extension .txt or .md", "value", hasExt)

}

// Example 3: File Information / 파일 정보
func example3FileInfo(logger *logging.Logger, tempDir string) {
	logger.Info("--- Example 3: File Information ---")

	// Create test file / 테스트 파일 생성
	testFile := filepath.Join(tempDir, "example3", "info.txt")
	if err := fileutil.WriteString(testFile, "File information test"); err != nil {
		logger.Fatalf("Error: %v", err)
	}

	// Check if file exists / 파일 존재 확인
	exists := fileutil.Exists(testFile)
	logger.Info("✓ File exists", "value", exists)

	// Check if it's a file / 파일인지 확인
	isFile := fileutil.IsFile(testFile)
	logger.Info("✓ Is file", "value", isFile)

	// Get file size / 파일 크기 가져오기
	size, err := fileutil.Size(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ File size", "bytes", size)

	// Get human-readable size / 사람이 읽기 쉬운 크기 가져오기
	sizeHuman, err := fileutil.SizeHuman(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ File size (human)", "value", sizeHuman)

	// Get modification time / 수정 시간 가져오기
	modTime, err := fileutil.ModTime(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Modified time", "value", modTime.Format("2006-01-02 15:04:05"))

	// Touch file (update modification time) / 파일 터치 (수정 시간 업데이트)
	if err := fileutil.Touch(testFile); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Touched file (updated modification time)")

}

// Example 4: File Copying / 파일 복사
func example4FileCopy(logger *logging.Logger, tempDir string) {
	logger.Info("--- Example 4: File Copying ---")

	// Create source file / 소스 파일 생성
	srcFile := filepath.Join(tempDir, "example4", "source.txt")
	if err := fileutil.WriteString(srcFile, "Content to copy"); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Created source file", "path", srcFile)

	// Copy file / 파일 복사
	dstFile := filepath.Join(tempDir, "example4", "destination.txt")
	if err := fileutil.CopyFile(srcFile, dstFile); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Copied to", "path", dstFile)

	// Verify copy / 복사 확인
	content, err := fileutil.ReadString(dstFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Verified content", "value", content)

	// Copy with progress callback / 진행 상황 콜백과 함께 복사
	dstFile2 := filepath.Join(tempDir, "example4", "destination2.txt")
	if err := fileutil.CopyFile(srcFile, dstFile2,
		fileutil.WithProgress(func(written, total int64) {
			percent := float64(written) / float64(total) * 100
			fmt.Printf("  Progress: %d/%d bytes (%.1f%%)\n", written, total, percent)
		})); err != nil {
		logger.Fatalf("Error: %v", err)
	}

	// Create directory to copy / 복사할 디렉토리 생성
	srcDir := filepath.Join(tempDir, "example4", "src-dir")
	fileutil.WriteString(filepath.Join(srcDir, "file1.txt"), "File 1")
	fileutil.WriteString(filepath.Join(srcDir, "file2.txt"), "File 2")
	logger.Info("✓ Created source directory", "path", srcDir)

	// Copy directory / 디렉토리 복사
	dstDir := filepath.Join(tempDir, "example4", "dst-dir")
	if err := fileutil.CopyDir(srcDir, dstDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Copied directory to", "path", dstDir)

}

// Example 5: File Hashing / 파일 해싱
func example5FileHash(logger *logging.Logger, tempDir string) {
	logger.Info("--- Example 5: File Hashing ---")

	// Create test file / 테스트 파일 생성
	testFile := filepath.Join(tempDir, "example5", "hash-test.txt")
	if err := fileutil.WriteString(testFile, "Content to hash"); err != nil {
		logger.Fatalf("Error: %v", err)
	}

	// Calculate MD5 hash / MD5 해시 계산
	md5Hash, err := fileutil.MD5(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ MD5 hash", "hash", md5Hash)

	// Calculate SHA1 hash / SHA1 해시 계산
	sha1Hash, err := fileutil.SHA1(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ SHA1 hash", "hash", sha1Hash)

	// Calculate SHA256 hash / SHA256 해시 계산
	sha256Hash, err := fileutil.SHA256(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ SHA256 hash", "hash", sha256Hash)

	// Calculate checksum / 체크섬 계산
	checksum, err := fileutil.Checksum(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Checksum", "checksum", checksum)

	// Verify checksum / 체크섬 검증
	valid, err := fileutil.VerifyChecksum(testFile, checksum)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Checksum valid", "valid", valid)

	// Compare files / 파일 비교
	file2 := filepath.Join(tempDir, "example5", "hash-test2.txt")
	fileutil.WriteString(file2, "Content to hash")
	same, err := fileutil.CompareHash(testFile, file2)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Files have same hash", "same", same)

}

// Example 6: Directory Operations / 디렉토리 작업
func example6DirectoryOperations(logger *logging.Logger, tempDir string) {
	logger.Info("--- Example 6: Directory Operations ---")

	// Create nested directory / 중첩 디렉토리 생성
	nestedDir := filepath.Join(tempDir, "example6", "a", "b", "c")
	if err := fileutil.MkdirAll(nestedDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Created nested directory", "path", nestedDir)

	// Create test files / 테스트 파일 생성
	testDir := filepath.Join(tempDir, "example6", "test-dir")
	fileutil.WriteString(filepath.Join(testDir, "file1.txt"), "File 1")
	fileutil.WriteString(filepath.Join(testDir, "file2.go"), "File 2")
	fileutil.WriteString(filepath.Join(testDir, "file3.txt"), "File 3")
	logger.Info("✓ Created test directory", "path", testDir)

	// List files / 파일 나열
	files, err := fileutil.ListFiles(testDir)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Found files", "count", len(files))
	for _, file := range files {
		logger.Info("  - File", "name", filepath.Base(file))
	}

	// Find .txt files / .txt 파일 찾기
	txtFiles, err := fileutil.FindFiles(testDir, func(path string, info os.FileInfo) bool {
		return fileutil.Ext(path) == ".txt"
	})
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Found .txt files", "count", len(txtFiles))

	// Check if directory is empty / 디렉토리가 비어 있는지 확인
	emptyDir := filepath.Join(tempDir, "example6", "empty")
	fileutil.MkdirAll(emptyDir)
	isEmpty, err := fileutil.IsEmpty(emptyDir)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Empty directory is empty", "isEmpty", isEmpty)

	// Calculate directory size / 디렉토리 크기 계산
	dirSize, err := fileutil.DirSize(testDir)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Directory size", "bytes", dirSize)

}

// Example 7: File Deletion / 파일 삭제
func example7FileDeletion(logger *logging.Logger, tempDir string) {
	logger.Info("--- Example 7: File Deletion ---")

	// Create test files / 테스트 파일 생성
	file1 := filepath.Join(tempDir, "example7", "delete-me.txt")
	fileutil.WriteString(file1, "Delete me")
	logger.Info("✓ Created file", "path", file1)

	// Delete file / 파일 삭제
	if err := fileutil.DeleteFile(file1); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Deleted file")

	// Verify deletion / 삭제 확인
	exists := fileutil.Exists(file1)
	logger.Info("✓ File exists after deletion", "exists", exists)

	// Create directory to delete / 삭제할 디렉토리 생성
	deleteDir := filepath.Join(tempDir, "example7", "delete-dir")
	fileutil.WriteString(filepath.Join(deleteDir, "file.txt"), "File in directory")
	logger.Info("✓ Created directory", "path", deleteDir)

	// Delete directory recursively / 디렉토리 재귀적으로 삭제
	if err := fileutil.DeleteRecursive(deleteDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Deleted directory recursively")

	// Clean directory (remove all contents but keep directory) / 디렉토리 정리 (모든 내용 제거하지만 디렉토리는 유지)
	cleanDir := filepath.Join(tempDir, "example7", "clean-dir")
	fileutil.WriteString(filepath.Join(cleanDir, "file1.txt"), "File 1")
	fileutil.WriteString(filepath.Join(cleanDir, "file2.txt"), "File 2")
	logger.Info("✓ Created directory with files", "path", cleanDir)

	if err := fileutil.Clean(cleanDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	logger.Info("✓ Cleaned directory (removed all contents)")

	// Verify directory still exists but is empty / 디렉토리가 여전히 존재하지만 비어 있는지 확인
	stillExists := fileutil.Exists(cleanDir)
	isEmpty, _ := fileutil.IsEmpty(cleanDir)
	logger.Info("✓ Directory status", "exists", stillExists, "isEmpty", isEmpty)

}
