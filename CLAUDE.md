# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

---

## Project Overview

**go-utils** is a production-ready collection of Go utility packages designed to reduce boilerplate code by 20-50 lines down to 1-2 lines. Current version: **v1.12.021**.

Key characteristics:
- **Extreme Simplicity**: Minimal code, maximum functionality
- **Type Safety**: Go 1.18+ generics where appropriate
- **Zero Config**: Sensible defaults for 99% of use cases
- **High Quality**: 80%+ test coverage target (100% for newer packages)
- **Bilingual**: All documentation in English/Korean

---

## 🚨 CRITICAL: Standard Work Cycle

**EVERY task must follow this exact 5-step sequence:**

```
1. Bump version in cfg/app.yaml (increase patch number)
2. Perform work (code or documentation)
3. Verify (go build, go test)
4. Update CHANGELOG (docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md)
5. Git commit and push
```

**Never skip steps. Never batch multiple tasks before updating CHANGELOG.**

---

## Build and Test Commands

### Build
```bash
# Build all packages
go build ./...

# Build specific package
go build ./errorutil
```

### Test
```bash
# Run all tests
go test ./... -v

# Run specific package tests
go test ./errorutil -v

# Run single test function
go test ./stringutil -run TestToSnakeCase -v

# Check coverage
go test ./... -cover

# Coverage for specific package
go test ./errorutil -cover

# Generate coverage report
go test ./errorutil -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Verify Before Commit
```bash
# Essential verification
go build ./... && go test ./... -v
```

---

## Version Management

### Single Source of Truth
- **File**: `cfg/app.yaml`
- **Format**: `version: v1.12.019` (MAJOR.MINOR.PATCH)
- **Strategy**: One patch increment = one function/task

### How to Bump Version
```bash
# Edit cfg/app.yaml
# Change: version: v1.12.019
# To:     version: v1.12.020

# Commit version change FIRST
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.12.020 / v1.12.020로 버전 증가"
git push
```

**Always increment version BEFORE starting any work.**

### 🚨 CRITICAL: Version Management Rule / 버전 관리 규칙

**NEVER hardcode versions in package version.go files!**
**패키지 version.go 파일에 버전을 하드코딩하지 마세요!**

All packages MUST read version from `cfg/app.yaml` using the internal/version utility:

모든 패키지는 internal/version 유틸리티를 사용하여 `cfg/app.yaml`에서 버전을 읽어야 합니다:

```go
// ❌ WRONG - Hardcoded version / 잘못됨 - 하드코딩된 버전
package mypackage

const Version = "v1.13.004"

// ✅ CORRECT - Dynamic version from app.yaml / 올바름 - app.yaml에서 동적으로
package mypackage

import "github.com/arkd0ng/go-utils/internal/version"

var Version = version.Get()
```

**Why / 이유:**
- Single source of truth (cfg/app.yaml)
- No sync issues between files
- Easier maintenance
- 단일 진실 소스 (cfg/app.yaml)
- 파일 간 동기화 문제 없음
- 유지보수 용이

---

## CHANGELOG Management

### Two-Level System

1. **Root CHANGELOG.md**: High-level overview of major/minor versions only
2. **docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md**: Detailed patch-by-patch history

### Required Format
Every change must be documented in `docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md`:

```markdown
## [v1.12.020] - 2025-10-17

### Added / 추가
- New function: `errorutil.WrapContext()` for context-aware error wrapping
- 새 함수: 컨텍스트 인식 에러 래핑을 위한 `errorutil.WrapContext()`

### Files Changed / 변경된 파일
- `errorutil/error.go` - Implemented WrapContext function
- `errorutil/error_test.go` - Added comprehensive tests for WrapContext
- `errorutil/README.md` - Updated API documentation

### Context / 컨텍스트
**User Request**: "Add context-aware error wrapping"
**Why**: Enable better error tracing in distributed systems
**Impact**: Developers can now track errors across service boundaries
```

**Update CHANGELOG before EVERY commit. No exceptions.**

---

## Git Commit Message Format

```
[Type]: [Description] / [한글 설명] (vX.Y.Z)
```

### Types
- **Feat**: New feature
- **Fix**: Bug fix
- **Docs**: Documentation only
- **Refactor**: Code refactoring
- **Chore**: Version bumps, config changes
- **Test**: Test additions/fixes

### Examples
```bash
git commit -m "Feat: Add WrapContext function to errorutil / errorutil에 WrapContext 함수 추가 (v1.12.020)"
git commit -m "Docs: Update errorutil USER_MANUAL with WrapContext examples / WrapContext 예제로 errorutil USER_MANUAL 업데이트 (v1.12.021)"
git commit -m "Chore: Bump version to v1.12.022 / v1.12.022로 버전 증가"
```

---

## Package Architecture

### Completed Packages (Production)
- **errorutil** (v1.12.x): Error handling with codes, context, stack traces - 99.2% coverage
- **stringutil** (v1.5.x): 53 string manipulation functions
- **timeutil** (v1.6.x): 114+ time/date functions, KST timezone support
- **sliceutil** (v1.7.x): 95 type-safe slice operations - 100% coverage
- **maputil** (v1.8.x): 99 type-safe map operations - 92.8% coverage
- **fileutil**: ~91 cross-platform file operations
- **httputil**: HTTP client reducing 30+ lines to 2-3 lines
- **logging**: Structured logging with auto-rotation
- **random**: Cryptographic random generation (14 methods)
- **database/mysql** (v1.3.x): MySQL client reducing 30+ lines to 2 lines
- **database/redis** (v1.4.x): Redis client reducing 20+ lines to 2 lines

### In Development
- **websvrutil** (v1.11.x): Web server utilities

### Package Structure Pattern
```
package-name/
├── implementation.go        # Core functions
├── impl_test.go            # Tests
├── types.go                # Type definitions
├── types_test.go           # Type tests
├── helper.go               # Helper functions
├── options.go              # Configuration options
└── README.md               # Package documentation
```

---

## Documentation Standards

### Required Documentation Levels
1. **Package README.md**: Quick start, API tables, examples
2. **docs/{package}/USER_MANUAL.md**: Comprehensive usage guide
3. **docs/{package}/DEVELOPER_GUIDE.md**: Architecture, internals
4. **examples/{package}/main.go**: Executable examples

### Bilingual Requirements
- All documentation must be in English AND Korean
- Format: English first, then Korean
- Code comments: English first, Korean as inline comment

Example:
```go
// WrapContext wraps an error with context information.
// WrapContext는 에러를 컨텍스트 정보로 래핑합니다.
func WrapContext(err error, ctx context.Context) error {
    // implementation
}
```

---

## Testing Standards

### Coverage Requirements
- **New packages**: 100% coverage target
- **Mature packages**: 80%+ minimum
- **Critical packages**: 95%+ (errorutil, database/mysql, database/redis)

### Test Organization
- Table-driven tests for parametric testing
- Separate test files by functionality (e.g., `case_test.go`, `manipulation_test.go`)
- Example functions for documentation (e.g., `ExampleToSnakeCase()`)
- Benchmark tests for performance-critical code

### Test Patterns
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"case1", "input1", "expected1"},
        {"case2", "input2", "expected2"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionName(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

---

## Common Development Tasks

### Adding a New Function

1. **Bump version** in `cfg/app.yaml`
2. **Implement function** in appropriate `.go` file
3. **Add tests** in corresponding `_test.go` file
4. **Add example** in `examples/{package}/main.go`
5. **Update documentation**:
   - Package README.md (API table)
   - docs/{package}/USER_MANUAL.md (detailed guide)
6. **Verify**: `go build ./... && go test ./... -v`
7. **Update CHANGELOG**: `docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md`
8. **Commit and push**

### Fixing a Bug

1. **Bump version** in `cfg/app.yaml`
2. **Write failing test** that reproduces the bug
3. **Fix the bug** in implementation
4. **Verify test passes**: `go test ./... -v`
5. **Update CHANGELOG** with "Fixed" section
6. **Commit and push**

### Updating Documentation Only

1. **Bump version** in `cfg/app.yaml`
2. **Update documentation** files
3. **Verify builds**: `go build ./...`
4. **Update CHANGELOG** with "Changed" section noting documentation updates
5. **Commit with "Docs:" prefix**

---

## Branch Strategy

### Main Branch
- **main**: Stable, production-ready code
- Direct commits allowed for docs, bug fixes, minor improvements

### Feature Branches
For new packages or major features:
```bash
# Format: feature/v{MAJOR}.{MINOR}.x-{package-name}
git checkout -b feature/v1.13.x-validation
```

When complete, merge to main.

---

## Code Quality Principles

### Design Philosophy
1. **Extreme Simplicity**: Reduce 20-50 lines to 1-2 lines
2. **Type Safety**: Use generics where appropriate (Go 1.18+)
3. **Zero Configuration**: Sensible defaults
4. **Standard Compatible**: Works with Go stdlib patterns
5. **Minimal Dependencies**: Use only stdlib (except specific needs like database drivers)

### Before Every Commit Checklist
- [ ] Version bumped in `cfg/app.yaml`
- [ ] Code implements complete feature
- [ ] All tests pass: `go test ./...`
- [ ] Build succeeds: `go build ./...`
- [ ] CHANGELOG updated in `docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md`
- [ ] Documentation complete (English + Korean)
- [ ] Examples added/updated
- [ ] Godoc comments present

---

## Special Notes

### KST Timezone (timeutil)
- Default timezone is KST (GMT+9) for all timeutil functions
- Use custom format tokens: `YYYY-MM-DD HH:mm:ss` instead of Go's `2006-01-02 15:04:05`
- Korean holiday support included

### Error Handling (errorutil)
- Compatible with standard library `errors.Is` and `errors.As`
- Supports error codes (string and numeric)
- Automatic stack trace capture
- Context-aware error wrapping

### Database Clients
- **mysql**: Zero-downtime credential rotation, auto-retry, no `defer rows.Close()` needed
- **redis**: Auto-retry with exponential backoff, connection pooling, health checks

### Generic Utilities (sliceutil, maputil)
- All functions are type-safe using Go 1.18+ generics
- Immutable operations (original collections unchanged)
- Functional programming style (Map, Filter, Reduce)

---

## Directory Structure

```
go-utils/
├── cfg/                           # Configuration (version source of truth)
│   └── app.yaml                   # Version: v1.12.021
├── docs/                          # Comprehensive documentation
│   ├── CHANGELOG/                 # Detailed version history
│   │   ├── CHANGELOG-v1.11.md    # websvrutil detailed history
│   │   ├── CHANGELOG-v1.12.md    # errorutil detailed history
│   │   └── ...
│   ├── errorutil/                 # Package-specific docs
│   │   ├── USER_MANUAL.md
│   │   ├── DEVELOPER_GUIDE.md
│   │   └── DESIGN_PLAN.md
│   ├── DEVELOPMENT_WORKFLOW_GUIDE.md  # This is critical - READ THIS
│   └── PACKAGE_DEVELOPMENT_GUIDE.md   # Package development process
├── examples/                      # Executable examples for each package
│   ├── errorutil/main.go
│   ├── stringutil/main.go
│   └── ...
├── errorutil/                     # Production package (v1.12.x)
├── stringutil/                    # Production package (v1.5.x)
├── timeutil/                      # Production package (v1.6.x)
├── sliceutil/                     # Production package (v1.7.x)
├── maputil/                       # Production package (v1.8.x)
├── fileutil/                      # Production package
├── httputil/                      # Production package
├── logging/                       # Production package
├── random/                        # Production package
├── websvrutil/                    # In development (v1.11.x)
├── database/
│   ├── mysql/                     # Production package (v1.3.x)
│   └── redis/                     # Production package (v1.4.x)
├── go.mod                         # Module definition
├── README.md                      # Main project documentation
├── CHANGELOG.md                   # High-level version overview
└── LICENSE                        # MIT License
```

---

## Quick Command Reference

```bash
# Version bump
vim cfg/app.yaml  # Increment patch version
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.12.XXX / v1.12.XXX로 버전 증가"

# Build and test
go build ./...
go test ./... -v

# Test specific package
go test ./errorutil -v
go test ./errorutil -cover

# Update CHANGELOG
vim docs/CHANGELOG/CHANGELOG-v1.12.md  # Add entry at top

# Commit work
git add .
git commit -m "Feat: Add new function / 새 함수 추가 (v1.12.XXX)"
git push origin main
```

---

## Important Files to Reference

When working on this codebase, consult these key documents:

1. **docs/DEVELOPMENT_WORKFLOW_GUIDE.md** - Complete workflow reference
2. **docs/PACKAGE_DEVELOPMENT_GUIDE.md** - How to develop packages
3. **docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md** - Version-specific change history
4. **cfg/app.yaml** - Current version number (single source of truth)
5. **README.md** - Project overview and package status

---

## Anti-Patterns to Avoid

❌ **Don't:**
- Start work without bumping version first
- Commit without updating CHANGELOG
- Push code with failing tests
- Write code-only without documentation
- Skip bilingual documentation
- Add external dependencies without justification
- Break backward compatibility in existing APIs
- Batch multiple features before CHANGELOG update

✅ **Do:**
- Follow the 5-step work cycle religiously
- Update CHANGELOG for every single commit
- Write tests alongside implementation
- Document in both English and Korean
- Use table-driven tests
- Aim for high test coverage (80%+ minimum)
- Keep functions simple and focused
- Follow existing code patterns in the package

---

## Current Status (v1.12.021)

- **Latest Package**: errorutil (v1.12.x) - COMPLETE with 99.2% coverage
- **In Development**: websvrutil (v1.11.x) - Advanced features being added
- **Next Target**: Achieve 100% coverage on errorutil
- **Git Status**: Main branch stable

### Common Debugging Commands

```bash
# Check test coverage for all packages
go test ./... -cover

# Find all test files missing coverage
find . -name "*_test.go" -type f

# View detailed coverage for specific package
go test ./errorutil -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html

# Run tests with race detection
go test ./... -race

# Check for unused dependencies
go mod tidy
go mod verify
```

---

## Support

For questions about this codebase or Claude Code:
- GitHub Issues: https://github.com/arkd0ng/go-utils/issues
- Claude Code Docs: https://docs.claude.com/en/docs/claude-code
