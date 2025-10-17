package validation

import (
	"fmt"
	"os"
	"path/filepath"
)

// FilePath validates that the value is a valid file path format.
// Checks path format validity without verifying actual file existence.
//
// FilePath는 값이 유효한 파일 경로 형식인지 검증합니다.
// 실제 파일 존재 여부를 확인하지 않고 경로 형식의 유효성만 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Normalizes path using filepath.Clean
//     filepath.Clean을 사용하여 경로 정규화
//   - Accepts both absolute and relative paths
//     절대 경로와 상대 경로 모두 허용
//   - Rejects empty paths and "." (current directory)
//     빈 경로와 "." (현재 디렉토리) 거부
//   - Does not check file existence
//     파일 존재 여부 확인 안 함
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - File path format validation / 파일 경로 형식 검증
//   - Configuration path validation / 구성 경로 검증
//   - Input path validation / 입력 경로 검증
//   - Path sanitization check / 경로 정제 확인
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = path length
//     시간 복잡도: O(n), n = 경로 길이
//   - No filesystem I/O operations
//     파일시스템 I/O 작업 없음
//
// Example / 예제:
//
//	// Valid file paths / 유효한 파일 경로
//	v := validation.New("/etc/config.yaml", "config_path")
//	v.FilePath()  // Passes (absolute path)
//
//	v = validation.New("./data/file.txt", "data_path")
//	v.FilePath()  // Passes (relative path)
//
//	// Invalid paths / 무효한 경로
//	v = validation.New("", "empty_path")
//	v.FilePath()  // Fails (empty string)
//
//	v = validation.New(".", "current_dir")
//	v.FilePath()  // Fails (current directory marker)
func (v *Validator) FilePath() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_path", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Clean the path to normalize it
	cleaned := filepath.Clean(str)

	// Check if path is absolute or relative
	if cleaned == "" || cleaned == "." {
		v.addError("file_path", fmt.Sprintf("%s must be a valid file path / %s은(는) 유효한 파일 경로여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// FileExists validates that a file or directory exists at the given path.
// Performs filesystem existence check using os.Stat.
//
// FileExists는 주어진 경로에 파일이나 디렉토리가 존재하는지 검증합니다.
// os.Stat을 사용하여 파일시스템 존재 여부를 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value as file path)
//     없음 (validator의 값을 파일 경로로 사용)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses os.Stat to check existence
//     os.Stat을 사용하여 존재 확인
//   - Accepts both files and directories
//     파일과 디렉토리 모두 허용
//   - Follows symlinks (os.Stat behavior)
//     심볼릭 링크 추적 (os.Stat 동작)
//   - Fails if path does not exist
//     경로가 존재하지 않으면 실패
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Configuration file validation / 구성 파일 검증
//   - Input file validation / 입력 파일 검증
//   - Resource existence check / 리소스 존재 확인
//   - Dependency validation / 의존성 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//   - Note: Filesystem state may change between check and use (TOCTOU)
//     참고: 확인과 사용 사이에 파일시스템 상태가 변경될 수 있음 (TOCTOU)
//
// Performance / 성능:
//   - Time complexity: O(1) filesystem operation
//     시간 복잡도: O(1) 파일시스템 작업
//   - Performs actual filesystem I/O
//     실제 파일시스템 I/O 수행
//
// Example / 예제:
//
//	// Existing file / 존재하는 파일
//	v := validation.New("/etc/hosts", "hosts_file")
//	v.FileExists()  // Passes if file exists
//
//	// Existing directory / 존재하는 디렉토리
//	v = validation.New("/usr/local", "local_dir")
//	v.FileExists()  // Passes (directories accepted)
//
//	// Non-existing file / 존재하지 않는 파일
//	v = validation.New("/nonexistent/file.txt", "missing_file")
//	v.FileExists()  // Fails
func (v *Validator) FileExists() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_exists", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if _, err := os.Stat(str); os.IsNotExist(err) {
		v.addError("file_exists", fmt.Sprintf("%s must be an existing file or directory / %s은(는) 존재하는 파일 또는 디렉토리여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// FileReadable validates that a file can be opened for reading.
// Attempts to open the file and verifies read permissions.
//
// FileReadable은 파일을 읽기 위해 열 수 있는지 검증합니다.
// 파일을 열어보고 읽기 권한을 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value as file path)
//     없음 (validator의 값을 파일 경로로 사용)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses os.Open to verify read access
//     os.Open을 사용하여 읽기 접근 확인
//   - Closes file immediately after check
//     확인 후 즉시 파일 닫기
//   - Checks both existence and permissions
//     존재와 권한 모두 확인
//   - Fails if file cannot be opened
//     파일을 열 수 없으면 실패
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Input file validation / 입력 파일 검증
//   - Configuration file access check / 구성 파일 접근 확인
//   - Data file validation / 데이터 파일 검증
//   - Permission validation / 권한 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//   - Note: File permissions may change between check and actual use
//     참고: 확인과 실제 사용 사이에 파일 권한이 변경될 수 있음
//
// Performance / 성능:
//   - Time complexity: O(1) filesystem operation
//     시간 복잡도: O(1) 파일시스템 작업
//   - Opens and closes file (I/O overhead)
//     파일 열기 및 닫기 (I/O 오버헤드)
//
// Example / 예제:
//
//	// Readable file / 읽기 가능한 파일
//	v := validation.New("/etc/hosts", "hosts_file")
//	v.FileReadable()  // Passes if readable
//
//	// Permission denied / 권한 거부
//	v = validation.New("/root/secret.txt", "secret_file")
//	v.FileReadable()  // Fails if no read permission
//
//	// Directory (not readable as file) / 디렉토리 (파일로 읽을 수 없음)
//	v = validation.New("/usr", "usr_dir")
//	v.FileReadable()  // May pass (os.Open can open directories)
func (v *Validator) FileReadable() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_readable", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Try to open the file for reading
	file, err := os.Open(str)
	if err != nil {
		v.addError("file_readable", fmt.Sprintf("%s must be a readable file / %s은(는) 읽기 가능한 파일이어야 합니다", v.fieldName, v.fieldName))
		return v
	}
	defer file.Close()

	return v
}

// FileWritable validates that a file is writable or its parent directory allows file creation.
// For existing files, checks write permissions. For new files, tests parent directory writability.
//
// FileWritable은 파일이 쓰기 가능하거나 부모 디렉토리가 파일 생성을 허용하는지 검증합니다.
// 기존 파일의 경우 쓰기 권한을 확인합니다. 새 파일의 경우 부모 디렉토리의 쓰기 가능성을 테스트합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value as file path)
//     없음 (validator의 값을 파일 경로로 사용)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - For existing files: Opens with O_WRONLY|O_APPEND flags
//     기존 파일: O_WRONLY|O_APPEND 플래그로 열기
//   - For new files: Tests parent directory by creating temp file
//     새 파일: 임시 파일을 생성하여 부모 디렉토리 테스트
//   - Cleans up temporary test file immediately
//     임시 테스트 파일 즉시 정리
//   - Fails if file/directory is not writable
//     파일/디렉토리가 쓰기 불가능하면 실패
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Output file validation / 출력 파일 검증
//   - Log file validation / 로그 파일 검증
//   - Configuration file update check / 구성 파일 업데이트 확인
//   - Data export validation / 데이터 내보내기 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//   - Note: Permissions may change between check and write
//     참고: 확인과 쓰기 사이에 권한이 변경될 수 있음
//
// Performance / 성능:
//   - Time complexity: O(1) filesystem operations
//     시간 복잡도: O(1) 파일시스템 작업
//   - Creates and deletes temp file for new files (I/O overhead)
//     새 파일의 경우 임시 파일 생성 및 삭제 (I/O 오버헤드)
//
// Example / 예제:
//
//	// Existing writable file / 기존 쓰기 가능한 파일
//	v := validation.New("/tmp/output.txt", "output_file")
//	v.FileWritable()  // Passes if writable
//
//	// New file in writable directory / 쓰기 가능한 디렉토리의 새 파일
//	v = validation.New("/tmp/newfile.log", "log_file")
//	v.FileWritable()  // Passes if /tmp is writable
//
//	// Read-only file / 읽기 전용 파일
//	v = validation.New("/etc/hosts", "hosts_file")
//	v.FileWritable()  // Fails (permission denied)
func (v *Validator) FileWritable() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_writable", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if file exists
	_, err := os.Stat(str)
	if err == nil {
		// File exists, check if we can write to it
		file, err := os.OpenFile(str, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			v.addError("file_writable", fmt.Sprintf("%s must be a writable file / %s은(는) 쓰기 가능한 파일이어야 합니다", v.fieldName, v.fieldName))
			return v
		}
		defer file.Close()
	} else if os.IsNotExist(err) {
		// File doesn't exist, check if parent directory is writable
		dir := filepath.Dir(str)
		dirInfo, err := os.Stat(dir)
		if err != nil || !dirInfo.IsDir() {
			v.addError("file_writable", fmt.Sprintf("%s parent directory must exist and be writable / %s의 부모 디렉토리가 존재하고 쓰기 가능해야 합니다", v.fieldName, v.fieldName))
			return v
		}

		// Try to create a temporary file in the directory to test writability
		tmpFile, err := os.CreateTemp(dir, ".write_test_*")
		if err != nil {
			v.addError("file_writable", fmt.Sprintf("%s parent directory must be writable / %s의 부모 디렉토리가 쓰기 가능해야 합니다", v.fieldName, v.fieldName))
			return v
		}
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	} else {
		// Other error (permission, etc.)
		// Note: os.Stat returns nil info when error occurs, so this is a catch-all for any stat errors
		// 참고: os.Stat은 에러 발생 시 nil info를 반환하므로, 이는 모든 stat 에러에 대한 포괄적 처리입니다
		v.addError("file_writable", fmt.Sprintf("%s must be accessible / %s에 접근할 수 있어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// FileSize validates that a file's size is within the specified byte range (inclusive).
// Checks both file existence and size constraints.
//
// FileSize는 파일 크기가 지정된 바이트 범위 내에 있는지 검증합니다 (포함).
// 파일 존재와 크기 제약을 모두 확인합니다.
//
// Parameters / 매개변수:
//   - min: Minimum file size in bytes (inclusive)
//     최소 파일 크기 (바이트, 포함)
//   - max: Maximum file size in bytes (inclusive)
//     최대 파일 크기 (바이트, 포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses os.Stat to get file info
//     os.Stat을 사용하여 파일 정보 가져오기
//   - Accepts sizes within [min, max] range (inclusive)
//     [min, max] 범위의 크기 허용 (포함)
//   - Rejects directories (only files)
//     디렉토리 거부 (파일만 허용)
//   - Fails if file does not exist
//     파일이 존재하지 않으면 실패
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Upload file size validation / 업로드 파일 크기 검증
//   - Configuration file size check / 구성 파일 크기 확인
//   - Data file validation / 데이터 파일 검증
//   - Resource quota enforcement / 리소스 할당량 적용
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//   - Note: File size may change between check and use
//     참고: 확인과 사용 사이에 파일 크기가 변경될 수 있음
//
// Performance / 성능:
//   - Time complexity: O(1) filesystem operation
//     시간 복잡도: O(1) 파일시스템 작업
//   - Single os.Stat call
//     단일 os.Stat 호출
//
// Example / 예제:
//
//	// Size constraints (1 KB - 10 MB) / 크기 제약 (1 KB - 10 MB)
//	v := validation.New("/tmp/upload.pdf", "upload")
//	v.FileSize(1024, 10*1024*1024)  // 1 KB min, 10 MB max
//
//	// Empty file allowed / 빈 파일 허용
//	v = validation.New("/tmp/empty.txt", "empty_file")
//	v.FileSize(0, 100)  // 0 - 100 bytes
//
//	// Exact size / 정확한 크기
//	v = validation.New("/tmp/data.bin", "data")
//	v.FileSize(1024, 1024)  // Must be exactly 1 KB
func (v *Validator) FileSize(min, max int64) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_size", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	info, err := os.Stat(str)
	if err != nil {
		v.addError("file_size", fmt.Sprintf("%s must be an existing file / %s은(는) 존재하는 파일이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if info.IsDir() {
		v.addError("file_size", fmt.Sprintf("%s must be a file, not a directory / %s은(는) 디렉토리가 아닌 파일이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	size := info.Size()
	if size < min || size > max {
		v.addError("file_size", fmt.Sprintf("%s file size must be between %d and %d bytes / %s 파일 크기는 %d와 %d 바이트 사이여야 합니다", v.fieldName, min, max, v.fieldName, min, max))
		return v
	}

	return v
}

// FileExtension validates that a file has one of the specified extensions.
// Case-sensitive extension matching with automatic dot prefix handling.
//
// FileExtension은 파일이 지정된 확장자 중 하나를 가지고 있는지 검증합니다.
// 자동 점 접두사 처리를 포함한 대소문자 구분 확장자 매칭입니다.
//
// Parameters / 매개변수:
//   - extensions: Allowed file extensions (with or without leading dot)
//     허용된 파일 확장자 (점 포함 또는 제외)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses filepath.Ext to extract extension
//     filepath.Ext를 사용하여 확장자 추출
//   - Automatically adds dot prefix if missing
//     점 접두사가 없으면 자동 추가
//   - Case-sensitive matching (.txt != .TXT)
//     대소문자 구분 매칭 (.txt != .TXT)
//   - Fails if file has no extension
//     파일에 확장자가 없으면 실패
//   - Fails if extension not in allowed list
//     확장자가 허용 목록에 없으면 실패
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Upload file type validation / 업로드 파일 유형 검증
//   - Configuration file format check / 구성 파일 형식 확인
//   - Input file type validation / 입력 파일 유형 검증
//   - Security whitelisting / 보안 화이트리스트
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n*m), n = extensions count, m = avg extension length
//     시간 복잡도: O(n*m), n = 확장자 개수, m = 평균 확장자 길이
//   - No filesystem I/O
//     파일시스템 I/O 없음
//
// Example / 예제:
//
//	// Image files / 이미지 파일
//	v := validation.New("/tmp/photo.jpg", "photo")
//	v.FileExtension("jpg", "jpeg", "png", "gif")  // Passes
//
//	// With or without dot / 점 포함 또는 제외
//	v = validation.New("/tmp/data.json", "data")
//	v.FileExtension(".json", ".xml")  // Passes (dot prefix)
//	v.FileExtension("json", "xml")    // Passes (auto-added dot)
//
//	// Case-sensitive / 대소문자 구분
//	v = validation.New("/tmp/file.TXT", "file")
//	v.FileExtension("txt")  // Fails (case mismatch)
//
//	// No extension / 확장자 없음
//	v = validation.New("/tmp/Makefile", "makefile")
//	v.FileExtension("txt", "md")  // Fails (no extension)
func (v *Validator) FileExtension(extensions ...string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_extension", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	ext := filepath.Ext(str)
	if ext == "" {
		v.addError("file_extension", fmt.Sprintf("%s must have a file extension / %s은(는) 파일 확장자가 있어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if extension matches any of the allowed extensions
	found := false
	for _, allowedExt := range extensions {
		// Ensure extension starts with a dot
		if allowedExt != "" && allowedExt[0] != '.' {
			allowedExt = "." + allowedExt
		}
		if ext == allowedExt {
			found = true
			break
		}
	}

	if !found {
		v.addError("file_extension", fmt.Sprintf("%s must have one of the allowed extensions / %s은(는) 허용된 확장자 중 하나여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}
