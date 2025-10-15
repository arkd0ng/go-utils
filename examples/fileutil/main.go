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
		logging.WithFilePath(fmt.Sprintf("logs/fileutil-examples-%s.log", time.Now().Format("20060102-150405"))),
		logging.WithLevel(logging.INFO),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("=== fileutil Package Examples Started ===")
	fmt.Println("=== fileutil Package Examples ===")
	fmt.Println()

	// Create temp directory for examples / 예제를 위한 임시 디렉토리 생성
	tempDir, err := fileutil.CreateTempDir("", "fileutil-examples-*")
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	defer fileutil.DeleteRecursive(tempDir)

	fmt.Printf("Using temp directory: %s\n\n", tempDir)

	// Example 1: File Writing and Reading / 파일 쓰기 및 읽기
	example1FileOperations(logger, tempDir)

	// Example 2: Path Operations / 경로 작업
	example2PathOperations()

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

	fmt.Println("\n=== All Examples Completed Successfully! ===")
}

// Example 1: File Writing and Reading / 파일 쓰기 및 읽기
func example1FileOperations(logger *logging.Logger, tempDir string) {
	fmt.Println("--- Example 1: File Writing and Reading ---")

	// Write string to file / 파일에 문자열 쓰기
	file1 := filepath.Join(tempDir, "example1", "hello.txt")
	if err := fileutil.WriteString(file1, "Hello, World!"); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Written to %s\n", file1)

	// Read string from file / 파일에서 문자열 읽기
	content, err := fileutil.ReadString(file1)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Read: %s\n", content)

	// Write lines to file / 파일에 줄 쓰기
	file2 := filepath.Join(tempDir, "example1", "lines.txt")
	lines := []string{"Line 1", "Line 2", "Line 3"}
	if err := fileutil.WriteLines(file2, lines); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Written %d lines to %s\n", len(lines), file2)

	// Read lines from file / 파일에서 줄 읽기
	readLines, err := fileutil.ReadLines(file2)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Read %d lines from %s\n", len(readLines), file2)

	// Append to file / 파일에 추가
	if err := fileutil.AppendString(file1, "\nAppended line"); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Appended to %s\n", file1)

	fmt.Println()
}

// Example 2: Path Operations / 경로 작업
func example2PathOperations() {
	fmt.Println("--- Example 2: Path Operations ---")

	// Join paths / 경로 결합
	path := fileutil.Join("home", "user", "documents", "file.txt")
	fmt.Printf("✓ Joined path: %s\n", path)

	// Get base name / 기본 이름 가져오기
	base := fileutil.Base(path)
	fmt.Printf("✓ Base name: %s\n", base)

	// Get directory / 디렉토리 가져오기
	dir := fileutil.Dir(path)
	fmt.Printf("✓ Directory: %s\n", dir)

	// Get extension / 확장자 가져오기
	ext := fileutil.Ext(path)
	fmt.Printf("✓ Extension: %s\n", ext)

	// Change extension / 확장자 변경
	newPath := fileutil.ChangeExt(path, ".md")
	fmt.Printf("✓ Changed extension: %s\n", newPath)

	// Check if has extension / 확장자 확인
	hasExt := fileutil.HasExt(path, ".txt", ".md")
	fmt.Printf("✓ Has extension .txt or .md: %t\n", hasExt)

	fmt.Println()
}

// Example 3: File Information / 파일 정보
func example3FileInfo(logger *logging.Logger, tempDir string) {
	fmt.Println("--- Example 3: File Information ---")

	// Create test file / 테스트 파일 생성
	testFile := filepath.Join(tempDir, "example3", "info.txt")
	if err := fileutil.WriteString(testFile, "File information test"); err != nil {
		logger.Fatalf("Error: %v", err)
	}

	// Check if file exists / 파일 존재 확인
	exists := fileutil.Exists(testFile)
	fmt.Printf("✓ File exists: %t\n", exists)

	// Check if it's a file / 파일인지 확인
	isFile := fileutil.IsFile(testFile)
	fmt.Printf("✓ Is file: %t\n", isFile)

	// Get file size / 파일 크기 가져오기
	size, err := fileutil.Size(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ File size: %d bytes\n", size)

	// Get human-readable size / 사람이 읽기 쉬운 크기 가져오기
	sizeHuman, err := fileutil.SizeHuman(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ File size (human): %s\n", sizeHuman)

	// Get modification time / 수정 시간 가져오기
	modTime, err := fileutil.ModTime(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Modified time: %s\n", modTime.Format("2006-01-02 15:04:05"))

	// Touch file (update modification time) / 파일 터치 (수정 시간 업데이트)
	if err := fileutil.Touch(testFile); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Touched file (updated modification time)\n")

	fmt.Println()
}

// Example 4: File Copying / 파일 복사
func example4FileCopy(logger *logging.Logger, tempDir string) {
	fmt.Println("--- Example 4: File Copying ---")

	// Create source file / 소스 파일 생성
	srcFile := filepath.Join(tempDir, "example4", "source.txt")
	if err := fileutil.WriteString(srcFile, "Content to copy"); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Created source file: %s\n", srcFile)

	// Copy file / 파일 복사
	dstFile := filepath.Join(tempDir, "example4", "destination.txt")
	if err := fileutil.CopyFile(srcFile, dstFile); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Copied to: %s\n", dstFile)

	// Verify copy / 복사 확인
	content, err := fileutil.ReadString(dstFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Verified content: %s\n", content)

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
	fmt.Printf("✓ Created source directory: %s\n", srcDir)

	// Copy directory / 디렉토리 복사
	dstDir := filepath.Join(tempDir, "example4", "dst-dir")
	if err := fileutil.CopyDir(srcDir, dstDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Copied directory to: %s\n", dstDir)

	fmt.Println()
}

// Example 5: File Hashing / 파일 해싱
func example5FileHash(logger *logging.Logger, tempDir string) {
	fmt.Println("--- Example 5: File Hashing ---")

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
	fmt.Printf("✓ MD5:    %s\n", md5Hash)

	// Calculate SHA1 hash / SHA1 해시 계산
	sha1Hash, err := fileutil.SHA1(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ SHA1:   %s\n", sha1Hash)

	// Calculate SHA256 hash / SHA256 해시 계산
	sha256Hash, err := fileutil.SHA256(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ SHA256: %s\n", sha256Hash)

	// Calculate checksum / 체크섬 계산
	checksum, err := fileutil.Checksum(testFile)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Checksum: %s\n", checksum)

	// Verify checksum / 체크섬 검증
	valid, err := fileutil.VerifyChecksum(testFile, checksum)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Checksum valid: %t\n", valid)

	// Compare files / 파일 비교
	file2 := filepath.Join(tempDir, "example5", "hash-test2.txt")
	fileutil.WriteString(file2, "Content to hash")
	same, err := fileutil.CompareHash(testFile, file2)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Files have same hash: %t\n", same)

	fmt.Println()
}

// Example 6: Directory Operations / 디렉토리 작업
func example6DirectoryOperations(logger *logging.Logger, tempDir string) {
	fmt.Println("--- Example 6: Directory Operations ---")

	// Create nested directory / 중첩 디렉토리 생성
	nestedDir := filepath.Join(tempDir, "example6", "a", "b", "c")
	if err := fileutil.MkdirAll(nestedDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Created nested directory: %s\n", nestedDir)

	// Create test files / 테스트 파일 생성
	testDir := filepath.Join(tempDir, "example6", "test-dir")
	fileutil.WriteString(filepath.Join(testDir, "file1.txt"), "File 1")
	fileutil.WriteString(filepath.Join(testDir, "file2.go"), "File 2")
	fileutil.WriteString(filepath.Join(testDir, "file3.txt"), "File 3")
	fmt.Printf("✓ Created test directory: %s\n", testDir)

	// List files / 파일 나열
	files, err := fileutil.ListFiles(testDir)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Found %d files:\n", len(files))
	for _, file := range files {
		fmt.Printf("  - %s\n", filepath.Base(file))
	}

	// Find .txt files / .txt 파일 찾기
	txtFiles, err := fileutil.FindFiles(testDir, func(path string, info os.FileInfo) bool {
		return fileutil.Ext(path) == ".txt"
	})
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Found %d .txt files\n", len(txtFiles))

	// Check if directory is empty / 디렉토리가 비어 있는지 확인
	emptyDir := filepath.Join(tempDir, "example6", "empty")
	fileutil.MkdirAll(emptyDir)
	isEmpty, err := fileutil.IsEmpty(emptyDir)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Empty directory is empty: %t\n", isEmpty)

	// Calculate directory size / 디렉토리 크기 계산
	dirSize, err := fileutil.DirSize(testDir)
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Directory size: %d bytes\n", dirSize)

	fmt.Println()
}

// Example 7: File Deletion / 파일 삭제
func example7FileDeletion(logger *logging.Logger, tempDir string) {
	fmt.Println("--- Example 7: File Deletion ---")

	// Create test files / 테스트 파일 생성
	file1 := filepath.Join(tempDir, "example7", "delete-me.txt")
	fileutil.WriteString(file1, "Delete me")
	fmt.Printf("✓ Created file: %s\n", file1)

	// Delete file / 파일 삭제
	if err := fileutil.DeleteFile(file1); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Deleted file\n")

	// Verify deletion / 삭제 확인
	exists := fileutil.Exists(file1)
	fmt.Printf("✓ File exists after deletion: %t\n", exists)

	// Create directory to delete / 삭제할 디렉토리 생성
	deleteDir := filepath.Join(tempDir, "example7", "delete-dir")
	fileutil.WriteString(filepath.Join(deleteDir, "file.txt"), "File in directory")
	fmt.Printf("✓ Created directory: %s\n", deleteDir)

	// Delete directory recursively / 디렉토리 재귀적으로 삭제
	if err := fileutil.DeleteRecursive(deleteDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Deleted directory recursively\n")

	// Clean directory (remove all contents but keep directory) / 디렉토리 정리 (모든 내용 제거하지만 디렉토리는 유지)
	cleanDir := filepath.Join(tempDir, "example7", "clean-dir")
	fileutil.WriteString(filepath.Join(cleanDir, "file1.txt"), "File 1")
	fileutil.WriteString(filepath.Join(cleanDir, "file2.txt"), "File 2")
	fmt.Printf("✓ Created directory with files: %s\n", cleanDir)

	if err := fileutil.Clean(cleanDir); err != nil {
		logger.Fatalf("Error: %v", err)
	}
	fmt.Printf("✓ Cleaned directory (removed all contents)\n")

	// Verify directory still exists but is empty / 디렉토리가 여전히 존재하지만 비어 있는지 확인
	stillExists := fileutil.Exists(cleanDir)
	isEmpty, _ := fileutil.IsEmpty(cleanDir)
	fmt.Printf("✓ Directory exists: %t, is empty: %t\n", stillExists, isEmpty)

	fmt.Println()
}
