# fileutil - Extreme Simplicity File and Path Utilities / 극도로 간단한 파일 및 경로 유틸리티

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](../LICENSE)
[![Version](https://img.shields.io/badge/version-v1.9.001-blue.svg)](https://github.com/arkd0ng/go-utils)

## Overview / 개요

The `fileutil` package provides extreme simplicity file and path utilities for Golang, reducing 20+ lines of repetitive file manipulation code to just 1-2 lines.

`fileutil` 패키지는 Golang을 위한 극도로 간단한 파일 및 경로 유틸리티를 제공하며, 20줄 이상의 반복적인 파일 조작 코드를 단 1-2줄로 줄입니다.

### Key Features / 주요 기능

- **Safe file operations with automatic directory creation** / 자동 디렉토리 생성을 사용한 안전한 파일 작업
- **Cross-platform compatibility** / 크로스 플랫폼 호환성
- **Comprehensive file/directory operations** / 포괄적인 파일/디렉토리 작업
- **Path manipulation and validation** / 경로 조작 및 검증
- **File hashing and checksums** / 파일 해싱 및 체크섬
- **Progress callbacks for long operations** / 긴 작업을 위한 진행 상황 콜백
- **Atomic file operations** / 원자적 파일 작업
- **Zero external dependencies** / 외부 의존성 없음 (yaml 제외)

### Design Philosophy / 설계 철학

**"20줄 → 1줄" (20 lines → 1 line)**

**Before / 이전:**
```go
// 20+ lines of code
dir := filepath.Dir(path)
if err := os.MkdirAll(dir, 0755); err != nil {
    return err
}
file, err := os.Create(path)
if err != nil {
    return err
}
defer file.Close()
if _, err := file.WriteString(content); err != nil {
    return err
}
// More boilerplate...
```

**After / 이후:**
```go
// 1 line of code
err := fileutil.WriteString("path/to/file.txt", "Hello, World!")
```

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/fileutil
```

## Quick Start / 빠른 시작

### Basic File Operations / 기본 파일 작업

```go
package main

import (
    "fmt"
    "log"

    "github.com/arkd0ng/go-utils/fileutil"
)

func main() {
    // Write file with automatic directory creation / 자동 디렉토리 생성과 함께 파일 쓰기
    err := fileutil.WriteString("path/to/file.txt", "Hello, World!")
    if err != nil {
        log.Fatal(err)
    }

    // Read file / 파일 읽기
    content, err := fileutil.ReadString("path/to/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(content) // Output: Hello, World!

    // Copy file / 파일 복사
    err = fileutil.CopyFile("source.txt", "destination.txt")
    if err != nil {
        log.Fatal(err)
    }
}
```

### Path Operations / 경로 작업

```go
// Join paths / 경로 결합
path := fileutil.Join("home", "user", "documents", "file.txt")
// Output: home/user/documents/file.txt

// Get base name / 기본 이름 가져오기
base := fileutil.Base(path)
// Output: file.txt

// Get extension / 확장자 가져오기
ext := fileutil.Ext(path)
// Output: .txt

// Change extension / 확장자 변경
newPath := fileutil.ChangeExt(path, ".md")
// Output: home/user/documents/file.md
```

### File Information / 파일 정보

```go
// Check if file exists / 파일 존재 확인
exists := fileutil.Exists("path/to/file.txt")

// Check if it's a file / 파일인지 확인
isFile := fileutil.IsFile("path/to/file.txt")

// Get file size / 파일 크기 가져오기
size, err := fileutil.Size("path/to/file.txt")

// Get human-readable size / 사람이 읽기 쉬운 크기 가져오기
sizeStr, err := fileutil.SizeHuman("path/to/file.txt")
// Output: "1.5 KB" or "2.3 MB"
```

### Directory Operations / 디렉토리 작업

```go
// Create directory / 디렉토리 생성
err := fileutil.MkdirAll("path/to/directory")

// List files / 파일 나열
files, err := fileutil.ListFiles("path/to/directory")

// Find .txt files recursively / .txt 파일 재귀적으로 찾기
txtFiles, err := fileutil.FindFiles(".", func(path string, info os.FileInfo) bool {
    return fileutil.Ext(path) == ".txt"
})

// Calculate directory size / 디렉토리 크기 계산
size, err := fileutil.DirSize("path/to/directory")
```

### File Hashing / 파일 해싱

```go
// Calculate MD5 hash / MD5 해시 계산
md5Hash, err := fileutil.MD5("path/to/file.txt")

// Calculate SHA256 hash / SHA256 해시 계산
sha256Hash, err := fileutil.SHA256("path/to/file.txt")

// Compare two files by hash / 두 파일을 해시로 비교
same, err := fileutil.CompareHash("file1.txt", "file2.txt")

// Verify checksum / 체크섬 검증
valid, err := fileutil.VerifyChecksum("file.txt", "expected-checksum")
```

### Copy with Progress / 진행 상황과 함께 복사

```go
// Copy with progress callback / 진행 상황 콜백과 함께 복사
err := fileutil.CopyFile("large-file.dat", "backup.dat",
    fileutil.WithProgress(func(written, total int64) {
        percent := float64(written) / float64(total) * 100
        fmt.Printf("\rProgress: %.1f%%", percent)
    }))
```

## Function Categories / 함수 카테고리

### 1. File Reading / 파일 읽기 (8 functions / 8개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `ReadFile(path string) ([]byte, error)` | Read entire file / 전체 파일 읽기 |
| `ReadString(path string) (string, error)` | Read file as string / 파일을 문자열로 읽기 |
| `ReadLines(path string) ([]string, error)` | Read file as lines / 파일을 줄로 읽기 |
| `ReadJSON(path string, v interface{}) error` | Read and unmarshal JSON / JSON 읽기 및 역직렬화 |
| `ReadYAML(path string, v interface{}) error` | Read and unmarshal YAML / YAML 읽기 및 역직렬화 |
| `ReadCSV(path string) ([][]string, error)` | Read CSV file / CSV 파일 읽기 |
| `ReadBytes(path string, offset, length int64) ([]byte, error)` | Read specific bytes / 특정 바이트 읽기 |
| `ReadChunk(path string, chunkSize int64, fn func([]byte) error) error` | Read file in chunks / 파일을 청크로 읽기 |

### 2. File Writing / 파일 쓰기 (11 functions / 11개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `WriteFile(path string, data []byte) error` | Write bytes to file / 바이트를 파일에 쓰기 |
| `WriteString(path string, s string) error` | Write string to file / 문자열을 파일에 쓰기 |
| `WriteLines(path string, lines []string) error` | Write lines to file / 줄을 파일에 쓰기 |
| `WriteJSON(path string, v interface{}) error` | Marshal and write JSON / JSON 직렬화 및 쓰기 |
| `WriteYAML(path string, v interface{}) error` | Marshal and write YAML / YAML 직렬화 및 쓰기 |
| `WriteCSV(path string, records [][]string) error` | Write CSV file / CSV 파일 쓰기 |
| `WriteAtomic(path string, data []byte) error` | Atomic write (temp + rename) / 원자적 쓰기 (임시 + 이름 변경) |
| `AppendFile(path string, data []byte) error` | Append bytes to file / 바이트를 파일에 추가 |
| `AppendString(path string, s string) error` | Append string to file / 문자열을 파일에 추가 |
| `AppendLines(path string, lines []string) error` | Append lines to file / 줄을 파일에 추가 |
| `AppendBytes(path string, data []byte) error` | Append bytes (alias) / 바이트 추가 (별칭) |

### 3. File Copying / 파일 복사 (4 functions / 4개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `CopyFile(src, dst string, opts ...CopyOption) error` | Copy single file / 단일 파일 복사 |
| `CopyDir(src, dst string, opts ...CopyOption) error` | Copy directory / 디렉토리 복사 |
| `CopyRecursive(src, dst string, opts ...CopyOption) error` | Copy recursively / 재귀적으로 복사 |
| `SyncDirs(src, dst string, opts ...CopyOption) error` | Sync two directories / 두 디렉토리 동기화 |

**Copy Options / 복사 옵션:**
- `WithOverwrite(bool)` - Overwrite existing files / 기존 파일 덮어쓰기
- `WithPreservePermissions(bool)` - Preserve file permissions / 파일 권한 보존
- `WithPreserveTimestamps(bool)` - Preserve timestamps / 타임스탬프 보존
- `WithProgress(func(written, total int64))` - Progress callback / 진행 상황 콜백
- `WithFilter(func(path string, info os.FileInfo) bool)` - File filter / 파일 필터

### 4. File Moving / 파일 이동 (5 functions / 5개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `MoveFile(src, dst string) error` | Move file / 파일 이동 |
| `MoveDir(src, dst string) error` | Move directory / 디렉토리 이동 |
| `Rename(oldPath, newPath string) error` | Rename file/directory / 파일/디렉토리 이름 변경 |
| `RenameExt(path, newExt string) (string, error)` | Change file extension / 파일 확장자 변경 |
| `SafeMove(src, dst string) error` | Move with existence check / 존재 확인과 함께 이동 |

### 5. File Deleting / 파일 삭제 (7 functions / 7개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `DeleteFile(path string) error` | Delete single file / 단일 파일 삭제 |
| `DeleteDir(path string) error` | Delete empty directory / 빈 디렉토리 삭제 |
| `DeleteRecursive(path string) error` | Delete recursively / 재귀적으로 삭제 |
| `DeletePattern(dir, pattern string) error` | Delete files by pattern / 패턴으로 파일 삭제 |
| `DeleteFiles(paths []string) error` | Delete multiple files / 여러 파일 삭제 |
| `Clean(path string) error` | Remove directory contents / 디렉토리 내용 제거 |
| `RemoveEmpty(path string) error` | Remove empty directories / 빈 디렉토리 제거 |

### 6. Directory Operations / 디렉토리 작업 (13 functions / 13개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `MkdirAll(path string) error` | Create directory tree / 디렉토리 트리 생성 |
| `CreateTemp(dir, pattern string) (*os.File, error)` | Create temp file / 임시 파일 생성 |
| `CreateTempDir(dir, pattern string) (string, error)` | Create temp directory / 임시 디렉토리 생성 |
| `IsEmpty(path string) (bool, error)` | Check if directory is empty / 디렉토리가 비어 있는지 확인 |
| `DirSize(path string) (int64, error)` | Calculate directory size / 디렉토리 크기 계산 |
| `ListFiles(dir string) ([]string, error)` | List files only / 파일만 나열 |
| `ListDirs(dir string) ([]string, error)` | List directories only / 디렉토리만 나열 |
| `ListAll(dir string) ([]string, error)` | List all entries / 모든 항목 나열 |
| `Walk(root string, fn filepath.WalkFunc) error` | Walk directory tree / 디렉토리 트리 순회 |
| `WalkFiles(root string, fn func(string, os.FileInfo) error) error` | Walk files only / 파일만 순회 |
| `WalkDirs(root string, fn func(string, os.FileInfo) error) error` | Walk directories only / 디렉토리만 순회 |
| `FindFiles(root string, filter func(string, interface{}) bool) ([]string, error)` | Find files by filter / 필터로 파일 찾기 |
| `FilterFiles(root string, patterns []string) ([]string, error)` | Filter files by patterns / 패턴으로 파일 필터링 |

### 7. Path Operations / 경로 작업 (18 functions / 18개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `Join(elem ...string) string` | Join path elements / 경로 요소 결합 |
| `Split(path string) (string, string)` | Split into dir and file / 디렉토리와 파일로 분할 |
| `Base(path string) string` | Get base name / 기본 이름 가져오기 |
| `Dir(path string) string` | Get directory / 디렉토리 가져오기 |
| `Ext(path string) string` | Get extension / 확장자 가져오기 |
| `Abs(path string) (string, error)` | Get absolute path / 절대 경로 가져오기 |
| `CleanPath(path string) string` | Clean path / 경로 정리 |
| `Normalize(path string) (string, error)` | Normalize path / 경로 정규화 |
| `ToSlash(path string) string` | Convert to forward slashes / 슬래시로 변환 |
| `FromSlash(path string) string` | Convert to OS-specific / OS 특정 변환 |
| `IsAbs(path string) bool` | Check if absolute / 절대 경로인지 확인 |
| `IsValid(path string) bool` | Validate path / 경로 검증 |
| `IsSafe(path, root string) bool` | Check path safety / 경로 안전성 확인 |
| `Match(pattern, name string) (bool, error)` | Match pattern / 패턴 매칭 |
| `Glob(pattern string) ([]string, error)` | Find by glob pattern / Glob 패턴으로 찾기 |
| `Rel(basepath, targpath string) (string, error)` | Get relative path / 상대 경로 가져오기 |
| `WithoutExt(path string) string` | Remove extension / 확장자 제거 |
| `ChangeExt(path, newExt string) string` | Change extension / 확장자 변경 |
| `HasExt(path string, exts ...string) bool` | Check extension / 확장자 확인 |

### 8. File Information / 파일 정보 (15 functions / 15개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `Exists(path string) bool` | Check existence / 존재 확인 |
| `IsFile(path string) bool` | Check if file / 파일인지 확인 |
| `IsDir(path string) bool` | Check if directory / 디렉토리인지 확인 |
| `IsSymlink(path string) bool` | Check if symlink / 심볼릭 링크인지 확인 |
| `Size(path string) (int64, error)` | Get file size / 파일 크기 가져오기 |
| `SizeHuman(path string) (string, error)` | Get human-readable size / 사람이 읽기 쉬운 크기 가져오기 |
| `Chmod(path string, mode os.FileMode) error` | Change permissions / 권한 변경 |
| `Chown(path string, uid, gid int) error` | Change owner / 소유자 변경 |
| `IsReadable(path string) bool` | Check if readable / 읽기 가능한지 확인 |
| `IsWritable(path string) bool` | Check if writable / 쓰기 가능한지 확인 |
| `IsExecutable(path string) bool` | Check if executable / 실행 가능한지 확인 |
| `ModTime(path string) (time.Time, error)` | Get modification time / 수정 시간 가져오기 |
| `AccessTime(path string) (time.Time, error)` | Get access time / 접근 시간 가져오기 |
| `ChangeTime(path string) (time.Time, error)` | Get change time / 변경 시간 가져오기 |
| `Touch(path string) error` | Update modification time / 수정 시간 업데이트 |

### 9. File Hashing / 파일 해싱 (10 functions / 10개 함수)

| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `MD5(path string) (string, error)` | Calculate MD5 hash / MD5 해시 계산 |
| `SHA1(path string) (string, error)` | Calculate SHA1 hash / SHA1 해시 계산 |
| `SHA256(path string) (string, error)` | Calculate SHA256 hash / SHA256 해시 계산 |
| `SHA512(path string) (string, error)` | Calculate SHA512 hash / SHA512 해시 계산 |
| `Hash(path, algorithm string) (string, error)` | Calculate hash by algorithm / 알고리즘별 해시 계산 |
| `HashBytes(data []byte, algorithm string) (string, error)` | Hash byte slice / 바이트 슬라이스 해시 |
| `CompareFiles(path1, path2 string) (bool, error)` | Compare files byte-by-byte / 파일을 바이트별로 비교 |
| `CompareHash(path1, path2 string) (bool, error)` | Compare files by hash / 해시로 파일 비교 |
| `Checksum(path string) (string, error)` | Calculate checksum (SHA256) / 체크섬 계산 (SHA256) |
| `VerifyChecksum(path, expected string) (bool, error)` | Verify checksum / 체크섬 검증 |

## Common Use Cases / 일반적인 사용 사례

### 1. Safe File Writing / 안전한 파일 쓰기

```go
// Automatically creates parent directories / 자동으로 상위 디렉토리 생성
err := fileutil.WriteString("deep/nested/path/file.txt", "content")
```

### 2. JSON/YAML Configuration / JSON/YAML 설정

```go
// Read config / 설정 읽기
var config Config
err := fileutil.ReadJSON("config.json", &config)

// Write config / 설정 쓰기
err = fileutil.WriteJSON("config.json", config)
```

### 3. File Backup / 파일 백업

```go
// Copy with timestamp / 타임스탬프와 함께 복사
timestamp := time.Now().Format("20060102-150405")
backupPath := fileutil.ChangeExt("file.txt", "."+timestamp+".bak")
err := fileutil.CopyFile("file.txt", backupPath)
```

### 4. Safe File Update / 안전한 파일 업데이트

```go
// Atomic write (writes to temp file, then renames) / 원자적 쓰기 (임시 파일에 쓰기, 그 다음 이름 변경)
err := fileutil.WriteAtomic("important-data.json", data)
```

### 5. File Integrity Check / 파일 무결성 확인

```go
// Generate checksum / 체크섬 생성
checksum, err := fileutil.Checksum("file.dat")

// Later, verify / 나중에 검증
valid, err := fileutil.VerifyChecksum("file.dat", checksum)
if !valid {
    log.Println("File has been modified!")
}
```

### 6. Recursive File Search / 재귀적 파일 검색

```go
// Find all Go files / 모든 Go 파일 찾기
goFiles, err := fileutil.FindFiles(".", func(path string, info interface{}) bool {
    return fileutil.Ext(path) == ".go"
})

for _, file := range goFiles {
    fmt.Println(file)
}
```

### 7. Large File Processing / 대용량 파일 처리

```go
// Process file in chunks / 파일을 청크로 처리
err := fileutil.ReadChunk("large-file.dat", 1024*1024, func(chunk []byte) error {
    // Process each 1MB chunk / 각 1MB 청크 처리
    return processChunk(chunk)
})
```

## Best Practices / 모범 사례

### 1. Always Check Errors / 항상 에러 확인

```go
// Good / 좋음
content, err := fileutil.ReadString("file.txt")
if err != nil {
    log.Printf("Failed to read file: %v", err)
    return err
}

// Bad / 나쁨
content, _ := fileutil.ReadString("file.txt")
```

### 2. Use Atomic Writes for Critical Data / 중요한 데이터에 원자적 쓰기 사용

```go
// For important data, use atomic writes / 중요한 데이터의 경우 원자적 쓰기 사용
err := fileutil.WriteAtomic("database.json", data)
```

### 3. Verify File Existence Before Operations / 작업 전 파일 존재 확인

```go
if !fileutil.Exists("input.txt") {
    return fmt.Errorf("input file does not exist")
}
```

### 4. Use Progress Callbacks for Large Operations / 대용량 작업에 진행 상황 콜백 사용

```go
err := fileutil.CopyFile("large-file.iso", "backup.iso",
    fileutil.WithProgress(func(written, total int64) {
        percent := float64(written) / float64(total) * 100
        fmt.Printf("\rCopying: %.1f%%", percent)
    }))
```

### 5. Clean Up Temporary Files / 임시 파일 정리

```go
tempFile, err := fileutil.CreateTemp("", "myapp-*")
if err != nil {
    return err
}
defer os.Remove(tempFile.Name()) // Clean up / 정리
```

### 6. Use Path Functions for Cross-Platform Compatibility / 크로스 플랫폼 호환성을 위해 경로 함수 사용

```go
// Good / 좋음
path := fileutil.Join("home", "user", "file.txt")

// Bad / 나쁨
path := "home/user/file.txt" // Unix-specific / Unix 전용
```

### 7. Validate Paths for Security / 보안을 위해 경로 검증

```go
// Check if path is safe (no directory traversal) / 경로가 안전한지 확인 (디렉토리 탐색 없음)
if !fileutil.IsSafe(userPath, rootDir) {
    return fmt.Errorf("unsafe path: %s", userPath)
}
```

## Error Handling / 에러 처리

The package provides custom error types for common scenarios:

패키지는 일반적인 시나리오에 대한 사용자 정의 에러 타입을 제공합니다:

```go
_, err := fileutil.ReadString("missing.txt")
if fileutil.IsNotFound(err) {
    // File doesn't exist / 파일이 존재하지 않음
}

_, err = fileutil.ReadString("/root/secure.txt")
if fileutil.IsPermission(err) {
    // Permission denied / 권한 거부
}

err = fileutil.WriteString("existing.txt", "data")
if fileutil.IsExist(err) {
    // File already exists / 파일이 이미 존재함
}
```

## Performance Considerations / 성능 고려사항

### 1. Buffered I/O / 버퍼링된 I/O

All file operations use buffered I/O with a default buffer size of 32KB for optimal performance.

모든 파일 작업은 최적의 성능을 위해 기본 버퍼 크기 32KB의 버퍼링된 I/O를 사용합니다.

### 2. Large Files / 대용량 파일

For large files, use `ReadChunk` to process data in chunks:

대용량 파일의 경우 `ReadChunk`를 사용하여 데이터를 청크로 처리합니다:

```go
// Process 1MB at a time / 한 번에 1MB 처리
err := fileutil.ReadChunk("huge-file.dat", fileutil.DefaultChunkSize, func(chunk []byte) error {
    return process(chunk)
})
```

### 3. Directory Operations / 디렉토리 작업

For recursive operations on large directory trees, use `Walk` functions:

대규모 디렉토리 트리의 재귀 작업의 경우 `Walk` 함수를 사용합니다:

```go
// More efficient than ListFiles for large directories / 대규모 디렉토리에는 ListFiles보다 효율적
err := fileutil.WalkFiles(".", func(path string, info os.FileInfo) error {
    // Process each file / 각 파일 처리
    return nil
})
```

## Comprehensive Documentation / 종합 문서

For more detailed information, please refer to:

더 자세한 정보는 다음을 참조하세요:

- **[User Manual](../docs/fileutil/USER_MANUAL.md)** - Complete usage guide with examples / 예제가 포함된 완전한 사용 가이드
- **[Developer Guide](../docs/fileutil/DEVELOPER_GUIDE.md)** - Architecture and implementation details / 아키텍처 및 구현 세부사항
- **[Examples](../examples/fileutil/)** - Comprehensive code examples / 종합적인 코드 예제

## Version / 버전

Current version: **v1.9.001**

현재 버전: **v1.9.001**

## License / 라이선스

MIT License - see [LICENSE](../LICENSE) file for details

MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE) 파일 참조

## Contributing / 기여

Contributions are welcome! Please see the [Developer Guide](../docs/fileutil/DEVELOPER_GUIDE.md) for guidelines.

기여를 환영합니다! 가이드라인은 [Developer Guide](../docs/fileutil/DEVELOPER_GUIDE.md)를 참조하세요.

## Support / 지원

- **GitHub Issues**: [https://github.com/arkd0ng/go-utils/issues](https://github.com/arkd0ng/go-utils/issues)
- **Documentation**: [https://github.com/arkd0ng/go-utils/tree/main/docs/fileutil](https://github.com/arkd0ng/go-utils/tree/main/docs/fileutil)

---

**Made with ❤️ for the Go community / Go 커뮤니티를 위해 ❤️로 만들었습니다**
