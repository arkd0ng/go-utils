# fileutil Package Work Plan / fileutil 패키지 작업 계획

**Version / 버전**: v1.9.001
**Date / 날짜**: 2025-10-15
**Status / 상태**: In Progress / 진행 중

---

## Table of Contents / 목차

1. [Implementation Tasks / 구현 작업](#implementation-tasks--구현-작업)
2. [File Structure / 파일 구조](#file-structure--파일-구조)
3. [Function List / 함수 목록](#function-list--함수-목록)
4. [Testing Plan / 테스트 계획](#testing-plan--테스트-계획)
5. [Documentation Plan / 문서화 계획](#documentation-plan--문서화-계획)
6. [Progress Tracking / 진행 상황 추적](#progress-tracking--진행-상황-추적)

---

## Implementation Tasks / 구현 작업

### Phase 1: Project Setup / 1단계: 프로젝트 설정

- [x] Create design plan / 설계 계획 생성
- [x] Create work plan / 작업 계획 생성
- [ ] Create package directory structure / 패키지 디렉토리 구조 생성
- [ ] Initialize package documentation / 패키지 문서 초기화

### Phase 2: Core File Operations / 2단계: 핵심 파일 작업

#### 2.1 File Reading (read.go)
- [ ] `ReadFile(path string) ([]byte, error)`
- [ ] `ReadString(path string) (string, error)`
- [ ] `ReadLines(path string) ([]string, error)`
- [ ] `ReadJSON(path string, v interface{}) error`
- [ ] `ReadYAML(path string, v interface{}) error`
- [ ] `ReadCSV(path string) ([][]string, error)`
- [ ] `ReadBytes(path string, offset, size int64) ([]byte, error)`
- [ ] `ReadChunk(path string, chunkSize int64, fn func([]byte) error) error`

#### 2.2 File Writing (write.go)
- [ ] `WriteFile(path string, data []byte, perm ...os.FileMode) error`
- [ ] `WriteString(path string, content string, perm ...os.FileMode) error`
- [ ] `WriteLines(path string, lines []string, perm ...os.FileMode) error`
- [ ] `WriteJSON(path string, v interface{}, perm ...os.FileMode) error`
- [ ] `WriteYAML(path string, v interface{}, perm ...os.FileMode) error`
- [ ] `WriteCSV(path string, records [][]string, perm ...os.FileMode) error`
- [ ] `WriteAtomic(path string, data []byte, perm ...os.FileMode) error`

#### 2.3 File Appending (write.go)
- [ ] `AppendFile(path string, data []byte) error`
- [ ] `AppendString(path string, content string) error`
- [ ] `AppendLines(path string, lines []string) error`
- [ ] `AppendBytes(path string, data []byte) error`

### Phase 3: File Operations / 3단계: 파일 작업

#### 3.1 File Copying (copy.go)
- [ ] `CopyFile(src, dst string, options ...CopyOption) error`
- [ ] `CopyDir(src, dst string, options ...CopyOption) error`
- [ ] `CopyRecursive(src, dst string, options ...CopyOption) error`
- [ ] `SyncDirs(src, dst string) error`

#### 3.2 File Moving (move.go)
- [ ] `MoveFile(src, dst string) error`
- [ ] `MoveDir(src, dst string) error`
- [ ] `Rename(old, new string) error`
- [ ] `RenameExt(path, newExt string) error`
- [ ] `SafeMove(src, dst string) error`

#### 3.3 File Deleting (delete.go)
- [ ] `DeleteFile(path string) error`
- [ ] `DeleteDir(path string) error`
- [ ] `DeleteRecursive(path string) error`
- [ ] `DeletePattern(pattern string) error`
- [ ] `DeleteFiles(paths ...string) error`
- [ ] `Clean(path string) error`

### Phase 4: Directory Operations / 4단계: 디렉토리 작업

#### 4.1 Directory Management (dir.go)
- [ ] `MkdirAll(path string, perm ...os.FileMode) error`
- [ ] `CreateTemp(dir, pattern string) (string, error)`
- [ ] `CreateTempDir(dir, pattern string) (string, error)`
- [ ] `IsEmpty(path string) (bool, error)`
- [ ] `DirSize(path string) (int64, error)`
- [ ] `ListFiles(dir string, recursive ...bool) ([]string, error)`
- [ ] `ListDirs(dir string, recursive ...bool) ([]string, error)`
- [ ] `ListAll(dir string, recursive ...bool) ([]string, error)`
- [ ] `Walk(root string, fn filepath.WalkFunc) error`
- [ ] `WalkFiles(root string, fn func(path string, info os.FileInfo) error) error`
- [ ] `WalkDirs(root string, fn func(path string, info os.FileInfo) error) error`
- [ ] `FindFiles(root string, predicate func(path string, info os.FileInfo) bool) ([]string, error)`
- [ ] `FilterFiles(files []string, predicate func(path string) bool) ([]string, error)`

### Phase 5: Path Operations / 5단계: 경로 작업

#### 5.1 Path Manipulation (path.go)
- [ ] `Join(elem ...string) string`
- [ ] `Split(path string) (dir, file string)`
- [ ] `Base(path string) string`
- [ ] `Dir(path string) string`
- [ ] `Ext(path string) string`
- [ ] `Abs(path string) (string, error)`
- [ ] `Clean(path string) string`
- [ ] `Normalize(path string) string`
- [ ] `ToSlash(path string) string`
- [ ] `FromSlash(path string) string`
- [ ] `IsAbs(path string) bool`
- [ ] `IsValid(path string) bool`
- [ ] `IsSafe(path, root string) bool`
- [ ] `Match(pattern, name string) (bool, error)`
- [ ] `Glob(pattern string) ([]string, error)`

### Phase 6: File Information / 6단계: 파일 정보

#### 6.1 Existence Checks (info.go)
- [ ] `Exists(path string) bool`
- [ ] `IsFile(path string) bool`
- [ ] `IsDir(path string) bool`
- [ ] `IsSymlink(path string) bool`

#### 6.2 Size Information (info.go)
- [ ] `Size(path string) (int64, error)`
- [ ] `SizeHuman(path string) (string, error)`

#### 6.3 Permissions (info.go)
- [ ] `Chmod(path string, mode os.FileMode) error`
- [ ] `Chown(path string, uid, gid int) error`
- [ ] `IsReadable(path string) bool`
- [ ] `IsWritable(path string) bool`
- [ ] `IsExecutable(path string) bool`

#### 6.4 Timestamps (info.go)
- [ ] `ModTime(path string) (time.Time, error)`
- [ ] `AccessTime(path string) (time.Time, error)`
- [ ] `ChangeTime(path string) (time.Time, error)`
- [ ] `Touch(path string) error`

#### 6.5 File Info (info.go)
- [ ] `FileInfo(path string) (os.FileInfo, error)`
- [ ] `Stat(path string) (os.FileInfo, error)`
- [ ] `Lstat(path string) (os.FileInfo, error)`

### Phase 7: File Hashing / 7단계: 파일 해싱

#### 7.1 Hash Functions (hash.go)
- [ ] `MD5(path string) (string, error)`
- [ ] `SHA1(path string) (string, error)`
- [ ] `SHA256(path string) (string, error)`
- [ ] `SHA512(path string) (string, error)`
- [ ] `Hash(path string, algorithm string) (string, error)`
- [ ] `HashBytes(path string) ([]byte, error)`

#### 7.2 Comparison (hash.go)
- [ ] `CompareFiles(path1, path2 string) (bool, error)`
- [ ] `CompareHash(path1, path2 string) (bool, error)`
- [ ] `Checksum(path string) (string, error)`
- [ ] `VerifyChecksum(path, checksum string) (bool, error)`

### Phase 8: File Compression / 8단계: 파일 압축

#### 8.1 Zip Operations (compress.go)
- [ ] `CreateZip(zipPath string, files []string) error`
- [ ] `ExtractZip(zipPath, dest string) error`
- [ ] `AddToZip(zipPath, filePath string) error`
- [ ] `ListZip(zipPath string) ([]string, error)`

#### 8.2 Tar Operations (compress.go)
- [ ] `CreateTar(tarPath string, files []string) error`
- [ ] `ExtractTar(tarPath, dest string) error`
- [ ] `ListTar(tarPath string) ([]string, error)`

#### 8.3 Gzip Operations (compress.go)
- [ ] `CreateGzip(gzPath, srcPath string) error`
- [ ] `ExtractGzip(gzPath, dest string) error`

#### 8.4 Archive Operations (compress.go)
- [ ] `ArchiveDir(dir, dest string) error`
- [ ] `Compress(src, dest string, format string) error`
- [ ] `Decompress(src, dest string) error`

### Phase 9: Temporary Files / 9단계: 임시 파일

#### 9.1 Temp Management (temp.go)
- [ ] `TempFile(dir, pattern string) (*os.File, error)`
- [ ] `TempDir(dir, pattern string) (string, error)`
- [ ] `TempFilePath(dir, pattern string) (string, error)`
- [ ] `CleanTemp(dir string) error`
- [ ] `AutoCleanTemp(dir string, maxAge time.Duration) error`

### Phase 10: File Watching / 10단계: 파일 감시

#### 10.1 Watch Operations (watch.go)
- [ ] `Watch(path string, fn func(event string) error) error`
- [ ] `WatchDir(dir string, fn func(event string) error) error`
- [ ] `OnChange(path string, fn func() error) error`

### Phase 11: Testing / 11단계: 테스트

#### 11.1 Unit Tests
- [ ] Test all read functions / 모든 읽기 함수 테스트
- [ ] Test all write functions / 모든 쓰기 함수 테스트
- [ ] Test all copy/move/delete functions / 모든 복사/이동/삭제 함수 테스트
- [ ] Test all directory functions / 모든 디렉토리 함수 테스트
- [ ] Test all path functions / 모든 경로 함수 테스트
- [ ] Test all info functions / 모든 정보 함수 테스트
- [ ] Test all hash functions / 모든 해시 함수 테스트
- [ ] Test all compress functions / 모든 압축 함수 테스트
- [ ] Test all temp functions / 모든 임시 파일 함수 테스트

#### 11.2 Integration Tests
- [ ] Test complete workflows / 완전한 워크플로우 테스트
- [ ] Test cross-platform behavior / 크로스 플랫폼 동작 테스트

#### 11.3 Benchmarks
- [ ] Benchmark read operations / 읽기 작업 벤치마크
- [ ] Benchmark write operations / 쓰기 작업 벤치마크
- [ ] Benchmark copy operations / 복사 작업 벤치마크
- [ ] Benchmark hash operations / 해시 작업 벤치마크

### Phase 12: Documentation / 12단계: 문서화

#### 12.1 Package Documentation
- [ ] Package README.md / 패키지 README.md
- [ ] API reference / API 참조
- [ ] Code examples / 코드 예제

#### 12.2 User Manual
- [ ] Installation guide / 설치 가이드
- [ ] Quick start / 빠른 시작
- [ ] Function reference / 함수 참조
- [ ] Common use cases / 일반적인 사용 사례
- [ ] Best practices / 모범 사례
- [ ] FAQ / 자주 묻는 질문

#### 12.3 Developer Guide
- [ ] Architecture overview / 아키텍처 개요
- [ ] Implementation details / 구현 세부사항
- [ ] Testing guide / 테스트 가이드
- [ ] Contributing guidelines / 기여 가이드라인

#### 12.4 Examples
- [ ] Basic file operations / 기본 파일 작업
- [ ] Directory traversal / 디렉토리 순회
- [ ] File hashing / 파일 해싱
- [ ] File compression / 파일 압축
- [ ] Complete application / 완전한 애플리케이션

---

## File Structure / 파일 구조

```
fileutil/
├── fileutil.go         # Package documentation and constants
├── read.go             # File reading operations (8 functions)
├── write.go            # File writing and appending (11 functions)
├── copy.go             # File/directory copying (4 functions)
├── move.go             # File/directory moving (5 functions)
├── delete.go           # File/directory deletion (6 functions)
├── dir.go              # Directory operations (13 functions)
├── path.go             # Path manipulation (15 functions)
├── info.go             # File information (15 functions)
├── hash.go             # File hashing (10 functions)
├── compress.go         # File compression (11 functions)
├── temp.go             # Temporary files (5 functions)
├── watch.go            # File watching (3 functions)
├── options.go          # Functional options
├── errors.go           # Error types
├── fileutil_test.go    # Comprehensive tests
├── examples_test.go    # Example tests
└── README.md           # Package documentation
```

---

## Function List / 함수 목록

### Total: 106 Functions / 총 106개 함수

#### 1. File Reading (8) - read.go
1. `ReadFile(path string) ([]byte, error)`
2. `ReadString(path string) (string, error)`
3. `ReadLines(path string) ([]string, error)`
4. `ReadJSON(path string, v interface{}) error`
5. `ReadYAML(path string, v interface{}) error`
6. `ReadCSV(path string) ([][]string, error)`
7. `ReadBytes(path string, offset, size int64) ([]byte, error)`
8. `ReadChunk(path string, chunkSize int64, fn func([]byte) error) error`

#### 2. File Writing (7) - write.go
1. `WriteFile(path string, data []byte, perm ...os.FileMode) error`
2. `WriteString(path string, content string, perm ...os.FileMode) error`
3. `WriteLines(path string, lines []string, perm ...os.FileMode) error`
4. `WriteJSON(path string, v interface{}, perm ...os.FileMode) error`
5. `WriteYAML(path string, v interface{}, perm ...os.FileMode) error`
6. `WriteCSV(path string, records [][]string, perm ...os.FileMode) error`
7. `WriteAtomic(path string, data []byte, perm ...os.FileMode) error`

#### 3. File Appending (4) - write.go
1. `AppendFile(path string, data []byte) error`
2. `AppendString(path string, content string) error`
3. `AppendLines(path string, lines []string) error`
4. `AppendBytes(path string, data []byte) error`

#### 4. File Copying (4) - copy.go
1. `CopyFile(src, dst string, options ...CopyOption) error`
2. `CopyDir(src, dst string, options ...CopyOption) error`
3. `CopyRecursive(src, dst string, options ...CopyOption) error`
4. `SyncDirs(src, dst string) error`

#### 5. File Moving (5) - move.go
1. `MoveFile(src, dst string) error`
2. `MoveDir(src, dst string) error`
3. `Rename(old, new string) error`
4. `RenameExt(path, newExt string) error`
5. `SafeMove(src, dst string) error`

#### 6. File Deleting (6) - delete.go
1. `DeleteFile(path string) error`
2. `DeleteDir(path string) error`
3. `DeleteRecursive(path string) error`
4. `DeletePattern(pattern string) error`
5. `DeleteFiles(paths ...string) error`
6. `Clean(path string) error`

#### 7. Directory Operations (13) - dir.go
1. `MkdirAll(path string, perm ...os.FileMode) error`
2. `CreateTemp(dir, pattern string) (string, error)`
3. `CreateTempDir(dir, pattern string) (string, error)`
4. `IsEmpty(path string) (bool, error)`
5. `DirSize(path string) (int64, error)`
6. `ListFiles(dir string, recursive ...bool) ([]string, error)`
7. `ListDirs(dir string, recursive ...bool) ([]string, error)`
8. `ListAll(dir string, recursive ...bool) ([]string, error)`
9. `Walk(root string, fn filepath.WalkFunc) error`
10. `WalkFiles(root string, fn func(path string, info os.FileInfo) error) error`
11. `WalkDirs(root string, fn func(path string, info os.FileInfo) error) error`
12. `FindFiles(root string, predicate func(path string, info os.FileInfo) bool) ([]string, error)`
13. `FilterFiles(files []string, predicate func(path string) bool) ([]string, error)`

#### 8. Path Operations (15) - path.go
1. `Join(elem ...string) string`
2. `Split(path string) (dir, file string)`
3. `Base(path string) string`
4. `Dir(path string) string`
5. `Ext(path string) string`
6. `Abs(path string) (string, error)`
7. `Clean(path string) string`
8. `Normalize(path string) string`
9. `ToSlash(path string) string`
10. `FromSlash(path string) string`
11. `IsAbs(path string) bool`
12. `IsValid(path string) bool`
13. `IsSafe(path, root string) bool`
14. `Match(pattern, name string) (bool, error)`
15. `Glob(pattern string) ([]string, error)`

#### 9. File Information (15) - info.go
1. `Exists(path string) bool`
2. `IsFile(path string) bool`
3. `IsDir(path string) bool`
4. `IsSymlink(path string) bool`
5. `Size(path string) (int64, error)`
6. `SizeHuman(path string) (string, error)`
7. `Chmod(path string, mode os.FileMode) error`
8. `Chown(path string, uid, gid int) error`
9. `IsReadable(path string) bool`
10. `IsWritable(path string) bool`
11. `IsExecutable(path string) bool`
12. `ModTime(path string) (time.Time, error)`
13. `AccessTime(path string) (time.Time, error)`
14. `ChangeTime(path string) (time.Time, error)`
15. `Touch(path string) error`

#### 10. File Hashing (10) - hash.go
1. `MD5(path string) (string, error)`
2. `SHA1(path string) (string, error)`
3. `SHA256(path string) (string, error)`
4. `SHA512(path string) (string, error)`
5. `Hash(path string, algorithm string) (string, error)`
6. `HashBytes(path string) ([]byte, error)`
7. `CompareFiles(path1, path2 string) (bool, error)`
8. `CompareHash(path1, path2 string) (bool, error)`
9. `Checksum(path string) (string, error)`
10. `VerifyChecksum(path, checksum string) (bool, error)`

#### 11. File Compression (11) - compress.go
1. `CreateZip(zipPath string, files []string) error`
2. `ExtractZip(zipPath, dest string) error`
3. `AddToZip(zipPath, filePath string) error`
4. `ListZip(zipPath string) ([]string, error)`
5. `CreateTar(tarPath string, files []string) error`
6. `ExtractTar(tarPath, dest string) error`
7. `ListTar(tarPath string) ([]string, error)`
8. `CreateGzip(gzPath, srcPath string) error`
9. `ExtractGzip(gzPath, dest string) error`
10. `ArchiveDir(dir, dest string) error`
11. `Compress(src, dest string, format string) error`

#### 12. Temporary Files (5) - temp.go
1. `TempFile(dir, pattern string) (*os.File, error)`
2. `TempDir(dir, pattern string) (string, error)`
3. `TempFilePath(dir, pattern string) (string, error)`
4. `CleanTemp(dir string) error`
5. `AutoCleanTemp(dir string, maxAge time.Duration) error`

#### 13. File Watching (3) - watch.go
1. `Watch(path string, fn func(event string) error) error`
2. `WatchDir(dir string, fn func(event string) error) error`
3. `OnChange(path string, fn func() error) error`

---

## Testing Plan / 테스트 계획

### Unit Tests / 단위 테스트

Each function will have:
- Basic functionality test / 기본 기능 테스트
- Edge case test / 엣지 케이스 테스트
- Error condition test / 에러 조건 테스트
- Cross-platform test (where applicable) / 크로스 플랫폼 테스트 (해당되는 경우)

### Integration Tests / 통합 테스트

- Complete file operation workflows / 완전한 파일 작업 워크플로우
- Directory traversal scenarios / 디렉토리 순회 시나리오
- File compression/decompression / 파일 압축/압축 해제
- Concurrent operations / 동시 작업

### Benchmarks / 벤치마크

- Read operations (small, medium, large files) / 읽기 작업 (작은, 중간, 큰 파일)
- Write operations (small, medium, large files) / 쓰기 작업 (작은, 중간, 큰 파일)
- Copy operations / 복사 작업
- Hash operations / 해시 작업
- Compression operations / 압축 작업

### Coverage Goal / 커버리지 목표

- Target: 90%+ code coverage / 목표: 90% 이상 코드 커버리지

---

## Documentation Plan / 문서화 계획

### 1. Package README.md (~500 lines)
- Quick start / 빠른 시작
- Installation / 설치
- Function categories / 함수 카테고리
- Common examples / 일반적인 예제

### 2. User Manual (~2500 lines)
- Introduction / 소개
- Installation / 설치
- Quick start / 빠른 시작
- Function reference (all 106 functions) / 함수 참조 (모든 106개 함수)
- Common use cases / 일반적인 사용 사례
- Best practices / 모범 사례
- Troubleshooting / 문제 해결
- FAQ / 자주 묻는 질문

### 3. Developer Guide (~2000 lines)
- Architecture overview / 아키텍처 개요
- Package structure / 패키지 구조
- Core components / 핵심 컴포넌트
- Implementation details / 구현 세부사항
- Testing guide / 테스트 가이드
- Performance / 성능
- Contributing / 기여하기

### 4. Examples (~1500 lines)
- Basic file operations / 기본 파일 작업
- Directory management / 디렉토리 관리
- Path manipulation / 경로 조작
- File hashing / 파일 해싱
- File compression / 파일 압축
- Complete applications / 완전한 애플리케이션

---

## Progress Tracking / 진행 상황 추적

### Phase 1: Setup ✅
- [x] Design plan created / 설계 계획 생성
- [x] Work plan created / 작업 계획 생성
- [ ] Package directory created / 패키지 디렉토리 생성

### Phase 2: Implementation ⏳
- [ ] File reading (0/8 functions)
- [ ] File writing (0/11 functions)
- [ ] File copying (0/4 functions)
- [ ] File moving (0/5 functions)
- [ ] File deleting (0/6 functions)
- [ ] Directory operations (0/13 functions)
- [ ] Path operations (0/15 functions)
- [ ] File information (0/15 functions)
- [ ] File hashing (0/10 functions)
- [ ] File compression (0/11 functions)
- [ ] Temporary files (0/5 functions)
- [ ] File watching (0/3 functions)

### Phase 3: Testing ⏳
- [ ] Unit tests (0/106 functions)
- [ ] Integration tests (0%)
- [ ] Benchmarks (0%)
- [ ] Coverage: 0%

### Phase 4: Documentation ⏳
- [ ] Package README.md (0%)
- [ ] User Manual (0%)
- [ ] Developer Guide (0%)
- [ ] Examples (0%)

---

## Milestones / 마일스톤

### Milestone 1: Core File Operations (Days 1-2)
- Read, write, append functions / 읽기, 쓰기, 추가 함수
- Copy, move, delete functions / 복사, 이동, 삭제 함수
- Basic tests / 기본 테스트

### Milestone 2: Directory & Path Operations (Day 3)
- Directory management / 디렉토리 관리
- Path manipulation / 경로 조작
- File information / 파일 정보

### Milestone 3: Advanced Features (Days 4-5)
- File hashing / 파일 해싱
- File compression / 파일 압축
- Temporary files / 임시 파일
- File watching / 파일 감시

### Milestone 4: Testing & Documentation (Days 6-7)
- Comprehensive tests / 포괄적인 테스트
- Complete documentation / 완전한 문서화
- Examples / 예제

---

## Next Steps / 다음 단계

1. Create package directory structure / 패키지 디렉토리 구조 생성
2. Implement core file reading functions / 핵심 파일 읽기 함수 구현
3. Implement core file writing functions / 핵심 파일 쓰기 함수 구현
4. Write tests for implemented functions / 구현된 함수 테스트 작성
5. Continue with remaining categories / 나머지 카테고리 계속 진행

---

## Notes / 참고사항

- Follow bilingual documentation standard (English/Korean) / 이중 언어 문서화 표준 준수 (영문/한글)
- Maintain consistency with existing packages (sliceutil, maputil) / 기존 패키지와 일관성 유지 (sliceutil, maputil)
- Ensure cross-platform compatibility / 크로스 플랫폼 호환성 보장
- Use standard library only (except gopkg.in/yaml.v3) / 표준 라이브러리만 사용 (gopkg.in/yaml.v3 제외)
- Target 90%+ test coverage / 90% 이상 테스트 커버리지 목표
- Write comprehensive documentation / 포괄적인 문서 작성

---

## Timeline / 타임라인

- **Day 1**: Design, work plan, core file operations (read, write, append) / 설계, 작업 계획, 핵심 파일 작업
- **Day 2**: File operations (copy, move, delete), directory operations / 파일 작업, 디렉토리 작업
- **Day 3**: Path operations, file information / 경로 작업, 파일 정보
- **Day 4**: File hashing, file compression / 파일 해싱, 파일 압축
- **Day 5**: Temporary files, file watching, tests / 임시 파일, 파일 감시, 테스트
- **Day 6**: Examples, documentation / 예제, 문서화
- **Day 7**: Review, polish, release / 검토, 다듬기, 릴리스

---

**Status / 상태**: Work plan completed. Ready to start implementation. / 작업 계획 완료. 구현 시작 준비 완료.
