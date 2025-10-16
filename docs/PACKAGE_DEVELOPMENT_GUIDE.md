# Package Development Guide / 패키지 개발 가이드

Complete guide for developing new packages in go-utils.

go-utils에서 새로운 패키지를 개발하기 위한 완전한 가이드입니다.

**Version / 버전**: v1.10.004
**Last Updated / 최종 업데이트**: 2025-10-16

---

## Table of Contents / 목차

- [Overview / 개요](#overview--개요)
- [Branch Strategy / 브랜치 전략](#branch-strategy--브랜치-전략)
- [Version Management / 버전 관리](#version-management--버전-관리)
- [Development Workflow / 개발 워크플로우](#development-workflow--개발-워크플로우)
- [Unit Task Workflow / 단위 작업 워크플로우](#unit-task-workflow--단위-작업-워크플로우)
- [Example Code Guidelines / 예제 코드 가이드라인](#example-code-guidelines--예제-코드-가이드라인)
- [Logging Guidelines / 로깅 가이드라인](#logging-guidelines--로깅-가이드라인)
- [Documentation Standards / 문서화 표준](#documentation-standards--문서화-표준)
- [Testing Standards / 테스트 표준](#testing-standards--테스트-표준)
- [Git Commit Guidelines / Git 커밋 가이드라인](#git-commit-guidelines--git-커밋-가이드라인)

---

## Overview / 개요

This guide ensures consistency and quality across all packages in the go-utils library.

이 가이드는 go-utils 라이브러리의 모든 패키지에서 일관성과 품질을 보장합니다.

**Core Principles / 핵심 원칙**:
1. **Extreme Simplicity** - Reduce 20-30 lines to 1-2 lines / 20-30줄을 1-2줄로 줄이기
2. **Type Safety** - Use Go 1.18+ generics where appropriate / 적절한 경우 Go 1.18+ 제네릭 사용
3. **Zero Configuration** - Sensible defaults for 99% of use cases / 99% 사용 사례에 대한 합리적인 기본값
4. **Comprehensive Documentation** - Bilingual (English/Korean) / 이중 언어 (영문/한글)
5. **Test Coverage** - Aim for 80%+ coverage / 80% 이상 커버리지 목표

---

## Branch Strategy / 브랜치 전략

### Creating a New Package / 새 패키지 생성

When starting a new package, follow these steps:

새 패키지를 시작할 때 다음 단계를 따르세요:

1. **Create Feature Branch / 기능 브랜치 생성**
   ```bash
   # Format: feature/v{MAJOR}.{MINOR}.x-{package-name}
   git checkout -b feature/v1.11.x-httpserver
   ```

2. **Update Version in cfg/app.yaml / cfg/app.yaml에서 버전 업데이트**
   ```yaml
   app:
     name: go-utils
     version: v1.11.001  # Minor version + 1
   ```

3. **Document in Root README.md / 루트 README.md에 문서화**
   ```markdown
   ### 🚧 [httpserver](./httpserver/) - HTTP Server Utilities (In Development / 개발 중)

   **Status**: v1.11.x - In Development / 개발 중
   **Branch**: feature/v1.11.x-httpserver
   ```

4. **Document in CHANGELOG.md / CHANGELOG.md에 문서화**
   ```markdown
   ## [v1.11.x] - HTTP Server Utilities Package / HTTP 서버 유틸리티 패키지 (개발 중 / In Development)

   **Focus / 초점**: Extreme simplicity HTTP server utilities
   **Status / 상태**: In Development / 개발 중
   **Branch / 브랜치**: feature/v1.11.x-httpserver
   ```

### Concurrent Development / 동시 개발

When multiple packages are being developed simultaneously:

여러 패키지가 동시에 개발될 때:

1. **Check Current Versions / 현재 버전 확인**
   - Read README.md and CHANGELOG.md / README.md 및 CHANGELOG.md 읽기
   - Find the latest version in development / 개발 중인 최신 버전 찾기

2. **Choose Next Minor Version / 다음 마이너 버전 선택**
   ```
   Current in development: v1.11.x (httpserver)
   Your new package: v1.12.x (validation)
   ```

3. **Document Your Version / 버전 문서화**
   - Add to README.md with "In Development" status / "개발 중" 상태로 README.md에 추가
   - Add to CHANGELOG.md with branch name / 브랜치 이름과 함께 CHANGELOG.md에 추가

---

## Version Management / 버전 관리

### Standard Release Workflow / 표준 작업 순서

Every patch cycle must follow the same order to keep history consistent:

패치 사이클마다 다음 순서를 지켜야 이력이 일관됩니다:

1. **Bump cfg/app.yaml** – increment the patch or minor version before modifying code.  
   **cfg/app.yaml 버전 증가** – 코드를 손대기 전에 패치/마이너 버전을 올립니다.
2. **Implement code & docs** – write code and update documentation in the same cycle.  
   **코드 및 문서 작업** – 코드 변경과 문서 수정을 같은 사이클에서 처리합니다.
3. **Verify changes** – run `go test ./...` (또는 필요한 범위), 포매터 및 정적 분석을 수행합니다.  
   **변경 검증** – `go test ./...` 등 필요한 검증을 반드시 실행합니다.
4. **Record updates** – log the changes in package-level changelog(s) and the root `CHANGELOG.md`.  
   **변경 기록** – 패키지별 체인지로그와 루트 `CHANGELOG.md`에 내용을 적습니다.
5. **Commit & push** – only after steps 1–4 succeed, commit and push the branch.  
   **커밋 및 푸시** – 1~4단계가 완료된 뒤에만 커밋하고 푸시합니다.

> Tip / 팁: Large features can be split into multiple micro cycles; repeat steps 1–5 for each to maintain clean history.  
> 큰 기능은 여러 마이크로 사이클로 나누어 진행하고, 각 사이클마다 1–5 단계를 반복하면 기록이 깔끔해집니다.

### Version Format / 버전 형식

```
vMAJOR.MINOR.PATCH
```

**Example / 예**: v1.11.005

### Version Rules / 버전 규칙

1. **MAJOR Version / 메이저 버전**
   - Breaking changes / 호환성을 깨는 변경
   - API redesign / API 재설계
   - **Rarely changed** / 거의 변경되지 않음

2. **MINOR Version / 마이너 버전**
   - New package / 새 패키지
   - New major feature / 새로운 주요 기능
   - **Increment by 1 for each new package** / 각 새 패키지마다 1씩 증가

3. **PATCH Version / 패치 버전**
   - Bug fixes / 버그 수정
   - Documentation updates / 문서 업데이트
   - Small feature additions / 작은 기능 추가
   - **Increment before every unit task** / 모든 단위 작업 전에 증가

### Patch Version Workflow / 패치 버전 워크플로우

Before every unit task (function implementation, test, documentation):

모든 단위 작업(함수 구현, 테스트, 문서화) 전에:

1. **Increment Patch Version / 패치 버전 증가**
   ```yaml
   # cfg/app.yaml
   app:
     version: v1.11.001  # → v1.11.002
   ```

2. **Commit the Version Change / 버전 변경 커밋**
   ```bash
   git add cfg/app.yaml
   git commit -m "Chore: Bump version to v1.11.002"
   ```

---

## Development Workflow / 개발 워크플로우

### Phase 1: Planning / 계획

**Step 1: Design Document / 설계 문서**

Create `docs/{package}/DESIGN_PLAN.md`:

```markdown
# {Package} Design Plan / 설계 계획

## Package Overview / 패키지 개요
- Purpose / 목적
- Target use cases / 목표 사용 사례
- Design principles / 설계 원칙

## Architecture / 아키텍처
- File structure / 파일 구조
- Core types / 핵심 타입
- Function categories / 함수 카테고리

## API Design / API 설계
- Function signatures / 함수 시그니처
- Options pattern / 옵션 패턴
- Error handling / 에러 처리

## Examples / 예제
- Common use cases / 일반적인 사용 사례
- Before vs After / 이전 vs 이후
```

**Step 2: Work Plan / 작업 계획서**

Create `docs/{package}/WORK_PLAN.md`:

```markdown
# {Package} Work Plan / 작업 계획

## Development Phases / 개발 단계

### Phase 1: Core Functions (v1.11.001-005)
- [ ] Function1 (v1.11.001)
- [ ] Function2 (v1.11.002)
- [ ] Function3 (v1.11.003)

### Phase 2: Advanced Features (v1.11.006-010)
- [ ] Feature1 (v1.11.006)
- [ ] Feature2 (v1.11.007)

### Phase 3: Documentation & Polish (v1.11.011-015)
- [ ] USER_MANUAL.md
- [ ] DEVELOPER_GUIDE.md
- [ ] Example code
```

### Phase 2: Implementation / 구현

Follow the **Unit Task Workflow** (see next section).

**단위 작업 워크플로우**를 따르세요 (다음 섹션 참조).

### Phase 3: Finalization / 마무리

After all planned features are implemented:

모든 계획된 기능이 구현된 후:

1. **Code Review / 코드 리뷰**
   - Review all code for consistency / 일관성을 위해 모든 코드 검토
   - Check for code duplication / 코드 중복 확인
   - Ensure error handling is consistent / 에러 처리가 일관되도록 확인

2. **Documentation Review / 문서 검토**
   - Update README.md / README.md 업데이트
   - Complete USER_MANUAL.md / USER_MANUAL.md 완성
   - Complete DEVELOPER_GUIDE.md / DEVELOPER_GUIDE.md 완성
   - Verify all examples work / 모든 예제가 작동하는지 확인

3. **Testing / 테스트**
   - Ensure 80%+ coverage / 80% 이상 커버리지 확인
   - Run all tests / 모든 테스트 실행
   - Run benchmarks / 벤치마크 실행

4. **Final Commit / 최종 커밋**
   ```bash
   git add .
   git commit -m "Feat: Complete {package} package (v1.11.015)"
   git push origin feature/v1.11.x-{package}
   ```

5. **Create Pull Request / Pull Request 생성**
   - Merge to main branch / main 브랜치에 병합
   - Tag the release / 릴리스 태그

---

## Unit Task Workflow / 단위 작업 워크플로우

Each unit task (e.g., implementing one function or feature) follows this workflow:

각 단위 작업(예: 하나의 함수 또는 기능 구현)은 다음 워크플로우를 따릅니다:

### Step 1: Increment Patch Version / 패치 버전 증가

```bash
# Edit cfg/app.yaml
# v1.11.001 → v1.11.002

git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.11.002"
```

### Step 2: Code Implementation / 코드 구현

Implement the function or feature in the appropriate file.

적절한 파일에 함수 또는 기능을 구현합니다.

**Guidelines / 가이드라인**:
- Follow existing code style / 기존 코드 스타일 따르기
- Add bilingual comments (English/Korean) / 이중 언어 주석 추가 (영문/한글)
- Use consistent naming / 일관된 명명 사용
- Handle errors properly / 에러를 적절히 처리

**Example / 예제**:
```go
// Get retrieves a value from the map by key.
// If the key does not exist, returns the zero value and false.
// Get은 키로 맵에서 값을 검색합니다.
// 키가 존재하지 않으면 제로 값과 false를 반환합니다.
func Get[K comparable, V any](m map[K]V, key K) (V, bool) {
    val, ok := m[key]
    return val, ok
}
```

### Step 3: Test Code / 테스트 코드

Write comprehensive tests in `{package}_test.go`.

`{package}_test.go`에 포괄적인 테스트를 작성합니다.

**Guidelines / 가이드라인**:
- Test all edge cases / 모든 엣지 케이스 테스트
- Use table-driven tests / 테이블 기반 테스트 사용
- Add sub-tests for clarity / 명확성을 위해 하위 테스트 추가
- Test error cases / 에러 케이스 테스트
- Add benchmarks for performance-critical functions / 성능 중요 함수에 벤치마크 추가

**Example / 예제**:
```go
func TestGet(t *testing.T) {
    tests := []struct {
        name     string
        input    map[string]int
        key      string
        wantVal  int
        wantOk   bool
    }{
        {
            name:    "existing key",
            input:   map[string]int{"a": 1, "b": 2},
            key:     "a",
            wantVal: 1,
            wantOk:  true,
        },
        {
            name:    "non-existing key",
            input:   map[string]int{"a": 1},
            key:     "b",
            wantVal: 0,
            wantOk:  false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotVal, gotOk := Get(tt.input, tt.key)
            if gotVal != tt.wantVal || gotOk != tt.wantOk {
                t.Errorf("Get() = (%v, %v), want (%v, %v)",
                    gotVal, gotOk, tt.wantVal, tt.wantOk)
            }
        })
    }
}
```

### Step 4: Documentation / 문서 작업

Update the following documents:

다음 문서들을 업데이트합니다:

1. **Package README.md**
   - Add function to API reference / API 참조에 함수 추가
   - Add usage example / 사용 예제 추가

2. **USER_MANUAL.md** (if exists / 존재하는 경우)
   - Add detailed usage guide / 상세 사용 가이드 추가
   - Add common use cases / 일반적인 사용 사례 추가

3. **DEVELOPER_GUIDE.md** (if exists / 존재하는 경우)
   - Add implementation details / 구현 세부사항 추가
   - Add design decisions / 설계 결정사항 추가

### Step 5: Example Code / 예제 코드

Add example to `examples/{package}/main.go`.

`examples/{package}/main.go`에 예제를 추가합니다.

**See [Example Code Guidelines](#example-code-guidelines--예제-코드-가이드라인) below.**

아래 [예제 코드 가이드라인](#example-code-guidelines--예제-코드-가이드라인)을 참조하세요.

### Step 6: CHANGELOG / 변경 기록

Update `docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md`:

`docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md`를 업데이트합니다:

```markdown
## [v1.11.002] - 2025-10-16

### Added
- Added `Get` function for retrieving values from maps
- Added `Get` 함수로 맵에서 값 검색

### Changed
- N/A

### Fixed
- N/A
```

### Step 7: Compile & Test / 컴파일 및 테스트

```bash
# Compile / 컴파일
go build ./...

# Run tests / 테스트 실행
go test ./{package} -v

# Check coverage / 커버리지 확인
go test ./{package} -cover

# Run example / 예제 실행
go run examples/{package}/main.go
```

**All must pass before proceeding** / 진행하기 전에 모두 통과해야 함

### Step 8: Git Commit & Push / Git 커밋 및 푸시

```bash
git add .
git commit -m "Feat: Add Get function to {package} (v1.11.002)"
git push origin feature/v1.11.x-{package}
```

**Repeat for Next Unit Task / 다음 단위 작업 반복**

---

## Example Code Guidelines / 예제 코드 가이드라인

### Structure / 구조

All examples must follow this structure:

모든 예제는 이 구조를 따라야 합니다:

```go
package main

import (
    "github.com/arkd0ng/go-utils/logging"
    "github.com/arkd0ng/go-utils/{package}"
)

func main() {
    // Initialize logger / 로거 초기화
    logger := initLogger()
    defer logger.Close()

    // Print banner / 배너 출력
    logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
    logger.Info("║         {Package} Package - Comprehensive Examples & Manual               ║")
    logger.Info("║         {Package} 패키지 - 종합 예제 및 매뉴얼                              ║")
    logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
    logger.Info("")

    // Package info / 패키지 정보
    logger.Info("📋 Package Information / 패키지 정보")
    logger.Info("   Package Name: github.com/arkd0ng/go-utils/{package}")
    logger.Info("   Version: " + {package}.Version)
    logger.Info("   Description: {description}")
    logger.Info("   설명: {korean description}")
    logger.Info("")

    // Example 1 / 예제 1
    example1(logger)

    // Example 2 / 예제 2
    example2(logger)

    // ... more examples
}

func initLogger() *logging.Logger {
    // Logger initialization with backup / 백업과 함께 로거 초기화
    // See logging section below / 아래 로깅 섹션 참조
}

func example1(logger *logging.Logger) {
    logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
    logger.Info("║  EXAMPLE 1: Function Name / 예제 1: 함수 이름                              ║")
    logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
    logger.Info("")

    logger.Info("📝 1.1 Description / 설명")
    logger.Info("   Function: {package}.FunctionName()")
    logger.Info("   Description: What it does / 무엇을 하는지")
    logger.Info("   Use Case: When to use it / 언제 사용하는지")

    // Implementation / 구현
    result := {package}.FunctionName(args)

    logger.Info("   ✅ Result / 결과: key=value")
    logger.Info("")
}
```

### Requirements / 요구사항

1. **Cover All Functions / 모든 함수 커버**
   - Every implemented function must have an example / 모든 구현된 함수는 예제가 있어야 함
   - Group related functions / 관련된 함수들을 그룹화

2. **Detailed Logging / 상세 로깅**
   - Log should be detailed enough to serve as a manual / 로그만으로도 매뉴얼 역할을 할 수 있을 만큼 상세해야 함
   - Explain what each example does / 각 예제가 무엇을 하는지 설명
   - Show input and output / 입력과 출력 표시
   - Explain any important concepts / 중요한 개념 설명

3. **Bilingual / 이중 언어**
   - All explanations in English and Korean / 모든 설명은 영문과 한글로
   - Use format: "English / 한글"

4. **Visual Structure / 시각적 구조**
   - Use box drawings for sections / 섹션에 박스 그림 사용
   - Use emojis for clarity (📝, ✅, ⚠️, 📋, etc.) / 명확성을 위해 이모지 사용

---

## Logging Guidelines / 로깅 가이드라인

### Logger Initialization / 로거 초기화

All examples must use the `logging` package:

모든 예제는 `logging` 패키지를 사용해야 합니다:

```go
func initLogger() *logging.Logger {
    // Create logs directory / logs 디렉토리 생성
    if err := os.MkdirAll("logs", 0755); err != nil {
        log.Fatal(err)
    }

    // Backup previous log file / 이전 로그 파일 백업
    logFile := "logs/{package}-example.log"
    if _, err := os.Stat(logFile); err == nil {
        backupName := fmt.Sprintf("logs/{package}-example-%s.log",
            time.Now().Format("20060102-150405"))
        os.Rename(logFile, backupName)
        fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
    }

    // Create logger / 로거 생성
    logger, err := logging.New(
        logging.WithFilePath(logFile),
        logging.WithLevel(logging.LevelInfo),
        logging.WithMaxSize(10),
        logging.WithMaxBackups(5),
        logging.WithMaxAge(30),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Print banner / 배너 출력
    cfg := config.LoadAppConfig()
    logger.Banner(cfg.App.Name, cfg.App.Version)

    return logger
}
```

### Log Directory / 로그 디렉토리

```
logs/
├── {package}-example.log              # Current log / 현재 로그
├── {package}-example-20251016-001.log # Backup 1 / 백업 1
├── {package}-example-20251016-002.log # Backup 2 / 백업 2
└── ...
```

### Log Format / 로그 형식

```
2025-10-16 00:37:28 [INFO] ╔════════════════════════════════════════════════════════════════════════════╗
2025-10-16 00:37:28 [INFO] ║  EXAMPLE 1: Function Name / 예제 1: 함수 이름                              ║
2025-10-16 00:37:28 [INFO] ╚════════════════════════════════════════════════════════════════════════════╝
2025-10-16 00:37:28 [INFO]
2025-10-16 00:37:28 [INFO] 📝 1.1 Description / 설명
2025-10-16 00:37:28 [INFO]    Function: package.FunctionName()
2025-10-16 00:37:28 [INFO]    Description: What it does / 무엇을 하는지
2025-10-16 00:37:28 [INFO]    ✅ Result / 결과: key=value
```

### Logging Best Practices / 로깅 모범 사례

1. **Be Extremely Detailed / 극도로 상세하게**
   - Log should be self-documenting / 로그는 자체 문서화되어야 함
   - User should understand without reading other docs / 사용자가 다른 문서를 읽지 않고도 이해할 수 있어야 함

2. **Show All Steps / 모든 단계 표시**
   - Log input values / 입력 값 로그
   - Log intermediate steps / 중간 단계 로그
   - Log results / 결과 로그
   - Log any errors / 모든 에러 로그

3. **Use Structured Logging / 구조화된 로깅 사용**
   ```go
   logger.Info("Processing user", "id", userID, "name", userName)
   ```

4. **Group Related Operations / 관련 작업 그룹화**
   - Use visual separators / 시각적 구분자 사용
   - Use consistent formatting / 일관된 형식 사용

---

## Documentation Standards / 문서화 표준

### Required Documents / 필수 문서

Every package must have:

모든 패키지는 다음을 가져야 합니다:

1. **{package}/README.md**
   - Package overview / 패키지 개요
   - Installation / 설치
   - Quick start / 빠른 시작
   - API reference / API 참조
   - Examples / 예제

2. **docs/{package}/DESIGN_PLAN.md**
   - Architecture / 아키텍처
   - Design decisions / 설계 결정사항
   - Trade-offs / 트레이드오프

3. **docs/{package}/WORK_PLAN.md**
   - Development phases / 개발 단계
   - Task breakdown / 작업 분류
   - Progress tracking / 진행 상황 추적

4. **docs/{package}/USER_MANUAL.md**
   - Comprehensive user guide / 포괄적인 사용자 가이드
   - All functions documented / 모든 함수 문서화
   - Usage patterns / 사용 패턴
   - Best practices / 모범 사례
   - Troubleshooting / 문제 해결

5. **docs/{package}/DEVELOPER_GUIDE.md**
   - Internal architecture / 내부 아키텍처
   - Implementation details / 구현 세부사항
   - Contributing guidelines / 기여 가이드라인
   - Testing guide / 테스트 가이드

6. **docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md**
   - Detailed patch-level changes / 상세 패치 레벨 변경사항
   - Date, version, changes / 날짜, 버전, 변경사항

### Bilingual Documentation / 이중 언어 문서화

**🚨 CRITICAL RULE: ALL DOCUMENTATION MUST BE BILINGUAL (ENGLISH/KOREAN)**
**🚨 핵심 규칙: 모든 문서는 반드시 영문/한글 병기**

All documentation must be bilingual (English/Korean):

모든 문서는 이중 언어(영문/한글)여야 합니다:

#### What Must Be Bilingual / 병기가 필요한 항목

1. **All Documentation Files / 모든 문서 파일**
   - README.md files / README.md 파일
   - DESIGN_PLAN.md / 설계 계획서
   - WORK_PLAN.md / 작업 계획서
   - USER_MANUAL.md / 사용자 매뉴얼
   - DEVELOPER_GUIDE.md / 개발자 가이드
   - CHANGELOG files / 변경 로그 파일
   - Any .md files in docs/ / docs/ 폴더의 모든 .md 파일

2. **Code Comments / 코드 주석**
   - Package-level comments / 패키지 레벨 주석
   - Function/method documentation / 함수/메서드 문서화
   - Important inline comments / 중요한 인라인 주석
   - Example code comments / 예제 코드 주석

3. **Git Commit Messages / Git 커밋 메시지**
   - Subject line must be bilingual / 제목은 반드시 병기
   - Body can be bilingual or English / 본문은 병기 또는 영문

4. **Log Messages / 로그 메시지**
   - All log messages must be bilingual / 모든 로그 메시지는 병기
   - Currently bilingual, will be separated later / 현재는 병기, 추후 분리 예정

5. **Error Messages / 에러 메시지**
   - User-facing error messages must be bilingual / 사용자 대상 에러 메시지는 병기
   - Internal error messages can be English / 내부 에러 메시지는 영문 가능

#### What Can Be English-Only / 영문만 사용 가능한 항목

1. **Personal Notes / 개인 노트**
   - CLAUDE.md (AI assistant guidance / AI 어시스턴트 가이드)
   - todo.md (personal task list / 개인 작업 목록)
   - Private development notes / 비공개 개발 노트

2. **Variable/Function Names / 변수/함수 이름**
   - All code identifiers must be in English / 모든 코드 식별자는 영문
   - Comments must be bilingual / 주석은 병기

#### Documentation Format Examples / 문서 형식 예제

**Section Headers / 섹션 헤더**:
```markdown
## Section Title / 섹션 제목
### Subsection / 하위 섹션
```

**Paragraphs / 문단**:
```markdown
## Overview / 개요

This package provides utility functions for string manipulation.

이 패키지는 문자열 조작을 위한 유틸리티 함수를 제공합니다.
```

**Lists / 목록**:
```markdown
**Features / 기능**:
- Feature one / 기능 1
- Feature two / 기능 2
```

**Tables / 테이블**:
```markdown
| Function / 함수 | Description / 설명 |
|-----------------|-------------------|
| `Get()` | Gets a value / 값을 가져옴 |
```

**Code Examples / 코드 예제**:
```go
// GetValue retrieves a value from the map.
// It returns the value and a boolean indicating if the key exists.
// GetValue는 맵에서 값을 검색합니다.
// 키가 존재하는지 나타내는 불리언과 함께 값을 반환합니다.
func GetValue(m map[string]int, key string) (int, bool) {
    // Check if key exists / 키 존재 여부 확인
    val, ok := m[key]
    return val, ok
}
```

### Code Comments / 코드 주석

```go
// FunctionName does something useful.
// It takes parameter X and returns Y.
// FunctionName은 유용한 작업을 수행합니다.
// X 매개변수를 받아 Y를 반환합니다.
func FunctionName(x int) int {
    // Implementation / 구현
    return x * 2
}
```

---

## Testing Standards / 테스트 표준

### Coverage Requirements / 커버리지 요구사항

- **Minimum**: 60% overall coverage / 최소: 전체 60% 커버리지
- **Target**: 80%+ coverage / 목표: 80% 이상 커버리지
- **Critical functions**: 100% coverage / 중요 함수: 100% 커버리지

### Test Categories / 테스트 카테고리

1. **Unit Tests / 단위 테스트**
   - Test each function independently / 각 함수를 독립적으로 테스트
   - Test all edge cases / 모든 엣지 케이스 테스트
   - Test error conditions / 에러 조건 테스트

2. **Integration Tests / 통합 테스트**
   - Test functions working together / 함수들이 함께 작동하는지 테스트
   - Test realistic scenarios / 현실적인 시나리오 테스트

3. **Benchmarks / 벤치마크**
   - Add benchmarks for performance-critical functions / 성능 중요 함수에 벤치마크 추가
   - Compare with standard library / 표준 라이브러리와 비교

### Test Structure / 테스트 구조

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    OutputType
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   ValidInput,
            want:    ExpectedOutput,
            wantErr: false,
        },
        {
            name:    "invalid input",
            input:   InvalidInput,
            want:    ZeroValue,
            wantErr: true,
        },
        // More test cases / 더 많은 테스트 케이스
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FunctionName(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("FunctionName() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

---

## Git Commit Guidelines / Git 커밋 가이드라인

### Commit Message Format / 커밋 메시지 형식

**🚨 CRITICAL: ALL COMMIT MESSAGES MUST BE BILINGUAL (ENGLISH/KOREAN)**
**🚨 핵심: 모든 커밋 메시지는 반드시 영문/한글 병기**

```
<type>: <subject in English / 한글 제목> (<version>)

[optional body in English / 한글 본문]

🤖 Generated with [Claude Code](https://claude.com/claude-code)

Co-Authored-By: Claude <noreply@anthropic.com>
```

### Commit Types / 커밋 타입

- **Feat**: New feature / 새 기능
- **Fix**: Bug fix / 버그 수정
- **Docs**: Documentation changes / 문서 변경
- **Refactor**: Code refactoring / 코드 리팩토링
- **Test**: Test additions/modifications / 테스트 추가/수정
- **Chore**: Build, configuration, version bumps / 빌드, 설정, 버전 증가
- **Perf**: Performance improvements / 성능 개선
- **Style**: Code style changes (formatting) / 코드 스타일 변경 (포맷팅)

### Examples / 예제

**✅ CORRECT - Bilingual / 올바른 예시 - 병기**:
```bash
# Version bump / 버전 증가
git commit -m "Chore: Bump version to v1.11.002 / v1.11.002로 버전 증가"

# New feature / 새 기능
git commit -m "Feat: Add Get function to maputil / maputil에 Get 함수 추가 (v1.11.002)"

# Bug fix / 버그 수정
git commit -m "Fix: Handle nil map in Get function / Get 함수에서 nil 맵 처리 (v1.11.003)"

# Documentation / 문서
git commit -m "Docs: Update maputil README with Get example / Get 예제로 maputil README 업데이트 (v1.11.004)"

# Test / 테스트
git commit -m "Test: Add comprehensive tests for Get function / Get 함수에 대한 포괄적인 테스트 추가 (v1.11.005)"

# Multiple line commit / 여러 줄 커밋
git commit -m "Feat: Add error handling utilities / 에러 처리 유틸리티 추가 (v1.12.010)

- Add Wrap function for error wrapping / 에러 래핑용 Wrap 함수 추가
- Add GetCode for error code extraction / 에러 코드 추출용 GetCode 추가
- Add comprehensive tests / 포괄적인 테스트 추가

🤖 Generated with Claude Code
Co-Authored-By: Claude <noreply@anthropic.com>"
```

**❌ INCORRECT - English only / 잘못된 예시 - 영문만**:
```bash
# Don't do this / 이렇게 하지 마세요
git commit -m "Chore: Bump version to v1.11.002"
git commit -m "Feat: Add Get function to maputil (v1.11.002)"
```

### Commit Message Best Practices / 커밋 메시지 모범 사례

1. **Keep it concise but descriptive / 간결하지만 설명적으로**
   - Subject line: 50-72 characters / 제목: 50-72자
   - Body: Wrap at 72 characters / 본문: 72자에서 줄바꿈

2. **Use imperative mood in English / 영문은 명령형으로**
   - "Add feature" not "Added feature" / "Add feature"이지 "Added feature"가 아님
   - "Fix bug" not "Fixed bug" / "Fix bug"이지 "Fixed bug"가 아님

3. **Always include version for code changes / 코드 변경시 항상 버전 포함**
   - (v1.11.002) at the end of subject / 제목 끝에 (v1.11.002)

4. **Use body for detailed explanation / 상세 설명은 본문에**
   - Why the change was made / 왜 변경했는지
   - What was changed / 무엇이 변경되었는지
   - Any breaking changes / 호환성 파괴 변경사항

### Commit Frequency / 커밋 빈도

Commit after each step in the unit task workflow:

단위 작업 워크플로우의 각 단계 후에 커밋:

1. After version bump / 버전 증가 후
2. After code implementation / 코드 구현 후
3. After test implementation / 테스트 구현 후
4. After documentation / 문서화 후
5. After example code / 예제 코드 후
6. Final commit with all changes / 모든 변경사항이 포함된 최종 커밋

---

## Summary Checklist / 요약 체크리스트

Before completing a package:

패키지를 완료하기 전에:

### Planning / 계획
- [ ] Created feature branch / 기능 브랜치 생성
- [ ] Updated version in cfg/app.yaml / cfg/app.yaml에서 버전 업데이트
- [ ] Documented in README.md and CHANGELOG.md / README.md 및 CHANGELOG.md에 문서화
- [ ] Created DESIGN_PLAN.md / DESIGN_PLAN.md 생성
- [ ] Created WORK_PLAN.md / WORK_PLAN.md 생성

### Implementation / 구현
- [ ] All planned functions implemented / 모든 계획된 함수 구현
- [ ] All functions have tests / 모든 함수에 테스트 있음
- [ ] Test coverage ≥ 60% / 테스트 커버리지 ≥ 60%
- [ ] All tests pass / 모든 테스트 통과
- [ ] Benchmarks added for critical functions / 중요 함수에 벤치마크 추가

### Documentation / 문서화
- [ ] Package README.md complete / 패키지 README.md 완성
- [ ] USER_MANUAL.md complete / USER_MANUAL.md 완성
- [ ] DEVELOPER_GUIDE.md complete / DEVELOPER_GUIDE.md 완성
- [ ] All functions documented / 모든 함수 문서화
- [ ] All comments bilingual / 모든 주석 이중 언어

### Examples / 예제
- [ ] All functions have examples / 모든 함수에 예제 있음
- [ ] Examples use logging package / 예제가 logging 패키지 사용
- [ ] Logs are detailed and bilingual / 로그가 상세하고 이중 언어
- [ ] Log backup implemented / 로그 백업 구현

### Finalization / 마무리
- [ ] CHANGELOG updated / CHANGELOG 업데이트
- [ ] Code review completed / 코드 리뷰 완료
- [ ] All commits follow guidelines / 모든 커밋이 가이드라인 따름
- [ ] Ready for merge / 병합 준비 완료

---

## Conclusion / 결론

Following this guide ensures:

이 가이드를 따르면 다음을 보장합니다:

1. **Consistency** across all packages / 모든 패키지에서 일관성
2. **Quality** with comprehensive tests and documentation / 포괄적인 테스트 및 문서화로 품질
3. **Maintainability** with clear structure and guidelines / 명확한 구조 및 가이드라인으로 유지보수성
4. **User Experience** with detailed examples and bilingual support / 상세한 예제 및 이중 언어 지원으로 사용자 경험

**Happy Coding! / 즐거운 코딩!** 🚀
