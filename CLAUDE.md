# CLAUDE.md

이 파일은 Claude Code(claude.ai/code)가 이 저장소에서 작업할 때 따라야 할 가이드를 제공합니다.

---

## 🚨 최우선 규칙 - 언어 사용 정책

### 📝 문서 언어 규칙 (반드시 준수)

#### 1. 공개 문서 (PUBLIC) - 영문/한글 병기 필수
- ✅ **모든 .md 파일** (README.md, USER_MANUAL.md, DEVELOPER_GUIDE.md 등)
- ✅ **CHANGELOG 파일**
- ✅ **공개 설계 문서** (DESIGN_PLAN.md, WORK_PLAN.md)

**형식**: "English / 한글" 또는 영문 문단 다음에 한글 문단

#### 2. 비공개 문서 (PRIVATE) - 한글만 사용
- ✅ **CLAUDE.md** (AI 어시스턴트 가이드 - 이 파일)
- ✅ **todo.md** (개인 작업 목록)
- ✅ **개인 노트 및 메모**

**형식**: 한글로만 작성 (영문 용어는 괄호로 병기 가능)

#### 3. 코드 주석 - 영문/한글 병기 필수 + 매우 상세하게
```go
// GetValue retrieves a value from the map by key.
// It returns the value and a boolean indicating whether the key exists.
// If the key is not found, returns the zero value for the value type.
// GetValue는 키로 맵에서 값을 검색합니다.
// 키가 존재하는지 여부를 나타내는 불리언과 함께 값을 반환합니다.
// 키를 찾지 못하면 값 타입의 제로값을 반환합니다.
func GetValue[K comparable, V any](m map[K]V, key K) (V, bool) {
    // Check if key exists in map / 맵에 키가 존재하는지 확인
    val, ok := m[key]
    return val, ok
}
```

**중요**: 주석은 **매우 상세하고 친절하게** 작성합니다. 초보자도 이해할 수 있도록!

#### 4. 로그 메시지 - 영문/한글 병기 필수
```go
logger.Info("Processing user data / 사용자 데이터 처리 중", "userID", userID)
logger.Error("Failed to connect to database / 데이터베이스 연결 실패", "error", err)
logger.Debug("Cache hit / 캐시 히트", "key", cacheKey, "ttl", ttl)
```

#### 5. Git 커밋 메시지 - 영문/한글 병기 필수
```bash
# ✅ 올바른 형식
git commit -m "Feat: Add Get function / Get 함수 추가 (v1.12.004)"
git commit -m "Fix: Handle nil pointer / nil 포인터 처리 (v1.12.005)"
git commit -m "Docs: Update USER_MANUAL / USER_MANUAL 업데이트 (v1.12.006)"

# ❌ 잘못된 형식 (한글 누락)
git commit -m "Feat: Add Get function (v1.12.004)"
```

#### 6. 변수/함수 이름 - 영문만 사용
```go
func CalculateTotal() int  // ✅ 올바름
func 합계계산() int          // ❌ 잘못됨
```

### ⚠️ 규칙 위반 시 조치

규칙 위반을 발견하면:
1. **즉시 중단** - 작업을 멈추고 사용자에게 알림
2. **전체 검토** - 관련된 모든 코드와 문서를 확인
3. **일괄 수정** - 발견된 모든 위반 사항을 수정
4. **CHANGELOG 기록** - 수정 내역을 CHANGELOG에 기록

---

## 🚨 반드시 먼저 읽어야 할 문서

**모든 작업 전에 다음 문서들을 반드시 읽으세요:**

1. **[DEVELOPMENT_WORKFLOW_GUIDE.md](./docs/DEVELOPMENT_WORKFLOW_GUIDE.md)** ⭐ 가장 중요
   - 모든 개발 작업의 완전한 워크플로우
   - 핵심 규칙 및 표준 작업 사이클
   - 모든 작업 시 가장 먼저 읽을 것

2. **[PACKAGE_DEVELOPMENT_GUIDE.md](./docs/PACKAGE_DEVELOPMENT_GUIDE.md)** ⭐ 필수
   - 패키지 개발 표준 및 워크플로우
   - 브랜치 전략, 버전 관리, 단위 작업 워크플로우
   - 예제 코드 및 로깅 가이드라인

3. **[CODE_TEST_MAKE_GUIDE.md](./docs/CODE_TEST_MAKE_GUIDE.md)**
   - 테스트 표준 및 가이드라인
   - 테스트 구조 및 커버리지 요구사항

4. **[EXAMPLE_CODE_GUIDE.md](./docs/EXAMPLE_CODE_GUIDE.md)**
   - 예제 코드 구조 및 요구사항
   - 로깅 모범 사례

---

## 🔄 표준 작업 사이클 (Standard Work Cycle)

**모든 작업은 다음 순서를 정확히 따라야 합니다:**

```
1. 버전 증가 (cfg/app.yaml)
   ↓
2. 작업 수행 (코드/문서)
   ↓
3. 테스트 및 검증 (go test ./...)
   ↓
4. CHANGELOG 업데이트 (커밋 전 필수!)
   ↓
5. Git 커밋 및 푸시
```

### 절대 하지 말아야 할 것 (NEVER)
- ❌ 버전 증가 없이 작업 시작
- ❌ CHANGELOG 업데이트 없이 커밋
- ❌ 테스트 실패 상태에서 푸시
- ❌ 문서화 없이 코드만 푸시

### 항상 해야 할 것 (ALWAYS)
- ✅ 작업 전 항상 버전 증가
- ✅ 작업 후 항상 CHANGELOG 업데이트
- ✅ 커밋 전 항상 테스트 실행
- ✅ 모든 것을 이중 언어로 문서화

---

## 📦 프로젝트 개요

- **저장소**: `github.com/arkd0ng/go-utils`
- **현재 버전**: v1.12.003 (cfg/app.yaml 참조)
- **Go 버전**: 1.24.6
- **라이선스**: MIT

### 목적
Golang 개발을 위한 모듈화된 유틸리티 패키지 모음입니다. 각 서브패키지는 독립적이며 개별적으로 import할 수 있습니다.

### 설계 원칙
1. **극단적 단순함** - 20-30줄 → 1-2줄로 줄이기
2. **독립성** - 패키지 간 의존성 없음
3. **이중 언어** - 모든 문서 영문/한글 병기
4. **타입 안전성** - Go 1.18+ 제네릭 적절히 사용
5. **제로 설정** - 99% 사용 사례에 대한 합리적 기본값

---

## 📚 패키지 구조

```
go-utils/
├── random/          # 암호학적으로 안전한 랜덤 문자열 (14개 메서드)
├── logging/         # 파일 로테이션 기능이 있는 구조화된 로깅
├── database/
│   ├── mysql/      # 극도로 간단한 MySQL 클라이언트 (30줄 → 2줄)
│   └── redis/      # 극도로 간단한 Redis 클라이언트 (20줄 → 2줄)
├── stringutil/     # 문자열 유틸리티 (53개 함수, 9개 카테고리)
├── timeutil/       # 시간/날짜 유틸리티 (114개 함수, 10개 카테고리)
├── sliceutil/      # 슬라이스 유틸리티 (95개 함수, 14개 카테고리)
├── maputil/        # 맵 유틸리티 (99개 함수, 14개 카테고리)
├── fileutil/       # 파일/경로 유틸리티 (~91개 함수, 12개 카테고리)
├── httputil/       # HTTP 클라이언트 유틸리티 (10개 메서드 + 12개 옵션)
└── websvrutil/     # 웹 서버 유틸리티 (종합 기능)
```

상세 패키지 정보는 [README.md](./README.md) 참조

---

## 🔢 버전 관리

### 버전 형식
```
vMAJOR.MINOR.PATCH
예: v1.12.003
```

- **MAJOR**: 호환성을 깨는 변경 (드물게)
- **MINOR**: 새 패키지
- **PATCH**: 모든 단위 작업

### 버전 증가 절차

**모든 작업 전에 버전을 증가시킵니다:**

```bash
# 1. cfg/app.yaml 편집
# version: v1.12.003 → v1.12.004

# 2. 버전 변경 커밋 (작업 전 먼저!)
git add cfg/app.yaml
git commit -m "Chore: Bump version to v1.12.004 / v1.12.004로 버전 증가"

# 3. 이제 작업 시작
```

---

## 📝 CHANGELOG 관리

### 🚨 매우 중요: 모든 커밋 전에 CHANGELOG 업데이트 필수

**예외**: 버전 증가만 하는 커밋

### 2단계 CHANGELOG 시스템

1. **루트 `CHANGELOG.md`**
   - 메이저/마이너 버전 개요만
   - 상세 파일로 링크

2. **`docs/CHANGELOG/CHANGELOG-v1.{MINOR}.md`**
   - 모든 패치 버전 상세 기록
   - 무엇이 변경되었는지
   - 어떤 파일이 변경되었는지
   - 왜 변경되었는지
   - 사용자 요청 컨텍스트

### CHANGELOG 템플릿

`docs/CHANGELOG/CHANGELOG-v1.12.md` 예시:

```markdown
## [v1.12.004] - 2025-10-17

### Added / 추가
- 추가된 기능 설명

### Changed / 변경
- 변경된 내용 설명

### Fixed / 수정
- 수정된 버그 설명

### Files Changed / 변경된 파일
- `path/to/file.go` - 변경 사항 설명
- `path/to/test.go` - 테스트 추가

### Context / 컨텍스트
**User Request / 사용자 요청**: "원본 요청 내용"

**Why / 이유**: 변경한 이유

**Impact / 영향**: 이 변경으로 가능해지는 것
```

---

## 🧪 테스트

### 커버리지 요구사항
- **최소**: 60%
- **목표**: 80%+
- **중요 함수**: 100%

### 테스트 명령어

```bash
# 모든 테스트 실행
go test ./... -v

# 특정 패키지 테스트
go test ./maputil -v
go test ./stringutil -v
go test ./database/mysql -v

# 커버리지 포함
go test ./maputil -cover
go test ./... -coverprofile=coverage.out

# 브라우저에서 커버리지 보기
go tool cover -html=coverage.out

# 특정 테스트만 실행
go test ./maputil -v -run TestGet

# 벤치마크 실행
go test ./maputil -bench=.
go test ./maputil -bench=BenchmarkGet -benchmem
```

---

## 🛠️ 개발 도구

### 빌드

```bash
# 모든 패키지 빌드
go build ./...

# 특정 패키지 빌드
go build ./maputil
go build ./database/mysql

# 컴파일 에러만 확인 (바이너리 생성 안 함)
go build -o /dev/null ./...
```

### 예제 실행

```bash
# 특정 패키지의 예제 실행
go run examples/maputil/main.go
go run examples/stringutil/main.go

# 예제 로그는 logs/{package}/ 에 저장됨
ls -la logs/maputil/
```

### 코드 찾기

```bash
# 패키지의 모든 export된 함수 찾기
grep "^func [A-Z]" maputil/*.go

# 특정 함수 사용처 검색
grep -r "maputil.Get" examples/

# 모든 테스트 파일 찾기
find . -name "*_test.go"

# 패키지의 코드 라인 수 세기 (테스트 제외)
find ./maputil -name "*.go" -not -name "*_test.go" | xargs wc -l
```

### Docker (MySQL/Redis 테스트용)

```bash
# MySQL 컨테이너 시작/중지
bash ./.docker/scripts/docker-mysql-start.sh
bash ./.docker/scripts/docker-mysql-stop.sh

# Redis 컨테이너 시작/중지
bash ./.docker/scripts/docker-redis-start.sh
bash ./.docker/scripts/docker-redis-stop.sh

# 실행 중인 컨테이너 확인
docker ps

# 로그 보기
docker logs go-utils-mysql
docker logs go-utils-redis
```

---

## 🔧 Git 워크플로우

### 커밋 메시지 형식

```
<type>: <영문 제목 / 한글 제목> (<version>)

[선택적 본문]

🤖 Generated with Claude Code
Co-Authored-By: Claude <noreply@anthropic.com>
```

### 커밋 타입
- **Feat**: 새 기능
- **Fix**: 버그 수정
- **Docs**: 문서
- **Test**: 테스트
- **Refactor**: 리팩토링
- **Chore**: 빌드, 버전
- **Perf**: 성능
- **Style**: 포맷팅

### 커밋 예시

```bash
# 버전 증가
git commit -m "Chore: Bump version to v1.12.004 / v1.12.004로 버전 증가"

# 새 기능
git commit -m "Feat: Add Get function to maputil / maputil에 Get 함수 추가 (v1.12.004)"

# 버그 수정
git commit -m "Fix: Handle nil pointer in Get / Get에서 nil 포인터 처리 (v1.12.005)"

# 문서
git commit -m "Docs: Update USER_MANUAL / USER_MANUAL 업데이트 (v1.12.006)"

# 테스트
git commit -m "Test: Add edge case tests for Get / Get 엣지 케이스 테스트 추가 (v1.12.007)"
```

---

## 📋 매 작업마다 체크리스트

```
□ 1. cfg/app.yaml에서 버전 증가
□ 2. 버전 증가 커밋
□ 3. 작업 수행 (코드/문서)
□ 4. 테스트 실행 (go test ./... -v)
□ 5. CHANGELOG 업데이트
□ 6. 적절한 메시지로 커밋
□ 7. GitHub에 푸시
```

---

## 📖 추가 자료

### 핵심 문서
- **[README.md](./README.md)** - 프로젝트 개요 및 패키지 목록
- **[DEVELOPMENT_WORKFLOW_GUIDE.md](./docs/DEVELOPMENT_WORKFLOW_GUIDE.md)** ⭐ 메인 워크플로우 가이드
- **[PACKAGE_DEVELOPMENT_GUIDE.md](./docs/PACKAGE_DEVELOPMENT_GUIDE.md)** ⭐ 패키지 개발 표준
- **[CODE_TEST_MAKE_GUIDE.md](./docs/CODE_TEST_MAKE_GUIDE.md)** - 테스트 가이드라인
- **[EXAMPLE_CODE_GUIDE.md](./docs/EXAMPLE_CODE_GUIDE.md)** - 예제 코드 표준

### 패키지별 문서

각 패키지는 자체 상세 문서를 가지고 있습니다:

```
{package}/README.md                      # 빠른 시작
docs/{package}/USER_MANUAL.md           # 사용자 가이드
docs/{package}/DEVELOPER_GUIDE.md       # 개발자 가이드
docs/{package}/DESIGN_PLAN.md           # 설계 (해당되는 경우)
docs/{package}/WORK_PLAN.md             # 작업 계획 (해당되는 경우)
examples/{package}/main.go               # 실행 가능한 예제
```

### 외부 의존성
- `github.com/go-sql-driver/mysql` - MySQL 드라이버
- `github.com/redis/go-redis/v9` - Redis 클라이언트
- `gopkg.in/natefinch/lumberjack.v2` - 로그 로테이션
- `gopkg.in/yaml.v3` - YAML 파싱
- `golang.org/x/text` - 유니코드 정규화
- `golang.org/x/exp` - 제네릭 제약조건

---

## ⚠️ 핵심 알림

1. **항상 DEVELOPMENT_WORKFLOW_GUIDE.md를 먼저 읽을 것**
2. **항상 작업 전에 버전 증가**
3. **항상 CHANGELOG 업데이트**
4. **항상 커밋 전에 테스트**
5. **항상 두 언어로 문서화**
6. **주석은 매우 상세하고 친절하게**

---

## 🎓 신규 기여자를 위한 학습 경로

1. [DEVELOPMENT_WORKFLOW_GUIDE.md](./docs/DEVELOPMENT_WORKFLOW_GUIDE.md) 읽기
2. [PACKAGE_DEVELOPMENT_GUIDE.md](./docs/PACKAGE_DEVELOPMENT_GUIDE.md) 읽기
3. 기존 패키지 README 둘러보기
4. `examples/`의 예제 코드 검토
5. `*_test.go` 파일의 패키지 테스트 확인
6. 참고용 패키지의 USER_MANUAL 및 DEVELOPER_GUIDE 읽기

---

**최종 업데이트**: 2025-10-17
**버전**: v1.12.003
**관리자**: go-utils team
