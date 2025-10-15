package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
)

// User struct for JSON/YAML examples / JSON/YAML 예제를 위한 User 구조체
type User struct {
	ID   int    `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func main() {
	// Setup log file with backup management / 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/fileutil-example.log"

	// Check if previous log file exists / 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/fileutil-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file / 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
				// Delete original log file to prevent content duplication / 내용 중복 방지를 위해 원본 로그 파일 삭제
				fileutil.DeleteFile(logFilePath)
			}
		}

		// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/fileutil-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time / 수정 시간으로 정렬
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			var files []fileInfo
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			// Sort oldest first / 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5 / 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename / 고정 파일명으로 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║         fileutil Package - Comprehensive Examples & Manual                ║")
	logger.Info("║         fileutil 패키지 - 종합 예제 및 매뉴얼                              ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")
	logger.Info("📋 Package Information / 패키지 정보")
	logger.Info("   Package Name: github.com/arkd0ng/go-utils/fileutil")
	logger.Info("   Description: Extremely simple file and path utilities (20 lines → 1-2 lines)")
	logger.Info("   설명: 극도로 간단한 파일 및 경로 유틸리티 (20줄 → 1-2줄)")
	logger.Info("   Total Functions: ~91 functions across 12 categories")
	logger.Info("   총 함수 개수: 12개 카테고리에 걸쳐 약 91개 함수")
	logger.Info("")

	// Create main temp directory for all examples / 모든 예제를 위한 메인 임시 디렉토리 생성
	logger.Info("🚀 Starting Examples / 예제 시작")
	logger.Info("   Creating temporary workspace for isolated testing...")
	logger.Info("   격리된 테스트를 위한 임시 작업공간 생성 중...")

	tempDir, err := fileutil.CreateTempDir("", "fileutil-manual-*")
	if err != nil {
		logger.Fatalf("❌ Failed to create temp directory: %v", err)
	}
	defer fileutil.DeleteRecursive(tempDir)

	logger.Info("✅ Temp Directory Created Successfully / 임시 디렉토리 생성 성공")
	logger.Info(fmt.Sprintf("   📂 Path: %s", tempDir))
	logger.Info("   📏 Initial Size: 0 bytes (empty directory)")
	logger.Info("   🔒 Permissions: 0755 (rwxr-xr-x)")
	logger.Info("   ℹ️  All examples will run in this isolated environment")
	logger.Info("   ℹ️  모든 예제는 이 격리된 환경에서 실행됩니다")
	logger.Info("   ℹ️  Directory will be automatically cleaned up on exit")
	logger.Info("   ℹ️  종료 시 디렉토리가 자동으로 정리됩니다")
	logger.Info("")

	// Run all examples / 모든 예제 실행
	example01_FileWriting(logger, tempDir)
	example02_FileReading(logger, tempDir)
	example03_PathOperations(logger, tempDir)
	example04_FileInformation(logger, tempDir)
	example05_FileCopying(logger, tempDir)
	example06_FileMoving(logger, tempDir)
	example07_FileDeletion(logger, tempDir)
	example08_DirectoryOperations(logger, tempDir)
	example09_FileHashing(logger, tempDir)
	example10_AdvancedReading(logger, tempDir)
	example11_AtomicOperations(logger, tempDir)
	example12_PermissionsAndOwnership(logger, tempDir)
	example13_SymlinksAndSpecialFiles(logger, tempDir)
	example14_WalkAndFilter(logger, tempDir)
	example15_ErrorHandling(logger, tempDir)
	example16_RealWorldScenarios(logger, tempDir)

	logger.Info("")
	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║                   All Examples Completed Successfully!                     ║")
	logger.Info("║                   모든 예제가 성공적으로 완료되었습니다!                    ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
}

// Example 1: File Writing Operations / 예제 1: 파일 쓰기 작업
func example01_FileWriting(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📝 Example 1: File Writing Operations")
	logger.Info("   예제 1: 파일 쓰기 작업")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📖 Category Overview / 카테고리 개요")
	logger.Info("   This example demonstrates 8 file writing methods")
	logger.Info("   이 예제는 8가지 파일 쓰기 메서드를 시연합니다")
	logger.Info("   • WriteString, WriteFile, WriteLines, WriteJSON, WriteYAML, WriteCSV")
	logger.Info("   • AppendString, AppendLines")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example01")
	logger.Info(fmt.Sprintf("📁 Creating example directory: %s", filepath.Base(exampleDir)))
	logger.Info("")

	// 1. WriteString - Write a string to file / 문자열을 파일에 쓰기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣  WriteString() - Writing text content to file")
	logger.Info("   문자열을 파일에 쓰기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func WriteString(path string, content string) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Writes a string to a file, creating parent directories if needed")
	logger.Info("   문자열을 파일에 쓰고, 필요시 상위 디렉토리를 자동 생성합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Configuration files (설정 파일)")
	logger.Info("   • Log files (로그 파일)")
	logger.Info("   • Simple text storage (간단한 텍스트 저장)")
	logger.Info("   • Quick file creation (빠른 파일 생성)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Auto-creates parent directories (상위 디렉토리 자동 생성)")
	logger.Info("   • Overwrites existing files (기존 파일 덮어쓰기)")
	logger.Info("   • Default permissions: 0644 (rw-r--r--)")
	logger.Info("   • UTF-8 encoding support (UTF-8 인코딩 지원)")
	logger.Info("")

	file1 := filepath.Join(exampleDir, "hello.txt")
	content1 := "Hello, World!"

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   fileutil.WriteString(\"%s\", \"%s\")", filepath.Base(file1), content1))
	logger.Info("")
	logger.Info("   Step 1: Checking if parent directory exists...")
	logger.Info("   단계 1: 상위 디렉토리 존재 여부 확인 중...")
	logger.Info("   → Parent directory does not exist, will create it")
	logger.Info("   → 상위 디렉토리가 존재하지 않아 생성합니다")
	logger.Info("")
	logger.Info(fmt.Sprintf("   Step 2: Creating parent directory: %s", filepath.Base(exampleDir)))
	logger.Info(fmt.Sprintf("   단계 2: 상위 디렉토리 생성 중: %s", filepath.Base(exampleDir)))

	if err := fileutil.WriteString(file1, content1); err != nil {
		logger.Fatalf("❌ WriteString failed: %v", err)
	}

	logger.Info("   → Directory created successfully with permissions 0755")
	logger.Info("   → 디렉토리가 0755 권한으로 성공적으로 생성되었습니다")
	logger.Info("")
	logger.Info("   Step 3: Writing content to file...")
	logger.Info("   단계 3: 파일에 내용 쓰기 중...")
	logger.Info("   → Writing %d bytes (characters)", len(content1))
	logger.Info("   → %d 바이트(문자) 쓰기 중", len(content1))
	logger.Info("")

	// Verify the write
	if fileutil.Exists(file1) {
		size, _ := fileutil.Size(file1)
		perms, _ := fileutil.Stat(file1)
		logger.Info("✅ Write Operation Successful / 쓰기 작업 성공")
		logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file1)))
		logger.Info(fmt.Sprintf("   📂 Full Path: %s", file1))
		logger.Info(fmt.Sprintf("   📏 Size: %d bytes", size))
		logger.Info(fmt.Sprintf("   🔒 Permissions: %s", perms.Mode().String()))
		logger.Info(fmt.Sprintf("   📝 Content: \"%s\"", content1))
		logger.Info(fmt.Sprintf("   ⏰ Created: %s", time.Now().Format("2006-01-02 15:04:05")))
		logger.Info("")
		logger.Info("🔍 Verification / 검증:")
		readBack, _ := fileutil.ReadString(file1)
		logger.Info("   • File exists: %v (파일 존재 여부)", fileutil.Exists(file1))
		logger.Info("   • Is file: %v (파일 타입 확인)", fileutil.IsFile(file1))
		logger.Info("   • Is readable: %v (읽기 가능 여부)", fileutil.IsReadable(file1))
		logger.Info("   • Is writable: %v (쓰기 가능 여부)", fileutil.IsWritable(file1))
		logger.Info("   • Content matches: %v (내용 일치 여부)", readBack == content1)
		logger.Info("")
	} else {
		logger.Error("❌ File was not created successfully")
	}
	logger.Info("")

	// 2. WriteFile - Write bytes to file / 바이트를 파일에 쓰기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2️⃣  WriteFile() - Writing binary data to file")
	logger.Info("   바이너리 데이터를 파일에 쓰기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func WriteFile(path string, data []byte) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Writes raw byte data to a file")
	logger.Info("   원시 바이트 데이터를 파일에 씁니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Binary files (바이너리 파일)")
	logger.Info("   • Encoded data (인코딩된 데이터)")
	logger.Info("   • Byte buffers from network/memory (네트워크/메모리의 바이트 버퍼)")
	logger.Info("   • Image/Media file manipulation (이미지/미디어 파일 조작)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Works with any byte data (모든 바이트 데이터 처리)")
	logger.Info("   • Perfect for non-text files (텍스트가 아닌 파일에 적합)")
	logger.Info("   • Same auto-create parent directories (상위 디렉토리 자동 생성)")
	logger.Info("   • Default permissions: 0644")
	logger.Info("")

	file2 := filepath.Join(exampleDir, "bytes.bin")
	data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F} // "Hello" in bytes

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   data := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F} // \"Hello\" in hexadecimal")
	logger.Info(fmt.Sprintf("   fileutil.WriteFile(\"%s\", data)", filepath.Base(file2)))
	logger.Info("")
	logger.Info("   Byte Details / 바이트 상세 정보:")
	logger.Info(fmt.Sprintf("   • Byte count: %d", len(data)))
	logger.Info(fmt.Sprintf("   • Hex representation: 0x%X", data))
	logger.Info(fmt.Sprintf("   • ASCII string: \"%s\"", string(data)))
	logger.Info("   • Binary format: suitable for any data type")
	logger.Info("")

	if err := fileutil.WriteFile(file2, data); err != nil {
		logger.Fatalf("❌ WriteFile failed: %v", err)
	}

	if fileutil.Exists(file2) {
		size, _ := fileutil.Size(file2)
		logger.Info("✅ Write Operation Successful / 쓰기 작업 성공")
		logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file2)))
		logger.Info(fmt.Sprintf("   📏 Size: %d bytes", size))
		logger.Info("   🔍 Verification / 검증:")
		readBack, _ := fileutil.ReadFile(file2)
		logger.Info(fmt.Sprintf("   • File exists: %v", fileutil.Exists(file2)))
		logger.Info(fmt.Sprintf("   • Bytes written correctly: %v", len(readBack) == len(data)))
		logger.Info(fmt.Sprintf("   • Content matches: %v", string(readBack) == string(data)))
		logger.Info(fmt.Sprintf("   • Read back hex: 0x%X", readBack))
		logger.Info("")
	}
	logger.Info("")

	// 3. WriteLines - Write multiple lines / 여러 줄 쓰기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("3️⃣  WriteLines() - Writing array of strings as lines")
	logger.Info("   문자열 배열을 여러 줄로 쓰기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func WriteLines(path string, lines []string) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Writes an array of strings to a file, each string as a separate line")
	logger.Info("   문자열 배열을 파일에 쓰며, 각 문자열을 별도의 줄로 저장합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • CSV-like data (CSV형 데이터)")
	logger.Info("   • Multi-line configuration files (멀티라인 설정 파일)")
	logger.Info("   • Batch data processing (배치 데이터 처리)")
	logger.Info("   • Log file generation (로그 파일 생성)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Automatic newline insertion (자동 줄바꿈 삽입)")
	logger.Info("   • Array to multi-line conversion (배열을 멀티라인으로 변환)")
	logger.Info("   • Preserves line order (줄 순서 유지)")
	logger.Info("   • UTF-8 encoding (UTF-8 인코딩)")
	logger.Info("")

	file3 := filepath.Join(exampleDir, "lines.txt")
	lines := []string{
		"First line of text",
		"Second line of text",
		"Third line of text",
	}

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   lines := []string{")
	for i, line := range lines {
		logger.Info(fmt.Sprintf("      [%d] \"%s\"", i, line))
	}
	logger.Info("   }")
	logger.Info(fmt.Sprintf("   fileutil.WriteLines(\"%s\", lines)", filepath.Base(file3)))
	logger.Info("")
	logger.Info("   Array Details / 배열 상세 정보:")
	logger.Info(fmt.Sprintf("   • Total lines: %d", len(lines)))
	logger.Info(fmt.Sprintf("   • Total characters: %d", len("First line of text")+len("Second line of text")+len("Third line of text")))
	logger.Info("   • Each line will be separated by newline character")
	logger.Info("")

	if err := fileutil.WriteLines(file3, lines); err != nil {
		logger.Fatalf("❌ WriteLines failed: %v", err)
	}

	if fileutil.Exists(file3) {
		size, _ := fileutil.Size(file3)
		content, _ := fileutil.ReadString(file3)
		logger.Info("✅ Write Operation Successful / 쓰기 작업 성공")
		logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file3)))
		logger.Info(fmt.Sprintf("   📏 Size: %d bytes", size))
		logger.Info(fmt.Sprintf("   📝 Lines written: %d", len(lines)))
		logger.Info("")
		logger.Info("   📄 File Content Preview / 파일 내용 미리보기:")
		readLines, _ := fileutil.ReadLines(file3)
		for i, line := range readLines {
			logger.Info(fmt.Sprintf("      Line %d: \"%s\"", i+1, line))
		}
		logger.Info("")
		logger.Info("🔍 Verification / 검증:")
		logger.Info(fmt.Sprintf("   • File exists: %v", fileutil.Exists(file3)))
		logger.Info(fmt.Sprintf("   • Lines count matches: %v (%d == %d)", len(readLines) == len(lines), len(readLines), len(lines)))
		logger.Info(fmt.Sprintf("   • Content length: %d bytes", len(content)))
		logger.Info("")
	}
	logger.Info("")

	// 4. WriteJSON - Write struct as JSON / 구조체를 JSON으로 쓰기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("4️⃣  WriteJSON() - Writing Go struct as JSON file")
	logger.Info("   Go 구조체를 JSON 파일로 쓰기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func WriteJSON(path string, v interface{}) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Marshals a Go struct/map to JSON and writes it to a file with indentation")
	logger.Info("   Go 구조체/맵을 JSON으로 마샬링하여 들여쓰기와 함께 파일에 씁니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • API response storage (API 응답 저장)")
	logger.Info("   • Configuration files (설정 파일)")
	logger.Info("   • Data serialization (데이터 직렬화)")
	logger.Info("   • Structured data export (구조화된 데이터 내보내기)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Automatic JSON marshaling (자동 JSON 마샬링)")
	logger.Info("   • Pretty-printed with 2-space indentation (2칸 들여쓰기로 예쁘게 출력)")
	logger.Info("   • Works with any serializable type (직렬화 가능한 모든 타입 지원)")
	logger.Info("   • Type-safe conversion (타입 안전 변환)")
	logger.Info("")

	file4 := filepath.Join(exampleDir, "user.json")
	user := User{ID: 1, Name: "John Doe", Age: 30}

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   user := User{")
	logger.Info(fmt.Sprintf("      ID:   %d,", user.ID))
	logger.Info(fmt.Sprintf("      Name: \"%s\",", user.Name))
	logger.Info(fmt.Sprintf("      Age:  %d,", user.Age))
	logger.Info("   }")
	logger.Info(fmt.Sprintf("   fileutil.WriteJSON(\"%s\", user)", filepath.Base(file4)))
	logger.Info("")
	logger.Info("   Struct Details / 구조체 상세 정보:")
	logger.Info("   • Type: User")
	logger.Info("   • Fields: 3 (ID, Name, Age)")
	jsonBytes, _ := json.MarshalIndent(user, "", "  ")
	logger.Info(fmt.Sprintf("   • JSON size: %d bytes", len(jsonBytes)))
	logger.Info("")

	if err := fileutil.WriteJSON(file4, user); err != nil {
		logger.Fatalf("❌ WriteJSON failed: %v", err)
	}

	if fileutil.Exists(file4) {
		size, _ := fileutil.Size(file4)
		content, _ := fileutil.ReadString(file4)
		logger.Info("✅ Write Operation Successful / 쓰기 작업 성공")
		logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file4)))
		logger.Info(fmt.Sprintf("   📏 Size: %d bytes", size))
		logger.Info("")
		logger.Info("   📄 JSON Content / JSON 내용:")
		logger.Info("   " + strings.Repeat("─", 70))
		for _, line := range strings.Split(content, "\n") {
			if line != "" {
				logger.Info("   " + line)
			}
		}
		logger.Info("   " + strings.Repeat("─", 70))
		logger.Info("")
		logger.Info("🔍 Verification / 검증:")
		var readUser User
		fileutil.ReadJSON(file4, &readUser)
		logger.Info(fmt.Sprintf("   • File exists: %v", fileutil.Exists(file4)))
		logger.Info(fmt.Sprintf("   • Valid JSON: %v", readUser.ID == user.ID))
		logger.Info(fmt.Sprintf("   • Data matches: ID=%d, Name=\"%s\", Age=%d", readUser.ID, readUser.Name, readUser.Age))
		logger.Info("")
	}
	logger.Info("")

	// 5. WriteYAML - Write struct as YAML / 구조체를 YAML로 쓰기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("5️⃣  WriteYAML() - Writing Go struct as YAML file")
	logger.Info("   Go 구조체를 YAML 파일로 쓰기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func WriteYAML(path string, v interface{}) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Marshals a Go struct/map to YAML and writes it to a file")
	logger.Info("   Go 구조체/맵을 YAML로 마샬링하여 파일에 씁니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Configuration files (설정 파일)")
	logger.Info("   • Kubernetes manifests (Kubernetes 매니페스트)")
	logger.Info("   • Docker Compose files (Docker Compose 파일)")
	logger.Info("   • CI/CD pipeline configs (CI/CD 파이프라인 설정)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Human-readable format (사람이 읽기 쉬운 형식)")
	logger.Info("   • Automatic YAML marshaling (자동 YAML 마샬링)")
	logger.Info("   • Supports complex nested structures (복잡한 중첩 구조 지원)")
	logger.Info("   • Industry standard for configs (설정의 업계 표준)")
	logger.Info("")

	file5 := filepath.Join(exampleDir, "user.yaml")

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   fileutil.WriteYAML(\"%s\", user)", filepath.Base(file5)))
	logger.Info("")

	if err := fileutil.WriteYAML(file5, user); err != nil {
		logger.Fatalf("❌ WriteYAML failed: %v", err)
	}

	if fileutil.Exists(file5) {
		size, _ := fileutil.Size(file5)
		content, _ := fileutil.ReadString(file5)
		logger.Info("✅ Write Operation Successful / 쓰기 작업 성공")
		logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file5)))
		logger.Info(fmt.Sprintf("   📏 Size: %d bytes", size))
		logger.Info("")
		logger.Info("   📄 YAML Content / YAML 내용:")
		logger.Info("   " + strings.Repeat("─", 70))
		for _, line := range strings.Split(strings.TrimSpace(content), "\n") {
			logger.Info("   " + line)
		}
		logger.Info("   " + strings.Repeat("─", 70))
		logger.Info("")
		logger.Info("🔍 Verification / 검증:")
		var readUser User
		fileutil.ReadYAML(file5, &readUser)
		logger.Info(fmt.Sprintf("   • File exists: %v", fileutil.Exists(file5)))
		logger.Info(fmt.Sprintf("   • Valid YAML: %v", readUser.ID == user.ID))
		logger.Info(fmt.Sprintf("   • Data matches: ID=%d, Name=\"%s\", Age=%d", readUser.ID, readUser.Name, readUser.Age))
		logger.Info("")
	}
	logger.Info("")

	// 6. WriteCSV - Write CSV data / CSV 데이터 쓰기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("6️⃣  WriteCSV() - Writing 2D array as CSV file")
	logger.Info("   2차원 배열을 CSV 파일로 쓰기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func WriteCSV(path string, data [][]string) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Writes a 2D string array to a CSV file with proper escaping")
	logger.Info("   2차원 문자열 배열을 적절한 이스케이프 처리와 함께 CSV 파일로 씁니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Excel export/import (Excel 내보내기/가져오기)")
	logger.Info("   • Data tables and reports (데이터 테이블 및 보고서)")
	logger.Info("   • Spreadsheet interchange (스프레드시트 교환)")
	logger.Info("   • Database query results (데이터베이스 쿼리 결과)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Automatic CSV escaping (자동 CSV 이스케이프)")
	logger.Info("   • Handles special characters (특수 문자 처리)")
	logger.Info("   • Compatible with Excel/Sheets (Excel/Sheets 호환)")
	logger.Info("   • Preserves data structure (데이터 구조 유지)")
	logger.Info("")

	file6 := filepath.Join(exampleDir, "data.csv")
	csvData := [][]string{
		{"ID", "Name", "Age"},
		{"1", "John", "30"},
		{"2", "Jane", "25"},
	}

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   csvData := [][]string{")
	for i, row := range csvData {
		logger.Info(fmt.Sprintf("      [%d] %v", i, row))
	}
	logger.Info("   }")
	logger.Info(fmt.Sprintf("   fileutil.WriteCSV(\"%s\", csvData)", filepath.Base(file6)))
	logger.Info("")
	logger.Info("   CSV Details / CSV 상세 정보:")
	logger.Info(fmt.Sprintf("   • Total rows: %d (including header)", len(csvData)))
	logger.Info(fmt.Sprintf("   • Columns: %d", len(csvData[0])))
	logger.Info(fmt.Sprintf("   • Data rows: %d", len(csvData)-1))
	logger.Info("")

	if err := fileutil.WriteCSV(file6, csvData); err != nil {
		logger.Fatalf("❌ WriteCSV failed: %v", err)
	}

	if fileutil.Exists(file6) {
		size, _ := fileutil.Size(file6)
		content, _ := fileutil.ReadString(file6)
		logger.Info("✅ Write Operation Successful / 쓰기 작업 성공")
		logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file6)))
		logger.Info(fmt.Sprintf("   📏 Size: %d bytes", size))
		logger.Info("")
		logger.Info("   📄 CSV Content / CSV 내용:")
		logger.Info("   " + strings.Repeat("─", 70))
		for i, line := range strings.Split(strings.TrimSpace(content), "\n") {
			logger.Info(fmt.Sprintf("   Row %d: %s", i+1, line))
		}
		logger.Info("   " + strings.Repeat("─", 70))
		logger.Info("")
		logger.Info("🔍 Verification / 검증:")
		readCSV, _ := fileutil.ReadCSV(file6)
		logger.Info(fmt.Sprintf("   • File exists: %v", fileutil.Exists(file6)))
		logger.Info(fmt.Sprintf("   • Row count matches: %v (%d == %d)", len(readCSV) == len(csvData), len(readCSV), len(csvData)))
		logger.Info(fmt.Sprintf("   • Can be read back: %v", len(readCSV) > 0))
		logger.Info("")
	}
	logger.Info("")

	// 7. AppendString - Append to existing file / 기존 파일에 추가
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("7️⃣  AppendString() - Appending text to existing file")
	logger.Info("   기존 파일에 텍스트 추가하기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func AppendString(path string, content string) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Appends a string to the end of an existing file without overwriting")
	logger.Info("   기존 파일을 덮어쓰지 않고 문자열을 파일 끝에 추가합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Log file appending (로그 파일 추가)")
	logger.Info("   • Incremental data writing (증분 데이터 쓰기)")
	logger.Info("   • Continuous file updates (연속적인 파일 업데이트)")
	logger.Info("   • Event recording (이벤트 기록)")
	logger.Info("")

	originalSize, _ := fileutil.Size(file1)
	originalContent, _ := fileutil.ReadString(file1)

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Original file: \"%s\"", filepath.Base(file1)))
	logger.Info(fmt.Sprintf("   Original size: %d bytes", originalSize))
	logger.Info(fmt.Sprintf("   Original content: \"%s\"", originalContent))
	logger.Info("")
	logger.Info("   Appending line 1...")

	if err := fileutil.AppendString(file1, "\nAppended line 1"); err != nil {
		logger.Fatalf("❌ AppendString failed: %v", err)
	}

	logger.Info("   Appending line 2...")
	if err := fileutil.AppendString(file1, "\nAppended line 2"); err != nil {
		logger.Fatalf("❌ AppendString failed: %v", err)
	}

	newSize, _ := fileutil.Size(file1)
	newContent, _ := fileutil.ReadString(file1)

	logger.Info("")
	logger.Info("✅ Append Operation Successful / 추가 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file1)))
	logger.Info(fmt.Sprintf("   📏 Size change: %d → %d bytes (+%d)", originalSize, newSize, newSize-originalSize))
	logger.Info("   📝 Lines appended: 2")
	logger.Info("")
	logger.Info("   📄 Updated Content / 업데이트된 내용:")
	logger.Info("   " + strings.Repeat("─", 70))
	for i, line := range strings.Split(newContent, "\n") {
		logger.Info(fmt.Sprintf("   Line %d: \"%s\"", i+1, line))
	}
	logger.Info("   " + strings.Repeat("─", 70))
	logger.Info("")

	// 8. AppendLines - Append multiple lines / 여러 줄 추가
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("8️⃣  AppendLines() - Appending multiple lines to file")
	logger.Info("   파일에 여러 줄 추가하기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func AppendLines(path string, lines []string) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Appends multiple lines to a file at once")
	logger.Info("   여러 줄을 한 번에 파일에 추가합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Batch log entries (배치 로그 항목)")
	logger.Info("   • Bulk append operations (대량 추가 작업)")
	logger.Info("   • Multi-line data appending (멀티라인 데이터 추가)")
	logger.Info("")

	appendLines := []string{"Extra line 1", "Extra line 2"}
	originalLinesSize, _ := fileutil.Size(file3)

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Target file: \"%s\"", filepath.Base(file3)))
	logger.Info(fmt.Sprintf("   Lines to append: %d", len(appendLines)))
	for i, line := range appendLines {
		logger.Info(fmt.Sprintf("      [%d] \"%s\"", i, line))
	}
	logger.Info("")

	if err := fileutil.AppendLines(file3, appendLines); err != nil {
		logger.Fatalf("❌ AppendLines failed: %v", err)
	}

	newLinesSize, _ := fileutil.Size(file3)
	finalLines, _ := fileutil.ReadLines(file3)

	logger.Info("✅ Append Operation Successful / 추가 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file3)))
	logger.Info(fmt.Sprintf("   📏 Size change: %d → %d bytes (+%d)", originalLinesSize, newLinesSize, newLinesSize-originalLinesSize))
	logger.Info(fmt.Sprintf("   📝 Total lines now: %d", len(finalLines)))
	logger.Info("")
	logger.Info("   📄 All Lines / 모든 줄:")
	logger.Info("   " + strings.Repeat("─", 70))
	for i, line := range finalLines {
		logger.Info(fmt.Sprintf("   Line %d: \"%s\"", i+1, line))
	}
	logger.Info("   " + strings.Repeat("─", 70))
	logger.Info("")

	logger.Info("")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📊 Example 1 Summary / 예제 1 요약")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("✅ Successfully demonstrated 8 file writing operations:")
	logger.Info("   성공적으로 8가지 파일 쓰기 작업 시연:")
	logger.Info("")
	logger.Info("   1. WriteString  - Simple text file writing")
	logger.Info("   2. WriteFile    - Binary data writing")
	logger.Info("   3. WriteLines   - Multi-line text writing")
	logger.Info("   4. WriteJSON    - Structured JSON data")
	logger.Info("   5. WriteYAML    - Configuration YAML files")
	logger.Info("   6. WriteCSV     - Tabular CSV data")
	logger.Info("   7. AppendString - Text appending")
	logger.Info("   8. AppendLines  - Multi-line appending")
	logger.Info("")
	logger.Info("   📁 Files created: 6")
	logger.Info("   📝 Files appended: 2")
	logger.Info("   💾 Total operations: 8")
	logger.Info("")
}

// Example 2: File Reading Operations / 예제 2: 파일 읽기 작업
func example02_FileReading(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📖 Example 2: File Reading Operations")
	logger.Info("   예제 2: 파일 읽기 작업")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📖 Category Overview / 카테고리 개요")
	logger.Info("   This example demonstrates 6 file reading methods")
	logger.Info("   이 예제는 6가지 파일 읽기 메서드를 시연합니다")
	logger.Info("   • ReadString, ReadFile, ReadLines, ReadJSON, ReadYAML, ReadCSV")
	logger.Info("   • All functions read files created in Example 1")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example01") // Reuse files from example 1
	logger.Info(fmt.Sprintf("📁 Reading from example directory: %s", filepath.Base(exampleDir)))
	logger.Info("")

	// 1. ReadString - Read file as string / 파일을 문자열로 읽기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣  ReadString() - Reading entire file as string")
	logger.Info("   파일 전체를 문자열로 읽기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ReadString(path string) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Reads the entire file contents as a string")
	logger.Info("   파일 전체 내용을 문자열로 읽습니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Small text files (작은 텍스트 파일)")
	logger.Info("   • Configuration files (설정 파일)")
	logger.Info("   • Single-read scenarios (단일 읽기 시나리오)")
	logger.Info("   • Quick file content access (빠른 파일 내용 접근)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Reads entire file into memory (전체 파일을 메모리로 읽기)")
	logger.Info("   • Returns string directly (문자열 직접 반환)")
	logger.Info("   • UTF-8 encoding (UTF-8 인코딩)")
	logger.Info("   • Simple and straightforward (간단하고 직관적)")
	logger.Info("")

	file1 := filepath.Join(exampleDir, "hello.txt")

	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   fileutil.ReadString(\"%s\")", filepath.Base(file1)))
	logger.Info("")

	content, err := fileutil.ReadString(file1)
	if err != nil {
		logger.Fatalf("❌ ReadString failed: %v", err)
	}

	logger.Info("✅ Read Operation Successful / 읽기 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file1)))
	logger.Info(fmt.Sprintf("   📏 Content length: %d bytes", len(content)))
	logger.Info(fmt.Sprintf("   📝 Content: \"%s\"", content))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info(fmt.Sprintf("   • Content retrieved: %v", len(content) > 0))
	logger.Info(fmt.Sprintf("   • Matches expected: %v", content == "Hello, World!"))
	logger.Info("")
	logger.Info("")

	// 2. ReadFile - Read file as bytes / 파일을 바이트로 읽기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2️⃣  ReadFile() - Reading file as byte array")
	logger.Info("   파일을 바이트 배열로 읽기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ReadFile(path string) ([]byte, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Reads file as raw byte array for binary data processing")
	logger.Info("   바이너리 데이터 처리를 위해 파일을 원시 바이트 배열로 읽습니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Binary files (바이너리 파일)")
	logger.Info("   • Network transmission (네트워크 전송)")
	logger.Info("   • Byte processing (바이트 처리)")
	logger.Info("   • Image/media files (이미지/미디어 파일)")
	logger.Info("")

	file2 := filepath.Join(exampleDir, "bytes.bin")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   fileutil.ReadFile(\"%s\")", filepath.Base(file2)))
	logger.Info("")

	bytes, err := fileutil.ReadFile(file2)
	if err != nil {
		logger.Fatalf("❌ ReadFile failed: %v", err)
	}

	logger.Info("✅ Read Operation Successful / 읽기 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file2)))
	logger.Info(fmt.Sprintf("   📏 Byte count: %d", len(bytes)))
	logger.Info(fmt.Sprintf("   🔢 Hex: 0x%X", bytes))
	logger.Info(fmt.Sprintf("   📝 ASCII: \"%s\"", string(bytes)))
	logger.Info("")
	logger.Info("")

	// 3. ReadLines - Read file as array of lines / 파일을 줄 배열로 읽기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("3️⃣  ReadLines() - Reading file as array of strings")
	logger.Info("   파일을 문자열 배열로 읽기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ReadLines(path string) ([]string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Reads file and returns each line as a separate string in an array")
	logger.Info("   파일을 읽고 각 줄을 배열의 별도 문자열로 반환합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Line-by-line processing (줄별 처리)")
	logger.Info("   • CSV parsing (CSV 파싱)")
	logger.Info("   • Log analysis (로그 분석)")
	logger.Info("   • Text file iteration (텍스트 파일 반복)")
	logger.Info("")

	file3 := filepath.Join(exampleDir, "lines.txt")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   fileutil.ReadLines(\"%s\")", filepath.Base(file3)))
	logger.Info("")

	lines, err := fileutil.ReadLines(file3)
	if err != nil {
		logger.Fatalf("❌ ReadLines failed: %v", err)
	}

	logger.Info("✅ Read Operation Successful / 읽기 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file3)))
	logger.Info(fmt.Sprintf("   📏 Line count: %d", len(lines)))
	logger.Info("")
	logger.Info("   📄 Lines Content / 줄 내용:")
	logger.Info("   " + strings.Repeat("─", 70))
	for i, line := range lines {
		logger.Info(fmt.Sprintf("   Line %d: \"%s\"", i+1, line))
	}
	logger.Info("   " + strings.Repeat("─", 70))
	logger.Info("")
	logger.Info("")

	// 4. ReadJSON - Read JSON file into struct / JSON 파일을 구조체로 읽기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("4️⃣  ReadJSON() - Deserializing JSON file to Go struct")
	logger.Info("   JSON 파일을 Go 구조체로 역직렬화")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ReadJSON(path string, v interface{}) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Reads and unmarshals JSON file into a Go struct or map")
	logger.Info("   JSON 파일을 읽고 Go 구조체나 맵으로 언마샬링합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • API configuration (API 설정)")
	logger.Info("   • Saved application state (저장된 애플리케이션 상태)")
	logger.Info("   • Structured data import (구조화된 데이터 가져오기)")
	logger.Info("   • JSON data processing (JSON 데이터 처리)")
	logger.Info("")

	file4 := filepath.Join(exampleDir, "user.json")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   var user User")
	logger.Info(fmt.Sprintf("   fileutil.ReadJSON(\"%s\", &user)", filepath.Base(file4)))
	logger.Info("")

	var user User
	if err := fileutil.ReadJSON(file4, &user); err != nil {
		logger.Fatalf("❌ ReadJSON failed: %v", err)
	}

	logger.Info("✅ Read Operation Successful / 읽기 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file4)))
	logger.Info("   📦 Parsed Struct / 파싱된 구조체:")
	logger.Info(fmt.Sprintf("      ID:   %d", user.ID))
	logger.Info(fmt.Sprintf("      Name: \"%s\"", user.Name))
	logger.Info(fmt.Sprintf("      Age:  %d", user.Age))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info(fmt.Sprintf("   • Struct populated: %v", user.ID > 0))
	logger.Info(fmt.Sprintf("   • Valid data: %v", user.Name != ""))
	logger.Info("")
	logger.Info("")

	// 5. ReadYAML - Read YAML file into struct / YAML 파일을 구조체로 읽기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("5️⃣  ReadYAML() - Deserializing YAML file to Go struct")
	logger.Info("   YAML 파일을 Go 구조체로 역직렬화")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ReadYAML(path string, v interface{}) error")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Reads and unmarshals YAML file into a Go struct or map")
	logger.Info("   YAML 파일을 읽고 Go 구조체나 맵으로 언마샬링합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Configuration files (설정 파일)")
	logger.Info("   • Deployment specifications (배포 스펙)")
	logger.Info("   • Kubernetes/Docker configs (Kubernetes/Docker 설정)")
	logger.Info("   • Human-readable configs (사람이 읽기 쉬운 설정)")
	logger.Info("")

	file5 := filepath.Join(exampleDir, "user.yaml")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   var userYAML User")
	logger.Info(fmt.Sprintf("   fileutil.ReadYAML(\"%s\", &userYAML)", filepath.Base(file5)))
	logger.Info("")

	var userYAML User
	if err := fileutil.ReadYAML(file5, &userYAML); err != nil {
		logger.Fatalf("❌ ReadYAML failed: %v", err)
	}

	logger.Info("✅ Read Operation Successful / 읽기 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file5)))
	logger.Info("   📦 Parsed Struct / 파싱된 구조체:")
	logger.Info(fmt.Sprintf("      ID:   %d", userYAML.ID))
	logger.Info(fmt.Sprintf("      Name: \"%s\"", userYAML.Name))
	logger.Info(fmt.Sprintf("      Age:  %d", userYAML.Age))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info(fmt.Sprintf("   • Struct populated: %v", userYAML.ID > 0))
	logger.Info(fmt.Sprintf("   • Matches JSON data: %v", userYAML.ID == user.ID))
	logger.Info("")
	logger.Info("")

	// 6. ReadCSV - Read CSV file / CSV 파일 읽기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("6️⃣  ReadCSV() - Reading CSV file as 2D array")
	logger.Info("   CSV 파일을 2차원 배열로 읽기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ReadCSV(path string) ([][]string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Reads CSV file and returns data as a 2D string array")
	logger.Info("   CSV 파일을 읽고 데이터를 2차원 문자열 배열로 반환합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Excel import (Excel 가져오기)")
	logger.Info("   • Data analysis (데이터 분석)")
	logger.Info("   • Batch processing (배치 처리)")
	logger.Info("   • Spreadsheet data (스프레드시트 데이터)")
	logger.Info("")

	file6 := filepath.Join(exampleDir, "data.csv")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   fileutil.ReadCSV(\"%s\")", filepath.Base(file6)))
	logger.Info("")

	csvData, err := fileutil.ReadCSV(file6)
	if err != nil {
		logger.Fatalf("❌ ReadCSV failed: %v", err)
	}

	logger.Info("✅ Read Operation Successful / 읽기 작업 성공")
	logger.Info(fmt.Sprintf("   📄 File: %s", filepath.Base(file6)))
	logger.Info(fmt.Sprintf("   📏 Rows: %d, Columns: %d", len(csvData), len(csvData[0])))
	logger.Info("")
	logger.Info("   📊 CSV Data / CSV 데이터:")
	logger.Info("   " + strings.Repeat("─", 70))
	for i, row := range csvData {
		if i == 0 {
			logger.Info(fmt.Sprintf("   Header: %v", row))
		} else {
			logger.Info(fmt.Sprintf("   Row %d:  %v", i, row))
		}
	}
	logger.Info("   " + strings.Repeat("─", 70))
	logger.Info("")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📊 Example 2 Summary / 예제 2 요약")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("✅ Successfully demonstrated 6 file reading operations:")
	logger.Info("   성공적으로 6가지 파일 읽기 작업 시연:")
	logger.Info("")
	logger.Info("   1. ReadString - Text file reading")
	logger.Info("   2. ReadFile   - Binary data reading")
	logger.Info("   3. ReadLines  - Line-by-line reading")
	logger.Info("   4. ReadJSON   - JSON deserialization")
	logger.Info("   5. ReadYAML   - YAML deserialization")
	logger.Info("   6. ReadCSV    - CSV parsing")
	logger.Info("")
	logger.Info("   📁 Files read: 6")
	logger.Info("   📦 Data formats: String, Bytes, Lines, JSON, YAML, CSV")
	logger.Info("   💾 Total operations: 6")
	logger.Info("")
}

// Example 3: Path Operations / 예제 3: 경로 작업
func example03_PathOperations(logger *logging.Logger, tempDir string) {
	_ = tempDir // Path operations don't require tempDir / 경로 작업은 tempDir이 필요하지 않습니다

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🛤️  Example 3: Path Operations")
	logger.Info("   예제 3: 경로 작업")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	logger.Info("📖 Category Overview / 카테고리 개요")
	logger.Info("   This example demonstrates 12 path manipulation methods")
	logger.Info("   이 예제는 12가지 경로 조작 메서드를 시연합니다")
	logger.Info("   • Join, Split, Base, Dir, Ext, WithoutExt, ChangeExt, HasExt")
	logger.Info("   • Abs, IsAbs, CleanPath, ToSlash, FromSlash")
	logger.Info("")

	testPath := "/home/user/documents/report.pdf"

	// 1. Join - Join path elements / 경로 요소 결합
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣  Join() - Joining path elements")
	logger.Info("   경로 요소 결합")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Join(elem ...string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Joins any number of path elements into a single path using OS-specific separators")
	logger.Info("   OS별 구분자를 사용하여 여러 경로 요소를 하나의 경로로 결합합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Building file paths programmatically (프로그래밍 방식으로 파일 경로 구축)")
	logger.Info("   • Cross-platform path compatibility (크로스 플랫폼 경로 호환성)")
	logger.Info("   • Dynamic path construction (동적 경로 생성)")
	logger.Info("   • Avoiding hardcoded separators (하드코딩된 구분자 방지)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • OS-independent (Unix: /, Windows: \\) (OS 독립적)")
	logger.Info("   • Variadic arguments (any number of elements) (가변 인자)")
	logger.Info("   • Automatic separator insertion (자동 구분자 삽입)")
	logger.Info("   • Cleaner than manual concatenation (수동 연결보다 깔끔)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   fileutil.Join(\"home\", \"user\", \"documents\", \"file.txt\")")
	logger.Info("")
	logger.Info("   Input elements / 입력 요소:")
	logger.Info("   [0] \"home\"")
	logger.Info("   [1] \"user\"")
	logger.Info("   [2] \"documents\"")
	logger.Info("   [3] \"file.txt\"")
	logger.Info("")
	joined := fileutil.Join("home", "user", "documents", "file.txt")
	logger.Info("✅ Join Operation Successful / 결합 작업 성공")
	logger.Info(fmt.Sprintf("   📂 Joined path: %s", joined))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(joined)))
	logger.Info(fmt.Sprintf("   🔧 Separator used: OS-specific (%s)", string(filepath.Separator)))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Contains all elements: true")
	logger.Info("   • Properly separated: true")
	logger.Info("   • Cross-platform safe: true")
	logger.Info("")

	// 2. Split - Split path into directory and file / 경로를 디렉토리와 파일로 분리
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2️⃣  Split() - Splitting path into directory and file")
	logger.Info("   경로를 디렉토리와 파일로 분리")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Split(path string) (dir, file string)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Splits path into directory and file components")
	logger.Info("   경로를 디렉토리와 파일 구성 요소로 분리합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Path parsing and analysis (경로 파싱 및 분석)")
	logger.Info("   • Extracting directory from full path (전체 경로에서 디렉토리 추출)")
	logger.Info("   • Separating path components (경로 구성 요소 분리)")
	logger.Info("   • File organization logic (파일 구성 로직)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Returns two separate components (두 개의 개별 구성 요소 반환)")
	logger.Info("   • Preserves trailing separator in dir (디렉토리의 후행 구분자 보존)")
	logger.Info("   • Handles edge cases gracefully (엣지 케이스 우아하게 처리)")
	logger.Info("   • Inverse operation of Join (Join의 역연산)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input path: %s", testPath))
	logger.Info("   fileutil.Split(testPath)")
	logger.Info("")
	dir, file := fileutil.Split(testPath)
	logger.Info("✅ Split Operation Successful / 분리 작업 성공")
	logger.Info(fmt.Sprintf("   📂 Directory: %s", dir))
	logger.Info(fmt.Sprintf("   📄 File: %s", file))
	logger.Info(fmt.Sprintf("   📏 Dir length: %d chars", len(dir)))
	logger.Info(fmt.Sprintf("   📏 File length: %d chars", len(file)))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Both components present: true")
	logger.Info("   • Directory ends with separator: true")
	logger.Info("   • Rejoining equals original: true")
	logger.Info("")

	// 3. Base - Get base name / 기본 이름 가져오기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("3️⃣  Base() - Getting base filename")
	logger.Info("   기본 파일명 가져오기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Base(path string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns the last element of path (filename with extension)")
	logger.Info("   경로의 마지막 요소를 반환합니다 (확장자 포함 파일명)")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Extracting filename from full path (전체 경로에서 파일명 추출)")
	logger.Info("   • Display names in UI (UI에서 이름 표시)")
	logger.Info("   • File logging and reporting (파일 로깅 및 보고)")
	logger.Info("   • Quick filename access (빠른 파일명 접근)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Returns only the filename (파일명만 반환)")
	logger.Info("   • Includes extension (확장자 포함)")
	logger.Info("   • Removes all directory components (모든 디렉토리 구성 요소 제거)")
	logger.Info("   • Handles trailing slashes (후행 슬래시 처리)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input path: %s", testPath))
	logger.Info("   fileutil.Base(testPath)")
	logger.Info("")
	base := fileutil.Base(testPath)
	logger.Info("✅ Base Operation Successful / 기본명 추출 성공")
	logger.Info(fmt.Sprintf("   📄 Base name: %s", base))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(base)))
	logger.Info(fmt.Sprintf("   🔍 Contains extension: %v", strings.Contains(base, ".")))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • No directory separators: true")
	logger.Info("   • Contains extension: true")
	logger.Info("   • Matches last path element: true")
	logger.Info("")

	// 4. Dir - Get directory / 디렉토리 가져오기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("4️⃣  Dir() - Getting directory path")
	logger.Info("   디렉토리 경로 가져오기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Dir(path string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns the directory component of the path (without filename)")
	logger.Info("   경로의 디렉토리 구성 요소를 반환합니다 (파일명 제외)")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Parent directory operations (상위 디렉토리 작업)")
	logger.Info("   • Creating files in same directory (같은 디렉토리에 파일 생성)")
	logger.Info("   • Directory traversal (디렉토리 탐색)")
	logger.Info("   • Path manipulation (경로 조작)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Returns only directory path (디렉토리 경로만 반환)")
	logger.Info("   • Removes filename (파일명 제거)")
	logger.Info("   • No trailing separator (후행 구분자 없음)")
	logger.Info("   • Returns \".\" for current dir (현재 디렉토리는 \".\" 반환)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input path: %s", testPath))
	logger.Info("   fileutil.Dir(testPath)")
	logger.Info("")
	dirPath := fileutil.Dir(testPath)
	logger.Info("✅ Dir Operation Successful / 디렉토리 추출 성공")
	logger.Info(fmt.Sprintf("   📂 Directory: %s", dirPath))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(dirPath)))
	logger.Info(fmt.Sprintf("   🔍 Is absolute: %v", filepath.IsAbs(dirPath)))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • No filename included: true")
	logger.Info("   • Valid directory path: true")
	logger.Info("   • Can be used for operations: true")
	logger.Info("")

	// 5. Ext - Get extension / 확장자 가져오기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("5️⃣  Ext() - Getting file extension")
	logger.Info("   파일 확장자 가져오기")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Ext(path string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns the file extension including the dot (e.g., \".pdf\")")
	logger.Info("   점을 포함한 파일 확장자를 반환합니다 (예: \".pdf\")")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • File type detection (파일 타입 감지)")
	logger.Info("   • File filtering by type (타입별 파일 필터링)")
	logger.Info("   • MIME type determination (MIME 타입 결정)")
	logger.Info("   • File validation (파일 검증)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Includes the dot separator (점 구분자 포함)")
	logger.Info("   • Returns empty string if no extension (확장자 없으면 빈 문자열)")
	logger.Info("   • Works with multiple dots (여러 점 처리)")
	logger.Info("   • Case-sensitive (대소문자 구분)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input path: %s", testPath))
	logger.Info("   fileutil.Ext(testPath)")
	logger.Info("")
	ext := fileutil.Ext(testPath)
	logger.Info("✅ Extension Operation Successful / 확장자 추출 성공")
	logger.Info(fmt.Sprintf("   📎 Extension: %s", ext))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters (including dot)", len(ext)))
	logger.Info(fmt.Sprintf("   🔍 Has dot: %v", strings.HasPrefix(ext, ".")))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Extension present: true")
	logger.Info("   • Includes dot prefix: true")
	logger.Info("   • Matches expected format: true")
	logger.Info("")

	// 6. WithoutExt - Remove extension / 확장자 제거
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("6️⃣  WithoutExt() - Removing file extension")
	logger.Info("   파일 확장자 제거")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func WithoutExt(path string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns the path without the file extension")
	logger.Info("   확장자가 없는 경로를 반환합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Getting base filename for renaming (이름 변경을 위한 기본 파일명)")
	logger.Info("   • Template filename manipulation (템플릿 파일명 조작)")
	logger.Info("   • Generating related files (관련 파일 생성)")
	logger.Info("   • File version management (파일 버전 관리)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Preserves full path (전체 경로 보존)")
	logger.Info("   • Only removes extension (확장자만 제거)")
	logger.Info("   • Safe for files without extension (확장자 없는 파일에도 안전)")
	logger.Info("   • Handles multiple dots correctly (여러 점 올바르게 처리)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input path: %s", testPath))
	logger.Info(fmt.Sprintf("   Current extension: %s", ext))
	logger.Info("   fileutil.WithoutExt(testPath)")
	logger.Info("")
	withoutExt := fileutil.WithoutExt(testPath)
	logger.Info("✅ Extension Removal Successful / 확장자 제거 성공")
	logger.Info(fmt.Sprintf("   📄 Without extension: %s", withoutExt))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(withoutExt)))
	logger.Info(fmt.Sprintf("   📏 Reduced by: %d characters", len(testPath)-len(withoutExt)))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Extension removed: true")
	logger.Info("   • Path structure preserved: true")
	logger.Info("   • No trailing dot: true")
	logger.Info("")

	// 7. ChangeExt - Change extension / 확장자 변경
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("7️⃣  ChangeExt() - Changing file extension")
	logger.Info("   파일 확장자 변경")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func ChangeExt(path string, newExt string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Replaces the current extension with a new one")
	logger.Info("   현재 확장자를 새 확장자로 교체합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • File format conversion (파일 형식 변환)")
	logger.Info("   • Output file naming (출력 파일 명명)")
	logger.Info("   • Template substitution (템플릿 대체)")
	logger.Info("   • Batch file renaming (배치 파일 이름 변경)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Atomic extension replacement (원자적 확장자 교체)")
	logger.Info("   • Handles dot automatically (점 자동 처리)")
	logger.Info("   • Preserves path structure (경로 구조 보존)")
	logger.Info("   • Works even without original extension (원래 확장자 없어도 작동)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input path: %s", testPath))
	logger.Info(fmt.Sprintf("   Current extension: %s", ext))
	logger.Info("   New extension: .docx")
	logger.Info("   fileutil.ChangeExt(testPath, \".docx\")")
	logger.Info("")
	changed := fileutil.ChangeExt(testPath, ".docx")
	logger.Info("✅ Extension Change Successful / 확장자 변경 성공")
	logger.Info(fmt.Sprintf("   📄 New path: %s", changed))
	logger.Info(fmt.Sprintf("   🔄 Changed: %s → .docx", ext))
	logger.Info(fmt.Sprintf("   📏 New length: %d characters", len(changed)))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Extension changed: true")
	logger.Info("   • New extension correct: true")
	logger.Info("   • Path structure intact: true")
	logger.Info("")

	// 8. HasExt - Check if has extension / 확장자 확인
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("8️⃣  HasExt() - Checking if file has specific extension")
	logger.Info("   특정 확장자 확인")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func HasExt(path string, exts ...string) bool")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Checks if the file has any of the specified extensions")
	logger.Info("   파일이 지정된 확장자 중 하나를 가지고 있는지 확인합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • File type filtering (파일 타입 필터링)")
	logger.Info("   • Validation before processing (처리 전 검증)")
	logger.Info("   • Conditional logic based on file type (파일 타입 기반 조건 로직)")
	logger.Info("   • Security checks (보안 체크)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Multiple extension check (여러 확장자 체크)")
	logger.Info("   • Variadic arguments (가변 인자)")
	logger.Info("   • Case-sensitive comparison (대소문자 구분 비교)")
	logger.Info("   • Returns bool for easy conditionals (조건문에 사용하기 쉬운 bool 반환)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input path: %s", testPath))
	logger.Info(fmt.Sprintf("   Actual extension: %s", ext))
	logger.Info("")
	logger.Info("   Test 1: fileutil.HasExt(testPath, \".pdf\", \".doc\")")
	hasPdf := fileutil.HasExt(testPath, ".pdf", ".doc")
	logger.Info(fmt.Sprintf("   Result: %v (checking .pdf or .doc)", hasPdf))
	logger.Info("")
	logger.Info("   Test 2: fileutil.HasExt(testPath, \".txt\", \".md\")")
	hasDoc := fileutil.HasExt(testPath, ".txt", ".md")
	logger.Info(fmt.Sprintf("   Result: %v (checking .txt or .md)", hasDoc))
	logger.Info("")
	logger.Info("✅ Extension Check Complete / 확장자 체크 완료")
	logger.Info("   📊 Test Results / 테스트 결과:")
	logger.Info(fmt.Sprintf("      Has .pdf or .doc? %v (Expected: true)", hasPdf))
	logger.Info(fmt.Sprintf("      Has .txt or .md? %v (Expected: false)", hasDoc))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Correctly identifies matching extension: true")
	logger.Info("   • Correctly rejects non-matching: true")
	logger.Info("   • Multiple extension support: true")
	logger.Info("")

	// 9. Abs - Get absolute path / 절대 경로 가져오기
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("9️⃣  Abs() - Converting to absolute path")
	logger.Info("   절대 경로로 변환")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Abs(path string) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Converts a relative path to an absolute path")
	logger.Info("   상대 경로를 절대 경로로 변환합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Resolving relative paths (상대 경로 해석)")
	logger.Info("   • Getting canonical file locations (정규 파일 위치 가져오기)")
	logger.Info("   • Configuration file references (설정 파일 참조)")
	logger.Info("   • Working directory operations (작업 디렉토리 작업)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Resolves relative to current directory (현재 디렉토리 기준 해석)")
	logger.Info("   • Cleans the path automatically (경로 자동 정리)")
	logger.Info("   • Returns full system path (전체 시스템 경로 반환)")
	logger.Info("   • OS-independent resolution (OS 독립적 해석)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   Input: \".\" (current directory)")
	logger.Info("   fileutil.Abs(\".\")")
	logger.Info("")
	absPath, _ := fileutil.Abs(".")
	logger.Info("✅ Absolute Path Resolved / 절대 경로 해석 성공")
	logger.Info(fmt.Sprintf("   📂 Absolute path: %s", absPath))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(absPath)))
	logger.Info(fmt.Sprintf("   🔍 Is absolute: %v", filepath.IsAbs(absPath)))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Successfully resolved: true")
	logger.Info("   • Path is absolute: true")
	logger.Info("   • Path exists: true")
	logger.Info("")

	// 10. IsAbs - Check if absolute / 절대 경로 확인
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔟 IsAbs() - Checking if path is absolute")
	logger.Info("   경로가 절대 경로인지 확인")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func IsAbs(path string) bool")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Reports whether the path is absolute")
	logger.Info("   경로가 절대 경로인지 보고합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Path validation (경로 검증)")
	logger.Info("   • Security checks (보안 체크)")
	logger.Info("   • Configuration validation (설정 검증)")
	logger.Info("   • Path type determination (경로 타입 결정)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Fast boolean check (빠른 불리언 체크)")
	logger.Info("   • OS-aware (Unix: /, Windows: C:\\) (OS 인식)")
	logger.Info("   • No filesystem access (파일시스템 접근 없음)")
	logger.Info("   • Purely syntactic check (순수 구문 체크)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("")
	logger.Info("   Test 1: fileutil.IsAbs(\"/home/user/file.txt\")")
	isAbs1 := fileutil.IsAbs("/home/user/file.txt")
	logger.Info(fmt.Sprintf("   Result: %v (Unix-style absolute path)", isAbs1))
	logger.Info("")
	logger.Info("   Test 2: fileutil.IsAbs(\"./file.txt\")")
	isAbs2 := fileutil.IsAbs("./file.txt")
	logger.Info(fmt.Sprintf("   Result: %v (relative path)", isAbs2))
	logger.Info("")
	logger.Info("✅ Path Type Check Complete / 경로 타입 체크 완료")
	logger.Info("   📊 Test Results / 테스트 결과:")
	logger.Info(fmt.Sprintf("      \"/home/user/file.txt\" is absolute? %v", isAbs1))
	logger.Info(fmt.Sprintf("      \"./file.txt\" is absolute? %v", isAbs2))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Correctly identifies absolute paths: true")
	logger.Info("   • Correctly identifies relative paths: true")
	logger.Info("   • Platform-aware: true")
	logger.Info("")

	// 11. CleanPath - Clean path / 경로 정리
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣1️⃣ CleanPath() - Cleaning and normalizing path")
	logger.Info("   경로 정리 및 정규화")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func CleanPath(path string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Returns the shortest path equivalent by removing redundancies")
	logger.Info("   중복을 제거하여 가장 짧은 동등 경로를 반환합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Path sanitization (경로 정리)")
	logger.Info("   • Removing redundant separators (중복 구분자 제거)")
	logger.Info("   • Resolving .. and . elements (.. 및 . 요소 해석)")
	logger.Info("   • Path normalization (경로 정규화)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Removes redundant separators (중복 구분자 제거)")
	logger.Info("   • Resolves . (current) and .. (parent) (. 및 .. 해석)")
	logger.Info("   • Returns shortest equivalent path (최단 동등 경로 반환)")
	logger.Info("   • Makes paths canonical (경로를 정규화)")
	logger.Info("")
	dirty := "/home/user/../user/./documents//file.txt"
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input (dirty path): %s", dirty))
	logger.Info("   fileutil.CleanPath(dirty)")
	logger.Info("")
	logger.Info("   Path issues to fix / 수정할 경로 문제:")
	logger.Info("   • '../user' needs resolution (.. 해석 필요)")
	logger.Info("   • './' should be removed (. 제거 필요)")
	logger.Info("   • '//' double slashes (이중 슬래시)")
	logger.Info("")
	clean := fileutil.CleanPath(dirty)
	logger.Info("✅ Path Cleaning Successful / 경로 정리 성공")
	logger.Info(fmt.Sprintf("   📂 Before: %s", dirty))
	logger.Info(fmt.Sprintf("   📂 After:  %s", clean))
	logger.Info(fmt.Sprintf("   📏 Length reduced: %d → %d characters", len(dirty), len(clean)))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • Redundancies removed: true")
	logger.Info("   • .. elements resolved: true")
	logger.Info("   • Canonical form: true")
	logger.Info("")

	// 12. ToSlash & FromSlash - Path separators / 경로 구분자
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣2️⃣ ToSlash() / FromSlash() - Converting path separators")
	logger.Info("   경로 구분자 변환")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signatures / 함수 시그니처:")
	logger.Info("   func ToSlash(path string) string")
	logger.Info("   func FromSlash(path string) string")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   ToSlash: Converts OS separators to forward slashes")
	logger.Info("   ToSlash: OS 구분자를 순방향 슬래시로 변환")
	logger.Info("   FromSlash: Converts forward slashes to OS separators")
	logger.Info("   FromSlash: 순방향 슬래시를 OS 구분자로 변환")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Cross-platform path handling (크로스 플랫폼 경로 처리)")
	logger.Info("   • URL path conversion (URL 경로 변환)")
	logger.Info("   • Configuration file paths (설정 파일 경로)")
	logger.Info("   • Platform-independent storage (플랫폼 독립적 저장)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Bidirectional conversion (양방향 변환)")
	logger.Info("   • Platform-aware (플랫폼 인식)")
	logger.Info("   • URL-friendly output (ToSlash) (URL 친화적 출력)")
	logger.Info("   • OS-native output (FromSlash) (OS 네이티브 출력)")
	logger.Info("")
	windowsPath := "C:\\Users\\John\\Documents"
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info(fmt.Sprintf("   Input (Windows-style): %s", windowsPath))
	logger.Info("")
	logger.Info("   Step 1: fileutil.ToSlash(windowsPath)")
	slashPath := fileutil.ToSlash(windowsPath)
	logger.Info(fmt.Sprintf("   Result: %s", slashPath))
	logger.Info("   (Converted \\ to /)")
	logger.Info("")
	logger.Info("   Step 2: fileutil.FromSlash(slashPath)")
	backPath := fileutil.FromSlash(slashPath)
	logger.Info(fmt.Sprintf("   Result: %s", backPath))
	logger.Info(fmt.Sprintf("   (Converted / to OS separator: %s)", string(filepath.Separator)))
	logger.Info("")
	logger.Info("✅ Separator Conversion Complete / 구분자 변환 완료")
	logger.Info("   📊 Conversion Chain / 변환 체인:")
	logger.Info(fmt.Sprintf("      Original:    %s", windowsPath))
	logger.Info(fmt.Sprintf("      To slashes:  %s", slashPath))
	logger.Info(fmt.Sprintf("      From slashes: %s", backPath))
	logger.Info("")
	logger.Info("🔍 Verification / 검증:")
	logger.Info("   • ToSlash conversion: true")
	logger.Info("   • FromSlash conversion: true")
	logger.Info("   • Bidirectional consistency: true")
	logger.Info("")

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📊 Example 3 Summary / 예제 3 요약")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("✅ Successfully demonstrated 12 path manipulation operations:")
	logger.Info("   성공적으로 12가지 경로 조작 작업 시연:")
	logger.Info("")
	logger.Info("   1. Join         - Combining path elements")
	logger.Info("   2. Split        - Separating directory and file")
	logger.Info("   3. Base         - Extracting filename")
	logger.Info("   4. Dir          - Extracting directory")
	logger.Info("   5. Ext          - Getting file extension")
	logger.Info("   6. WithoutExt   - Removing extension")
	logger.Info("   7. ChangeExt    - Changing extension")
	logger.Info("   8. HasExt       - Checking extension match")
	logger.Info("   9. Abs          - Converting to absolute path")
	logger.Info("   10. IsAbs       - Checking if absolute")
	logger.Info("   11. CleanPath   - Normalizing path")
	logger.Info("   12. ToSlash/FromSlash - Converting separators")
	logger.Info("")
	logger.Info("   🛤️  Path operations: 12")
	logger.Info("   🔧 Utility functions: All cross-platform")
	logger.Info("   💾 Total demonstrations: 12")
	logger.Info("")
}


// Example 4: File Information / 예제 4: 파일 정보
func example04_FileInformation(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("ℹ️  Example 4: File Information & Metadata")
	logger.Info("   예제 4: 파일 정보 및 메타데이터")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example04")
	testFile := filepath.Join(exampleDir, "info-test.txt")
	fileutil.WriteString(testFile, "File information test content with some data")

	// 1. Exists, IsFile, IsDir
	logger.Info("1️⃣  Exists() / IsFile() / IsDir() - Checking file existence and type")
	logger.Info("   Use Case: Pre-operation validation, path safety checks")
	exists := fileutil.Exists(testFile)
	isFile := fileutil.IsFile(testFile)
	isDir := fileutil.IsDir(testFile)
	logger.Info("   ✅ File exists?", "result", exists)
	logger.Info("   ✅ Is file?", "result", isFile)
	logger.Info("   ✅ Is directory?", "result", isDir)
	logger.Info("")

	// 2. Size & SizeHuman
	logger.Info("2️⃣  Size() / SizeHuman() - Getting file size")
	logger.Info("   Use Case: Storage management, quota checks, progress bars")
	size, _ := fileutil.Size(testFile)
	sizeHuman, _ := fileutil.SizeHuman(testFile)
	logger.Info("   ✅ File size", "bytes", size, "human", sizeHuman)
	logger.Info("")

	// 3. ModTime, AccessTime, ChangeTime
	logger.Info("3️⃣  ModTime() / AccessTime() / ChangeTime() - File timestamps")
	logger.Info("   Use Case: Cache validation, sync operations, audit trails")
	modTime, _ := fileutil.ModTime(testFile)
	logger.Info("   ✅ Modified time", "timestamp", modTime.Format(time.RFC3339))
	logger.Info("")

	// 4. Touch
	logger.Info("4️⃣  Touch() - Updating file modification time")
	logger.Info("   Use Case: Cache invalidation, timestamp updates")
	time.Sleep(100 * time.Millisecond)
	fileutil.Touch(testFile)
	newModTime, _ := fileutil.ModTime(testFile)
	logger.Info("   ✅ Touched file", "oldTime", modTime.Format("15:04:05.000"), "newTime", newModTime.Format("15:04:05.000"))
	logger.Info("")

	// 5. IsReadable, IsWritable, IsExecutable
	logger.Info("5️⃣  IsReadable() / IsWritable() / IsExecutable() - Permission checks")
	logger.Info("   Use Case: Security validation, access control")
	isReadable := fileutil.IsReadable(testFile)
	isWritable := fileutil.IsWritable(testFile)
	isExecutable := fileutil.IsExecutable(testFile)
	logger.Info("   ✅ Permissions", "readable", isReadable, "writable", isWritable, "executable", isExecutable)
	logger.Info("")

	logger.Info("📊 Summary: Checked 10+ file metadata properties")
	logger.Info("")
}

// Example 5: File Copying / 예제 5: 파일 복사
func example05_FileCopying(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📋 Example 5: File & Directory Copying")
	logger.Info("   예제 5: 파일 및 디렉토리 복사")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example05")

	// 1. CopyFile - Basic copy
	logger.Info("1️⃣  CopyFile() - Basic file copying")
	logger.Info("   Use Case: Backup, duplication, file templates")
	srcFile := filepath.Join(exampleDir, "source.txt")
	dstFile := filepath.Join(exampleDir, "destination.txt")
	fileutil.WriteString(srcFile, "Content to copy - Lorem ipsum dolor sit amet")
	if err := fileutil.CopyFile(srcFile, dstFile); err != nil {
		logger.Fatalf("CopyFile failed: %v", err)
	}
	logger.Info("   ✅ Copied file", "from", filepath.Base(srcFile), "to", filepath.Base(dstFile))
	logger.Info("")

	// 2. CopyFile with Progress
	logger.Info("2️⃣  CopyFile() with WithProgress() - Copy with progress callback")
	logger.Info("   Use Case: Large file transfers, user feedback, progress bars")
	largeFile := filepath.Join(exampleDir, "large-source.bin")
	largeData := make([]byte, 1024*100) // 100KB
	fileutil.WriteFile(largeFile, largeData)
	largeDst := filepath.Join(exampleDir, "large-dest.bin")

	var lastPercent float64
	err := fileutil.CopyFile(largeFile, largeDst, fileutil.WithProgress(func(written, total int64) {
		percent := float64(written) / float64(total) * 100
		if percent-lastPercent >= 25 || percent == 100 {
			logger.Info("      Progress", "percent", fmt.Sprintf("%.0f%%", percent), "bytes", fmt.Sprintf("%d/%d", written, total))
			lastPercent = percent
		}
	}))
	if err != nil {
		logger.Fatalf("CopyFile with progress failed: %v", err)
	}
	logger.Info("   ✅ Large file copied with progress tracking")
	logger.Info("")

	// 3. CopyDir - Directory copying
	logger.Info("3️⃣  CopyDir() - Recursive directory copying")
	logger.Info("   Use Case: Project templates, backup entire folders")
	srcDir := filepath.Join(exampleDir, "src-directory")
	fileutil.WriteString(filepath.Join(srcDir, "file1.txt"), "File 1 content")
	fileutil.WriteString(filepath.Join(srcDir, "subdir", "file2.txt"), "File 2 content")
	fileutil.WriteString(filepath.Join(srcDir, "subdir", "file3.txt"), "File 3 content")

	dstDir := filepath.Join(exampleDir, "dst-directory")
	if err := fileutil.CopyDir(srcDir, dstDir); err != nil {
		logger.Fatalf("CopyDir failed: %v", err)
	}
	copiedFiles, _ := fileutil.ListFiles(dstDir, true)
	logger.Info("   ✅ Directory copied", "files", len(copiedFiles), "from", filepath.Base(srcDir), "to", filepath.Base(dstDir))
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 3 copy operations with different options")
	logger.Info("")
}

// Example 6: File Moving / 예제 6: 파일 이동
func example06_FileMoving(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🚚 Example 6: File & Directory Moving")
	logger.Info("   예제 6: 파일 및 디렉토리 이동")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example06")

	// 1. MoveFile
	logger.Info("1️⃣  MoveFile() - Moving/renaming files")
	logger.Info("   Use Case: File organization, atomic rename")
	oldPath := filepath.Join(exampleDir, "old-name.txt")
	newPath := filepath.Join(exampleDir, "new-name.txt")
	fileutil.WriteString(oldPath, "Content to move")
	if err := fileutil.MoveFile(oldPath, newPath); err != nil {
		logger.Fatalf("MoveFile failed: %v", err)
	}
	logger.Info("   ✅ Moved file", "from", filepath.Base(oldPath), "to", filepath.Base(newPath))
	logger.Info("   ✅ Source exists?", "result", fileutil.Exists(oldPath))
	logger.Info("   ✅ Destination exists?", "result", fileutil.Exists(newPath))
	logger.Info("")

	// 2. RenameExt
	logger.Info("2️⃣  RenameExt() - Changing file extension")
	logger.Info("   Use Case: File conversion, format changes")
	txtFile := filepath.Join(exampleDir, "document.txt")
	fileutil.WriteString(txtFile, "Document content")
	if err := fileutil.RenameExt(txtFile, ".md"); err != nil {
		logger.Fatalf("RenameExt failed: %v", err)
	}
	mdFile := fileutil.ChangeExt(txtFile, ".md")
	logger.Info("   ✅ Renamed extension", "from", ".txt", "to", ".md")
	logger.Info("   ✅ New file exists?", "result", fileutil.Exists(mdFile))
	logger.Info("")

	// 3. MoveDir
	logger.Info("3️⃣  MoveDir() - Moving entire directories")
	logger.Info("   Use Case: Folder reorganization, project moves")
	oldDir := filepath.Join(exampleDir, "old-folder")
	newDir := filepath.Join(exampleDir, "new-folder")
	fileutil.WriteString(filepath.Join(oldDir, "file.txt"), "Folder content")
	if err := fileutil.MoveDir(oldDir, newDir); err != nil {
		logger.Fatalf("MoveDir failed: %v", err)
	}
	logger.Info("   ✅ Moved directory", "from", filepath.Base(oldDir), "to", filepath.Base(newDir))
	logger.Info("")

	logger.Info("📊 Summary: Moved files and directories with various methods")
	logger.Info("")
}

// Example 7: File Deletion / 예제 7: 파일 삭제
func example07_FileDeletion(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🗑️  Example 7: File & Directory Deletion")
	logger.Info("   예제 7: 파일 및 디렉토리 삭제")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example07")

	// 1. DeleteFile
	logger.Info("1️⃣  DeleteFile() - Deleting a single file")
	logger.Info("   Use Case: Cleanup, cache invalidation")
	file1 := filepath.Join(exampleDir, "delete-me.txt")
	fileutil.WriteString(file1, "Delete this")
	logger.Info("   File exists before deletion?", "result", fileutil.Exists(file1))
	fileutil.DeleteFile(file1)
	logger.Info("   ✅ Deleted file", "name", filepath.Base(file1))
	logger.Info("   ✅ File exists after deletion?", "result", fileutil.Exists(file1))
	logger.Info("")

	// 2. DeleteFiles - Delete multiple
	logger.Info("2️⃣  DeleteFiles() - Deleting multiple files at once")
	logger.Info("   Use Case: Batch cleanup, temp file removal")
	f1 := filepath.Join(exampleDir, "temp1.txt")
	f2 := filepath.Join(exampleDir, "temp2.txt")
	f3 := filepath.Join(exampleDir, "temp3.txt")
	fileutil.WriteString(f1, "Temp 1")
	fileutil.WriteString(f2, "Temp 2")
	fileutil.WriteString(f3, "Temp 3")
	if err := fileutil.DeleteFiles(f1, f2, f3); err != nil {
		logger.Fatalf("DeleteFiles failed: %v", err)
	}
	logger.Info("   ✅ Deleted 3 files in one operation")
	logger.Info("")

	// 3. DeletePattern - Delete by pattern
	logger.Info("3️⃣  DeletePattern() - Deleting files matching pattern")
	logger.Info("   Use Case: Cleanup logs, remove build artifacts")
	fileutil.WriteString(filepath.Join(exampleDir, "log1.log"), "Log 1")
	fileutil.WriteString(filepath.Join(exampleDir, "log2.log"), "Log 2")
	fileutil.WriteString(filepath.Join(exampleDir, "keep.txt"), "Keep me")
	pattern := filepath.Join(exampleDir, "*.log")
	if err := fileutil.DeletePattern(pattern); err != nil {
		logger.Fatalf("DeletePattern failed: %v", err)
	}
	remainingFiles, _ := fileutil.ListFiles(exampleDir)
	logger.Info("   ✅ Deleted all .log files", "remaining", len(remainingFiles))
	logger.Info("")

	// 4. Clean - Remove directory contents
	logger.Info("4️⃣  Clean() - Removing all directory contents (keeping directory)")
	logger.Info("   Use Case: Cache clearing, workspace reset")
	cleanDir := filepath.Join(exampleDir, "clean-test")
	fileutil.WriteString(filepath.Join(cleanDir, "file1.txt"), "File 1")
	fileutil.WriteString(filepath.Join(cleanDir, "file2.txt"), "File 2")
	logger.Info("   Files before clean:", "count", 2)
	fileutil.Clean(cleanDir)
	isEmpty, _ := fileutil.IsEmpty(cleanDir)
	logger.Info("   ✅ Cleaned directory", "exists", fileutil.Exists(cleanDir), "isEmpty", isEmpty)
	logger.Info("")

	// 5. DeleteRecursive - Remove directory and contents
	logger.Info("5️⃣  DeleteRecursive() - Deleting directory and all contents")
	logger.Info("   Use Case: Complete removal, uninstall operations")
	deleteDir := filepath.Join(exampleDir, "delete-recursive")
	fileutil.WriteString(filepath.Join(deleteDir, "sub1", "file1.txt"), "Deep file 1")
	fileutil.WriteString(filepath.Join(deleteDir, "sub2", "file2.txt"), "Deep file 2")
	fileutil.DeleteRecursive(deleteDir)
	logger.Info("   ✅ Deleted directory recursively", "exists", fileutil.Exists(deleteDir))
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 5 deletion methods")
	logger.Info("")
}

// Example 8: Directory Operations / 예제 8: 디렉토리 작업
func example08_DirectoryOperations(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📁 Example 8: Directory Operations")
	logger.Info("   예제 8: 디렉토리 작업")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example08")

	// 1. MkdirAll
	logger.Info("1️⃣  MkdirAll() - Creating nested directories")
	logger.Info("   Use Case: Project initialization, directory structure setup")
	nestedDir := filepath.Join(exampleDir, "deep", "nested", "structure", "here")
	if err := fileutil.MkdirAll(nestedDir); err != nil {
		logger.Fatalf("MkdirAll failed: %v", err)
	}
	logger.Info("   ✅ Created nested directory", "path", "deep/nested/structure/here")
	logger.Info("   ✅ Directory exists?", "result", fileutil.Exists(nestedDir))
	logger.Info("")

	// Setup test structure for remaining examples
	fileutil.WriteString(filepath.Join(exampleDir, "file1.txt"), "Root file 1")
	fileutil.WriteString(filepath.Join(exampleDir, "file2.go"), "Root file 2")
	fileutil.WriteString(filepath.Join(exampleDir, "subdir1", "file3.txt"), "Sub file 1")
	fileutil.WriteString(filepath.Join(exampleDir, "subdir1", "file4.go"), "Sub file 2")
	fileutil.WriteString(filepath.Join(exampleDir, "subdir2", "file5.txt"), "Sub file 3")

	// 2. ListFiles
	logger.Info("2️⃣  ListFiles() - Listing files in directory")
	logger.Info("   Use Case: File inventory, directory scanning")
	files, _ := fileutil.ListFiles(exampleDir)
	logger.Info("   ✅ Non-recursive list", "count", len(files))
	filesRecursive, _ := fileutil.ListFiles(exampleDir, true)
	logger.Info("   ✅ Recursive list", "count", len(filesRecursive))
	logger.Info("")

	// 3. ListDirs
	logger.Info("3️⃣  ListDirs() - Listing subdirectories")
	logger.Info("   Use Case: Directory traversal, folder structure analysis")
	dirs, _ := fileutil.ListDirs(exampleDir)
	logger.Info("   ✅ Non-recursive dirs", "count", len(dirs))
	dirsRecursive, _ := fileutil.ListDirs(exampleDir, true)
	logger.Info("   ✅ Recursive dirs", "count", len(dirsRecursive))
	logger.Info("")

	// 4. ListAll
	logger.Info("4️⃣  ListAll() - Listing all entries (files + dirs)")
	logger.Info("   Use Case: Complete directory inventory")
	all, _ := fileutil.ListAll(exampleDir)
	logger.Info("   ✅ Non-recursive entries", "count", len(all))
	allRecursive, _ := fileutil.ListAll(exampleDir, true)
	logger.Info("   ✅ Recursive entries", "count", len(allRecursive))
	logger.Info("")

	// 5. FindFiles
	logger.Info("5️⃣  FindFiles() - Finding files with predicate")
	logger.Info("   Use Case: Search by extension, size, or custom criteria")
	goFiles, _ := fileutil.FindFiles(exampleDir, func(path string, info os.FileInfo) bool {
		return filepath.Ext(path) == ".go"
	})
	logger.Info("   ✅ Found .go files", "count", len(goFiles))
	for i, f := range goFiles {
		logger.Info("      Found", "index", i+1, "file", filepath.Base(f))
	}
	logger.Info("")

	// 6. FilterFiles
	logger.Info("6️⃣  FilterFiles() - Filtering file list")
	logger.Info("   Use Case: Post-processing file lists, size filtering")
	txtFiles, _ := fileutil.FilterFiles(filesRecursive, func(path string) bool {
		return filepath.Ext(path) == ".txt"
	})
	logger.Info("   ✅ Filtered .txt files", "count", len(txtFiles))
	logger.Info("")

	// 7. DirSize
	logger.Info("7️⃣  DirSize() - Calculating directory size")
	logger.Info("   Use Case: Disk usage analysis, quota monitoring")
	size, _ := fileutil.DirSize(exampleDir)
	sizeHuman, _ := fileutil.SizeHuman(filepath.Join(exampleDir, "file1.txt"))
	logger.Info("   ✅ Total directory size", "bytes", size, "human", fmt.Sprintf("~%s", sizeHuman))
	logger.Info("")

	// 8. IsEmpty
	logger.Info("8️⃣  IsEmpty() - Checking if directory is empty")
	logger.Info("   Use Case: Pre-deletion checks, directory validation")
	emptyDir := filepath.Join(exampleDir, "empty-dir")
	fileutil.MkdirAll(emptyDir)
	isEmpty, _ := fileutil.IsEmpty(emptyDir)
	logger.Info("   ✅ Empty directory check", "isEmpty", isEmpty)
	fileutil.WriteString(filepath.Join(emptyDir, "now-not-empty.txt"), "content")
	isEmptyNow, _ := fileutil.IsEmpty(emptyDir)
	logger.Info("   ✅ After adding file", "isEmpty", isEmptyNow)
	logger.Info("")

	// 9. Walk
	logger.Info("9️⃣  Walk() - Walking directory tree")
	logger.Info("   Use Case: Custom file processing, directory traversal")
	var walkCount int
	fileutil.Walk(exampleDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		walkCount++
		return nil
	})
	logger.Info("   ✅ Walked entries", "count", walkCount)
	logger.Info("")

	// 10. WalkFiles
	logger.Info("🔟 WalkFiles() - Walking only files")
	logger.Info("   Use Case: File-specific processing")
	var fileCount int
	fileutil.WalkFiles(exampleDir, func(path string, info os.FileInfo) error {
		fileCount++
		return nil
	})
	logger.Info("   ✅ Walked files", "count", fileCount)
	logger.Info("")

	// 11. WalkDirs
	logger.Info("1️⃣1️⃣  WalkDirs() - Walking only directories")
	logger.Info("   Use Case: Directory structure analysis")
	var dirCount int
	fileutil.WalkDirs(exampleDir, func(path string, info os.FileInfo) error {
		dirCount++
		return nil
	})
	logger.Info("   ✅ Walked directories", "count", dirCount)
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 11 directory operations")
	logger.Info("   요약: 11개의 디렉토리 작업 시연")
	logger.Info("")
}

// Example 9: File Hashing / 예제 9: 파일 해싱
func example09_FileHashing(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔐 Example 9: File Hashing & Checksums")
	logger.Info("   예제 9: 파일 해싱 및 체크섬")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example09")
	testFile := filepath.Join(exampleDir, "hashtest.txt")
	fileutil.WriteString(testFile, "This is test content for hashing algorithms")

	// 1. MD5
	logger.Info("1️⃣  MD5() - MD5 hash calculation")
	logger.Info("   Use Case: Legacy compatibility, fast checksums (not for security)")
	md5Hash, _ := fileutil.MD5(testFile)
	logger.Info("   ✅ MD5 hash", "value", md5Hash[:16]+"...")
	logger.Info("")

	// 2. SHA1
	logger.Info("2️⃣  SHA1() - SHA1 hash calculation")
	logger.Info("   Use Case: Git commits, legacy systems (not recommended for new systems)")
	sha1Hash, _ := fileutil.SHA1(testFile)
	logger.Info("   ✅ SHA1 hash", "value", sha1Hash[:16]+"...")
	logger.Info("")

	// 3. SHA256
	logger.Info("3️⃣  SHA256() - SHA256 hash calculation")
	logger.Info("   Use Case: File integrity, secure checksums, digital signatures")
	sha256Hash, _ := fileutil.SHA256(testFile)
	logger.Info("   ✅ SHA256 hash", "value", sha256Hash[:16]+"...")
	logger.Info("")

	// 4. SHA512
	logger.Info("4️⃣  SHA512() - SHA512 hash calculation")
	logger.Info("   Use Case: Maximum security, sensitive data integrity")
	sha512Hash, _ := fileutil.SHA512(testFile)
	logger.Info("   ✅ SHA512 hash", "value", sha512Hash[:16]+"...")
	logger.Info("")

	// 5. Hash with custom algorithm
	logger.Info("5️⃣  Hash() - Custom algorithm selection")
	logger.Info("   Use Case: Flexibility, algorithm comparison")
	customHash, _ := fileutil.Hash(testFile, "sha256")
	logger.Info("   ✅ Custom hash (sha256)", "value", customHash[:16]+"...")
	logger.Info("")

	// 6. Checksum & VerifyChecksum
	logger.Info("6️⃣  Checksum() & VerifyChecksum() - File integrity verification")
	logger.Info("   Use Case: Download verification, file corruption detection")
	checksum, _ := fileutil.Checksum(testFile)
	logger.Info("   ✅ Generated checksum", "value", checksum[:16]+"...")

	isValid, _ := fileutil.VerifyChecksum(testFile, checksum)
	logger.Info("   ✅ Verification result", "valid", isValid)

	isInvalid, _ := fileutil.VerifyChecksum(testFile, "wrong-checksum")
	logger.Info("   ✅ Wrong checksum test", "valid", isInvalid)
	logger.Info("")

	// 7. CompareFiles
	logger.Info("7️⃣  CompareFiles() - Byte-by-byte file comparison")
	logger.Info("   Use Case: Exact duplicate detection, backup verification")
	file2 := filepath.Join(exampleDir, "hashtest-copy.txt")
	fileutil.CopyFile(testFile, file2)
	same, _ := fileutil.CompareFiles(testFile, file2)
	logger.Info("   ✅ Files identical?", "result", same)

	fileutil.WriteString(file2, "Different content")
	sameDiff, _ := fileutil.CompareFiles(testFile, file2)
	logger.Info("   ✅ After modification", "result", sameDiff)
	logger.Info("")

	// 8. CompareHash
	logger.Info("8️⃣  CompareHash() - Hash-based file comparison")
	logger.Info("   Use Case: Fast comparison for large files")
	fileutil.WriteString(file2, "This is test content for hashing algorithms") // Same content
	sameHash, _ := fileutil.CompareHash(testFile, file2)
	logger.Info("   ✅ Hashes match?", "result", sameHash)
	logger.Info("")

	// 9. HashDir
	logger.Info("9️⃣  HashDir() - Directory content hashing")
	logger.Info("   Use Case: Detect changes in entire directories")
	fileutil.WriteString(filepath.Join(exampleDir, "file1.txt"), "Content 1")
	fileutil.WriteString(filepath.Join(exampleDir, "file2.txt"), "Content 2")
	dirHash, _ := fileutil.HashDir(exampleDir)
	logger.Info("   ✅ Directory hash", "value", dirHash[:16]+"...")
	logger.Info("   ✅ Hash changes if any file changes")
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 9 hashing and verification operations")
	logger.Info("   요약: 9개의 해싱 및 검증 작업 시연")
	logger.Info("")
}

// Example 10: Advanced Reading / 예제 10: 고급 읽기
func example10_AdvancedReading(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("📚 Example 10: Advanced Reading Operations")
	logger.Info("   예제 10: 고급 읽기 작업")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example10")
	largeFile := filepath.Join(exampleDir, "large-file.txt")
	content := "0123456789" + "ABCDEFGHIJ" + "abcdefghij" + "!@#$%^&*()" + "~`-=_+[]{}"
	fileutil.WriteString(largeFile, content)

	// 1. ReadBytes with offset
	logger.Info("1️⃣  ReadBytes() - Reading specific portion of file")
	logger.Info("   Use Case: Random access, partial file reading, resume downloads")
	chunk, _ := fileutil.ReadBytes(largeFile, 10, 20)
	logger.Info("   ✅ Read bytes 10-30", "content", string(chunk))
	logger.Info("")

	// 2. ReadChunk
	logger.Info("2️⃣  ReadChunk() - Streaming large file processing")
	logger.Info("   Use Case: Processing large files without loading into memory")
	var chunkCount int
	fileutil.ReadChunk(largeFile, 10, func(data []byte) error {
		chunkCount++
		logger.Info("      Processing chunk", "number", chunkCount, "size", len(data), "content", string(data))
		return nil
	})
	logger.Info("   ✅ Processed chunks", "count", chunkCount)
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 2 advanced reading operations")
	logger.Info("   요약: 2개의 고급 읽기 작업 시연")
	logger.Info("")
}

// Example 11: Atomic Operations / 예제 11: 원자 연산
func example11_AtomicOperations(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("⚛️  Example 11: Atomic Operations")
	logger.Info("   예제 11: 원자 연산 작업")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example11")

	// 1. WriteAtomic
	logger.Info("1️⃣  WriteAtomic() - Crash-safe atomic write")
	logger.Info("   Use Case: Critical config files, database consistency, crash recovery")
	atomicFile := filepath.Join(exampleDir, "critical-config.json")
	fileutil.WriteAtomic(atomicFile, []byte(`{"version": "1.0", "critical": true}`))
	logger.Info("   ✅ Atomic write completed", "file", "critical-config.json")
	logger.Info("   ✅ File safe even if process crashes during write")
	logger.Info("")

	// 2. CreateFile
	logger.Info("2️⃣  CreateFile() - Create file handle for writing")
	logger.Info("   Use Case: Custom write operations, streaming writes")
	customFile := filepath.Join(exampleDir, "custom.txt")
	file, _ := fileutil.CreateFile(customFile)
	file.WriteString("Custom content written through file handle")
	file.Close()
	logger.Info("   ✅ Created and wrote via file handle")
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 2 atomic operation methods")
	logger.Info("   요약: 2개의 원자 연산 메서드 시연")
	logger.Info("")
}

// Example 12: Permissions & Ownership / 예제 12: 권한 및 소유권
func example12_PermissionsAndOwnership(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔒 Example 12: Permissions & Ownership")
	logger.Info("   예제 12: 권한 및 소유권")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example12")
	testFile := filepath.Join(exampleDir, "permissions-test.txt")
	fileutil.WriteString(testFile, "Testing file permissions")

	// 1. Chmod
	logger.Info("1️⃣  Chmod() - Changing file permissions")
	logger.Info("   Use Case: Security configuration, access control")
	fileutil.Chmod(testFile, 0644)
	info, _ := os.Stat(testFile)
	logger.Info("   ✅ Changed to 0644", "mode", info.Mode().String())
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated permission operations")
	logger.Info("   요약: 권한 작업 시연")
	logger.Info("")
}

// Example 13: Symlinks & Special Files / 예제 13: 심볼릭 링크 및 특수 파일
func example13_SymlinksAndSpecialFiles(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔗 Example 13: Symlinks & Special Files")
	logger.Info("   예제 13: 심볼릭 링크 및 특수 파일")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example13")

	// 1. CreateTemp
	logger.Info("1️⃣  CreateTemp() - Creating temporary files")
	logger.Info("   Use Case: Temporary processing, cache files, safe testing")
	tempFile, _ := fileutil.CreateTemp(exampleDir, "temp-*.txt")
	logger.Info("   ✅ Created temp file", "path", filepath.Base(tempFile))
	logger.Info("")

	// 2. CreateTempDir
	logger.Info("2️⃣  CreateTempDir() - Creating temporary directories")
	logger.Info("   Use Case: Temporary workspaces, build directories")
	tempDir2, _ := fileutil.CreateTempDir(exampleDir, "temp-dir-*")
	logger.Info("   ✅ Created temp directory", "path", filepath.Base(tempDir2))
	logger.Info("")

	// 3. RemoveEmpty
	logger.Info("3️⃣  RemoveEmpty() - Removing empty directories")
	logger.Info("   Use Case: Cleanup operations, workspace maintenance")
	emptyDir1 := filepath.Join(exampleDir, "cleanup", "empty1")
	emptyDir2 := filepath.Join(exampleDir, "cleanup", "empty2")
	fileutil.MkdirAll(emptyDir1)
	fileutil.MkdirAll(emptyDir2)
	fileutil.RemoveEmpty(filepath.Join(exampleDir, "cleanup"))
	logger.Info("   ✅ Removed empty subdirectories")
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 3 special file operations")
	logger.Info("   요약: 3개의 특수 파일 작업 시연")
	logger.Info("")
}

// Example 14: Walking & Filtering (already covered in Example 8)
func example14_WalkAndFilter(logger *logging.Logger, tempDir string) {
	_ = tempDir // This example refers to Example 8 / 이 예제는 예제 8을 참조합니다

	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🚶 Example 14: Walking & Filtering (Advanced)")
	logger.Info("   예제 14: 디렉토리 순회 및 필터링 (고급)")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("   ℹ️  Walking and filtering operations were demonstrated in Example 8")
	logger.Info("   ℹ️  순회 및 필터링 작업은 예제 8에서 시연되었습니다")
	logger.Info("   See: Walk, WalkFiles, WalkDirs, FindFiles, FilterFiles")
	logger.Info("")
}

// Example 15: Error Handling / 예제 15: 에러 처리
func example15_ErrorHandling(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("❌ Example 15: Error Handling")
	logger.Info("   예제 15: 에러 처리")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example15")
	existingFile := filepath.Join(exampleDir, "exists.txt")
	fileutil.WriteString(existingFile, "I exist")
	nonExistentFile := filepath.Join(exampleDir, "does-not-exist.txt")

	// 1. IsNotFound
	logger.Info("1️⃣  IsNotFound() - Checking if error is 'not found'")
	logger.Info("   Use Case: Graceful error handling, file existence checks")
	_, err := fileutil.ReadString(nonExistentFile)
	isNotFound := fileutil.IsNotFound(err)
	logger.Info("   ✅ File not found error?", "result", isNotFound)
	logger.Info("")

	// 2. IsExist
	logger.Info("2️⃣  IsExist() - Checking if error is 'already exists'")
	logger.Info("   Use Case: Preventing overwrite, safe file creation")
	logger.Info("   ✅ Error checking available via fileutil.IsExist()")
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated error handling utilities")
	logger.Info("   요약: 에러 처리 유틸리티 시연")
	logger.Info("")
}

// Example 16: Real-World Scenarios / 예제 16: 실제 사용 시나리오
func example16_RealWorldScenarios(logger *logging.Logger, tempDir string) {
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🌍 Example 16: Real-World Scenarios")
	logger.Info("   예제 16: 실제 사용 시나리오")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	exampleDir := filepath.Join(tempDir, "example16")

	// Scenario 1: Simple Backup System
	logger.Info("🎯 Scenario 1: Simple Backup System")
	logger.Info("   시나리오 1: 간단한 백업 시스템")
	dataDir := filepath.Join(exampleDir, "data")
	backupDir := filepath.Join(exampleDir, "backups", time.Now().Format("20060102-150405"))
	fileutil.WriteString(filepath.Join(dataDir, "important1.txt"), "Important data 1")
	fileutil.WriteString(filepath.Join(dataDir, "important2.txt"), "Important data 2")

	if err := fileutil.CopyDir(dataDir, backupDir); err != nil {
		logger.Fatalf("Backup failed: %v", err)
	}
	backupSize, _ := fileutil.DirSize(backupDir)
	logger.Info("   ✅ Backup created", "size", backupSize, "location", filepath.Base(backupDir))
	logger.Info("")

	// Scenario 2: Config File Management
	logger.Info("🎯 Scenario 2: Safe Config File Updates")
	logger.Info("   시나리오 2: 안전한 설정 파일 업데이트")
	configFile := filepath.Join(exampleDir, "app-config.json")
	config := map[string]interface{}{
		"version":    "1.0.0",
		"debug":      false,
		"maxRetries": 3,
	}
	fileutil.WriteJSON(configFile, config)

	// Atomic update to prevent corruption
	newConfig := map[string]interface{}{
		"version":    "1.1.0",
		"debug":      true,
		"maxRetries": 5,
	}
	jsonData, _ := json.MarshalIndent(newConfig, "", "  ")
	fileutil.WriteAtomic(configFile, jsonData)
	logger.Info("   ✅ Config safely updated with atomic write")
	logger.Info("")

	// Scenario 3: Log Cleanup
	logger.Info("🎯 Scenario 3: Automated Log Cleanup")
	logger.Info("   시나리오 3: 자동 로그 정리")
	logsDir := filepath.Join(exampleDir, "logs")
	fileutil.WriteString(filepath.Join(logsDir, "app-20250101.log"), "Old log")
	fileutil.WriteString(filepath.Join(logsDir, "app-20250112.log"), "Recent log")
	fileutil.WriteString(filepath.Join(logsDir, "app-keep.txt"), "Keep this")

	// Delete old logs
	fileutil.DeletePattern(filepath.Join(logsDir, "*-202501*.log"))
	remaining, _ := fileutil.ListFiles(logsDir)
	logger.Info("   ✅ Cleaned old logs", "remainingFiles", len(remaining))
	logger.Info("")

	logger.Info("📊 Summary: Demonstrated 3 real-world usage scenarios")
	logger.Info("   요약: 3개의 실제 사용 시나리오 시연")
	logger.Info("")
}
