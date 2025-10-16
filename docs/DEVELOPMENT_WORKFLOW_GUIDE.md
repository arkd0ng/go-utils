# Development Workflow Guide / 개발 워크플로우 가이드

Complete guide for all development work in go-utils.

go-utils의 모든 개발 작업을 위한 완전한 가이드입니다.

**Version / 버전**: v1.11.045  
**Last Updated / 최종 업데이트**: 2025-10-16

---

## 📋 Table of Contents / 목차

- [Critical Rules / 핵심 규칙](#critical-rules--핵심-규칙)
- [Standard Work Cycle / 표준 작업 사이클](#standard-work-cycle--표준-작업-사이클)
- [Branch Strategy / 브랜치 전략](#branch-strategy--브랜치-전략)
- [Version Management / 버전 관리](#version-management--버전-관리)
- [Package Development / 패키지 개발](#package-development--패키지-개발)
- [Documentation Work / 문서 작업](#documentation-work--문서-작업)
- [Testing Standards / 테스트 표준](#testing-standards--테스트-표준)
- [Example Code / 예제 코드](#example-code--예제-코드)
- [Git Workflow / Git 워크플로우](#git-workflow--git-워크플로우)
- [Quick Reference / 빠른 참조](#quick-reference--빠른-참조)

---

## 🚨 Critical Rules / 핵심 규칙

### ⚠️ MUST FOLLOW / 반드시 준수

**모든 작업은 다음 순서를 엄격히 따라야 합니다:**

```
1. 패치 버전 증가 (cfg/app.yaml)
   ↓
2. 작업 수행 (코딩 또는 문서화)
   ↓
3. 테스트 및 검증 (go test, go build)
   ↓
4. CHANGELOG 업데이트
   ↓
5. Git 커밋 및 푸시
```

**❌ 절대 하지 말 것:**
- 버전 증가 없이 작업 시작
- CHANGELOG 업데이트 없이 커밋
- 테스트 실패 상태에서 푸시
- 문서화 없이 코드만 푸시

**✅ 항상 할 것:**
- 작업 전에 항상 버전 증가
- 작업 후에 항상 CHANGELOG 업데이트
- 커밋 전에 항상 테스트 실행
- 이중 언어 문서화 (영문/한글)

---

## 🔄 Standard Work Cycle / 표준 작업 사이클

Every task (coding, documentation, bug fix) follows this exact cycle:

모든 작업(코딩, 문서화, 버그 수정)은 이 사이클을 정확히 따릅니다:

### Step 1: Version Bump / 버전 증가

```bash
# Edit cfg/app.yaml
# v1.11.044 → v1.11.045

# Commit version change
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.11.045"
```

**Why / 이유**: Version must be incremented BEFORE any work to maintain clean history.

버전은 깔끔한 이력 유지를 위해 모든 작업 전에 증가되어야 합니다.

### Step 2: Perform Work / 작업 수행

**Coding Work / 코딩 작업**:
- Implement function / 함수 구현
- Add tests / 테스트 추가
- Add examples / 예제 추가

**Documentation Work / 문서 작업**:
- Update README / README 업데이트
- Update USER_MANUAL / USER_MANUAL 업데이트
- Update DEVELOPER_GUIDE / DEVELOPER_GUIDE 업데이트

### Step 3: Verify / 검증

```bash
# Build
go build ./...

# Test
go test ./... -v

# Test specific package
go test ./{package} -v

# Check coverage
go test ./{package} -cover
```

**All must pass / 모두 통과해야 함**: ✅

### Step 4: Update CHANGELOG / CHANGELOG 업데이트

Edit `docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md`:

```markdown
## [v1.11.045] - 2025-10-16

### Added
- Added new development workflow guide
- 새로운 개발 워크플로우 가이드 추가

### Changed
- N/A

### Fixed
- N/A
```

### Step 5: Git Commit & Push / Git 커밋 및 푸시

```bash
# Stage all changes
git add .

# Commit with proper message
git commit -m "Docs: Add comprehensive development workflow guide (v1.11.045)"

# Push to repository
git push origin main
```

**🎉 Cycle Complete / 사이클 완료!**

---

## 🌿 Branch Strategy / 브랜치 전략

### Main Branch / 메인 브랜치

- **main**: Stable releases / 안정 릴리스
- Direct commits allowed for: / 직접 커밋 허용:
  - Documentation updates / 문서 업데이트
  - Bug fixes / 버그 수정
  - Minor improvements / 사소한 개선

### Feature Branches / 기능 브랜치

For new packages, create a feature branch:

새 패키지의 경우 기능 브랜치를 생성:

```bash
# Format: feature/v{MAJOR}.{MINOR}.x-{package-name}
git checkout -b feature/v1.11.x-httpserver
```

**Example / 예제**:
- `feature/v1.11.x-httpserver` - HTTP server utilities
- `feature/v1.12.x-validation` - Validation utilities
- `feature/v1.13.x-crypto` - Cryptography utilities

### Branch Workflow / 브랜치 워크플로우

```
main
  │
  ├─ feature/v1.11.x-httpserver
  │   └─ (develop package)
  │   └─ (merge to main when complete)
  │
  ├─ feature/v1.12.x-validation
  │   └─ (develop package)
  │   └─ (merge to main when complete)
  │
  └─ (continue main development)
```

### Concurrent Development / 동시 개발

When multiple packages are being developed:

여러 패키지가 동시에 개발될 때:

1. **Check Current Versions / 현재 버전 확인**
   ```bash
   # Check README.md and CHANGELOG.md
   cat README.md | grep "In Development"
   ```

2. **Choose Next Minor Version / 다음 마이너 버전 선택**
   ```
   v1.11.x - httpserver (in development)
   v1.12.x - validation (your new package)
   ```

3. **Document Your Branch / 브랜치 문서화**
   - Add to README.md with status / README.md에 상태와 함께 추가
   - Add to CHANGELOG.md with branch name / CHANGELOG.md에 브랜치 이름과 함께 추가

---

## 📊 Version Management / 버전 관리

### Version Format / 버전 형식

```
vMAJOR.MINOR.PATCH
```

**Example / 예**: v1.11.045

### Version Semantic / 버전 의미

| Type / 타입 | When / 언제 | Example / 예시 |
|-------------|------------|---------------|
| **MAJOR** | Breaking changes / 호환성 깨짐 | v1.0.0 → v2.0.0 |
| **MINOR** | New package / 새 패키지 | v1.10.0 → v1.11.0 |
| **PATCH** | Every task / 모든 작업 | v1.11.044 → v1.11.045 |

### Patch Version Rules / 패치 버전 규칙

**Increment before EVERY task / 모든 작업 전에 증가:**

- ✅ Implementing a function / 함수 구현
- ✅ Adding a test / 테스트 추가
- ✅ Writing documentation / 문서 작성
- ✅ Fixing a bug / 버그 수정
- ✅ Updating example / 예제 업데이트
- ✅ Refactoring code / 코드 리팩토링

**Format / 형식**: Always 3 digits / 항상 3자리

```
v1.11.001
v1.11.002
...
v1.11.099
v1.11.100
```

### Version Bump Process / 버전 증가 프로세스

```bash
# 1. Edit cfg/app.yaml
# Change: version: v1.11.044
# To:     version: v1.11.045

# 2. Commit version change FIRST
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.11.045"

# 3. NOW you can start your work
```

---

## 📦 Package Development / 패키지 개발

### Phase 1: Planning / 계획 단계

**Step 1: Create Branch / 브랜치 생성**

```bash
git checkout -b feature/v1.11.x-{package}
```

**Step 2: Design Document / 설계 문서**

Create `docs/{package}/DESIGN_PLAN.md`:

```markdown
# {Package} Design Plan / 설계 계획

## 1. Package Overview / 패키지 개요
- Purpose / 목적
- Target users / 대상 사용자
- Key features / 주요 기능

## 2. Architecture / 아키텍처
- File structure / 파일 구조
- Core components / 핵심 구성요소
- Dependencies / 의존성

## 3. API Design / API 설계
- Function signatures / 함수 시그니처
- Type definitions / 타입 정의
- Options pattern / 옵션 패턴

## 4. Examples / 예제
- Use case 1 / 사용 사례 1
- Use case 2 / 사용 사례 2
- Before vs After / 이전 vs 이후
```

**Step 3: Work Plan / 작업 계획**

Create `docs/{package}/WORK_PLAN.md`:

```markdown
# {Package} Work Plan / 작업 계획

## Phase 1: Core Functions (v1.11.001-010)
- [ ] Function1 (v1.11.001)
- [ ] Function2 (v1.11.002)
- [ ] Function3 (v1.11.003)
- [ ] Tests for Core Functions (v1.11.004-006)
- [ ] Examples for Core Functions (v1.11.007-009)
- [ ] README.md (v1.11.010)

## Phase 2: Advanced Features (v1.11.011-020)
- [ ] Feature1 (v1.11.011)
- [ ] Feature2 (v1.11.012)
- [ ] Tests (v1.11.013-015)
- [ ] Examples (v1.11.016-018)
- [ ] Documentation (v1.11.019-020)

## Phase 3: Finalization (v1.11.021-030)
- [ ] USER_MANUAL.md (v1.11.021-025)
- [ ] DEVELOPER_GUIDE.md (v1.11.026-030)
```

### Phase 2: Implementation / 구현 단계

Follow the **Unit Task Workflow** for each function:

각 함수에 대해 **단위 작업 워크플로우**를 따릅니다:

#### Unit Task Workflow / 단위 작업 워크플로우

```
1. Bump Version → 2. Code → 3. Test → 4. Example → 5. Docs → 6. CHANGELOG → 7. Commit
```

**Detailed Steps / 상세 단계:**

**1. Bump Version / 버전 증가**
```bash
# cfg/app.yaml: v1.11.001 → v1.11.002
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.11.002"
```

**2. Implement Code / 코드 구현**
```go
// Get retrieves a value from the map by key.
// Get은 키로 맵에서 값을 검색합니다.
func Get[K comparable, V any](m map[K]V, key K) (V, bool) {
    val, ok := m[key]
    return val, ok
}
```

**3. Write Tests / 테스트 작성**
```go
func TestGet(t *testing.T) {
    tests := []struct {
        name    string
        input   map[string]int
        key     string
        wantVal int
        wantOk  bool
    }{
        {
            name:    "existing key",
            input:   map[string]int{"a": 1},
            key:     "a",
            wantVal: 1,
            wantOk:  true,
        },
        {
            name:    "missing key",
            input:   map[string]int{"a": 1},
            key:     "b",
            wantVal: 0,
            wantOk:  false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, ok := Get(tt.input, tt.key)
            if got != tt.wantVal || ok != tt.wantOk {
                t.Errorf("Get() = (%v, %v), want (%v, %v)",
                    got, ok, tt.wantVal, tt.wantOk)
            }
        })
    }
}
```

**4. Add Example / 예제 추가**
```go
// examples/{package}/main.go
func example1(logger *logging.Logger) {
    logger.Info("╔═══════════════════════════════════════╗")
    logger.Info("║  EXAMPLE 1: Get Function              ║")
    logger.Info("╚═══════════════════════════════════════╝")
    
    m := map[string]int{"a": 1, "b": 2}
    val, ok := Get(m, "a")
    
    logger.Info("✅ Result", "value", val, "found", ok)
}
```

**5. Update Documentation / 문서 업데이트**
```markdown
### Get

Retrieves a value from the map by key.

키로 맵에서 값을 검색합니다.

**Signature / 시그니처:**
```go
func Get[K comparable, V any](m map[K]V, key K) (V, bool)
```

**Example / 예제:**
```go
val, ok := maputil.Get(myMap, "key")
if ok {
    fmt.Println(val)
}
```
```

**6. Update CHANGELOG / CHANGELOG 업데이트**
```markdown
## [v1.11.002] - 2025-10-16

### Added
- Added Get function for map value retrieval
- Get 함수로 맵 값 검색 추가
```

**7. Commit & Push / 커밋 및 푸시**
```bash
git add .
git commit -m "Feat: Add Get function to maputil (v1.11.002)"
git push origin feature/v1.11.x-maputil
```

### Phase 3: Finalization / 마무리 단계

**Step 1: Comprehensive Documentation / 종합 문서화**

Create complete user and developer manuals:

완전한 사용자 및 개발자 매뉴얼 생성:

- `docs/{package}/USER_MANUAL.md` (1000+ lines)
- `docs/{package}/DEVELOPER_GUIDE.md` (1000+ lines)

**Step 2: Final Review / 최종 검토**

```bash
# Run all tests
go test ./... -v

# Check coverage
go test ./{package} -cover

# Run examples
go run examples/{package}/main.go

# Verify documentation
# - All functions documented
# - All examples working
# - Bilingual throughout
```

**Step 3: Merge to Main / 메인에 병합**

```bash
# Switch to main
git checkout main

# Merge feature branch
git merge feature/v1.11.x-{package}

# Push to main
git push origin main

# Tag the release
git tag v1.11.030
git push origin v1.11.030
```

---

## 📚 Documentation Work / 문서 작업

### Documentation Types / 문서 유형

1. **Package README.md**
   - Overview / 개요
   - Installation / 설치
   - Quick start / 빠른 시작
   - API reference / API 참조

2. **USER_MANUAL.md**
   - Comprehensive guide / 포괄적 가이드
   - All functions / 모든 함수
   - Usage patterns / 사용 패턴
   - Best practices / 모범 사례

3. **DEVELOPER_GUIDE.md**
   - Architecture / 아키텍처
   - Implementation details / 구현 세부사항
   - Contributing / 기여 방법
   - Testing guide / 테스트 가이드

4. **CHANGELOG**
   - All version changes / 모든 버전 변경
   - Detailed patch history / 상세 패치 이력

### Bilingual Format / 이중 언어 형식

**🚨 CRITICAL: ALL DOCUMENTATION MUST BE BILINGUAL (ENGLISH/KOREAN)**
**🚨 핵심: 모든 문서는 반드시 영문/한글 병기**

**All documentation MUST be bilingual (English/Korean):**

모든 문서는 이중 언어(영문/한글)여야 합니다:

#### What Must Be Bilingual / 반드시 병기해야 하는 항목

1. **All .md files / 모든 .md 파일**
2. **Code comments / 코드 주석**
3. **Git commit messages / Git 커밋 메시지**
4. **Log messages / 로그 메시지**
5. **Error messages / 에러 메시지**

#### Exceptions (English Only) / 예외 (영문만)

- **CLAUDE.md** (personal AI guidance / 개인 AI 가이드)
- **todo.md** (personal task list / 개인 작업 목록)
- **Variable/function names / 변수/함수 이름**

#### Format Examples / 형식 예제

```markdown
## Section Title / 섹션 제목

English description first.

한글 설명 다음.

**Example / 예제:**
```go
// English comment / 한글 주석
code here
```

**Note / 참고**: Important information / 중요한 정보
```

### Documentation Workflow / 문서 작업 워크플로우

```
1. Bump Version
   ↓
2. Write/Update Documentation
   ↓
3. Review for Bilingual Completeness
   ↓
4. Update CHANGELOG
   ↓
5. Commit & Push
```

---

## 🧪 Testing Standards / 테스트 표준

### Coverage Requirements / 커버리지 요구사항

| Level / 레벨 | Coverage / 커버리지 | Status / 상태 |
|--------------|-------------------|--------------|
| Minimum / 최소 | 60% | 🟡 Acceptable / 허용 |
| Target / 목표 | 80% | 🟢 Recommended / 권장 |
| Critical / 중요 | 100% | ⭐ Required / 필수 |

### Test Structure / 테스트 구조

**Use table-driven tests / 테이블 기반 테스트 사용:**

```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    OutputType
        wantErr bool
    }{
        {
            name:    "valid case",
            input:   ValidInput,
            want:    ExpectedOutput,
            wantErr: false,
        },
        {
            name:    "error case",
            input:   InvalidInput,
            want:    ZeroValue,
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Function(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Test Categories / 테스트 카테고리

1. **Unit Tests / 단위 테스트**
   - Test each function independently / 각 함수를 독립적으로 테스트
   - All edge cases / 모든 엣지 케이스
   - Error conditions / 에러 조건

2. **Integration Tests / 통합 테스트**
   - Functions working together / 함수들이 함께 작동
   - Realistic scenarios / 현실적인 시나리오

3. **Benchmarks / 벤치마크**
   - Performance-critical functions / 성능 중요 함수
   - Memory allocations / 메모리 할당

### Running Tests / 테스트 실행

```bash
# All tests with verbose output
go test ./... -v

# Specific package
go test ./{package} -v

# With coverage
go test ./{package} -cover

# Coverage report
go test ./{package} -coverprofile=coverage.out
go tool cover -html=coverage.out

# Benchmarks
go test ./{package} -bench=.

# Specific test
go test ./{package} -v -run TestFunction
```

---

## 💡 Example Code / 예제 코드

### Example Structure / 예제 구조

**All examples follow this template:**

모든 예제는 이 템플릿을 따릅니다:

```go
package main

import (
    "github.com/arkd0ng/go-utils/logging"
    "github.com/arkd0ng/go-utils/{package}"
)

func main() {
    // Initialize logger with backup
    logger := initLogger()
    defer logger.Close()
    
    // Print banner
    logger.Info("╔════════════════════════════════════════════════════╗")
    logger.Info("║  {Package} - Examples & Manual                     ║")
    logger.Info("║  {Package} - 예제 및 매뉴얼                         ║")
    logger.Info("╚════════════════════════════════════════════════════╝")
    logger.Info("")
    
    // Package information
    logger.Info("📋 Package: github.com/arkd0ng/go-utils/{package}")
    logger.Info("   Version: " + {package}.Version)
    logger.Info("")
    
    // Run examples
    example1(logger)
    example2(logger)
    example3(logger)
}

func initLogger() *logging.Logger {
    // Create logs directory
    os.MkdirAll("logs/{package}", 0755)
    
    // Backup previous log
    logFile := "logs/{package}/example.log"
    if _, err := os.Stat(logFile); err == nil {
        backup := fmt.Sprintf("logs/{package}/example-%s.log",
            time.Now().Format("20060102-150405"))
        os.Rename(logFile, backup)
        fmt.Printf("✅ Backup: %s\n", backup)
    }
    
    // Create logger
    logger, _ := logging.New(
        logging.WithFilePath(logFile),
        logging.WithLevel(logging.LevelInfo),
        logging.WithMaxSize(10),
    )
    
    return logger
}

func example1(logger *logging.Logger) {
    logger.Info("╔════════════════════════════════════════════════════╗")
    logger.Info("║  EXAMPLE 1: Function Name                          ║")
    logger.Info("╚════════════════════════════════════════════════════╝")
    logger.Info("")
    
    logger.Info("📝 Description / 설명")
    logger.Info("   What this example demonstrates")
    logger.Info("   이 예제가 시연하는 것")
    logger.Info("")
    
    logger.Info("📥 Input / 입력")
    logger.Info("   param1:", param1)
    logger.Info("   param2:", param2)
    logger.Info("")
    
    // Execute function
    result := {package}.Function(param1, param2)
    
    logger.Info("📤 Output / 출력")
    logger.Info("   result:", result)
    logger.Info("")
    
    logger.Info("✅ Example 1 Complete / 예제 1 완료")
    logger.Info("")
}
```

### Logging Best Practices / 로깅 모범 사례

1. **Be Extremely Detailed / 극도로 상세하게**
   - Log should serve as a manual / 로그가 매뉴얼 역할을 해야 함
   - User should understand without reading docs / 문서 없이도 이해 가능해야 함

2. **Use Visual Structure / 시각적 구조 사용**
   - Box drawings for sections / 섹션에 박스 그림
   - Emojis for clarity (📝, ✅, ⚠️, 📋, 📥, 📤) / 명확성을 위한 이모지

3. **Show All Steps / 모든 단계 표시**
   - Input values / 입력 값
   - Intermediate steps / 중간 단계
   - Output values / 출력 값
   - Any errors / 모든 에러

4. **Bilingual Throughout / 전체 이중 언어**
   - All explanations in both languages / 모든 설명을 두 언어로
   - Format: "English / 한글"

### Log Directory Structure / 로그 디렉토리 구조

```
logs/
├── {package}/
│   ├── example.log                    # Current log
│   ├── example-20251016-001.log       # Backup 1
│   ├── example-20251016-002.log       # Backup 2
│   └── ...
```

---

## 🔧 Git Workflow / Git 워크플로우

### Commit Message Format / 커밋 메시지 형식

**🚨 CRITICAL: ALL COMMIT MESSAGES MUST BE BILINGUAL (ENGLISH/KOREAN)**
**🚨 핵심: 모든 커밋 메시지는 반드시 영문/한글 병기**

```
<type>: <subject in English / 한글 제목> (<version>)

[optional body in English / 한글 본문]

🤖 Generated with Claude Code
Co-Authored-By: Claude <noreply@anthropic.com>
```

### Commit Types / 커밋 타입

| Type / 타입 | Usage / 사용 | Example / 예시 |
|-------------|-------------|---------------|
| **Feat** | New feature / 새 기능 | `Feat: Add Get function (v1.11.002)` |
| **Fix** | Bug fix / 버그 수정 | `Fix: Handle nil map (v1.11.003)` |
| **Docs** | Documentation / 문서 | `Docs: Update README (v1.11.004)` |
| **Test** | Tests / 테스트 | `Test: Add Get tests (v1.11.005)` |
| **Refactor** | Refactoring / 리팩토링 | `Refactor: Optimize Get (v1.11.006)` |
| **Chore** | Build, version / 빌드, 버전 | `Chore: Bump to v1.11.007` |
| **Perf** | Performance / 성능 | `Perf: Improve Get speed (v1.11.008)` |
| **Style** | Formatting / 포맷팅 | `Style: Format code (v1.11.009)` |

### Commit Examples / 커밋 예제

**✅ CORRECT - Bilingual / 올바른 예시 - 병기:**

```bash
# Version bump (ALWAYS FIRST)
git commit -m "Chore: Bump version to v1.11.045 / v1.11.045로 버전 증가"

# New feature
git commit -m "Feat: Add Get function to maputil / maputil에 Get 함수 추가 (v1.11.045)"

# Bug fix
git commit -m "Fix: Handle nil pointer in Get function / Get 함수에서 nil 포인터 처리 (v1.11.046)"

# Documentation
git commit -m "Docs: Add comprehensive workflow guide / 포괄적인 워크플로우 가이드 추가 (v1.11.047)"

# Test
git commit -m "Test: Add edge case tests for Get / Get 함수 엣지 케이스 테스트 추가 (v1.11.048)"

# Multiple changes with body
git commit -m "Feat: Complete maputil basic operations / maputil 기본 연산 완료 (v1.11.049)

- Added Get, Set, Delete functions / Get, Set, Delete 함수 추가
- Added comprehensive tests / 포괄적인 테스트 추가
- Added examples and documentation / 예제 및 문서 추가
- Updated README and CHANGELOG / README 및 CHANGELOG 업데이트

🤖 Generated with Claude Code
Co-Authored-By: Claude <noreply@anthropic.com>"
```

**❌ INCORRECT - English only / 잘못된 예시 - 영문만:**

```bash
# Don't do this / 이렇게 하지 마세요
git commit -m "Chore: Bump version to v1.11.045"
git commit -m "Feat: Add Get function to maputil (v1.11.045)"
```

### Push Workflow / 푸시 워크플로우

```bash
# 1. Ensure all tests pass
go test ./... -v

# 2. Ensure all changes are committed
git status

# 3. Push to remote
git push origin main

# 4. Verify on GitHub
# Check that all files are updated correctly
```

---

## ⚡ Quick Reference / 빠른 참조

### Every Task Checklist / 모든 작업 체크리스트

```
□ 1. Bump version in cfg/app.yaml
□ 2. Commit version bump
□ 3. Perform your work (code/docs)
□ 4. Run tests (go test ./... -v)
□ 5. Update CHANGELOG
□ 6. Commit with proper message
□ 7. Push to GitHub
```

### Common Commands / 자주 쓰는 명령어

```bash
# Version management
vi cfg/app.yaml  # Edit version

# Testing
go test ./... -v                      # All tests
go test ./{package} -v                # Package tests
go test ./{package} -cover            # With coverage
go test ./{package} -run TestFunc     # Specific test

# Building
go build ./...                        # Build all
go build ./{package}                  # Build package

# Examples
go run examples/{package}/main.go     # Run example

# Git
git add .
git commit -m "Type: Message (vX.Y.Z)"
git push origin main
```

### File Locations / 파일 위치

```
cfg/app.yaml                              # Version
docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md    # Patch changelog
{package}/README.md                       # Package docs
docs/{package}/USER_MANUAL.md             # User guide
docs/{package}/DEVELOPER_GUIDE.md         # Dev guide
examples/{package}/main.go                # Examples
logs/{package}/                           # Example logs
```

### Version Bump Quick Guide / 버전 증가 빠른 가이드

```yaml
# cfg/app.yaml

# BEFORE:
version: v1.11.044

# AFTER:
version: v1.11.045
```

```bash
# Commit it FIRST
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.11.045"

# NOW start your work
```

---

## 📝 Summary / 요약

### Golden Rules / 황금 규칙

1. **Always bump version FIRST / 항상 버전을 먼저 증가**
2. **Always update CHANGELOG / 항상 CHANGELOG 업데이트**
3. **Always test before commit / 항상 커밋 전에 테스트**
4. **Always document in both languages / 항상 두 언어로 문서화**

### Work Order / 작업 순서

```
Version Bump → Work → Test → CHANGELOG → Commit → Push
```

### Quality Standards / 품질 표준

- ✅ 60%+ test coverage / 60% 이상 테스트 커버리지
- ✅ Bilingual documentation / 이중 언어 문서화
- ✅ Comprehensive examples / 포괄적인 예제
- ✅ Detailed logging / 상세한 로깅

---

## 🎯 Conclusion / 결론

Following this guide ensures:

이 가이드를 따르면 다음을 보장합니다:

1. **Clean History / 깔끔한 이력** - Every version is properly tracked / 모든 버전이 적절히 추적됨
2. **Quality Code / 품질 코드** - All changes are tested / 모든 변경사항이 테스트됨
3. **Complete Documentation / 완전한 문서화** - Everything is documented / 모든 것이 문서화됨
4. **User Experience / 사용자 경험** - Examples and guides are comprehensive / 예제와 가이드가 포괄적

**Remember / 기억하세요**: The workflow is not a burden, it's a guarantee of quality.

워크플로우는 부담이 아니라 품질의 보증입니다.

**Happy Coding! / 즐거운 코딩!** 🚀

---

**Document Version / 문서 버전**: v1.11.045  
**Last Updated / 최종 업데이트**: 2025-10-16  
**Maintained By / 관리자**: go-utils team
