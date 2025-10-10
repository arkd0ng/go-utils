# CHANGELOG - v1.3.x

This document tracks all changes made in version 1.3.x of the go-utils library.

이 문서는 go-utils 라이브러리의 버전 1.3.x에서 이루어진 모든 변경사항을 추적합니다.

---

## [v1.3.001] - 2025-10-10

### Added / 추가
- **Design Documents / 설계 문서**:
  - Created `docs/database/mysql/DESIGN_PLAN.md` - Comprehensive design plan for database/mysql package
  - Created `docs/database/mysql/WORK_PLAN.md` - Detailed work plan with 5 phases
  - database/mysql 패키지에 대한 종합 설계 계획서 작성
  - 5단계로 구성된 상세 작업 계획서 작성

- **Key Features Planned / 주요 기획 기능**:
  - Extreme simplicity: 30 lines → 2 lines of code / 극도의 간결함: 30줄 → 2줄 코드
  - Auto connection management with reconnection / 자동 재연결을 포함한 연결 관리
  - Auto retry on transient errors / 일시적 에러 자동 재시도
  - Auto resource cleanup (no defer rows.Close()) / 자동 리소스 정리
  - Three-layer API: Simple, Query Builder, Raw SQL / 3계층 API
  - **Dynamic credential rotation support / 동적 자격 증명 순환 지원**:
    - User-provided credential refresh function / 사용자 제공 자격 증명 갱신 함수
    - Multiple connection pools with rolling rotation / 순환 교체 방식의 다중 연결 풀
    - Zero-downtime credential updates / 무중단 자격 증명 업데이트
    - Compatible with Vault, AWS Secrets Manager, etc. / Vault, AWS Secrets Manager 등과 호환

- **Design Philosophy / 설계 철학**:
  - Zero Mental Overhead: Connect once, forget about DB state / 한 번 연결하면 DB 상태를 잊어버려도 됨
  - SQL-Like API: Close to actual SQL syntax / SQL 문법에 가까운 API
  - Auto Everything: All tedious tasks handled automatically / 모든 번거로운 작업 자동 처리

### Changed / 변경
- **Version / 버전**: Updated from v1.2.004 to v1.3.001
- **Focus / 초점**: Starting database utility development / 데이터베이스 유틸리티 개발 시작

### Design Highlights / 설계 주요 사항

**File Structure (15 files) / 파일 구조 (15개 파일)**:
```
database/mysql/
├── client.go          # Client struct, New(), Close()
├── connection.go      # Auto connection management
├── rotation.go        # Credential rotation (optional)
├── simple.go          # Simple API (SelectAll, Insert, etc.)
├── builder.go         # Query builder API
├── transaction.go     # Transaction support
├── retry.go           # Auto retry logic
├── scan.go            # Result scanning
├── config.go          # Configuration
├── options.go         # Functional options
├── errors.go          # Error types
├── types.go           # Common types
├── client_test.go     # Unit tests
├── rotation_test.go   # Rotation tests
└── README.md          # Documentation
```

**Usage Example / 사용 예시**:
```go
// Static credentials / 정적 자격 증명
db, _ := mysql.New(mysql.WithDSN("user:pass@tcp(localhost:3306)/db"))

// Dynamic credentials (Vault, etc.) / 동적 자격 증명 (Vault 등)
db, _ := mysql.New(
    mysql.WithCredentialRefresh(
        func() (string, error) {
            // User fetches credentials from Vault, file, etc.
            // 사용자가 Vault, 파일 등에서 자격 증명 가져오기
            return "user:pass@tcp(localhost:3306)/db", nil
        },
        3,              // 3 connection pools / 3개 연결 풀
        1*time.Hour,    // Rotate one per hour / 1시간마다 하나씩 교체
    ),
)

// Simple queries / 간단한 쿼리
users, _ := db.SelectAll("users", "age > ?", 18)
db.Insert("users", map[string]interface{}{"name": "John", "age": 30})
```

**Zero-Downtime Credential Rotation / 무중단 자격 증명 순환**:
```
Time 0:00 - [Session1, Session2, Session3] (all with Credential A)
Time 1:00 - [Session1, Session2, Session3-NEW] (Session3 rotated to Credential B)
Time 2:00 - [Session1, Session2-NEW, Session3-NEW] (Session2 rotated to Credential B)
            ↑ Credential A expires, but Session2 & Session3 still work!
Time 3:00 - [Session1-NEW, Session2-NEW, Session3-NEW] (Session1 rotated to Credential C)
```

### Notes / 참고사항
- This version contains **design documents only** / 이 버전은 **설계 문서만** 포함
- Implementation will follow in subsequent patches / 구현은 후속 패치에서 진행
- Vault integration is **user's responsibility** (not built-in) / Vault 통합은 **사용자 책임** (내장 아님)
- Package follows extreme simplicity principle: "If not 10x simpler, don't build it" / 극도의 간결함 원칙 준수: "10배 간단하지 않으면 만들지 마세요"

---

**Version History / 버전 히스토리**:
- v1.3.001: Design documents for database/mysql package / database/mysql 패키지 설계 문서
