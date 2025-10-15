# fileutil Package Design Plan / fileutil 패키지 설계 계획

**Version / 버전**: v1.9.001
**Date / 날짜**: 2025-10-15
**Status / 상태**: Draft / 초안

---

## Table of Contents / 목차

1. [Overview / 개요](#overview--개요)
2. [Design Philosophy / 설계 철학](#design-philosophy--설계-철학)
3. [Core Features / 핵심 기능](#core-features--핵심-기능)
4. [Package Architecture / 패키지 아키텍처](#package-architecture--패키지-아키텍처)
5. [Function Categories / 함수 카테고리](#function-categories--함수-카테고리)
6. [API Design / API 설계](#api-design--api-설계)
7. [Error Handling / 에러 처리](#error-handling--에러-처리)
8. [Performance Considerations / 성능 고려사항](#performance-considerations--성능-고려사항)
9. [Security Considerations / 보안 고려사항](#security-considerations--보안-고려사항)
10. [Dependencies / 의존성](#dependencies--의존성)

---

## Overview / 개요

### Purpose / 목적

The `fileutil` package provides extreme simplicity file and path utilities for Golang, reducing 20+ lines of repetitive file manipulation code to just 1-2 lines.

`fileutil` 패키지는 Golang을 위한 극도로 간단한 파일 및 경로 유틸리티를 제공하여, 20줄 이상의 반복적인 파일 조작 코드를 단 1-2줄로 줄입니다.

### Goals / 목표

1. **Simplicity**: Reduce boilerplate code for common file operations / 일반적인 파일 작업을 위한 보일러플레이트 코드 감소
2. **Safety**: Provide safe file operations with proper error handling / 적절한 에러 처리와 함께 안전한 파일 작업 제공
3. **Comprehensive**: Cover 90%+ of common file/path use cases / 일반적인 파일/경로 사용 사례의 90% 이상 커버
4. **Cross-platform**: Work seamlessly on Windows, macOS, Linux / Windows, macOS, Linux에서 원활하게 작동
5. **Type-safe**: Use Go 1.18+ generics where appropriate / 적절한 경우 Go 1.18+ 제네릭 사용
6. **Zero external dependencies**: Only use standard library / 표준 라이브러리만 사용

### Target Use Cases / 대상 사용 사례

- Reading/writing files with automatic directory creation / 자동 디렉토리 생성과 함께 파일 읽기/쓰기
- File/directory existence checks / 파일/디렉토리 존재 확인
- File copying, moving, and deletion / 파일 복사, 이동, 삭제
- Directory traversal and filtering / 디렉토리 순회 및 필터링
- Path manipulation and normalization / 경로 조작 및 정규화
- File size and permission management / 파일 크기 및 권한 관리
- Temporary file/directory handling / 임시 파일/디렉토리 처리
- File hashing and checksums / 파일 해싱 및 체크섬
- File compression (zip, tar.gz) / 파일 압축
- File watching and monitoring / 파일 감시 및 모니터링

---

## Design Philosophy / 설계 철학

### 1. "20 Lines → 1-2 Lines"

**Before / 이전**:
```go
// 20+ lines to safely write a file
dir := filepath.Dir(path)
if err := os.MkdirAll(dir, 0755); err != nil {
    return err
}
file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
if err != nil {
    return err
}
defer file.Close()
if _, err := file.Write(data); err != nil {
    return err
}
// ... more error handling
```

**After / 이후**:
```go
// 1 line with fileutil
err := fileutil.WriteFile(path, data)
```

### 2. Safe by Default / 기본적으로 안전

- Automatic directory creation / 자동 디렉토리 생성
- Proper error handling and propagation / 적절한 에러 처리 및 전파
- File permission management / 파일 권한 관리
- Path validation and sanitization / 경로 검증 및 정리
- Atomic file operations where possible / 가능한 경우 원자적 파일 작업

### 3. Cross-Platform Compatibility / 크로스 플랫폼 호환성

- Use `filepath` package for path operations / 경로 작업에 `filepath` 패키지 사용
- Handle Windows/Unix path separators / Windows/Unix 경로 구분자 처리
- Respect OS-specific file permissions / OS별 파일 권한 준수
- Platform-specific optimizations / 플랫폼별 최적화

### 4. Functional Programming Style / 함수형 프로그래밍 스타일

- Filter files by predicates / 조건자로 파일 필터링
- Map operations over files / 파일에 대한 Map 작업
- Chainable operations / 체이닝 가능한 작업

### 5. Performance Considerations / 성능 고려사항

- Lazy evaluation for directory traversal / 디렉토리 순회를 위한 지연 평가
- Buffered I/O for large files / 큰 파일을 위한 버퍼링된 I/O
- Concurrent operations where beneficial / 유익한 경우 동시 작업
- Memory-efficient streaming / 메모리 효율적인 스트리밍

---

## Core Features / 핵심 기능

### 1. File Operations / 파일 작업

- **Reading**: ReadFile, ReadLines, ReadString, ReadJSON, ReadYAML / 읽기
- **Writing**: WriteFile, WriteLines, WriteString, WriteJSON, WriteYAML / 쓰기
- **Appending**: AppendFile, AppendLines / 추가
- **Copying**: CopyFile, CopyDir, CopyRecursive / 복사
- **Moving**: MoveFile, MoveDir, Rename / 이동
- **Deleting**: DeleteFile, DeleteDir, DeleteRecursive / 삭제

### 2. Directory Operations / 디렉토리 작업

- **Creation**: MkdirAll, CreateTemp / 생성
- **Listing**: ListFiles, ListDirs, ListAll, Walk / 목록
- **Traversal**: WalkFiles, WalkDirs, FindFiles / 순회
- **Filtering**: FilterFiles, FilterDirs / 필터링

### 3. Path Operations / 경로 작업

- **Manipulation**: Join, Split, Base, Dir, Ext / 조작
- **Normalization**: Abs, Clean, Normalize / 정규화
- **Validation**: IsAbs, IsValid, IsSafe / 검증
- **Conversion**: ToSlash, FromSlash / 변환

### 4. File Information / 파일 정보

- **Existence**: Exists, IsFile, IsDir, IsSymlink / 존재
- **Size**: Size, SizeHuman, DirSize / 크기
- **Permissions**: Chmod, Chown, IsReadable, IsWritable, IsExecutable / 권한
- **Timestamps**: ModTime, AccessTime, ChangeTime / 타임스탬프

### 5. File Hashing / 파일 해싱

- **Checksums**: MD5, SHA1, SHA256, SHA512 / 체크섬
- **Comparison**: CompareFiles, CompareHash / 비교

### 6. File Compression / 파일 압축

- **Zip**: CreateZip, ExtractZip / Zip
- **Tar**: CreateTar, ExtractTar / Tar
- **Gzip**: CreateGzip, ExtractGzip / Gzip

### 7. Temporary Files / 임시 파일

- **Creation**: TempFile, TempDir / 생성
- **Management**: CleanTemp, AutoCleanTemp / 관리

### 8. File Watching / 파일 감시

- **Monitoring**: Watch, WatchDir, OnChange / 모니터링

---

## Package Architecture / 패키지 아키텍처

### File Organization / 파일 구성

```
fileutil/
├── fileutil.go         # Package documentation and version / 패키지 문서 및 버전
├── read.go             # File reading operations (10+ functions)
├── write.go            # File writing operations (10+ functions)
├── copy.go             # File/directory copying (5+ functions)
├── move.go             # File/directory moving (5+ functions)
├── delete.go           # File/directory deletion (5+ functions)
├── dir.go              # Directory operations (10+ functions)
├── path.go             # Path manipulation (15+ functions)
├── info.go             # File information (15+ functions)
├── hash.go             # File hashing (10+ functions)
├── compress.go         # File compression (10+ functions)
├── temp.go             # Temporary files (5+ functions)
├── watch.go            # File watching (5+ functions)
├── fileutil_test.go    # Comprehensive tests
└── README.md           # Package documentation (bilingual)
```

### Target Function Count / 목표 함수 개수

**Total: ~100 functions across 12 categories**

총 100개 함수를 12개 카테고리에 걸쳐 제공:

1. **File Reading (10)**: ReadFile, ReadLines, ReadString, ReadJSON, ReadYAML, ReadCSV, ReadBytes, ReadAt, ReadChunk, ReadStream
2. **File Writing (10)**: WriteFile, WriteLines, WriteString, WriteJSON, WriteYAML, WriteCSV, WriteBytes, WriteAt, WriteAtomic, WriteStream
3. **File Appending (5)**: AppendFile, AppendLines, AppendString, AppendBytes, AppendCSV
4. **File Copying (5)**: CopyFile, CopyDir, CopyRecursive, CopyWithProgress, SyncDirs
5. **File Moving (5)**: MoveFile, MoveDir, Rename, RenameExt, SafeMove
6. **File Deleting (5)**: DeleteFile, DeleteDir, DeleteRecursive, DeletePattern, Clean
7. **Directory Operations (10)**: MkdirAll, ListFiles, ListDirs, ListAll, Walk, WalkFiles, WalkDirs, FindFiles, CreateTemp, IsEmpty
8. **Path Operations (15)**: Join, Split, Base, Dir, Ext, Abs, Clean, Normalize, ToSlash, FromSlash, IsAbs, IsValid, IsSafe, Match, Glob
9. **File Information (15)**: Exists, IsFile, IsDir, IsSymlink, Size, SizeHuman, DirSize, Chmod, Chown, IsReadable, IsWritable, IsExecutable, ModTime, AccessTime, FileInfo
10. **File Hashing (10)**: MD5, SHA1, SHA256, SHA512, Hash, CompareFiles, CompareHash, Checksum, VerifyChecksum, HashDir
11. **File Compression (10)**: CreateZip, ExtractZip, AddToZip, CreateTar, ExtractTar, CreateGzip, ExtractGzip, Compress, Decompress, ArchiveDir
12. **Temporary & Watching (5)**: TempFile, TempDir, CleanTemp, Watch, WatchDir

---

## API Design / API 설계

### 1. Consistent Naming / 일관된 네이밍

- **Verbs first**: ReadFile, WriteFile, CopyFile / 동사 우선
- **Clear intent**: DeleteRecursive, CopyWithProgress / 명확한 의도
- **No abbreviations**: Use full words / 약어 사용 안함

### 2. Error Handling / 에러 처리

All functions return `error` as the last return value / 모든 함수는 마지막 반환값으로 `error` 반환:

```go
// Read operations / 읽기 작업
content, err := fileutil.ReadFile("path/to/file.txt")
if err != nil {
    // Handle error / 에러 처리
}

// Write operations / 쓰기 작업
err := fileutil.WriteFile("path/to/file.txt", data)
if err != nil {
    // Handle error / 에러 처리
}
```

### 3. Functional Options Pattern / 함수형 옵션 패턴

For complex operations with multiple configurations / 여러 설정이 있는 복잡한 작업의 경우:

```go
// Copy with options / 옵션과 함께 복사
err := fileutil.CopyFile(src, dst,
    fileutil.WithOverwrite(true),
    fileutil.WithPreservePermissions(true),
    fileutil.WithProgress(func(written, total int64) {
        fmt.Printf("Progress: %d/%d\n", written, total)
    }),
)
```

### 4. Variadic Parameters / 가변 인자

For flexible API / 유연한 API:

```go
// Join multiple path segments / 여러 경로 세그먼트 결합
path := fileutil.Join("home", "user", "documents", "file.txt")

// Delete multiple files / 여러 파일 삭제
err := fileutil.DeleteFiles("file1.txt", "file2.txt", "file3.txt")
```

### 5. Predicate Functions / 조건자 함수

For filtering and searching / 필터링 및 검색:

```go
// Find all .go files / 모든 .go 파일 찾기
files, err := fileutil.FindFiles(".", func(path string, info os.FileInfo) bool {
    return filepath.Ext(path) == ".go"
})

// Filter files larger than 1MB / 1MB보다 큰 파일 필터링
large, err := fileutil.FilterFiles(files, func(path string) bool {
    size, _ := fileutil.Size(path)
    return size > 1024*1024
})
```

---

## Error Handling / 에러 처리

### 1. Custom Error Types / 커스텀 에러 타입

```go
var (
    ErrNotFound      = errors.New("fileutil: file not found")
    ErrNotDirectory  = errors.New("fileutil: not a directory")
    ErrPermission    = errors.New("fileutil: permission denied")
    ErrInvalidPath   = errors.New("fileutil: invalid path")
    ErrAlreadyExists = errors.New("fileutil: file already exists")
)
```

### 2. Error Wrapping / 에러 래핑

Use `fmt.Errorf` with `%w` to wrap errors / `%w`와 함께 `fmt.Errorf` 사용:

```go
func ReadFile(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("fileutil.ReadFile: %w", err)
    }
    return data, nil
}
```

### 3. Error Checking Helpers / 에러 확인 헬퍼

```go
// Check if error is "file not found" / "파일을 찾을 수 없음" 에러 확인
func IsNotFound(err error) bool {
    return errors.Is(err, os.ErrNotExist)
}

// Check if error is "permission denied" / "권한 거부됨" 에러 확인
func IsPermission(err error) bool {
    return errors.Is(err, os.ErrPermission)
}
```

---

## Performance Considerations / 성능 고려사항

### 1. Buffered I/O / 버퍼링된 I/O

Use `bufio` for large file operations / 큰 파일 작업에 `bufio` 사용:

```go
// ReadLines uses buffered I/O / ReadLines는 버퍼링된 I/O 사용
func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
```

### 2. Lazy Evaluation / 지연 평가

Use iterators for directory traversal / 디렉토리 순회에 반복자 사용:

```go
// Walk uses lazy evaluation / Walk는 지연 평가 사용
func Walk(root string, fn filepath.WalkFunc) error {
    return filepath.Walk(root, fn)
}
```

### 3. Concurrent Operations / 동시 작업

Use goroutines for independent operations / 독립적인 작업에 고루틴 사용:

```go
// CopyFiles concurrently / 파일 동시 복사
func CopyFiles(files []string, dest string) error {
    var wg sync.WaitGroup
    errChan := make(chan error, len(files))

    for _, file := range files {
        wg.Add(1)
        go func(f string) {
            defer wg.Done()
            if err := CopyFile(f, filepath.Join(dest, filepath.Base(f))); err != nil {
                errChan <- err
            }
        }(file)
    }

    wg.Wait()
    close(errChan)

    if len(errChan) > 0 {
        return <-errChan
    }
    return nil
}
```

### 4. Memory Efficiency / 메모리 효율성

Use streaming for large files / 큰 파일에 스트리밍 사용:

```go
// CopyFile uses streaming / CopyFile은 스트리밍 사용
func CopyFile(src, dst string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()

    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    return err
}
```

---

## Security Considerations / 보안 고려사항

### 1. Path Traversal Prevention / 경로 순회 방지

```go
// IsSafe checks if path is safe (no ".." or absolute paths outside root)
// IsSafe는 경로가 안전한지 확인합니다 (".." 또는 루트 외부의 절대 경로 없음)
func IsSafe(path string, root string) bool {
    absPath, err := filepath.Abs(path)
    if err != nil {
        return false
    }
    absRoot, err := filepath.Abs(root)
    if err != nil {
        return false
    }
    return strings.HasPrefix(absPath, absRoot)
}
```

### 2. File Permission Validation / 파일 권한 검증

```go
// Ensure safe permissions / 안전한 권한 보장
const (
    DefaultFileMode = 0644  // rw-r--r--
    DefaultDirMode  = 0755  // rwxr-xr-x
)
```

### 3. Atomic File Operations / 원자적 파일 작업

```go
// WriteAtomic writes to temp file first, then renames
// WriteAtomic는 먼저 임시 파일에 쓴 다음 이름을 변경합니다
func WriteAtomic(path string, data []byte) error {
    temp := path + ".tmp"
    if err := WriteFile(temp, data); err != nil {
        return err
    }
    return os.Rename(temp, path)
}
```

---

## Dependencies / 의존성

### Standard Library Only / 표준 라이브러리만 사용

- `os` - File system operations / 파일 시스템 작업
- `io` - I/O primitives / I/O 기본 요소
- `io/fs` - File system interfaces / 파일 시스템 인터페이스
- `path/filepath` - Path manipulation / 경로 조작
- `bufio` - Buffered I/O / 버퍼링된 I/O
- `crypto/md5` - MD5 hashing / MD5 해싱
- `crypto/sha1` - SHA1 hashing / SHA1 해싱
- `crypto/sha256` - SHA256 hashing / SHA256 해싱
- `crypto/sha512` - SHA512 hashing / SHA512 해싱
- `archive/zip` - Zip compression / Zip 압축
- `archive/tar` - Tar archiving / Tar 아카이빙
- `compress/gzip` - Gzip compression / Gzip 압축
- `encoding/json` - JSON encoding/decoding / JSON 인코딩/디코딩
- `gopkg.in/yaml.v3` - YAML encoding/decoding / YAML 인코딩/디코딩

**Note**: Only `gopkg.in/yaml.v3` is an external dependency (already used in other packages).

**참고**: `gopkg.in/yaml.v3`만 외부 의존성입니다 (이미 다른 패키지에서 사용 중).

---

## Implementation Phases / 구현 단계

### Phase 1: Core File Operations / 1단계: 핵심 파일 작업
- File reading/writing/appending / 파일 읽기/쓰기/추가
- File copying/moving/deleting / 파일 복사/이동/삭제
- Basic directory operations / 기본 디렉토리 작업

### Phase 2: Path & Information / 2단계: 경로 및 정보
- Path manipulation / 경로 조작
- File information / 파일 정보
- Permission management / 권한 관리

### Phase 3: Advanced Features / 3단계: 고급 기능
- File hashing / 파일 해싱
- File compression / 파일 압축
- Temporary file management / 임시 파일 관리

### Phase 4: Specialized Features / 4단계: 특수 기능
- File watching / 파일 감시
- Progress tracking / 진행 상황 추적
- Concurrent operations / 동시 작업

---

## Testing Strategy / 테스트 전략

### 1. Unit Tests / 단위 테스트
- Test each function independently / 각 함수를 독립적으로 테스트
- Cover edge cases / 엣지 케이스 커버
- Test error conditions / 에러 조건 테스트

### 2. Integration Tests / 통합 테스트
- Test file operation workflows / 파일 작업 워크플로우 테스트
- Test cross-platform behavior / 크로스 플랫폼 동작 테스트

### 3. Benchmarks / 벤치마크
- Measure I/O performance / I/O 성능 측정
- Compare with standard library / 표준 라이브러리와 비교

### 4. Coverage Goal / 커버리지 목표
- Target: 90%+ code coverage / 목표: 90% 이상 코드 커버리지

---

## Documentation Plan / 문서화 계획

### 1. Package README / 패키지 README
- Quick start examples / 빠른 시작 예제
- Function categories / 함수 카테고리
- Common use cases / 일반적인 사용 사례

### 2. User Manual / 사용자 매뉴얼
- Comprehensive guide (~2000 lines) / 포괄적인 가이드 (~2000줄)
- All functions with examples / 모든 함수와 예제
- Best practices / 모범 사례

### 3. Developer Guide / 개발자 가이드
- Architecture overview / 아키텍처 개요
- Implementation details / 구현 세부사항
- Contributing guidelines / 기여 가이드라인

### 4. Examples / 예제
- Complete working examples / 완전한 작동 예제
- Real-world use cases / 실제 사용 사례

---

## Success Criteria / 성공 기준

1. ✅ 100+ functions implemented / 100개 이상 함수 구현
2. ✅ 90%+ test coverage / 90% 이상 테스트 커버리지
3. ✅ Zero external dependencies (except yaml) / 외부 의존성 없음 (yaml 제외)
4. ✅ Cross-platform compatibility / 크로스 플랫폼 호환성
5. ✅ Comprehensive documentation / 포괄적인 문서화
6. ✅ All examples working / 모든 예제 작동
7. ✅ Performance benchmarks / 성능 벤치마크
8. ✅ Security review passed / 보안 검토 통과

---

## Timeline / 타임라인

- **Day 1**: Design plan, work plan, core file operations (read, write, append) / 설계 계획, 작업 계획, 핵심 파일 작업
- **Day 2**: File operations (copy, move, delete), directory operations / 파일 작업, 디렉토리 작업
- **Day 3**: Path operations, file information / 경로 작업, 파일 정보
- **Day 4**: File hashing, file compression / 파일 해싱, 파일 압축
- **Day 5**: Temporary files, file watching, tests / 임시 파일, 파일 감시, 테스트
- **Day 6**: Examples, documentation / 예제, 문서화
- **Day 7**: Review, polish, release / 검토, 다듬기, 릴리스

---

## Conclusion / 결론

The `fileutil` package will provide extreme simplicity file and path utilities for Golang, following the same design philosophy as `sliceutil` and `maputil`. With 100+ functions across 12 categories, it will cover 90%+ of common file/path use cases and reduce boilerplate code from 20+ lines to just 1-2 lines.

`fileutil` 패키지는 `sliceutil` 및 `maputil`과 동일한 설계 철학을 따라 Golang을 위한 극도로 간단한 파일 및 경로 유틸리티를 제공합니다. 12개 카테고리에 걸쳐 100개 이상의 함수로 일반적인 파일/경로 사용 사례의 90% 이상을 커버하고, 보일러플레이트 코드를 20줄 이상에서 단 1-2줄로 줄입니다.

**Next Steps / 다음 단계**:
1. Create work plan / 작업 계획 생성
2. Implement core functions / 핵심 함수 구현
3. Write comprehensive tests / 포괄적인 테스트 작성
4. Create documentation / 문서 작성
5. Release v1.9.001 / v1.9.001 릴리스
