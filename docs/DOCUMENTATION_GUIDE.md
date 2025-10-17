# go-utils 문서화 및 코드 작성 가이드
# Documentation and Code Writing Guide

**프로젝트**: arkd0ng/go-utils  
**최종 업데이트**: 2025년 10월 17일  
**버전**: 1.1.0

---

## 📋 목차 / Table of Contents

1. [개요 / Overview](#개요--overview)
2. [문서화 철학 / Documentation Philosophy](#문서화-철학--documentation-philosophy)
3. [주석 작성 표준 / Comment Standards](#주석-작성-표준--comment-standards)
4. [코드 작성 가이드 / Code Writing Guide](#코드-작성-가이드--code-writing-guide)
5. [스크립트 작성 가이드 / Script Writing Guide](#스크립트-작성-가이드--script-writing-guide)
6. [문서 작성 가이드 / Documentation Writing Guide](#문서-작성-가이드--documentation-writing-guide)
7. [품질 기준 / Quality Standards](#품질-기준--quality-standards)
8. [실전 예제 / Practical Examples](#실전-예제--practical-examples)
9. [체크리스트 / Checklist](#체크리스트--checklist)

---

## 개요 / Overview

### 목적 / Purpose

이 가이드는 go-utils 프로젝트의 **일관되고 고품질의 문서화**를 보장하기 위한 표준을 정의합니다.

**주요 목표**:
- ✅ **엔터프라이즈 레벨**의 문서화 품질
- ✅ 초보자도 쉽게 이해할 수 있는 상세한 설명

### 적용 범위 / Scope

이 가이드는 다음에 적용됩니다:

#### Go 코드
- 모든 Go 소스 파일 (.go)
- 패키지 레벨 문서
- 함수/메서드 주석
- 타입/구조체 주석
- 상수/변수 주석

#### 문서 파일
- README 파일 (README.md, 각 패키지 README)
- 기술 문서 (docs/ 디렉토리)
- 가이드 문서 (GUIDE, TUTORIAL 등)
- CHANGELOG 파일
- 설계 문서 (DESIGN, ARCHITECTURE 등)

#### 스크립트 및 설정 파일
- Shell 스크립트 (.sh)
- Build 스크립트 (Makefile)
- YAML 설정 파일 (.yaml, .yml)
- JSON 설정 파일 (.json)
- 환경 설정 파일 (.env.example)

---

## 문서화 철학 / Documentation Philosophy

### 핵심 원칙 / Core Principles

#### 1. 📖 충분히 자세하게 (Sufficiently Detailed)

**원칙**: 코드를 보지 않아도 동작을 완전히 이해할 수 있어야 합니다.

```go
// ❌ 나쁜 예 (Bad Example)
// Add adds two numbers
func Add(a, b int) int {
    return a + b
}

// ✅ 좋은 예 (Good Example)
// Add performs integer addition and returns the sum of two numbers.
// This function handles standard integer arithmetic with Go's built-in
// overflow behavior (wraps around at max/min int values).
//
// Add는 정수 덧셈을 수행하고 두 숫자의 합을 반환합니다.
// 이 함수는 Go의 내장 오버플로우 동작을 사용한 표준 정수 산술을 처리합니다
// (최대/최소 int 값에서 순환).
//
// Parameters / 매개변수:
//   - a: First integer operand (any valid int value)
//     첫 번째 정수 피연산자 (유효한 모든 int 값)
//   - b: Second integer operand (any valid int value)
//     두 번째 정수 피연산자 (유효한 모든 int 값)
//
// Returns / 반환값:
//   - int: Sum of a and b. Note that overflow wraps around.
//     a와 b의 합. 오버플로우 시 순환됩니다.
//
// Example / 예제:
//   result := Add(10, 20)  // returns 30
//   overflow := Add(math.MaxInt, 1)  // wraps to math.MinInt
func Add(a, b int) int {
    return a + b
}
```

#### 2. 👨‍🎓 매우 친절하게 (Very User-Friendly)

**원칙**: Go 언어 초보자도 쉽게 이해할 수 있어야 합니다.

```go
// ✅ 초보자 친화적 예제
// NewClient creates a new HTTP client with recommended default settings.
// It automatically configures timeout (30s), retry logic (3 attempts),
// and connection pooling for optimal performance.
//
// NewClient는 권장 기본 설정으로 새 HTTP 클라이언트를 생성합니다.
// 최적의 성능을 위해 타임아웃(30초), 재시도 로직(3회 시도) 및
// 연결 풀링을 자동으로 구성합니다.
//
// When to use / 사용 시기:
//   - Making HTTP requests to external APIs
//     외부 API에 HTTP 요청을 할 때
//   - Need automatic retry on temporary failures
//     임시 실패 시 자동 재시도가 필요할 때
//   - Want optimized connection reuse
//     최적화된 연결 재사용을 원할 때
//
// Parameters / 매개변수:
//   - opts: Optional configuration functions (can be nil for defaults)
//     선택적 구성 함수 (기본값을 사용하려면 nil 가능)
//
// Returns / 반환값:
//   - *Client: Ready-to-use HTTP client instance
//     바로 사용 가능한 HTTP 클라이언트 인스턴스
//
// Example / 예제:
//   // Basic usage with defaults / 기본 사용법
//   client := NewClient()
//   
//   // Custom configuration / 사용자 정의 구성
//   client := NewClient(
//       WithTimeout(60*time.Second),
//       WithRetry(5),
//   )
func NewClient(opts ...Option) *Client {
    // implementation
}
```

#### 3. 🔍 포괄적으로 (Comprehensive)

**원칙**: 엣지 케이스, 에러 상황, 성능 특성을 모두 설명합니다.

```go
// ✅ 포괄적인 문서화 예제
// ProcessData validates and processes user input data with comprehensive
// error handling and edge case management.
//
// ProcessData는 포괄적인 에러 처리와 엣지 케이스 관리를 통해
// 사용자 입력 데이터를 검증하고 처리합니다.
//
// Parameters / 매개변수:
//   - data: Input data to process (cannot be nil)
//     처리할 입력 데이터 (nil 불가)
//     * Empty slices are allowed and return empty result
//       빈 슬라이스는 허용되며 빈 결과 반환
//     * Duplicate values are automatically removed
//       중복 값은 자동으로 제거됨
//
// Returns / 반환값:
//   - []string: Processed data, sorted and deduplicated
//     처리된 데이터, 정렬 및 중복 제거됨
//   - error: Error if validation fails
//     검증 실패 시 에러
//
// Errors / 에러:
//   - ErrNilData: when data parameter is nil
//     data 매개변수가 nil일 때
//   - ErrInvalidFormat: when data contains invalid characters
//     데이터에 유효하지 않은 문자가 포함될 때
//   - ErrTooLarge: when data exceeds 10,000 items
//     데이터가 10,000개 항목을 초과할 때
//
// Performance / 성능:
//   - Time complexity: O(n log n) due to sorting
//     정렬로 인한 시간 복잡도: O(n log n)
//   - Space complexity: O(n) for deduplication map
//     중복 제거 맵을 위한 공간 복잡도: O(n)
//   - Memory allocation: One allocation for result slice
//     메모리 할당: 결과 슬라이스를 위한 1회 할당
//
// Notes / 주의사항:
//   - Thread-safe: Yes, no shared state
//     스레드 안전: 예, 공유 상태 없음
//   - Large datasets: Consider batching for >100,000 items
//     대용량 데이터셋: 100,000개 이상 항목은 배치 처리 고려
//
// Example / 예제:
//   data := []string{"apple", "banana", "apple", "cherry"}
//   result, err := ProcessData(data)
//   // result: ["apple", "banana", "cherry"]
//   
//   // Edge case: empty input / 엣지 케이스: 빈 입력
//   result, err := ProcessData([]string{})
//   // result: []string{}, err: nil
func ProcessData(data []string) ([]string, error) {
    // implementation
}
```

#### 4. 💡 실용적으로 (Practical)

**원칙**: 실제 사용 예시와 주의사항을 포함합니다.

#### 5. 🌐 이중 언어 (Bilingual)

**원칙**: 영문과 한글 모두 동일한 수준의 상세함을 유지합니다.

---

## 주석 작성 표준 / Comment Standards

### 1. 패키지 레벨 주석 / Package-Level Comments

**형식**:

```go
// Package [name] provides [core functionality].
// It offers [key features] with [capabilities].
//
// [name] 패키지는 [핵심 기능]을 제공합니다.
// [주요 기능]과 [기능들]을 제공합니다.
//
// Key Features / 주요 기능:
//   - Feature 1: Description
//     기능 1: 설명
//   - Feature 2: Description
//     기능 2: 설명
//
// Performance / 성능:
//   - Characteristic 1
//     특성 1
//   - Characteristic 2
//     특성 2
//
// Thread Safety / 스레드 안전성:
//   - Safety information
//     안전성 정보
//
// Usage / 사용법:
//   [example code]
//
// See also / 참고:
//   - Related package
//     관련 패키지
package packagename
```

**실제 예제** (httputil 패키지):

```go
// Package httputil provides extreme simplicity HTTP utilities for Go.
// 패키지 httputil은 Go를 위한 극도로 간단한 HTTP 유틸리티를 제공합니다.
//
// This package reduces 30+ lines of repetitive HTTP code to just 2-3 lines
// with automatic JSON handling, retry logic, and type-safe operations.
//
// 이 패키지는 30줄 이상의 반복적인 HTTP 코드를 자동 JSON 처리, 재시도 로직,
// 타입 안전 작업을 통해 단 2-3줄로 줄입니다.
//
// Key Features / 주요 기능:
//
// - Simple HTTP methods (GET, POST, PUT, PATCH, DELETE)
//   간단한 HTTP 메서드
// - Automatic JSON encoding/decoding
//   자동 JSON 인코딩/디코딩
// - Automatic retry with exponential backoff
//   지수 백오프를 사용한 자동 재시도
//
// Quick Start / 빠른 시작:
//
//   var result MyStruct
//   err := httputil.Get("https://api.example.com/data", &result)
package httputil
```

### 2. 함수/메서드 주석 / Function Comments

**완전한 형식**:

```go
// FunctionName performs [specific action] with [behavior].
// It [detailed explanation of purpose and use cases].
//
// FunctionName은 [특정 동작]을 [방식]으로 수행합니다.
// [목적과 사용 사례에 대한 상세 설명].
//
// Parameters / 매개변수:
//   - param1: [Detailed description]
//     [상세 설명]
//     * Expected values: [범위/형식]
//     * Constraints: [제약사항]
//     * Special values: [특수 값 의미]
//
// Returns / 반환값:
//   - type: [Description]
//     [설명]
//     * Success case: [성공 시]
//     * Failure case: [실패 시]
//
// Errors / 에러:
//   - ErrType1: [condition]
//     [조건]
//   - ErrType2: [condition]
//     [조건]
//
// Performance / 성능:
//   - Time complexity: O(n)
//     시간 복잡도: O(n)
//   - Space complexity: O(1)
//     공간 복잡도: O(1)
//
// Notes / 주의사항:
//   - Thread-safe: Yes/No
//     스레드 안전: 예/아니오
//   - [Other important notes]
//     [기타 중요 사항]
//
// Example / 예제:
//   [code example]
//   [코드 예제]
func FunctionName(param1 type) (type, error) {
    // implementation
}
```

### 3. 타입/구조체 주석 / Type Comments

**형식**:

```go
// TypeName represents [concept/entity].
// It is used for [purpose] and provides [capabilities].
//
// TypeName은 [개념/엔티티]를 나타냅니다.
// [목적]에 사용되며 [기능]을 제공합니다.
//
// Lifecycle / 생명주기:
//   - Creation: [how to create]
//     생성: [생성 방법]
//   - Usage: [how to use]
//     사용: [사용 방법]
//   - Cleanup: [if needed]
//     정리: [필요한 경우]
//
// Thread Safety / 스레드 안전성:
//   - [safety information]
//     [안전성 정보]
//
// Example / 예제:
//   [code]
type TypeName struct {
    // Field1 stores [purpose].
    // Special values: [if any]
    //
    // Field1은 [목적]을 저장합니다.
    // 특수 값: [있는 경우]
    Field1 string
    
    // Field2 contains [purpose].
    // Valid range: [range]
    //
    // Field2는 [목적]을 포함합니다.
    // 유효 범위: [범위]
    Field2 int
}
```

### 4. 상수/변수 주석 / Constant/Variable Comments

```go
// ConstantName defines [purpose].
// It is used [when/where] for [reason].
// Value: [value and meaning]
//
// ConstantName은 [목적]을 정의합니다.
// [시기/장소]에서 [이유]로 사용됩니다.
// 값: [값과 의미]
const ConstantName = value
```

---

## 코드 작성 가이드 / Code Writing Guide

### 1. 네이밍 규칙 / Naming Conventions

**함수명**:
- 동사로 시작 (Get, Set, Create, Update, Delete, Process, Validate)
- 명확하고 설명적인 이름 사용
- 약어 피하기 (GetHTTPClient ✅, GetHC ❌)

**변수명**:
- 의미 있는 이름 사용
- 단일 문자 변수는 루프나 짧은 범위에서만
- 타입 정보 포함하지 않기 (userString ❌, userName ✅)

**상수명**:
- 대문자와 언더스코어 사용 또는 CamelCase
- 명확한 의미 전달

### 2. 에러 처리 / Error Handling

**원칙**:
- 모든 에러는 명시적으로 처리
- 에러 메시지는 컨텍스트 포함
- 에러는 wrap하여 스택 추적 가능하게

```go
// ✅ 좋은 에러 처리
func ProcessFile(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("failed to read file %s: %w", path, err)
    }
    
    if err := validate(data); err != nil {
        return fmt.Errorf("validation failed for %s: %w", path, err)
    }
    
    return nil
}
```

### 3. 성능 고려사항 / Performance Considerations

모든 함수는 다음을 문서화해야 합니다:
- ⏱️ **시간 복잡도** (Time Complexity)
- 💾 **공간 복잡도** (Space Complexity)
- 🔄 **메모리 할당** (Memory Allocation)
- 🚀 **최적화 팁** (Optimization Tips)

```go
// FindDuplicates finds duplicate elements in a slice.
// It uses a map for O(n) lookup instead of nested loops.
//
// Performance / 성능:
//   - Time complexity: O(n) where n is slice length
//     시간 복잡도: O(n) (n은 슬라이스 길이)
//   - Space complexity: O(n) for the map
//     공간 복잡도: O(n) (맵을 위한)
//   - Memory allocation: One map allocation
//     메모리 할당: 맵 1회 할당
//
// For large slices (>10,000 elements), consider:
// 대용량 슬라이스(>10,000개 요소)의 경우 고려사항:
//   - Pre-allocating the map with make(map[T]bool, len(slice))
//     make(map[T]bool, len(slice))로 맵 사전 할당
//   - Processing in batches if memory is constrained
//     메모리 제약이 있는 경우 배치로 처리
func FindDuplicates(slice []string) []string {
    seen := make(map[string]bool)
    duplicates := []string{}
    
    for _, item := range slice {
        if seen[item] {
            duplicates = append(duplicates, item)
        }
        seen[item] = true
    }
    
    return duplicates
}
```

---

## 스크립트 작성 가이드 / Script Writing Guide

### 1. Shell 스크립트 / Shell Scripts

**헤더 형식**:

```bash
#!/bin/bash
#
# Script Name: script_name.sh
# Description: Brief description of what this script does
#              스크립트 설명 - 이 스크립트가 하는 일
#
# Usage: ./script_name.sh [options] [arguments]
#        사용법: ./script_name.sh [옵션] [인수]
#
# Options / 옵션:
#   -h, --help     Show this help message
#                  도움말 메시지 표시
#   -v, --verbose  Enable verbose output
#                  상세 출력 활성화
#
# Examples / 예제:
#   ./script_name.sh --verbose input.txt
#   ./script_name.sh -h
#
# Author: arkd0ng
# Created: 2025-10-17
# Modified: 2025-10-17
#

# Exit on error / 에러 시 종료
set -e

# Exit on undefined variable / 미정의 변수 사용 시 종료
set -u

# Pipe failure causes exit / 파이프 실패 시 종료
set -o pipefail
```

**함수 주석**:

```bash
# Function: validate_input
# Description: Validates user input and checks required conditions
#              사용자 입력을 검증하고 필수 조건을 확인합니다
#
# Parameters / 매개변수:
#   $1 - Input file path (must exist and be readable)
#        입력 파일 경로 (존재하고 읽기 가능해야 함)
#
# Returns / 반환값:
#   0 - Success / 성공
#   1 - Invalid input / 유효하지 않은 입력
#
# Example / 예제:
#   if validate_input "$input_file"; then
#       echo "Valid input"
#   fi
validate_input() {
    local input_file="$1"
    
    if [[ ! -f "$input_file" ]]; then
        echo "Error: File not found: $input_file" >&2
        echo "에러: 파일을 찾을 수 없음: $input_file" >&2
        return 1
    fi
    
    return 0
}
```

**변수 주석**:

```bash
# Configuration / 구성
# Database connection timeout in seconds
# 데이터베이스 연결 타임아웃 (초)
readonly DB_TIMEOUT=30

# Maximum retry attempts for API calls
# API 호출 최대 재시도 횟수
readonly MAX_RETRIES=3

# Log file location
# 로그 파일 위치
LOG_FILE="/var/log/myapp.log"
```

### 2. Makefile

**헤더 형식**:

```makefile
# Makefile for go-utils project
# go-utils 프로젝트용 Makefile
#
# Available targets / 사용 가능한 타겟:
#   make build    - Build all packages / 모든 패키지 빌드
#   make test     - Run all tests / 모든 테스트 실행
#   make clean    - Remove build artifacts / 빌드 산출물 제거
#   make help     - Show this help / 도움말 표시
#
# Examples / 예제:
#   make build
#   make test
#

.PHONY: all build test clean help

# Default target / 기본 타겟
all: build test

# Build all Go packages
# 모든 Go 패키지 빌드
build:
	@echo "Building all packages..."
	@echo "모든 패키지 빌드 중..."
	go build ./...

# Run all tests with coverage
# 커버리지와 함께 모든 테스트 실행
test:
	@echo "Running tests..."
	@echo "테스트 실행 중..."
	go test -v -cover ./...
```

### 3. YAML 설정 파일

**주석 형식**:

```yaml
# Application Configuration
# 애플리케이션 구성
#
# This file contains the main application settings.
# Production values should be set via environment variables.
#
# 이 파일은 주요 애플리케이션 설정을 포함합니다.
# 프로덕션 값은 환경 변수로 설정해야 합니다.

# Application metadata / 애플리케이션 메타데이터
app:
  # Application name (used in logs and metrics)
  # 애플리케이션 이름 (로그 및 메트릭에 사용)
  name: "go-utils"
  
  # Semantic version / 시맨틱 버전
  version: "1.12.021"
  
  # Environment: development, staging, production
  # 환경: development, staging, production
  environment: "development"

# Database configuration / 데이터베이스 구성
database:
  mysql:
    # Connection settings / 연결 설정
    host: "localhost"      # Database host / 데이터베이스 호스트
    port: 3306            # MySQL default port / MySQL 기본 포트
    
    # Connection pool / 연결 풀
    # Maximum number of open connections
    # 최대 열린 연결 수
    max_open_conns: 100
    
    # Maximum number of idle connections
    # 최대 유휴 연결 수
    max_idle_conns: 10
```

### 4. JSON 설정 파일

```json
{
  "comment": "Application Configuration - 애플리케이션 구성",
  "comment_note": "This file uses 'comment' fields for documentation as JSON doesn't support native comments",
  
  "server": {
    "comment": "HTTP server settings - HTTP 서버 설정",
    "host": "0.0.0.0",
    "port": 8080,
    "timeout_seconds": 30
  },
  
  "logging": {
    "comment": "Logging configuration - 로깅 구성",
    "level": "info",
    "comment_levels": "Available: debug, info, warn, error - 가능한 값: debug, info, warn, error",
    "format": "json"
  }
}
```

---

## 문서 작성 가이드 / Documentation Writing Guide

### 1. README 파일

**표준 구조**:

```markdown
# Package/Project Name

## Overview / 개요

Brief description in English.
영문 간단 설명.

Detailed description in Korean.
한글 상세 설명.

## Features / 주요 기능

- Feature 1: Description
  기능 1: 설명
- Feature 2: Description
  기능 2: 설명

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/packagename
```

## Quick Start / 빠른 시작

```go
// Code example
// 코드 예제
package main

import "github.com/arkd0ng/go-utils/packagename"

func main() {
    // Example usage
    // 사용 예제
}
```

## API Reference / API 참조

### FunctionName

Description in English.
영문 설명.

**Parameters / 매개변수:**
- `param`: Description / 설명

**Returns / 반환값:**
- Description / 설명

**Example / 예제:**
```go
// Code
```

## Contributing / 기여하기

Contribution guidelines.
기여 가이드라인.

## License / 라이선스

License information.
라이선스 정보.
```

### 2. 기술 문서 (Technical Documentation)

**구조 템플릿**:

```markdown
# Document Title
# 문서 제목

**작성일**: 2025-10-17  
**작성자**: arkd0ng  
**버전**: 1.0.0

---

## 목적 / Purpose

Why this document exists.
이 문서가 존재하는 이유.

## 배경 / Background

Context and history.
맥락과 배경.

## 상세 내용 / Details

### 섹션 1 / Section 1

Content in both languages.
양쪽 언어로 된 내용.

**예제 / Example:**

```
Code or example
코드 또는 예제
```

## 결론 / Conclusion

Summary and next steps.
요약 및 다음 단계.
```

### 3. CHANGELOG 파일

**형식** (docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md):

```markdown
# CHANGELOG v1.12

## [v1.12.021] - 2025-10-17

### Added / 추가
- New feature description
  새 기능 설명
- Another feature
  다른 기능

### Changed / 변경
- Modified behavior in X
  X의 동작 수정
- Updated Y to Z
  Y를 Z로 업데이트

### Fixed / 수정
- Fixed bug in A
  A의 버그 수정
- Resolved issue with B
  B의 문제 해결

### Deprecated / 지원 중단
- Function X is deprecated, use Y instead
  함수 X는 지원 중단됨, 대신 Y 사용

### Removed / 제거
- Removed obsolete feature
  더 이상 사용하지 않는 기능 제거

### Security / 보안
- Security fix for vulnerability
  취약점에 대한 보안 수정
```

### 4. 가이드 문서 (Guide/Tutorial)

**구조**:

```markdown
# How to [Task] Guide
# [작업] 하는 방법 가이드

## Prerequisites / 사전 요구사항

- Requirement 1 / 요구사항 1
- Requirement 2 / 요구사항 2

## Step-by-Step / 단계별 가이드

### Step 1: [Action]
### 1단계: [작업]

Description of what to do.
무엇을 할지 설명.

```bash
# Command to run
# 실행할 명령
```

**Expected result / 예상 결과:**
What you should see.
보게 될 내용.

### Step 2: [Action]
### 2단계: [작업]

Continue with more steps...
더 많은 단계 계속...

## Troubleshooting / 문제 해결

**Problem / 문제:** Description
**Solution / 해결:** Solution

## Next Steps / 다음 단계

What to do after completing this guide.
이 가이드 완료 후 할 일.
```

---

## 품질 기준 / Quality Standards

### 현재 달성 수준 (2025-10-17 기준)

| 항목 | 목표 | 현재 | 상태 |
|------|------|------|------|
| **주석 비율** | ≥ 30% | **51.42%** | 🏆 초과 달성 |
| **이중언어 비율** | ≥ 40% | **~45%** | ✅ 달성 |
| **우수 등급 패키지** | ≥ 70% | **79%** (10/11) | 🏆 초과 달성 |
| **총 주석 라인** | - | **33,211줄** | 📊 통계 |

### 파일별 체크리스트

모든 파일은 다음 기준을 충족해야 합니다:

- [ ] **완전성**: 모든 public 함수/타입/상수에 주석
- [ ] **상세성**: 초보자가 이해할 수 있는 수준
- [ ] **정확성**: 코드와 주석 일치
- [ ] **이중언어**: 영문/한글 동일 수준
- [ ] **예제**: 복잡한 함수에 사용 예제
- [ ] **에러**: 모든 에러 케이스 문서화
- [ ] **성능**: 성능 특성 명시
- [ ] **안전성**: Thread-safety 명시

### 주석 품질 등급

| 등급 | 이중언어 비율 | 설명 |
|------|--------------|------|
| 🏆 **최우수** | ≥ 60% | 업계 최고 수준 |
| ✅ **우수** | ≥ 40% | 목표 달성 |
| ⚠️ **보통** | 20-40% | 개선 권장 |
| ❌ **미흡** | < 20% | 개선 필요 |

---

## 실전 예제 / Practical Examples

### 예제 1: 간단한 함수

```go
// IsEmpty checks if a string is empty or contains only whitespace.
// It trims leading and trailing spaces before checking.
//
// IsEmpty는 문자열이 비어있거나 공백만 포함하는지 확인합니다.
// 확인 전에 앞뒤 공백을 제거합니다.
//
// Parameters / 매개변수:
//   - s: String to check
//     확인할 문자열
//
// Returns / 반환값:
//   - bool: true if empty or whitespace-only, false otherwise
//     비어있거나 공백만 있으면 true, 아니면 false
//
// Example / 예제:
//   IsEmpty("")        // true
//   IsEmpty("  ")      // true
//   IsEmpty("hello")   // false
//   IsEmpty("  hi  ")  // false
func IsEmpty(s string) bool {
    return len(strings.TrimSpace(s)) == 0
}
```

### 예제 2: 복잡한 함수

```go
// ParseDate parses a date string in various formats and returns a time.Time.
// It attempts multiple common formats automatically and returns the first match.
// All parsed times are converted to KST (Asia/Seoul) timezone.
//
// ParseDate는 다양한 형식의 날짜 문자열을 파싱하여 time.Time을 반환합니다.
// 여러 일반적인 형식을 자동으로 시도하고 첫 번째 일치 항목을 반환합니다.
// 모든 파싱된 시간은 KST (Asia/Seoul) 타임존으로 변환됩니다.
//
// Supported formats / 지원 형식:
//   - ISO 8601: "2006-01-02T15:04:05Z"
//   - RFC 3339: "2006-01-02T15:04:05+09:00"
//   - Common: "2006-01-02", "2006/01/02", "01/02/2006"
//   - Custom: "YYYY-MM-DD HH:mm:ss"
//
// Parameters / 매개변수:
//   - dateStr: Date string to parse (cannot be empty)
//     파싱할 날짜 문자열 (비어있을 수 없음)
//     * Whitespace is trimmed automatically
//       공백은 자동으로 제거됨
//     * Case-insensitive for month names
//       월 이름은 대소문자 구분 안 함
//
// Returns / 반환값:
//   - time.Time: Parsed date in KST timezone
//     KST 타임존의 파싱된 날짜
//   - error: Error if parsing fails for all formats
//     모든 형식에서 파싱 실패 시 에러
//
// Errors / 에러:
//   - ErrEmptyString: when dateStr is empty or whitespace-only
//     dateStr이 비어있거나 공백만 있을 때
//   - ErrInvalidFormat: when no format matches the input
//     입력과 일치하는 형식이 없을 때
//     * Returns list of attempted formats in error message
//       에러 메시지에 시도한 형식 목록 포함
//
// Performance / 성능:
//   - Time complexity: O(n) where n is number of formats
//     시간 복잡도: O(n) (n은 형식 개수)
//   - Typically completes in <1ms for valid dates
//     유효한 날짜의 경우 일반적으로 <1ms 내 완료
//   - No memory allocation if first format matches
//     첫 번째 형식이 일치하면 메모리 할당 없음
//
// Notes / 주의사항:
//   - Thread-safe: Yes, no shared state
//     스레드 안전: 예, 공유 상태 없음
//   - Timezone: Always returns KST, regardless of input timezone
//     타임존: 입력 타임존과 관계없이 항상 KST 반환
//   - Ambiguous dates: "01/02/2006" parsed as MM/DD/YYYY (US format)
//     모호한 날짜: "01/02/2006"은 MM/DD/YYYY로 파싱 (미국 형식)
//
// Example / 예제:
//   // ISO 8601 format / ISO 8601 형식
//   date, err := ParseDate("2024-03-15T14:30:00Z")
//   // Returns: 2024-03-15 23:30:00 +0900 KST
//   
//   // Simple date / 간단한 날짜
//   date, err := ParseDate("2024-03-15")
//   // Returns: 2024-03-15 00:00:00 +0900 KST
//   
//   // US format / 미국 형식
//   date, err := ParseDate("03/15/2024")
//   // Returns: 2024-03-15 00:00:00 +0900 KST
//   
//   // Error case / 에러 케이스
//   date, err := ParseDate("invalid")
//   // Returns: zero time, ErrInvalidFormat
func ParseDate(dateStr string) (time.Time, error) {
    // implementation
}
```

### 예제 3: 구조체

```go
// Client represents an HTTP client with automatic retry and timeout handling.
// It manages connection pooling, request lifecycle, and error handling.
//
// Client는 자동 재시도 및 타임아웃 처리를 갖춘 HTTP 클라이언트를 나타냅니다.
// 연결 풀링, 요청 생명주기 및 에러 처리를 관리합니다.
//
// Lifecycle / 생명주기:
//   1. Creation: Use NewClient() with optional configuration
//      생성: 선택적 구성과 함께 NewClient() 사용
//   2. Usage: Call Get/Post/etc methods as needed
//      사용: 필요에 따라 Get/Post 등의 메서드 호출
//   3. Cleanup: No explicit cleanup needed, connections auto-close
//      정리: 명시적 정리 불필요, 연결 자동 종료
//
// Thread Safety / 스레드 안전성:
//   - Safe for concurrent use by multiple goroutines
//     여러 고루틴에서 동시 사용 안전
//   - Internal connection pool is protected by mutex
//     내부 연결 풀은 뮤텍스로 보호됨
//   - Do not modify configuration after creation
//     생성 후 구성 수정하지 말 것
//
// Performance / 성능:
//   - Connection pooling: Reuses connections for better performance
//     연결 풀링: 더 나은 성능을 위해 연결 재사용
//   - Default pool size: 100 connections
//     기본 풀 크기: 100개 연결
//   - Idle timeout: 90 seconds
//     유휴 타임아웃: 90초
//
// Example / 예제:
//   // Basic usage / 기본 사용
//   client := NewClient()
//   var result MyData
//   err := client.Get("/api/data", &result)
//   
//   // Custom configuration / 사용자 정의 구성
//   client := NewClient(
//       WithTimeout(60*time.Second),
//       WithRetry(5),
//       WithBaseURL("https://api.example.com"),
//   )
type Client struct {
    // client is the underlying http.Client for making requests.
    // It manages connection pooling and transport settings.
    //
    // client는 요청을 위한 기본 http.Client입니다.
    // 연결 풀링 및 전송 설정을 관리합니다.
    client *http.Client
    
    // config holds the client configuration.
    // It is immutable after Client creation.
    //
    // config는 클라이언트 구성을 보유합니다.
    // Client 생성 후 불변입니다.
    config *config
    
    // baseURL is prepended to all request URLs.
    // It can be empty for absolute URLs.
    //
    // baseURL은 모든 요청 URL 앞에 추가됩니다.
    // 절대 URL의 경우 비어있을 수 있습니다.
    baseURL string
}
```

---

## 체크리스트 / Checklist

### 코드 작성 전 / Before Writing Code

- [ ] 함수/타입의 목적을 명확히 이해
- [ ] 유사한 기존 코드 검토
- [ ] 에러 처리 전략 계획
- [ ] 성능 요구사항 확인

### 코드 작성 중 / During Writing

- [ ] 의미 있는 변수/함수명 사용
- [ ] 에러는 항상 처리
- [ ] 복잡한 로직에 인라인 주석 추가
- [ ] 매직 넘버는 상수로 정의

### 문서화 작성 / Writing Documentation

- [ ] 패키지 레벨 주석 작성
- [ ] 모든 public 함수/타입에 주석
- [ ] Parameters/Returns 섹션 완성
- [ ] 에러 케이스 모두 문서화
- [ ] 실제 사용 예제 포함
- [ ] 영문/한글 모두 상세하게
- [ ] 성능 특성 명시
- [ ] Thread-safety 명시

### 스크립트 작성 / Writing Scripts

- [ ] 헤더에 설명/사용법 포함
- [ ] 모든 함수에 주석
- [ ] 에러 처리 포함 (set -e, set -u)
- [ ] 변수에 의미 있는 이름 사용
- [ ] 중요 변수에 주석 추가
- [ ] 사용 예제 제공
- [ ] 영문/한글 이중언어

### 별도 문서 작성 / Writing Standalone Documentation

- [ ] 문서 목적 명확히 설명
- [ ] 표준 구조 따르기 (Overview → Details → Examples)
- [ ] 코드 예제 포함
- [ ] 단계별 설명 (필요 시)
- [ ] 문제 해결 섹션 (필요 시)
- [ ] 영문/한글 이중언어
- [ ] 메타데이터 포함 (작성일, 작성자, 버전)

### 완료 전 / Before Completion

- [ ] `go test` 실행 및 통과
- [ ] `go build` 실행 및 통과
- [ ] 스크립트 실행 테스트 (해당 시)
- [ ] 코드 리뷰 (self-review)
- [ ] 문서 오타 확인
- [ ] 예제 코드 동작 확인
- [ ] 링크 유효성 확인 (문서에 포함된 경우)

---

## 참고 자료 / References

### 내부 문서
- [BILINGUAL_AUDIT.md](BILINGUAL_AUDIT.md) - 이중언어 감사 현황
- [temp/Status-Code-Comment.md](temp/Status-Code-Comment.md) - 문서화 상태 전수 검사 보고서
- [temp/todo-codex.md](temp/todo-codex.md) - 작업 마스터 체크리스트

### 외부 참고
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

---

## 버전 히스토리 / Version History

| 버전 | 날짜 | 변경사항 |
|------|------|---------|
| 1.1.0 | 2025-10-17 | 스크립트 작성 가이드 추가, 문서 작성 가이드 추가 |
| 1.0.0 | 2025-10-17 | 초기 버전 생성 |

---

**문의 / Contact**: arkd0ng  
**프로젝트**: [go-utils](https://github.com/arkd0ng/go-utils)
