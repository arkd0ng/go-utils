# MySQL Package Examples

This directory contains comprehensive examples demonstrating all features of the go-utils MySQL package.

이 디렉토리는 go-utils MySQL 패키지의 모든 기능을 시연하는 포괄적인 예제를 포함합니다.

## Running Examples / 예제 실행

```bash
cd examples/mysql
go run .
```

The examples will automatically:
- Start Docker MySQL if not running
- Run all 35 examples
- Stop Docker MySQL (if it was started by the examples)

예제는 자동으로:
- MySQL이 실행 중이 아니면 Docker MySQL 시작
- 35개 예제 모두 실행
- Docker MySQL 중지 (예제에서 시작한 경우)

## Output Directory Structure / 출력 디렉토리 구조

All example outputs are organized in the `results/` directory:

모든 예제 출력은 `results/` 디렉토리에 정리됩니다:

```
examples/mysql/
├── main.go
├── README.md
└── results/                    # Created automatically / 자동 생성
    ├── logs/                   # Log files / 로그 파일
    │   └── mysql_example_20251014_120000.log
    └── mysql_export/           # CSV export files / CSV 내보내기 파일
        └── users_export_20251014_120000.csv
```

**Benefits / 장점:**
- Clean directory structure / 깨끗한 디렉토리 구조
- All outputs in one place / 모든 출력을 한 곳에
- Easy to clean up / 쉬운 정리
- Files named with timestamps / 타임스탬프로 파일 이름 지정

## Examples Covered / 포함된 예제

### Basic Operations (Examples 1-9) / 기본 작업 (예제 1-9)
1. SelectAll - Select all records
2. SelectOne - Select single record
3. Insert - Insert new record
4. Update - Update record
5. Count - Count records
6. Exists - Check if record exists
7. Transaction - Use transactions
8. Delete - Delete record
9. RawSQL - Execute raw SQL queries

### Query Builder (Examples 10-12) / 쿼리 빌더 (예제 10-12)
10. Simple SELECT with WHERE, ORDER BY, LIMIT
11. GROUP BY with COUNT
12. Complex query with multiple conditions

### SelectWhere API (Examples 13-15) / SelectWhere API (예제 13-15)
13. Simple query with options
14. GROUP BY with options
15. Multiple conditions and options

### Column Selection (Examples 16-17) / 컬럼 선택 (예제 16-17)
16. SelectColumn - Select single column
17. SelectColumns - Select multiple columns

### Batch Operations (Examples 18-20) / 배치 작업 (예제 18-20)
18. BatchInsert - Insert multiple rows in single query
19. BatchUpdate - Update multiple rows in transaction
20. BatchDelete - Delete multiple rows by IDs

### Upsert Operations (Examples 21-22) / Upsert 작업 (예제 21-22)
21. Upsert - Insert or update on duplicate
22. UpsertBatch - Batch upsert operations

### Pagination (Examples 23-24) / 페이지네이션 (예제 23-24)
23. Basic pagination with metadata
24. Pagination with WHERE and ORDER BY

### Soft Delete (Examples 25-27) / 소프트 삭제 (예제 25-27)
24.5. Prepare table for soft delete (auto migration)
25. SoftDelete - Mark record as deleted
26. Restore - Restore soft-deleted record
27. SelectTrashed - Query trashed and all records

### Query Statistics (Examples 28-29) / 쿼리 통계 (예제 28-29)
28. QueryStats - Query execution statistics
29. SlowQueryLog - Slow query detection

### Pool Metrics (Example 30) / 풀 메트릭 (예제 30)
30. PoolMetrics - Connection pool metrics

### Schema Inspector (Examples 31-32) / 스키마 검사기 (예제 31-32)
31. GetTables - List all tables
32. InspectTable - Comprehensive table inspection

### Migration (Examples 33-34) / 마이그레이션 (예제 33-34)
33. CreateTable - Create new table
34. Migration operations - Add/modify/drop columns

### CSV Export (Example 35) / CSV 내보내기 (예제 35)
35. ExportCSV - Export table to CSV file

## Prerequisites / 전제 조건

- Docker Desktop installed and running
- Go 1.24.6 or higher

Docker Desktop이 설치되어 실행 중이어야 합니다.

## Cleaning Up / 정리

To remove all example outputs:

모든 예제 출력을 제거하려면:

```bash
rm -rf results/
```

The `results/` directory will be automatically recreated when you run the examples again.

`results/` 디렉토리는 예제를 다시 실행할 때 자동으로 재생성됩니다.
