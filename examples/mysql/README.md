# MySQL Package Examples / MySQL 패키지 예제

This example demonstrates all features of the `database/mysql` package.

이 예제는 `database/mysql` 패키지의 모든 기능을 시연합니다.

## Features / 기능

This example includes:

이 예제에는 다음이 포함됩니다:

1. **Auto MySQL Management / 자동 MySQL 관리**: Automatically starts/stops MySQL daemon if needed
2. **SelectAll**: Select multiple rows with conditions / 조건과 함께 여러 행 선택
3. **SelectOne**: Select single row / 단일 행 선택
4. **Insert**: Insert new records / 새 레코드 삽입
5. **Update**: Update existing records / 기존 레코드 업데이트
6. **Count**: Count rows with conditions / 조건과 함께 행 개수 계산
7. **Exists**: Check if record exists / 레코드 존재 확인
8. **Transaction**: Multi-operation transaction with auto commit/rollback / 자동 커밋/롤백이 있는 다중 작업 트랜잭션
9. **Delete**: Delete records / 레코드 삭제
10. **Raw SQL**: Execute raw SQL queries / Raw SQL 쿼리 실행

## Prerequisites / 전제 조건

1. **Homebrew** must be installed / Homebrew가 설치되어 있어야 합니다
2. **MySQL** must be installed via Homebrew / MySQL이 Homebrew를 통해 설치되어 있어야 합니다:
   ```bash
   brew install mysql
   ```

3. **Test database** must be set up / 테스트 데이터베이스가 설정되어 있어야 합니다:
   ```bash
   # Start MySQL
   brew services start mysql

   # Create database and sample data
   mysql -u root <<'EOF'
   CREATE DATABASE IF NOT EXISTS testdb;
   USE testdb;

   CREATE TABLE IF NOT EXISTS users (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       email VARCHAR(100) NOT NULL UNIQUE,
       age INT NOT NULL,
       city VARCHAR(100),
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );

   INSERT INTO users (name, email, age, city) VALUES
       ('John Doe', 'john@example.com', 30, 'Seoul'),
       ('Jane Smith', 'jane@example.com', 25, 'Busan'),
       ('Bob Johnson', 'bob@example.com', 35, 'Seoul'),
       ('Alice Williams', 'alice@example.com', 28, 'Incheon'),
       ('Charlie Brown', 'charlie@example.com', 22, 'Seoul')
   ON DUPLICATE KEY UPDATE name=name;
   EOF
   ```

## Running the Example / 예제 실행

### Option 1: Run directly / 직접 실행
```bash
cd examples/mysql
go run main.go
```

### Option 2: Build and run / 빌드 후 실행
```bash
cd examples/mysql
go build -o mysql-example
./mysql-example
```

## How It Works / 작동 방식

The example program:

예제 프로그램은:

1. **Checks MySQL status** / **MySQL 상태 확인**: Detects if MySQL is already running
2. **Starts MySQL if needed** / **필요시 MySQL 시작**: Automatically starts the daemon if not running
3. **Runs all examples** / **모든 예제 실행**: Demonstrates all package features
4. **Stops MySQL on exit** / **종료 시 MySQL 중지**: Uses `defer` to stop the daemon if it was started by this program

**Important**: If MySQL was already running before starting this program, it will remain running after the program exits.

**중요**: 이 프로그램을 시작하기 전에 MySQL이 이미 실행 중이었다면, 프로그램 종료 후에도 계속 실행됩니다.

## Expected Output / 예상 출력

```
✅ MySQL is already running
✅ MySQL이 이미 실행 중입니다

======================================================================
MySQL Package Examples - go-utils/database/mysql
======================================================================

📋 Example 1: SelectAll - Select all users
📋 예제 1: SelectAll - 모든 사용자 선택

Found 3 users from Seoul:
서울에서 3명의 사용자를 찾았습니다:
  1. John Doe (age: 30, email: john@example.com)
  2. Bob Johnson (age: 35, email: bob@example.com)
  3. Charlie Brown (age: 22, email: charlie@example.com)

👤 Example 2: SelectOne - Select single user
👤 예제 2: SelectOne - 단일 사용자 선택

Found user: John Doe
  - Email: john@example.com
  - Age: 30
  - City: Seoul

[... more examples ...]

======================================================================
✅ All examples completed successfully!
✅ 모든 예제가 성공적으로 완료되었습니다!
======================================================================
```

## Key Features Demonstrated / 시연된 주요 기능

### 1. Extreme Simplicity / 극도의 간결함
```go
// Just 2 lines! / 단 2줄!
db, _ := mysql.New(mysql.WithDSN(dsn))
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)
```

### 2. No defer rows.Close() / defer rows.Close() 불필요
```go
// No need to manually close rows
// 수동으로 rows를 닫을 필요 없음
users, _ := db.SelectAll(ctx, "users")
// Automatic cleanup handled internally
// 내부적으로 자동 정리 처리
```

### 3. Transaction Support / 트랜잭션 지원
```go
db.Transaction(ctx, func(tx *mysql.Tx) error {
    tx.Insert(ctx, "users", map[string]interface{}{"name": "Emily"})
    tx.Insert(ctx, "users", map[string]interface{}{"name": "Frank"})
    return nil // Auto commit / 자동 커밋
})
```

### 4. Auto Daemon Management / 자동 데몬 관리
```go
// Automatically starts MySQL if not running
// 실행 중이 아니면 자동으로 MySQL 시작
wasRunning := isMySQLRunning()
if !wasRunning {
    startMySQL()
    defer stopMySQL() // Stop on exit / 종료 시 중지
}
```

## Troubleshooting / 문제 해결

### MySQL won't start / MySQL이 시작되지 않음
```bash
# Check MySQL status
brew services list | grep mysql

# View MySQL error log
tail -f /opt/homebrew/var/mysql/*.err

# Restart MySQL
brew services restart mysql
```

### Connection refused / 연결 거부
```bash
# Ensure MySQL is running
mysql -u root -e "SELECT VERSION();"

# Check if socket file exists
ls -la /tmp/mysql.sock
```

### Permission denied / 권한 거부
```bash
# Ensure you have permission to start/stop services
brew services list

# Try running with sudo (not recommended)
sudo brew services start mysql
```

## Clean Up / 정리

To remove the test database:

테스트 데이터베이스를 제거하려면:

```bash
mysql -u root -e "DROP DATABASE testdb;"
```

To stop MySQL:

MySQL을 중지하려면:

```bash
brew services stop mysql
```

## License / 라이선스

MIT
