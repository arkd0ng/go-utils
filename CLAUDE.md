# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 프로젝트 개요

`go-utils`는 Golang 개발을 위한 모듈화된 유틸리티 패키지 모음입니다. 라이브러리는 서브패키지 구조로 설계되어 사용자가 전체 라이브러리가 아닌 특정 유틸리티만 import할 수 있습니다.

**GitHub 저장소**: `github.com/arkd0ng/go-utils`
**현재 버전**: v0.2.0
**Go 버전**: 1.24.6
**라이선스**: MIT

## 아키텍처

### 서브패키지 구조

라이브러리는 각 유틸리티 카테고리가 자체 디렉토리에 존재하는 서브패키지 아키텍처를 따릅니다:

```
go-utils/
├── random/              # 암호학적으로 안전한 랜덤 문자열 생성
│   ├── string.go       # 핵심 랜덤 문자열 생성 로직
│   ├── string_test.go  # 충돌 확률 테스트를 포함한 종합 테스트
│   └── README.md       # 패키지별 문서 (이중 언어: 영문/한글)
├── examples/
│   └── random_string/  # 모든 메서드를 시연하는 실행 예제
└── (향후 패키지)        # stringutil, sliceutil, maputil, fileutil 등
```

### 설계 원칙

1. **독립성**: 각 서브패키지는 교차 의존성 없이 자체 포함됩니다
2. **이중 언어 문서화**: 모든 문서, 주석, 테스트는 영문과 한글로 작성됩니다
3. **보안 우선**: 암호학적으로 안전한 랜덤 생성을 위해 `crypto/rand` 사용
4. **에러 처리**: 모든 메서드는 적절한 에러 처리를 위해 `(결과, error)`를 반환합니다
5. **가변 인자**: 메서드는 유연성을 위해 가변 `length` 인자를 받습니다:
   - 1개 인자: 고정 길이 (예: `Alnum(32)`는 정확히 32자 생성)
   - 2개 인자: 범위 (예: `Alnum(32, 128)`는 32-128자 생성)

### Random 패키지 아키텍처

`random` 패키지는 전역 싱글톤 패턴을 사용합니다:
- `stringGenerator` 구조체가 모든 생성 메서드를 제공
- `GenString`은 `random.GenString.메서드명()`으로 접근하는 전역 인스턴스
- 핵심 헬퍼 함수 `generateRandomString()`이 검증 및 생성 로직을 처리
- 문자 집합은 패키지 레벨 상수로 정의됨

**14개의 사용 가능한 메서드**:
- 기본: `Letters`, `Alnum`, `Digits`, `Complex`, `Standard`
- 대소문자 구분: `AlphaUpper`, `AlphaLower`, `AlnumUpper`, `AlnumLower`
- 16진수: `Hex`, `HexLower`
- 인코딩: `Base64`, `Base64URL`
- 사용자 정의: `Custom(charset string, length ...int)`

## 개발 워크플로우

### 빌드 및 테스트

```bash
# 모든 테스트를 상세 출력으로 실행
go test ./... -v

# 특정 패키지에 대한 테스트 실행
go test ./random -v

# 단일 테스트 실행
go test ./random -v -run TestLetters

# 벤치마크 실행
go test ./... -bench=.
go test ./random -bench=BenchmarkAlnum

# 커버리지와 함께 테스트 실행
go test ./... -cover
go test ./random -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### 예제 실행

```bash
# random string 예제 실행
go run examples/random_string/main.go

# 예제 바이너리 빌드
go build -o bin/random_example examples/random_string/main.go
```

### 새로운 기능 추가

`random` 패키지에 새로운 메서드를 추가할 때:

1. 필요한 경우 `random/string.go` 상단에 문자 집합 상수 추가
2. `stringGenerator` 구조체에 메서드 생성
3. 포괄적인 이중 언어 문서 포함 (영문/한글)
4. 적절한 charset과 함께 `generateRandomString(charset, length...)` 호출
5. `random/string_test.go`에 해당 테스트 추가:
   - 기능 테스트 (길이, charset 정확성 검증)
   - 해당되는 경우 엣지 케이스 테스트
   - 벤치마크 테스트
6. `random/README.md`에 메서드 문서 업데이트
7. `examples/random_string/main.go`에 사용 예제 추가
8. 새 패키지나 주요 기능을 추가하는 경우 루트 `README.md` 업데이트

### 새로운 유틸리티 패키지 생성

새로운 유틸리티 패키지(예: `stringutil`, `sliceutil`)를 추가할 때:

1. 새 디렉토리 생성: `mkdir packagename`
2. 이중 언어 주석과 함께 패키지 파일 생성
3. 종합 테스트 파일 생성: `packagename_test.go`
4. 패키지 README 생성: `packagename/README.md` (이중 언어)
5. `examples/packagename/main.go`에 예제 추가
6. 새 패키지를 반영하도록 루트 `README.md` 업데이트
7. 패키지가 교차 의존성 없이 자체 포함되도록 보장

## 테스트 요구사항

모든 테스트는 다음을 포함해야 합니다:

1. **기능 테스트**: 예상 동작 및 출력 특성 검증
2. **엣지 케이스 테스트**: 음수 값, 잘못된 입력, 경계 조건
3. **랜덤성 테스트**: 고유성 및 적절한 분포 확인
4. **충돌 확률 테스트**: 랜덤 생성의 경우, 이론적 대 실제 충돌률 계산 및 검증
5. **벤치마크 테스트**: 모든 공개 메서드에 대한 성능 벤치마크
6. **이중 언어 주석**: 영문과 한글 설명 모두 포함

## 문서화 표준

### 코드 주석

모든 코드 주석은 이중 언어여야 합니다 (영문 먼저, 한글 두 번째):

```go
// Letters generates a random string containing only alphabetic characters (a-z, A-Z)
// Letters는 알파벳 문자(a-z, A-Z)만 포함하는 랜덤 문자열을 생성합니다
```

### README 파일

모든 README 파일은 병렬 구조로 이중 언어여야 합니다:
- 영문 문장 다음에 한글 번역
- 기술 용어는 두 언어로 표시
- 이중 언어 주석이 있는 코드 예제

## Import 경로

사용자는 특정 패키지를 import합니다:

```go
import "github.com/arkd0ng/go-utils/random"
```

루트 패키지를 import하지 않습니다:

```go
import "github.com/arkd0ng/go-utils"  // ❌ 이렇게 하지 마세요
```

## 에러 처리 패턴

모든 메서드는 `(결과, error)`를 반환합니다:

```go
str, err := random.GenString.Alnum(32)
if err != nil {
    log.Fatal(err)
}
```

에러 메시지는 영문으로만 작성됩니다 (이중 언어 아님).

## 버전 히스토리 컨텍스트

- **v0.1.0**: 루트 레벨 `GenRandomString`으로 첫 릴리스
- **v0.2.0** (현재): Breaking change - 서브패키지 구조로 리팩토링
  - `GenRandomString`에서 `random.GenString`으로 변경
  - 9개의 새로운 메서드 추가 (총 14개)
  - 가변 인자 및 에러 처리 추가
  - 이중 언어 문서 추가

## 향후 로드맵

계획된 유틸리티 패키지 (README에 참조됨):
- `stringutil` - 문자열 처리 유틸리티
- `sliceutil` - 슬라이스/배열 헬퍼
- `maputil` - 맵 유틸리티
- `fileutil` - 파일/경로 유틸리티
- `httputil` - HTTP 헬퍼
- `timeutil` - 시간/날짜 유틸리티
- `validation` - 검증 유틸리티
- `errorutil` - 에러 처리 헬퍼
