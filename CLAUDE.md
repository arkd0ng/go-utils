# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

## 🚨 CRITICAL - MUST READ FIRST / 필수 읽기

**BEFORE ANY WORK, READ THESE DOCUMENTS / 모든 작업 전에 이 문서들을 읽으세요:**

### 📋 Core Development Guides / 핵심 개발 가이드

1. **[DEVELOPMENT_WORKFLOW_GUIDE.md](./docs/DEVELOPMENT_WORKFLOW_GUIDE.md)** ⭐ **MOST IMPORTANT / 가장 중요**
   - Complete workflow for all development tasks / 모든 개발 작업의 완전한 워크플로우
   - Critical rules and standard work cycle / 핵심 규칙 및 표준 작업 사이클
   - **READ THIS FIRST for any task** / 모든 작업 시 가장 먼저 읽을 것

2. **[PACKAGE_DEVELOPMENT_GUIDE.md](./docs/PACKAGE_DEVELOPMENT_GUIDE.md)** ⭐ **ESSENTIAL / 필수**
   - Package development standards and workflow / 패키지 개발 표준 및 워크플로우
   - Branch strategy, version management, unit task workflow / 브랜치 전략, 버전 관리, 단위 작업 워크플로우
   - Example code and logging guidelines / 예제 코드 및 로깅 가이드라인

3. **[CODE_TEST_MAKE_GUIDE.md](./docs/CODE_TEST_MAKE_GUIDE.md)**
   - Testing standards and guidelines / 테스트 표준 및 가이드라인
   - Test structure and coverage requirements / 테스트 구조 및 커버리지 요구사항

4. **[EXAMPLE_CODE_GUIDE.md](./docs/EXAMPLE_CODE_GUIDE.md)**
   - Example code structure and requirements / 예제 코드 구조 및 요구사항
   - Logging best practices / 로깅 모범 사례

### 🔄 Standard Work Cycle / 표준 작업 사이클

**EVERY task follows this exact order / 모든 작업은 이 순서를 정확히 따름:**

```
1. Version Bump (cfg/app.yaml) / 버전 증가
   ↓
2. Perform Work (Code/Docs) / 작업 수행
   ↓
3. Test & Verify (go test ./...) / 테스트 및 검증
   ↓
4. Update CHANGELOG / CHANGELOG 업데이트
   ↓
5. Git Commit & Push / Git 커밋 및 푸시
```

**❌ NEVER / 절대 금지:**
- Skip version bump before work / 작업 전 버전 증가 생략
- Skip CHANGELOG update / CHANGELOG 업데이트 생략
- Push without testing / 테스트 없이 푸시
- Skip documentation / 문서화 생략

---

## 📦 Project Overview / 프로젝트 개요

**Repository**: `github.com/arkd0ng/go-utils`  
**Current Version**: v1.11.046 (from cfg/app.yaml)  
**Go Version**: 1.24.6  
**License**: MIT

### Purpose / 목적

Modular collection of utility packages for Golang development. Each subpackage is independent and can be imported individually.

Golang 개발을 위한 모듈화된 유틸리티 패키지 모음입니다. 각 서브패키지는 독립적이며 개별적으로 import할 수 있습니다.

### Design Principles / 설계 원칙

1. **Extreme Simplicity** - 20-30 lines → 1-2 lines / 20-30줄 → 1-2줄
2. **Independence** - No cross-package dependencies / 패키지 간 의존성 없음
3. **Bilingual** - All docs in English/Korean / 모든 문서 영문/한글
4. **Type Safety** - Go 1.18+ generics where appropriate / 적절한 경우 제네릭 사용
5. **Zero Config** - Sensible defaults for 99% cases / 99% 사례에 대한 합리적 기본값

---

## 📚 Package Architecture / 패키지 구조

### Current Packages / 현재 패키지

```
go-utils/
├── random/          # Cryptographically secure random strings (14 methods)
├── logging/         # Structured logging with file rotation
├── database/
│   ├── mysql/      # Extremely simple MySQL client (30 lines → 2 lines)
│   └── redis/      # Extremely simple Redis client (20 lines → 2 lines)
├── stringutil/     # String utilities (53 functions, 9 categories)
├── timeutil/       # Time/date utilities (114 functions, 10 categories)
├── sliceutil/      # Slice utilities (95 functions, 14 categories)
├── maputil/        # Map utilities (99 functions, 14 categories)
├── fileutil/       # File/path utilities (~91 functions, 12 categories)
└── websvrutil/     # Web server utilities (comprehensive features)
```

### Package Overview / 패키지 개요

| Package | Version | Functions | Description |
|---------|---------|-----------|-------------|
| **random** | v1.0.x | 14 methods | Crypto-safe random string generation |
| **logging** | v1.1.x | Full logging | Structured logging + file rotation |
| **mysql** | v1.3.x | 3 API levels | Simple API, Query Builder, Raw SQL |
| **redis** | v1.4.x | 60+ methods | String, Hash, List, Set, ZSet, Key ops |
| **stringutil** | v1.5.x | 53 functions | Unicode-safe string manipulation |
| **timeutil** | v1.6.x | 114 functions | KST-default time utilities |
| **sliceutil** | v1.7.x | 95 functions | Type-safe generic slice operations |
| **maputil** | v1.8.x | 99 functions | Type-safe generic map operations |
| **fileutil** | v1.9.x | ~91 functions | Cross-platform file/path utilities |
| **websvrutil** | v1.10.x | Comprehensive | HTTP server framework |

**For detailed package architecture, see:** / 상세한 패키지 아키텍처는 다음 참조:
- Each package's `README.md` / 각 패키지의 `README.md`
- `docs/{package}/USER_MANUAL.md` / 사용자 매뉴얼
- `docs/{package}/DEVELOPER_GUIDE.md` / 개발자 가이드

---

## 🔢 Version Management / 버전 관리

### Version Format / 버전 형식

```
vMAJOR.MINOR.PATCH
```

- **MAJOR**: Breaking changes (rarely) / 호환성 깨짐 (드물게)
- **MINOR**: New package / 새 패키지
- **PATCH**: Every unit task / 모든 단위 작업

### Version Rules / 버전 규칙

**Increment BEFORE every task / 모든 작업 전에 증가:**

```bash
# Edit cfg/app.yaml
version: v1.11.046  # Increment this / 이것을 증가

# Commit version bump FIRST
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.11.047"

# NOW start your work
```

### Version History / 버전 히스토리

- **v1.0.x** - Random package
- **v1.1.x** - Logging package
- **v1.2.x** - Documentation (Random, Logging)
- **v1.3.x** - MySQL package
- **v1.4.x** - Redis package
- **v1.5.x** - Stringutil package
- **v1.6.x** - Timeutil package
- **v1.7.x** - Sliceutil package
- **v1.8.x** - Maputil package
- **v1.9.x** - Fileutil package
- **v1.10.x** - Websvrutil package
- **v1.11.x** - Current development / 현재 개발

---

## 📝 Documentation Standards / 문서화 표준

### Required Documents / 필수 문서

Every package must have / 모든 패키지는 다음을 가져야 함:

1. **{package}/README.md** - Quick start and API reference
2. **docs/{package}/USER_MANUAL.md** - Comprehensive user guide
3. **docs/{package}/DEVELOPER_GUIDE.md** - Architecture and internals
4. **docs/{package}/DESIGN_PLAN.md** - Design decisions (for new packages)
5. **docs/{package}/WORK_PLAN.md** - Development phases (for new packages)
6. **examples/{package}/main.go** - Executable examples with logging

### Bilingual Format / 이중 언어 형식

**All documentation MUST be bilingual:**

```markdown
## Section Title / 섹션 제목

English description first.

한글 설명 다음.

**Example / 예제:**
```go
// English comment / 한글 주석
code here
```
```

---

## 🧪 Testing Standards / 테스트 표준

### Coverage Requirements / 커버리지 요구사항

- **Minimum**: 60% / 최소: 60%
- **Target**: 80%+ / 목표: 80% 이상
- **Critical functions**: 100% / 중요 함수: 100%

### Test Categories / 테스트 카테고리

1. **Unit Tests** - Each function independently / 각 함수 독립적으로
2. **Integration Tests** - Functions working together / 함수들의 협동
3. **Benchmarks** - Performance-critical functions / 성능 중요 함수

**See details in:** [CODE_TEST_MAKE_GUIDE.md](./docs/CODE_TEST_MAKE_GUIDE.md)

---

## 💡 Example Code Standards / 예제 코드 표준

### Structure / 구조

All examples follow this template / 모든 예제는 이 템플릿을 따름:

```go
package main

import (
    "github.com/arkd0ng/go-utils/logging"
    "github.com/arkd0ng/go-utils/{package}"
)

func main() {
    // 1. Initialize logger with backup
    logger := initLogger()
    defer logger.Close()
    
    // 2. Print banner
    printBanner(logger)
    
    // 3. Run examples
    example1(logger)
    example2(logger)
}
```

### Logging Requirements / 로깅 요구사항

- Use `logging` package / logging 패키지 사용
- Save to `logs/{package}/` / logs/{package}/에 저장
- Backup previous logs / 이전 로그 백업
- Extremely detailed / 극도로 상세하게
- Bilingual throughout / 전체 이중 언어

**See details in:** [EXAMPLE_CODE_GUIDE.md](./docs/EXAMPLE_CODE_GUIDE.md)

---

## 🔄 Git Workflow / Git 워크플로우

### Commit Message Format / 커밋 메시지 형식

```
<type>: <subject> (<version>)

[optional body]

🤖 Generated with Claude Code
Co-Authored-By: Claude <noreply@anthropic.com>
```

### Commit Types / 커밋 타입

- **Feat**: New feature / 새 기능
- **Fix**: Bug fix / 버그 수정
- **Docs**: Documentation / 문서
- **Test**: Tests / 테스트
- **Refactor**: Refactoring / 리팩토링
- **Chore**: Build, version / 빌드, 버전
- **Perf**: Performance / 성능
- **Style**: Formatting / 포맷팅

### Example Commits / 커밋 예제

```bash
git commit -m "Chore: Bump version to v1.11.046"
git commit -m "Feat: Add Get function to maputil (v1.11.046)"
git commit -m "Docs: Update USER_MANUAL (v1.11.046)"
git commit -m "Test: Add comprehensive tests (v1.11.046)"
```

---

## 📂 CHANGELOG Management / CHANGELOG 관리

### File Structure / 파일 구조

```
CHANGELOG.md                              # Major/Minor overview
docs/CHANGELOG/
    ├── CHANGELOG-v1.0.md                # v1.0.x detailed changes
    ├── CHANGELOG-v1.1.md                # v1.1.x detailed changes
    └── CHANGELOG-v1.11.md               # v1.11.x detailed changes
```

### Update Rules / 업데이트 규칙

**MUST update BEFORE every commit:**

```markdown
## [v1.11.046] - 2025-10-16

### Added
- Added new feature X / 새 기능 X 추가

### Changed
- Modified feature Y / 기능 Y 수정

### Fixed
- Fixed bug Z / 버그 Z 수정
```

---

## 🛠️ Development Tools / 개발 도구

### Testing / 테스트

```bash
go test ./... -v              # All tests
go test ./{package} -v        # Package tests
go test ./{package} -cover    # With coverage
```

### Building / 빌드

```bash
go build ./...                # Build all
go build ./{package}          # Build package
```

### Examples / 예제

```bash
go run examples/{package}/main.go
```

### Docker (MySQL/Redis) / Docker (MySQL/Redis)

```bash
# MySQL
./.docker/scripts/docker-mysql-start.sh
./.docker/scripts/docker-mysql-stop.sh

# Redis
./.docker/scripts/docker-redis-start.sh
./.docker/scripts/docker-redis-stop.sh
```

---

## 🎯 Quick Reference / 빠른 참조

### Every Task Checklist / 모든 작업 체크리스트

```
□ 1. Read DEVELOPMENT_WORKFLOW_GUIDE.md
□ 2. Bump version in cfg/app.yaml
□ 3. Commit version bump
□ 4. Perform work (code/docs)
□ 5. Run tests (go test ./... -v)
□ 6. Update CHANGELOG
□ 7. Commit with proper message
□ 8. Push to GitHub
```

### Import Pattern / Import 패턴

**✅ Correct:**
```go
import "github.com/arkd0ng/go-utils/random"
import "github.com/arkd0ng/go-utils/logging"
```

**❌ Incorrect:**
```go
import "github.com/arkd0ng/go-utils"  // Don't import root
```

### Error Handling / 에러 처리

All methods return `(result, error)`:

```go
str, err := random.GenString.Alnum(32)
if err != nil {
    log.Fatal(err)
}
```

---

## 📖 Additional Resources / 추가 자료

### Core Documentation / 핵심 문서

- **[README.md](./README.md)** - Project overview and package list
- **[DEVELOPMENT_WORKFLOW_GUIDE.md](./docs/DEVELOPMENT_WORKFLOW_GUIDE.md)** ⭐ Main workflow guide
- **[PACKAGE_DEVELOPMENT_GUIDE.md](./docs/PACKAGE_DEVELOPMENT_GUIDE.md)** ⭐ Package development standards
- **[CODE_TEST_MAKE_GUIDE.md](./docs/CODE_TEST_MAKE_GUIDE.md)** - Testing guidelines
- **[EXAMPLE_CODE_GUIDE.md](./docs/EXAMPLE_CODE_GUIDE.md)** - Example code standards

### Package Documentation / 패키지 문서

Each package has its own detailed documentation:

각 패키지는 자체 상세 문서를 가지고 있습니다:

```
{package}/README.md                      # Quick start
docs/{package}/USER_MANUAL.md           # User guide
docs/{package}/DEVELOPER_GUIDE.md       # Developer guide
docs/{package}/DESIGN_PLAN.md           # Design (if applicable)
docs/{package}/WORK_PLAN.md             # Work plan (if applicable)
examples/{package}/main.go               # Executable examples
```

### External Dependencies / 외부 의존성

- `github.com/go-sql-driver/mysql` - MySQL driver
- `github.com/redis/go-redis/v9` - Redis client
- `gopkg.in/natefinch/lumberjack.v2` - Log rotation
- `gopkg.in/yaml.v3` - YAML parsing
- `golang.org/x/text` - Unicode normalization
- `golang.org/x/exp` - Generic constraints

---

## ⚠️ Critical Reminders / 중요 알림

1. **ALWAYS read DEVELOPMENT_WORKFLOW_GUIDE.md first** / 항상 DEVELOPMENT_WORKFLOW_GUIDE.md를 먼저 읽을 것
2. **ALWAYS bump version before work** / 항상 작업 전 버전 증가
3. **ALWAYS update CHANGELOG** / 항상 CHANGELOG 업데이트
4. **ALWAYS test before commit** / 항상 커밋 전 테스트
5. **ALWAYS document in both languages** / 항상 두 언어로 문서화

---

## 🎓 Learning Path / 학습 경로

**For new contributors / 새로운 기여자를 위한:**

1. Read [DEVELOPMENT_WORKFLOW_GUIDE.md](./docs/DEVELOPMENT_WORKFLOW_GUIDE.md)
2. Read [PACKAGE_DEVELOPMENT_GUIDE.md](./docs/PACKAGE_DEVELOPMENT_GUIDE.md)
3. Browse existing package READMEs
4. Review example code in `examples/`
5. Check package tests in `*_test.go` files
6. Read USER_MANUAL and DEVELOPER_GUIDE for reference packages

---

**Last Updated / 최종 업데이트**: 2025-10-16  
**Version / 버전**: v1.11.046  
**Maintained By / 관리자**: go-utils team
