# CHANGELOG - v1.3.x

This document tracks all changes made in version 1.3.x of the go-utils library.

이 문서는 go-utils 라이브러리의 버전 1.3.x에서 이루어진 모든 변경사항을 추적합니다.

---

## [v1.3.014] - 2025-10-14 (Bug Fixes: Soft Delete Schema & MySQL 8.0+ Compatibility)

### Fixed / 수정
- **Soft Delete Schema Preparation** / 소프트 삭제 스키마 준비:
  - Added Example 24.5 to automatically create `deleted_at` column before soft delete examples
  - Prevents "Unknown column 'deleted_at'" error when running examples
  - Checks if column exists and skips migration if already present
  - Verifies table structure after column addition

- **MySQL 8.0+ Compatibility** / MySQL 8.0+ 호환성:
  - Fixed `schema.go` to handle MySQL 8.0+ `SHOW INDEX` output (15 columns vs 14)
  - Added `expression` field to handle Expression column introduced in MySQL 8.0
  - Fixes "sql: expected 15 destination arguments in Scan, not 14" error

### Changed / 변경
- **examples/mysql/main.go**:
  - Added `example24_5PrepareForSoftDelete()` function (~60 lines)
  - Automatically prepares users table with deleted_at column
  - Integrated between Example 24 and Example 25

- **database/mysql/schema.go**:
  - Updated `GetIndexes()` method to scan 15 columns for MySQL 8.0+ compatibility
  - Added `expression sql.NullString` variable for Expression column

### Testing / 테스트
- All 35 examples now run successfully without errors
- Soft delete examples (25-27) work correctly after schema preparation
- Schema inspection (Example 32) works with MySQL 8.0+

---

## [v1.3.013] - 2025-10-14 (MySQL Documentation Update)

### Changed / 변경
- **Updated MySQL Package Documentation** / MySQL 패키지 문서 업데이트:
  - `docs/database/mysql/USER_MANUAL.md`: Added comprehensive "Advanced Features" section with 9 new feature categories
  - `docs/database/mysql/DEVELOPER_GUIDE.md`: Added "Advanced Features Architecture" section with detailed implementation explanations
  - Updated version numbers to v1.3.013 in all documentation files
  - Updated file structure table to include all new feature files

### Documentation / 문서화
- **USER_MANUAL.md Updates** (~900 new lines):
  - Section 6: Advanced Features with 9 subsections
  - Batch Operations: BatchInsert, BatchUpdate, BatchDelete, BatchSelectByIDs with complete examples
  - Upsert Operations: Upsert, UpsertBatch, Replace with usage patterns
  - Pagination: Paginate, PaginateQuery with helper methods and navigation examples
  - Soft Delete: SoftDelete, Restore, SelectAllWithTrashed, SelectAllOnlyTrashed, PermanentDelete
  - Query Statistics: EnableQueryStats, EnableSlowQueryLog, GetQueryStats with monitoring examples
  - Pool Metrics: GetPoolMetrics, GetPoolHealthInfo, GetConnectionUtilization, GetWaitStatistics
  - Schema Inspector: GetTables, GetColumns, GetIndexes, InspectTable with introspection examples
  - Migration Helpers: CreateTable, AddColumn, AddIndex, AddForeignKey and 20+ DDL operations
  - CSV Export/Import: ExportTableToCSV, ImportFromCSV, ExportQueryToCSV with options

- **DEVELOPER_GUIDE.md Updates** (~300 new lines):
  - Advanced Features Architecture section with implementation details for all 9 features
  - Design goals and philosophy for each feature
  - Key design decisions explained with rationale
  - Performance characteristics and optimization tips
  - Thread safety and concurrency considerations
  - Architecture diagrams and flow explanations
  - Updated package structure table with new files (~6,000+ total lines)

### Features Documented / 문서화된 기능
1. **Batch Operations**: Performance optimization for bulk operations
2. **Upsert Operations**: Idempotent inserts with MySQL's ON DUPLICATE KEY UPDATE
3. **Pagination**: Comprehensive pagination with rich metadata
4. **Soft Delete**: Data recovery and audit trail support
5. **Query Statistics**: Zero-overhead performance monitoring
6. **Pool Metrics**: Connection pool health monitoring
7. **Schema Inspector**: Programmatic access to database metadata
8. **Migration Helpers**: High-level API for schema changes
9. **CSV Export/Import**: Efficient data exchange with external systems

### Best Practices / 모범 사례
- Added "When to use" sections for each advanced feature
- Performance tips and common pitfalls documented
- Cross-references between USER_MANUAL and DEVELOPER_GUIDE
- Dual-language format maintained throughout (English/Korean)
- Complete code examples with inline comments

---

## [v1.3.012] - 2025-10-14 (MySQL Comprehensive Test Suite)

### Added / 추가
- **Comprehensive Test Files for MySQL Advanced Features** / MySQL 고급 기능을 위한 종합 테스트 파일:
  - `database/mysql/batch_test.go`: Batch operations tests (insert, update, delete, select by IDs)
  - `database/mysql/upsert_test.go`: Upsert and replace operations tests
  - `database/mysql/pagination_test.go`: Pagination functionality tests with metadata validation
  - `database/mysql/softdelete_test.go`: Soft delete workflow tests (delete, restore, query trashed)
  - `database/mysql/stats_test.go`: Query statistics and slow query logging tests
  - `database/mysql/metrics_test.go`: Pool metrics and health monitoring tests
  - `database/mysql/schema_test.go`: Schema inspection tests (tables, columns, indexes)
  - `database/mysql/migration_test.go`: Database migration tests (DDL operations)
  - `database/mysql/export_test.go`: CSV export/import functionality tests

### Features / 기능
- **Test Coverage** / 테스트 커버리지:
  - Table-driven testing for comprehensive coverage
  - Edge case testing (empty data, invalid inputs, boundary conditions)
  - Dual-language comments (English/Korean) in all tests
  - Benchmark tests for performance-critical operations
  - Structure validation tests for data types
  - Skip markers for tests requiring actual database connection

### Files Added / 추가된 파일
- `database/mysql/batch_test.go` (228 lines)
- `database/mysql/upsert_test.go` (263 lines)
- `database/mysql/pagination_test.go` (321 lines)
- `database/mysql/softdelete_test.go` (306 lines)
- `database/mysql/stats_test.go` (384 lines)
- `database/mysql/metrics_test.go` (308 lines)
- `database/mysql/schema_test.go` (251 lines)
- `database/mysql/migration_test.go` (340 lines)
- `database/mysql/export_test.go` (333 lines)

### Testing / 테스트
- All test files follow existing test patterns from `client_test.go`
- Tests are designed to work without actual database where possible
- Validation errors are tested without database connection
- Database-dependent tests are properly skipped with descriptive messages
- Benchmark tests included for performance monitoring

---

## [v1.3.012] - 2025-10-14 (MySQL Advanced Features Examples)

### Added / 추가
- **Examples for 9 Advanced Features** / 9개 고급 기능 예제:
  - Added 18 new example functions (example18-example35) in `examples/mysql/main.go`
  - **Batch Operations** (3 examples):
    - `example18BatchInsert`: Demonstrate BatchInsert with 4 users
    - `example19BatchUpdate`: Demonstrate BatchUpdate in transaction
    - `example20BatchDelete`: Demonstrate BatchDelete by IDs
  - **Upsert** (2 examples):
    - `example21Upsert`: Demonstrate Upsert with insert and update
    - `example22UpsertBatch`: Demonstrate batch upsert operations
  - **Pagination** (2 examples):
    - `example23Pagination`: Demonstrate basic pagination with metadata
    - `example24PaginationWithOptions`: Demonstrate pagination with WHERE and ORDER BY
  - **Soft Delete** (3 examples):
    - `example25SoftDelete`: Demonstrate soft delete operation
    - `example26RestoreSoftDeleted`: Demonstrate restore operation
    - `example27SelectTrashed`: Demonstrate querying trashed users
  - **Query Statistics** (2 examples):
    - `example28QueryStats`: Demonstrate GetQueryStats
    - `example29SlowQueryLog`: Demonstrate slow query logging with threshold
  - **Pool Metrics** (1 example):
    - `example30PoolMetrics`: Demonstrate GetPoolMetrics and utilization
  - **Schema Inspector** (2 examples):
    - `example31GetTables`: Demonstrate GetTables to list all tables
    - `example32InspectTable`: Demonstrate comprehensive table inspection
  - **Migration** (2 examples):
    - `example33CreateTable`: Demonstrate CreateTable with schema
    - `example34AddColumn`: Demonstrate AddColumn, ModifyColumn, and AddIndex
  - **CSV Export** (1 example):
    - `example35ExportCSV`: Demonstrate ExportTableToCSV with options

### Changed / 변경
- **Example Updates** / 예제 업데이트:
  - Updated `runExamples()` function to call all 35 examples (17 basic + 18 advanced)
  - Added helper function `min()` for CSV line truncation
  - All examples include dual-language comments (English/Korean)
  - All examples include comprehensive error handling and logging

### Files Updated / 업데이트된 파일
- `examples/mysql/main.go` (Added 18 new example functions + helper function)
- `cfg/app.yaml` (Version v1.3.012)

---

## [v1.3.010] - 2025-10-14 (MySQL Advanced Features Documentation Update)

### Changed / 변경
- **Documentation Updates** / 문서 업데이트:
  - Updated `database/mysql/README.md` with all 9 new advanced features (batch operations, upsert, pagination, soft delete, query stats, pool metrics, schema inspector, migration helpers, CSV export/import)
  - Updated root `README.md` MySQL section with advanced features list
  - Updated `CLAUDE.md` MySQL architecture section with comprehensive feature list and file structure
  - Added detailed examples for all new features in dual-language format (English/Korean)
  - Added 5 new best practices related to advanced features
  - Version references updated from v1.3.008/009 to v1.3.010

### Files Updated / 업데이트된 파일
- `database/mysql/README.md` (Added 9 advanced feature sections with examples)
- `README.md` (Updated MySQL package feature list)
- `CLAUDE.md` (Updated MySQL architecture with 9 advanced features)
- `cfg/app.yaml` (Version v1.3.010)

---

## [v1.3.009] - 2025-10-14 (MySQL Advanced Features Implementation)

### Added / 추가
- **Batch Operations** / 배치 작업:
  - `batch.go`: Batch operations for improved performance
  - `BatchInsert()`: Insert multiple rows in a single query
  - `BatchUpdate()`: Update multiple rows with different values
  - `BatchDelete()`: Delete multiple rows by IDs
  - `BatchSelectByIDs()`: Select multiple rows by IDs

- **Upsert Operations** / Upsert 작업:
  - `upsert.go`: Insert or update operations
  - `Upsert()`: INSERT ... ON DUPLICATE KEY UPDATE
  - `UpsertBatch()`: Batch upsert for multiple rows
  - `Replace()`: REPLACE INTO (DELETE + INSERT)

- **Pagination** / 페이지네이션:
  - `pagination.go`: Pagination with metadata
  - `Paginate()`: Simple table pagination with page metadata
  - `PaginateQuery()`: Paginate custom queries
  - `PaginationResult` struct: TotalRows, TotalPages, HasNext, HasPrev, Page, PageSize

- **Soft Delete** / 소프트 삭제:
  - `softdelete.go`: Soft delete functionality (requires deleted_at column)
  - `SoftDelete()`: Mark rows as deleted by setting deleted_at
  - `Restore()`: Restore soft-deleted rows
  - `SelectAllWithTrashed()`: Include soft-deleted rows in results
  - `SelectAllOnlyTrashed()`: Select only soft-deleted rows
  - `PermanentDelete()`: Permanently delete rows (hard delete)

- **Query Statistics Tracking** / 쿼리 통계 추적:
  - `stats.go`: Query execution statistics with `QueryStats` struct
  - `GetQueryStats()`: Returns comprehensive statistics (total, success, failed queries, durations)
  - `ResetQueryStats()`: Reset all statistics
  - `EnableSlowQueryLog()`: Configure slow query logging with custom threshold and handler
  - `GetSlowQueries()`: Retrieve recent slow queries
  - `EnableQueryStats()` / `DisableQueryStats()`: Control stats tracking
  - Stats tracking integrated into Query() and Exec() methods

- **Connection Pool Metrics** / 연결 풀 메트릭:
  - `metrics.go`: Comprehensive connection pool monitoring
  - `GetPoolMetrics()`: Returns detailed pool statistics (connections, utilization, wait stats)
  - `GetPoolHealthInfo()`: Health check for all connection pools
  - `GetConnectionUtilization()`: Per-pool utilization percentages
  - `GetWaitStatistics()`: Wait counts and average wait times per pool
  - `GetCurrentConnectionIndex()` / `GetRotationIndex()`: Pool status inspection

- **Schema Inspector** / 스키마 검사:
  - `schema.go`: Database schema inspection utilities
  - `GetTables()`: List all tables with metadata (engine, row count, comment)
  - `GetColumns()`: Retrieve column information (type, nullable, default, key, extra)
  - `GetIndexes()`: Index information (name, columns, unique, type)
  - `TableExists()`: Check table existence
  - `GetTableSchema()`: Get CREATE TABLE statement
  - `GetPrimaryKey()`: Retrieve primary key columns
  - `GetForeignKeys()`: Foreign key relationships
  - `GetTableSize()` / `GetDatabaseSize()`: Size information in bytes
  - `InspectTable()`: Comprehensive table analysis with `TableInspection` struct

- **Migration Helpers** / 마이그레이션 도우미:
  - `migration.go`: Schema migration utilities
  - `CreateTable()` / `CreateTableIfNotExists()`: Table creation with schema
  - `DropTable()`: Table deletion with optional IF EXISTS
  - `TruncateTable()`: Remove all rows from table
  - `AddColumn()` / `DropColumn()` / `ModifyColumn()` / `RenameColumn()`: Column operations
  - `AddIndex()` / `DropIndex()`: Index management (supports unique indexes)
  - `RenameTable()`: Rename tables
  - `AddForeignKey()` / `DropForeignKey()`: Foreign key constraints with ON DELETE/UPDATE
  - `CopyTable()`: Table duplication (structure only or with data)
  - `AlterTableEngine()`: Change storage engine
  - `AlterTableCharset()`: Update character set and collation

- **CSV Export/Import** / CSV 내보내기/가져오기:
  - `export.go`: CSV data exchange utilities
  - `ExportTableToCSV()`: Export table data to CSV file with options (headers, delimiter, where clause, columns)
  - `ImportFromCSV()`: Import CSV data to table with batch insert (supports headers, skip rows, duplicate handling)
  - `ExportQueryToCSV()`: Export custom query results to CSV
  - `CSVExportOptions` / `CSVImportOptions`: Flexible configuration structs
  - Support for NULL value handling, custom delimiters, batch sizes

### Changed / 변경
- **Config** (`config.go`):
  - Added `enableStats` field for query statistics tracking configuration

- **Client** (`client.go`):
  - Added `statsTracker` field to store query statistics
  - Integrated stats recording into Query() and Exec() methods
  - Stats tracking enabled when `WithQueryStats(true)` option is used

- **Options** (`options.go`):
  - Added `WithQueryStats()` option to enable/disable query statistics tracking

- **Bug Fixes** / 버그 수정:
  - Fixed duplicate `IsHealthy()` method (removed from metrics.go, kept in connection.go)
  - Fixed `pagination.go` CountContext call to properly pass variadic arguments

### Files Created / 생성된 파일
- `database/mysql/stats.go` (~280 lines)
- `database/mysql/metrics.go` (~280 lines)
- `database/mysql/schema.go` (~530 lines)
- `database/mysql/migration.go` (~500 lines)
- `database/mysql/export.go` (~420 lines)

### Verification / 확인
- ✅ Build successful: `go build ./database/mysql`
- ✅ All tests passing: `go test ./database/mysql/... -v`
- ✅ All code follows dual-language comment standard (English/Korean)
- ✅ All methods include comprehensive usage examples

---

## [v1.3.008] - 2025-10-10 (Documentation Update)

### Changed / 변경
- **Complete documentation update for non-context API** / Non-context API에 대한 전체 문서 업데이트:
  - Updated `docs/database/mysql/USER_MANUAL.md` with non-context API examples / 모든 예제를 non-context API로 업데이트
  - Updated `docs/database/mysql/DEVELOPER_GUIDE.md` with dual API pattern guidance / 이중 API 패턴 가이드로 업데이트
  - All Simple API examples now use non-context versions by default / 모든 Simple API 예제가 기본적으로 non-context 버전 사용
  - Added explanatory sections about dual API pattern (non-context vs *Context) / 이중 API 패턴 설명 섹션 추가

### Updated Sections / 업데이트된 섹션

**USER_MANUAL.md**:
- Added "API Versions" section explaining dual API pattern / 이중 API 패턴을 설명하는 "API Versions" 섹션 추가
- Updated all method signatures: SelectAll, SelectOne, SelectColumn, SelectColumns, Insert, Update, Delete, Count, Exists
- Converted all examples from `db.Method(ctx, ...)` to `db.Method(...)`
- Updated Quick Start, Usage Patterns, Common Use Cases, and Best Practices sections / 빠른 시작, 사용 패턴, 일반 사용 사례, 모범 사례 섹션 업데이트
- ~80 example code blocks updated across the entire manual / 매뉴얼 전체에 걸쳐 약 80개의 예제 코드 블록 업데이트

**DEVELOPER_GUIDE.md**:
- Added "Dual API Pattern" section in "Adding New Features" / "새 기능 추가"에 "이중 API 패턴" 섹션 추가
- Updated all example implementations to show both non-context and *Context versions / 모든 예제 구현을 non-context 및 *Context 버전으로 업데이트
- Updated query execution flow diagram / 쿼리 실행 흐름 다이어그램 업데이트
- ~20 example code blocks updated / 약 20개의 예제 코드 블록 업데이트

### Pattern / 패턴

All documentation now follows this pattern:
```go
// Non-context (recommended for most cases)
users, _ := db.SelectAll("users")

// Context version (for timeout/cancellation control)
users, _ := db.SelectAllContext(ctx, "users")
```

모든 문서는 이제 이 패턴을 따릅니다:
- Non-context 버전: 대부분의 경우 권장
- Context 버전: 타임아웃/취소 제어가 필요한 경우

### Verification / 확인
- ✅ All old API patterns (`db.Method(ctx, ...)`) removed from documentation
- ✅ Build successful: `go build ./database/mysql` and `go build ./examples/mysql`
- ✅ Documentation consistency verified across USER_MANUAL and DEVELOPER_GUIDE

---

## [v1.3.008] - 2025-10-10 (Code Implementation)

### Added / 추가
- **Non-Context API Methods** - Simplified API without context parameter / Context 매개변수 없는 간소화된 API:
  - All Simple API methods now have non-context versions / 모든 Simple API 메서드에 non-context 버전 추가
  - Methods: `SelectAll`, `SelectColumn`, `SelectColumns`, `SelectOne`, `Insert`, `Update`, `Delete`, `Count`, `Exists`
  - Context versions renamed with `*Context` suffix / Context 버전은 `*Context` 접미사로 renamed
  - Transaction methods also updated / Transaction 메서드도 업데이트

### Breaking Changes / 호환성 변경
- **API Signature Changes** / API 서명 변경:
  - Old: `db.SelectAll(ctx, "users")`
  - New: `db.SelectAll("users")` (non-context) or `db.SelectAllContext(ctx, "users")` (with context)
  - All Simple API methods follow this pattern / 모든 Simple API 메서드가 이 패턴을 따름

### Motivation / 동기
- **Simplified usage for common cases** / 일반적인 경우의 사용 간소화:
  - Most CRUD operations don't need timeout/cancellation control / 대부분의 CRUD 작업은 timeout/cancellation 제어가 필요 없음
  - Non-context versions use `context.Background()` internally / Non-context 버전은 내부적으로 `context.Background()` 사용
  - `*Context` versions available when explicit control needed / 명시적 제어가 필요할 때는 `*Context` 버전 사용 가능

### Examples / 예제

**Before (v1.3.007)**:
```go
ctx := context.Background()
users, err := db.SelectAll(ctx, "users")
user, err := db.SelectOne(ctx, "users", "id = ?", 123)
result, err := db.Insert(ctx, "users", data)
```

**After (v1.3.008)**:
```go
// Simple usage without context / Context 없이 간단한 사용
users, err := db.SelectAll("users")
user, err := db.SelectOne("users", "id = ?", 123)
result, err := db.Insert("users", data)

// With timeout control / Timeout 제어가 필요한 경우
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
users, err := db.SelectAllContext(ctx, "users")
```

### Files Modified / 수정된 파일
```
database/mysql/simple.go         (modified) - Added non-context wrappers for all methods
database/mysql/transaction.go    (modified) - Added non-context wrappers for Tx methods
examples/mysql/main.go            (modified) - Updated examples to use non-context versions
```

### Migration Guide / 마이그레이션 가이드

**Option 1**: Use non-context versions (recommended for most cases)
```go
// Change this:
db.SelectAll(ctx, "users")

// To this:
db.SelectAll("users")
```

**Option 2**: Use *Context versions (for timeout/cancellation control)
```go
// Change this:
db.SelectAll(ctx, "users")

// To this:
db.SelectAllContext(ctx, "users")
```

### Verification / 확인
- ✅ Build successful: `go build ./database/mysql/...`
- ✅ All tests passed: `go test ./database/mysql -v`
- ✅ All 17 examples tested with non-context versions

---

## [v1.3.007] - 2025-10-10

### Added / 추가
- **New Simple API Methods** - SelectColumn and SelectColumns for column-specific queries / 컬럼 특정 쿼리를 위한 SelectColumn과 SelectColumns:
  - `SelectColumn(ctx, table, column, conditionAndArgs...)` - Select single column / 단일 컬럼 선택
  - `SelectColumns(ctx, table, []columns, conditionAndArgs...)` - Select multiple columns / 여러 컬럼 선택
  - Available in both `Client` and `Tx` (Transaction) / Client와 Tx(트랜잭션) 모두에서 사용 가능

### Motivation / 동기
- Query Builder의 `Select()` 메서드와 충돌을 피하기 위해 `SelectColumn`, `SelectColumns`로 명명
- Simple API를 더욱 명확하고 사용하기 쉽게 만들기 위함
- `SELECT * FROM table` (SelectAll), `SELECT column FROM table` (SelectColumn), `SELECT col1, col2 FROM table` (SelectColumns)로 구분

### API Examples / API 예제

**SelectColumn - Single column selection / 단일 컬럼 선택**:
```go
// SELECT email FROM users
emails, _ := db.SelectColumn(ctx, "users", "email")

// SELECT name FROM users WHERE age > 25
names, _ := db.SelectColumn(ctx, "users", "name", "age > ?", 25)

// Process results / 결과 처리
for _, row := range emails {
    fmt.Println(row["email"])
}
```

**SelectColumns - Multiple columns selection / 여러 컬럼 선택**:
```go
// SELECT name, email FROM users
users, _ := db.SelectColumns(ctx, "users", []string{"name", "email"})

// SELECT name, age, city FROM users WHERE age > 25
users, _ := db.SelectColumns(ctx, "users", []string{"name", "age", "city"}, "age > ?", 25)

// Process results / 결과 처리
for _, user := range users {
    fmt.Printf("%s <%s>\n", user["name"], user["email"])
}
```

### Files Modified / 수정된 파일
```
database/mysql/simple.go          (+120 lines) - Added SelectColumn and SelectColumns to Client
database/mysql/transaction.go     (+61 lines)  - Added SelectColumn and SelectColumns to Tx
examples/mysql/main.go             (+118 lines) - Added Example 16 and 17
docs/database/mysql/USER_MANUAL.md      (updated) - Added SelectColumn and SelectColumns documentation
docs/database/mysql/DEVELOPER_GUIDE.md  (updated) - Updated file structure table
```

### Examples / 예제
- **Example 16**: SelectColumn - Single column selection / 단일 컬럼 선택
- **Example 17**: SelectColumns - Multiple columns selection / 여러 컬럼 선택
- Total 17 examples now (1-9: Simple API, 10-12: Query Builder, 13-15: SelectWhere, 16-17: SelectColumn/Columns) / 총 17개 예제

### Verification / 확인
- ✅ Build successful: `go build ./database/mysql/...`
- ✅ All tests passed: `go test ./database/mysql -v`
- ✅ All 17 examples tested

---

## [v1.3.006] - 2025-10-10

### Added / 추가
- **Comprehensive Documentation for MySQL Package** / MySQL 패키지 종합 문서:
  - Created `docs/database/mysql/USER_MANUAL.md` - Complete user manual with all usage patterns, API reference, troubleshooting, FAQ
  - Created `docs/database/mysql/DEVELOPER_GUIDE.md` - Complete developer guide with architecture, internal implementation, design patterns, testing guide

### Documentation Contents / 문서 내용

**USER_MANUAL.md** (~1000+ lines):
- Introduction and Design Philosophy / 소개 및 설계 철학
- Installation and Quick Start / 설치 및 빠른 시작
- Configuration Reference (all options with defaults) / 설정 참조
- Complete API Reference:
  - Simple API (SelectAll, Insert, Update, Delete, Exec) / 간단한 API
  - SelectWhere API (Functional Options) / SelectWhere API (함수형 옵션)
  - Query Builder API (Fluent API) / 쿼리 빌더 API
  - Raw SQL API / Raw SQL API
- Usage Patterns (10 common patterns) / 사용 패턴
- Common Use Cases:
  - User Authentication System / 사용자 인증 시스템
  - E-commerce Order Management / 전자상거래 주문 관리
  - Blog Post System / 블로그 게시물 시스템
  - Analytics Dashboard / 분석 대시보드
  - Notification System / 알림 시스템
- Best Practices (15 recommendations) / 모범 사례
- Troubleshooting (10 common problems) / 문제 해결
- FAQ (15 frequently asked questions) / 자주 묻는 질문

**DEVELOPER_GUIDE.md** (~900+ lines):
- Architecture Overview / 아키텍처 개요
- Package Structure (14 files, ~2,648 lines) / 패키지 구조
- Core Components:
  - Client struct with connection pooling
  - Config with validation
  - Tx (Transaction) with auto rollback
  - QueryBuilder for fluent API
  - Error types classification
- Internal Implementation:
  - Connection Management (multi-pool strategy)
  - Query Execution Flow
  - Auto-Retry with Exponential Backoff
  - Health Check System
  - Credential Rotation Support
- Design Patterns:
  - Functional Options Pattern
  - Builder Pattern
  - Singleton Pattern
  - Strategy Pattern
  - Decorator Pattern
  - Template Method Pattern
- Adding New Features (step-by-step guides)
- Testing Guide (setup, test cases, benchmarking)
- Performance characteristics and optimization
- Contributing Guidelines
- Code Style conventions

### Files Created / 생성된 파일
```
docs/database/mysql/USER_MANUAL.md      (~1000+ lines) - Complete user manual
docs/database/mysql/DEVELOPER_GUIDE.md  (~900+ lines) - Complete developer guide
```

### Verification / 확인
- ✅ Build successful: `go build ./database/mysql/...`
- ✅ All tests passed: `go test ./database/mysql -v`

---

## [v1.3.005] - 2025-10-10

### Added / 추가
- **Query Builder API** - Fluent API for complex queries / 복잡한 쿼리를 위한 Fluent API:
  - Created `database/mysql/builder.go` - Query Builder implementation (~285 lines)
  - Fluent method chaining: `Select().From().Join().Where().OrderBy().Limit().All()`
  - Support for INNER JOIN, LEFT JOIN, RIGHT JOIN / INNER JOIN, LEFT JOIN, RIGHT JOIN 지원
  - Support for GROUP BY, HAVING, ORDER BY, LIMIT, OFFSET / GROUP BY, HAVING, ORDER BY, LIMIT, OFFSET 지원
  - Works with both Client and Transaction / Client와 Transaction 모두에서 사용 가능
  - Query Builder로 복잡한 JOIN, 서브쿼리 패턴 해결

- **SelectWhere API** - Functional options for simple queries / 간단한 쿼리를 위한 함수형 옵션:
  - Created `database/mysql/select_options.go` - SelectWhere with functional options (~360 lines)
  - **One-liner complex queries**: Single function call with multiple options / 한 줄로 복잡한 쿼리 작성
  - Functional options pattern (consistent with `WithDSN()`, etc.) / 함수형 옵션 패턴 (WithDSN() 등과 일관성)
  - Available options / 사용 가능한 옵션:
    - `WithColumns(...string)` - SELECT specific columns / 특정 컬럼 선택
    - `WithOrderBy(string)` - Add ORDER BY / ORDER BY 추가
    - `WithLimit(int)` - Add LIMIT / LIMIT 추가
    - `WithOffset(int)` - Add OFFSET / OFFSET 추가
    - `WithGroupBy(...string)` - Add GROUP BY / GROUP BY 추가
    - `WithHaving(string, ...interface{})` - Add HAVING / HAVING 추가
    - `WithJoin/WithLeftJoin/WithRightJoin(table, condition)` - Add JOINs / JOIN 추가
    - `WithDistinct()` - Add DISTINCT keyword / DISTINCT 키워드 추가
  - Two new methods: `SelectWhere()` and `SelectOneWhere()` / 두 개의 새 메서드

### Technical Details / 기술 세부사항

**New Files / 새 파일**:
```
database/mysql/builder.go         (~285 lines) - Query Builder with fluent API
database/mysql/select_options.go  (~360 lines) - SelectWhere with functional options
examples/mysql/main.go             (+190 lines) - 6 new examples (10-15)
```

**API Examples / API 예제**:

1. **Query Builder (Fluent API)**:
```go
// Complex query with JOIN / JOIN을 사용한 복잡한 쿼리
users, _ := db.Select("u.name", "o.total").
    From("users u").
    Join("orders o", "u.id = o.user_id").
    Where("o.status = ?", "completed").
    OrderBy("o.total DESC").
    Limit(10).
    All(ctx)

// GROUP BY with HAVING
results, _ := db.Select("city", "COUNT(*) as count").
    From("users").
    GroupBy("city").
    Having("COUNT(*) > ?", 10).
    OrderBy("count DESC").
    All(ctx)
```

2. **SelectWhere (Functional Options)**:
```go
// One-liner with multiple options / 여러 옵션을 사용한 한 줄 쿼리
users, _ := db.SelectWhere(ctx, "users", "age > ?", 25,
    mysql.WithColumns("name", "email", "age"),
    mysql.WithOrderBy("age DESC"),
    mysql.WithLimit(3))

// GROUP BY in one line / 한 줄로 GROUP BY
results, _ := db.SelectWhere(ctx, "users", "",
    mysql.WithColumns("city", "COUNT(*) as count"),
    mysql.WithGroupBy("city"),
    mysql.WithHaving("COUNT(*) > ?", 2),
    mysql.WithOrderBy("count DESC"))

// DISTINCT query / DISTINCT 쿼리
cities, _ := db.SelectWhere(ctx, "users", "age > ?", 25,
    mysql.WithColumns("city"),
    mysql.WithDistinct(),
    mysql.WithOrderBy("city ASC"))
```

**Example Summary / 예제 요약**:
- Total 15 examples / 총 15개 예제:
  - Examples 1-9: Simple API (SelectAll, Insert, etc.) / 간단한 API
  - Examples 10-12: Query Builder (Fluent API) / 쿼리 빌더
  - Examples 13-15: SelectWhere (Functional Options) / SelectWhere 함수형 옵션
- All examples tested successfully / 모든 예제 테스트 성공

### Why Both Query Builder and SelectWhere? / 왜 둘 다 제공하는가?

**Query Builder (Fluent API)**:
- ✅ For complex queries with multiple JOINs / 여러 JOIN이 있는 복잡한 쿼리용
- ✅ IDE autocomplete support / IDE 자동완성 지원
- ✅ Step-by-step query building / 단계별 쿼리 빌드

**SelectWhere (Functional Options)**:
- ✅ For simple to moderate queries / 간단~중간 복잡도 쿼리용
- ✅ **One-liner**: Entire query in single function call / 한 줄로 전체 쿼리 작성
- ✅ Consistent with package option pattern / 패키지 옵션 패턴과 일관성
- ✅ Closer to "30 lines → 2 lines" goal / "30줄 → 2줄" 목표에 더 가까움

### Notes / 참고사항
- Both APIs support transactions / 두 API 모두 트랜잭션 지원
- Both APIs have auto-retry, auto-reconnect / 두 API 모두 자동 재시도, 자동 재연결
- Users can choose based on preference / 사용자가 선호에 따라 선택 가능
- Query Builder for complex queries / 복잡한 쿼리는 Query Builder
- SelectWhere for simple one-liners / 간단한 쿼리는 SelectWhere

---

## [v1.3.004] - 2025-10-10

### Added / 추가
- **Configuration Management / 설정 관리**:
  - Created `cfg/database.yaml` - Database configuration file with MySQL settings
  - Supports host, port, user, password, database, connection pool settings
  - 데이터베이스 설정 파일 생성 (MySQL 설정 포함)
  - 호스트, 포트, 사용자, 패스워드, 데이터베이스, 연결 풀 설정 지원

- **MySQL Examples / MySQL 예제**:
  - Created `examples/mysql/main.go` - Comprehensive example demonstrating all package features
  - Created `examples/mysql/README.md` - Example documentation
  - 모든 패키지 기능을 시연하는 종합 예제 생성
  - 예제 문서 생성

- **Example Features / 예제 기능**:
  - YAML configuration loading with `gopkg.in/yaml.v3` / YAML 설정 로딩
  - Integrated logging package for structured logs / 구조화된 로그를 위한 로깅 패키지 통합
  - Auto MySQL daemon management (start/stop) / MySQL 데몬 자동 관리 (시작/중지)
  - 9 working examples: SelectAll, SelectOne, Insert, Update, Count, Exists, Transaction, Delete, Raw SQL
  - 9개 작동 예제: SelectAll, SelectOne, Insert, Update, Count, Exists, Transaction, Delete, Raw SQL

- **MySQL Setup / MySQL 설정**:
  - Set MySQL root password to `test1234` / MySQL root 패스워드 설정
  - Created `testdb` database with sample `users` table / 샘플 users 테이블이 있는 testdb 데이터베이스 생성
  - Populated with 5 initial sample records / 5개 초기 샘플 레코드 삽입

### Technical Details / 기술 세부사항

**New Files / 새 파일**:
```
cfg/database.yaml         - Database configuration
examples/mysql/main.go    (~470 lines) - Complete examples with logging
examples/mysql/README.md  - Example documentation
```

**Configuration Structure / 설정 구조**:
```yaml
mysql:
  host: localhost
  port: 3306
  user: root
  password: "test1234"
  database: testdb
  max_open_conns: 25
  max_idle_conns: 10
  conn_max_lifetime: 300
  params:
    parseTime: true
    charset: utf8mb4
    loc: Local
```

**Example Highlights / 예제 주요 사항**:
- Database configuration loaded from YAML file / YAML 파일에서 데이터베이스 설정 로드
- **Full logging integration**: Replaced all `fmt.Print` statements with logging package / 완전한 로깅 통합: 모든 `fmt.Print` 문을 로깅 패키지로 교체
- All output logged to both console and file / 모든 출력이 콘솔과 파일 모두에 로깅됨
- Auto MySQL start/stop if not running / 실행 중이 아니면 MySQL 자동 시작/중지
- **Repeatable examples**: Uses timestamp-based unique emails / 반복 실행 가능한 예제: 타임스탬프 기반 고유 이메일 사용
- No hardcoded credentials in helper functions / 헬퍼 함수에 하드코딩된 자격 증명 없음
- All 9 examples executed successfully / 모든 9개 예제가 성공적으로 실행됨

### Fixed / 수정
- **Duplicate key errors**: Insert examples now generate timestamp-based unique emails / 중복 키 에러 수정: Insert 예제가 이제 타임스탬프 기반 고유 이메일 생성
- **Delete example**: Changed to delete existing sample user (charlie@example.com) instead of non-existent user / Delete 예제를 존재하지 않는 사용자 대신 기존 샘플 사용자 삭제로 변경
- **Hardcoded credentials**: Removed from `isMySQLRunning()` function / isMySQLRunning() 함수에서 하드코딩된 자격 증명 제거
- Error message capitalization to follow Go conventions (ST1005) / Go 규칙에 따라 에러 메시지 대문자 수정

### Dependencies / 의존성
- Added `gopkg.in/yaml.v3` for YAML configuration parsing / YAML 설정 파싱용

### Notes / 참고사항
- Examples demonstrate "30 lines → 2 lines" simplicity goal / 예제가 "30줄 → 2줄" 간결함 목표를 시연
- All examples include bilingual output (English/Korean) / 모든 예제가 이중 언어 출력 포함 (영문/한글)
- Examples tested on macOS with Homebrew MySQL 9.4.0 / macOS Homebrew MySQL 9.4.0에서 테스트됨
- Examples are fully repeatable with no duplicate key errors / 예제는 중복 키 에러 없이 완전히 반복 실행 가능

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

## [v1.3.002] - 2025-10-10

### Added / 추가
- **Core Implementation / 핵심 구현**:
  - Implemented Phase 1 (Foundation): errors.go, types.go, config.go
  - Implemented Phase 2 (Core Features): options.go, client.go, connection.go, rotation.go
  - 7 core files with bilingual comments
  - Phase 1 (기초) 구현: 에러 타입, 공통 타입, 설정 구조체
  - Phase 2 (핵심 기능) 구현: 옵션, 클라이언트, 연결 관리, 순환 로직

- **Features Implemented / 구현된 기능**:
  - Client struct with multiple connection pools / 다중 연결 풀을 갖춘 클라이언트 구조체
  - Functional options pattern (20+ options) / 함수형 옵션 패턴 (20개 이상 옵션)
  - Auto health check goroutine / 자동 헬스 체크 goroutine
  - Credential rotation goroutine / 자격 증명 순환 goroutine
  - Round-robin connection selection / Round-robin 연결 선택
  - Configuration validation / 설정 검증
  - Comprehensive error types / 포괄적인 에러 타입

- **Testing / 테스팅**:
  - Basic unit tests for config, options, and client creation
  - All tests passing (100% pass rate)
  - 설정, 옵션, 클라이언트 생성에 대한 기본 유닛 테스트
  - 모든 테스트 통과 (100% 통과율)

### Technical Details / 기술 세부사항

**Files Created / 생성된 파일**:
```
database/mysql/
├── errors.go        (130 lines) - Error types and classification
├── types.go         (73 lines) - Common types (CredentialRefreshFunc, Tx)
├── config.go        (130 lines) - Configuration structure and validation
├── options.go       (230 lines) - 20+ functional options
├── client.go        (260 lines) - Main client with connection management
├── connection.go    (75 lines) - Health check and connection monitoring
├── rotation.go      (85 lines) - Credential rotation logic
└── client_test.go   (120 lines) - Unit tests
```

**Dependencies / 의존성**:
- Added `github.com/go-sql-driver/mysql v1.9.3`
- Added `filippo.io/edwards25519 v1.1.0` (MySQL driver dependency)

### Changed / 변경
- go.mod: Added MySQL driver dependency / MySQL 드라이버 의존성 추가

### Notes / 참고사항
- Compilation successful / 컴파일 성공
- All basic tests passing / 모든 기본 테스트 통과
- Remaining work: simple.go, builder.go, transaction.go, retry.go, scan.go, README
- 남은 작업: simple.go, builder.go, transaction.go, retry.go, scan.go, README

---

## [v1.3.003] - 2025-10-10

### Added / 추가
- **Complete Implementation / 완전한 구현**:
  - Implemented Phase 3+: simple.go, transaction.go, retry.go, scan.go
  - All API methods fully functional
  - Comprehensive README.md with examples
  - Phase 3+ 구현: 간단한 API, 트랜잭션, 재시도, 스캔
  - 모든 API 메서드 완전 작동
  - 예제가 포함된 종합 README.md

- **Simple API (simple.go)** - Core feature! / 핵심 기능!:
  - SelectAll() - Select all rows with optional conditions
  - SelectOne() - Select single row
  - Insert() - Insert with map
  - Update() - Update with map
  - Delete() - Delete with conditions
  - Count() - Count rows
  - Exists() - Check existence
  - 30 lines → 2 lines code reduction achieved! / 30줄 → 2줄 코드 감소 달성!

- **Transaction API (transaction.go)**:
  - Transaction() - Execute function within transaction
  - Auto commit/rollback
  - All simple.go methods available within Tx
  - Panic recovery with automatic rollback
  - 트랜잭션 함수 실행
  - 자동 커밋/롤백
  - Tx 내에서 모든 simple.go 메서드 사용 가능

- **Auto Retry (retry.go)**:
  - executeWithRetry() with exponential backoff
  - Automatic retry for transient errors
  - Context-aware cancellation
  - 지수 백오프로 재시도
  - 일시적 에러 자동 재시도

- **Result Scanning (scan.go)**:
  - scanRows() - Scan multiple rows to []map
  - scanRow() - Scan single row to map
  - scanCount() - Scan COUNT(*) result
  - Automatic type conversion ([]byte → string)
  - 자동 타입 변환

### Features Completed / 완료된 기능

✅ **30 lines → 2 lines**: Extreme simplicity achieved / 극도의 간결함 달성
✅ **No defer rows.Close()**: Automatic resource cleanup / 자동 리소스 정리
✅ **SQL-like API**: Close to actual SQL syntax / SQL 문법에 가까운 API
✅ **Auto retry**: Transient errors handled automatically / 일시적 에러 자동 처리
✅ **Transaction support**: Auto commit/rollback / 자동 커밋/롤백
✅ **Type conversion**: []byte → string, etc. / 타입 변환

### Technical Details / 기술 세부사항

**New Files / 새 파일**:
```
database/mysql/
├── retry.go         (120 lines) - Auto retry with exponential backoff
├── scan.go          (180 lines) - Result scanning and type conversion
├── simple.go        (370 lines) - Simple API (7 methods)
├── transaction.go   (230 lines) - Transaction helpers
└── README.md        (380 lines) - Comprehensive documentation
```

**Total Package Size / 전체 패키지 크기**:
- 13 files / 13개 파일
- ~2,300 lines of code / ~2,300줄 코드
- 100% bilingual comments / 100% 이중 언어 주석
- Compilation successful / 컴파일 성공

### Example Usage / 사용 예시

**Before (standard database/sql) / 이전 (표준 database/sql)**:
```go
// ❌ 30+ lines
db, _ := sql.Open("mysql", dsn)
defer db.Close()
rows, _ := db.Query("SELECT * FROM users WHERE age > ?", 18)
defer rows.Close()
// ... 20+ more lines for scanning / 스캔을 위한 20줄 이상
```

**After (this package) / 이후 (이 패키지)**:
```go
// ✅ 2 lines
db, _ := mysql.New(mysql.WithDSN(dsn))
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)
```

### Notes / 참고사항
- **Goal achieved**: "If not 10x simpler, don't build it" → We did it! / 목표 달성: "10배 간단하지 않으면 만들지 마세요" → 달성!
- Package is production-ready / 패키지는 프로덕션 준비 완료
- All core features implemented / 모든 핵심 기능 구현됨
- Query builder (builder.go) can be added later / 쿼리 빌더는 나중에 추가 가능

---

**Version History / 버전 히스토리**:
- v1.3.003: Complete implementation (simple API, transaction, retry, scan, README) / 완전한 구현
- v1.3.002: Core implementation (Phase 1 & 2) / 핵심 구현 (Phase 1 & 2)
- v1.3.001: Design documents for database/mysql package / database/mysql 패키지 설계 문서
