package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/arkd0ng/go-utils/database/mysql"
	"github.com/arkd0ng/go-utils/logging"
	"gopkg.in/yaml.v3"
)

// DatabaseConfig represents database configuration from YAML
// DatabaseConfig는 YAML에서 읽은 데이터베이스 설정을 나타냅니다
type DatabaseConfig struct {
	MySQL MySQLConfig `yaml:"mysql"`
}

// MySQLConfig represents MySQL connection settings
// MySQLConfig는 MySQL 연결 설정을 나타냅니다
type MySQLConfig struct {
	Host            string            `yaml:"host"`
	Port            int               `yaml:"port"`
	User            string            `yaml:"user"`
	Password        string            `yaml:"password"`
	Database        string            `yaml:"database"`
	MaxOpenConns    int               `yaml:"max_open_conns"`
	MaxIdleConns    int               `yaml:"max_idle_conns"`
	ConnMaxLifetime int               `yaml:"conn_max_lifetime"`
	Params          map[string]string `yaml:"params"`
}

func main() {
	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(fmt.Sprintf("./mysql_example_%s.log", time.Now().Format("20060102_150405"))),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Banner("MySQL Package Examples", "go-utils/database/mysql")

	// Load database configuration / 데이터베이스 설정 로드
	logger.Info("Loading database configuration from cfg/database.yaml")
	logger.Info("cfg/database.yaml에서 데이터베이스 설정 로드 중")
	config, err := loadDatabaseConfig()
	if err != nil {
		logger.Error("Failed to load database config", "error", err)
		os.Exit(1)
	}

	// Build DSN from config / 설정에서 DSN 빌드
	dsn := buildDSN(config.MySQL)
	logger.Info("Configuration loaded successfully",
		"host", config.MySQL.Host,
		"port", config.MySQL.Port,
		"database", config.MySQL.Database,
		"user", config.MySQL.User)

	// Check if MySQL is running / MySQL 실행 여부 확인
	wasRunning := isMySQLRunning(config.MySQL)

	if !wasRunning {
		logger.Info("MySQL is not running, starting daemon...")
		logger.Info("MySQL이 실행 중이 아닙니다. 데몬을 시작합니다...")
		if err := startMySQL(); err != nil {
			logger.Error("Failed to start MySQL", "error", err)
			os.Exit(1)
		}

		// Wait for MySQL to be ready / MySQL 준비 대기
		logger.Info("Waiting for MySQL to be ready...")
		logger.Info("MySQL 준비 중...")
		time.Sleep(3 * time.Second)

		// Ensure MySQL stops when program exits / 프로그램 종료 시 MySQL 중지 보장
		defer func() {
			logger.Info("Stopping MySQL daemon...")
			logger.Info("MySQL 데몬 중지 중...")
			if err := stopMySQL(); err != nil {
				logger.Warn("Failed to stop MySQL", "error", err)
			} else {
				logger.Info("MySQL daemon stopped successfully")
				logger.Info("MySQL 데몬이 성공적으로 중지되었습니다")
			}
		}()
	} else {
		logger.Info("MySQL is already running")
		logger.Info("MySQL이 이미 실행 중입니다")
	}

	logger.Info("")
	logger.Info(strings.Repeat("=", 70))
	logger.Info("Running Examples / 예제 실행 중")
	logger.Info(strings.Repeat("=", 70))
	logger.Info("")

	// Run all examples / 모든 예제 실행
	if err := runExamples(dsn, config.MySQL, logger); err != nil {
		logger.Error("Examples failed", "error", err)
		os.Exit(1)
	}

	logger.Info("")
	logger.Info(strings.Repeat("=", 70))
	logger.Info("All examples completed successfully!")
	logger.Info("모든 예제가 성공적으로 완료되었습니다!")
	logger.Info(strings.Repeat("=", 70))
	logger.Info("")
}

// loadDatabaseConfig loads database configuration from YAML file
// loadDatabaseConfig는 YAML 파일에서 데이터베이스 설정을 로드합니다
func loadDatabaseConfig() (*DatabaseConfig, error) {
	// Get project root directory / 프로젝트 루트 디렉토리 가져오기
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	// Navigate to project root (examples/mysql -> go-utils)
	// 프로젝트 루트로 이동
	projectRoot := filepath.Join(wd, "..", "..")
	configPath := filepath.Join(projectRoot, "cfg", "database.yaml")

	// Read YAML file / YAML 파일 읽기
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML / YAML 파싱
	var config DatabaseConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// buildDSN builds MySQL DSN string from configuration
// buildDSN은 설정에서 MySQL DSN 문자열을 빌드합니다
func buildDSN(cfg MySQLConfig) string {
	// Format: user:password@tcp(host:port)/database?params
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	// Add parameters / 파라미터 추가
	if len(cfg.Params) > 0 {
		params := []string{}
		for key, value := range cfg.Params {
			params = append(params, fmt.Sprintf("%s=%s", key, value))
		}
		dsn += "?" + strings.Join(params, "&")
	}

	return dsn
}

// isMySQLRunning checks if MySQL service is running
// isMySQLRunning은 MySQL 서비스가 실행 중인지 확인합니다
func isMySQLRunning(cfg MySQLConfig) bool {
	// Try to connect to MySQL using configured credentials
	// 설정된 자격 증명을 사용하여 MySQL 연결 시도
	cmd := exec.Command("mysql",
		"-h", cfg.Host,
		"-P", fmt.Sprintf("%d", cfg.Port),
		"-u", cfg.User,
		fmt.Sprintf("-p%s", cfg.Password),
		"-e", "SELECT 1")
	err := cmd.Run()
	return err == nil
}

// startMySQL starts the MySQL service
// startMySQL은 MySQL 서비스를 시작합니다
func startMySQL() error {
	cmd := exec.Command("brew", "services", "start", "mysql")
	return cmd.Run()
}

// stopMySQL stops the MySQL service
// stopMySQL은 MySQL 서비스를 중지합니다
func stopMySQL() error {
	cmd := exec.Command("brew", "services", "stop", "mysql")
	return cmd.Run()
}

// runExamples runs all MySQL package examples
// runExamples는 모든 MySQL 패키지 예제를 실행합니다
func runExamples(dsn string, cfg MySQLConfig, logger *logging.Logger) error {
	logger.Info("Creating MySQL client with connection pool settings...")
	logger.Info("연결 풀 설정과 함께 MySQL 클라이언트 생성 중...")

	// Create MySQL client with connection pool settings
	// 연결 풀 설정과 함께 MySQL 클라이언트 생성
	db, err := mysql.New(
		mysql.WithDSN(dsn),
		mysql.WithMaxOpenConns(cfg.MaxOpenConns),
		mysql.WithMaxIdleConns(cfg.MaxIdleConns),
		mysql.WithConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime)*time.Second),
	)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	defer db.Close()

	logger.Info("MySQL client created successfully",
		"max_open_conns", cfg.MaxOpenConns,
		"max_idle_conns", cfg.MaxIdleConns,
		"conn_max_lifetime", cfg.ConnMaxLifetime)

	ctx := context.Background()

	// Example 1: SelectAll - Select all users / 모든 사용자 선택
	if err := example1SelectAll(ctx, db, logger); err != nil {
		return err
	}

	// Example 2: SelectOne - Select single user / 단일 사용자 선택
	if err := example2SelectOne(ctx, db, logger); err != nil {
		return err
	}

	// Example 3: Insert - Insert new user / 새 사용자 삽입
	if err := example3Insert(ctx, db, logger); err != nil {
		return err
	}

	// Example 4: Update - Update user / 사용자 업데이트
	if err := example4Update(ctx, db, logger); err != nil {
		return err
	}

	// Example 5: Count - Count users / 사용자 수 계산
	if err := example5Count(ctx, db, logger); err != nil {
		return err
	}

	// Example 6: Exists - Check if user exists / 사용자 존재 확인
	if err := example6Exists(ctx, db, logger); err != nil {
		return err
	}

	// Example 7: Transaction - Insert with transaction / 트랜잭션으로 삽입
	if err := example7Transaction(ctx, db, logger); err != nil {
		return err
	}

	// Example 8: Delete - Delete user / 사용자 삭제
	if err := example8Delete(ctx, db, logger); err != nil {
		return err
	}

	// Example 9: Raw SQL - Use raw SQL queries / Raw SQL 쿼리 사용
	if err := example9RawSQL(ctx, db, logger); err != nil {
		return err
	}

	// Example 10: Query Builder - Simple SELECT with WHERE, ORDER BY, LIMIT
	if err := example10QueryBuilderSimple(ctx, db, logger); err != nil {
		return err
	}

	// Example 11: Query Builder - GROUP BY with COUNT
	if err := example11QueryBuilderGroupBy(ctx, db, logger); err != nil {
		return err
	}

	// Example 12: Query Builder - Complex query with multiple conditions
	if err := example12QueryBuilderComplex(ctx, db, logger); err != nil {
		return err
	}

	// Example 13: SelectWhere - Simple query with options
	if err := example13SelectWhereSimple(ctx, db, logger); err != nil {
		return err
	}

	// Example 14: SelectWhere - GROUP BY with options
	if err := example14SelectWhereGroupBy(ctx, db, logger); err != nil {
		return err
	}

	// Example 15: SelectWhere - Complex with multiple options
	if err := example15SelectWhereComplex(ctx, db, logger); err != nil {
		return err
	}

	return nil
}

// example1SelectAll demonstrates SelectAll method
// example1SelectAll은 SelectAll 메서드를 시연합니다
func example1SelectAll(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("📋 Example 1: SelectAll - Select all users")
	logger.Info("📋 예제 1: SelectAll - 모든 사용자 선택")
	logger.Info("")

	// Select all users from Seoul / 서울의 모든 사용자 선택
	users, err := db.SelectAll(ctx, "users", "city = ?", "Seoul")
	if err != nil {
		return fmt.Errorf("selectAll failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d users from Seoul:", len(users)))
	logger.Info(fmt.Sprintf("서울에서 %d명의 사용자를 찾았습니다:", len(users)))
	for i, user := range users {
		logger.Info(fmt.Sprintf("  %d. %s (age: %v, email: %s)",
			i+1, user["name"], user["age"], user["email"]))
	}
	logger.Info("")
	return nil
}

// example2SelectOne demonstrates SelectOne method
// example2SelectOne은 SelectOne 메서드를 시연합니다
func example2SelectOne(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("👤 Example 2: SelectOne - Select single user")
	logger.Info("👤 예제 2: SelectOne - 단일 사용자 선택")
	logger.Info("")

	user, err := db.SelectOne(ctx, "users", "email = ?", "john@example.com")
	if err != nil {
		return fmt.Errorf("selectOne failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Found user: %s", user["name"]))
	logger.Info(fmt.Sprintf("  - Email: %s", user["email"]))
	logger.Info(fmt.Sprintf("  - Age: %v", user["age"]))
	logger.Info(fmt.Sprintf("  - City: %s", user["city"]))
	logger.Info("")
	return nil
}

// example3Insert demonstrates Insert method
// example3Insert는 Insert 메서드를 시연합니다
func example3Insert(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("➕ Example 3: Insert - Insert new user")
	logger.Info("➕예제 3: Insert - 새 사용자 삽입")
	logger.Info("")

	// Generate unique email with timestamp / 타임스탬프로 유니크한 이메일 생성
	timestamp := time.Now().Unix()
	email := fmt.Sprintf("david.kim.%d@example.com", timestamp)

	result, err := db.Insert(ctx, "users", map[string]any{
		"name":  "David Kim",
		"email": email,
		"age":   32,
		"city":  "Daegu",
	})
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}

	id, _ := result.LastInsertId()
	logger.Info(fmt.Sprintf("✅ Inserted user 'David Kim' (%s) with ID: %d", email, id))
	logger.Info(fmt.Sprintf("✅ 'David Kim' (%s) 사용자를 ID %d로 삽입했습니다", email, id))
	logger.Info("")
	return nil
}

// example4Update demonstrates Update method
// example4Update는 Update 메서드를 시연합니다
func example4Update(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔄 Example 4: Update - Update user")
	logger.Info("🔄 예제 4: Update - 사용자 업데이트")
	logger.Info("")

	// Update Jane Smith's age / Jane Smith의 나이 업데이트
	result, err := db.Update(ctx, "users",
		map[string]any{
			"age": 26,
		},
		"email = ?", "jane@example.com")
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Updated %d user(s) - Jane's age changed to 26", rows))
	logger.Info(fmt.Sprintf("✅ %d명의 사용자를 업데이트했습니다 - Jane의 나이를 26으로 변경", rows))
	logger.Info("")
	return nil
}

// example5Count demonstrates Count method
// example5Count는 Count 메서드를 시연합니다
func example5Count(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔢 Example 5: Count - Count users")
	logger.Info("🔢 예제 5: Count - 사용자 수 계산")
	logger.Info("")

	// Count all users / 모든 사용자 수
	totalCount, err := db.Count(ctx, "users")
	if err != nil {
		return fmt.Errorf("count failed: %w", err)
	}
	logger.Info(fmt.Sprintf("Total users: %d", totalCount))
	logger.Info(fmt.Sprintf("전체 사용자: %d명", totalCount))

	// Count users older than 25 / 25세 이상 사용자 수
	adultCount, err := db.Count(ctx, "users", "age > ?", 25)
	if err != nil {
		return fmt.Errorf("count with condition failed: %w", err)
	}
	logger.Info(fmt.Sprintf("Users older than 25: %d", adultCount))
	logger.Info(fmt.Sprintf("25세 이상 사용자: %d명", adultCount))
	logger.Info("")
	return nil
}

// example6Exists demonstrates Exists method
// example6Exists는 Exists 메서드를 시연합니다
func example6Exists(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔍 Example 6: Exists - Check if user exists")
	logger.Info("🔍 예제 6: Exists - 사용자 존재 확인")
	logger.Info("")

	// Check if John Doe exists / John Doe 존재 확인
	exists, err := db.Exists(ctx, "users", "email = ?", "john@example.com")
	if err != nil {
		return fmt.Errorf("exists failed: %w", err)
	}

	if exists {
		logger.Info("✅ User 'john@example.com' exists")
		logger.Info("✅ 사용자 'john@example.com'이 존재합니다")
	} else {
		logger.Info("❌ User 'john@example.com' does not exist")
		logger.Info("❌ 사용자 'john@example.com'이 존재하지 않습니다")
	}
	logger.Info("")
	return nil
}

// example7Transaction demonstrates Transaction method
// example7Transaction은 Transaction 메서드를 시연합니다
func example7Transaction(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("💳 Example 7: Transaction - Insert with transaction")
	logger.Info("💳 예제 7: Transaction - 트랜잭션으로 삽입")
	logger.Info("")

	// Generate unique emails with timestamp / 타임스탬프로 유니크한 이메일 생성
	timestamp := time.Now().Unix()

	err := db.Transaction(ctx, func(tx *mysql.Tx) error {
		// Insert first user / 첫 번째 사용자 삽입
		email1 := fmt.Sprintf("emily.park.%d@example.com", timestamp)
		result1, err := tx.Insert(ctx, "users", map[string]any{
			"name":  "Emily Park",
			"email": email1,
			"age":   27,
			"city":  "Gwangju",
		})
		if err != nil {
			return err // Auto rollback / 자동 롤백
		}
		id1, _ := result1.LastInsertId()
		logger.Info(fmt.Sprintf("  - Inserted Emily Park (ID: %d)", id1))

		// Insert second user / 두 번째 사용자 삽입
		email2 := fmt.Sprintf("frank.lee.%d@example.com", timestamp+1)
		result2, err := tx.Insert(ctx, "users", map[string]any{
			"name":  "Frank Lee",
			"email": email2,
			"age":   29,
			"city":  "Ulsan",
		})
		if err != nil {
			return err // Auto rollback / 자동 롤백
		}
		id2, _ := result2.LastInsertId()
		logger.Info(fmt.Sprintf("  - Inserted Frank Lee (ID: %d)", id2))

		return nil // Auto commit / 자동 커밋
	})

	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	logger.Info("✅ Transaction completed successfully")
	logger.Info("✅ 트랜잭션이 성공적으로 완료되었습니다")
	logger.Info("")
	return nil
}

// example8Delete demonstrates Delete method
// example8Delete는 Delete 메서드를 시연합니다
func example8Delete(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🗑️  Example 8: Delete - Delete user")
	logger.Info("🗑️  예제 8: Delete - 사용자 삭제")
	logger.Info("")

	// Delete Charlie Brown (one of the sample users) / 샘플 사용자 중 한 명인 Charlie Brown 삭제
	result, err := db.Delete(ctx, "users", "email = ?", "charlie@example.com")
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Deleted %d user(s) - Charlie Brown removed", rows))
	logger.Info(fmt.Sprintf("✅ %d명의 사용자를 삭제했습니다 - Charlie Brown 제거됨", rows))
	logger.Info("")
	return nil
}

// example9RawSQL demonstrates raw SQL queries
// example9RawSQL은 raw SQL 쿼리를 시연합니다
func example9RawSQL(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔧 Example 9: Raw SQL - Use raw SQL queries")
	logger.Info("🔧 예제 9: Raw SQL - Raw SQL 쿼리 사용")
	logger.Info("")

	// Execute raw query / Raw 쿼리 실행
	rows, err := db.Query(ctx, "SELECT city, COUNT(*) as count FROM users GROUP BY city ORDER BY count DESC")
	if err != nil {
		return fmt.Errorf("raw query failed: %w", err)
	}
	defer rows.Close()

	logger.Info("Users per city:")
	logger.Info("도시별 사용자 수:")

	for rows.Next() {
		var city string
		var count int
		if err := rows.Scan(&city, &count); err != nil {
			return err
		}
		logger.Info(fmt.Sprintf("  - %s: %d users", city, count))
	}

	logger.Info("")
	return nil
}

// example10QueryBuilderSimple demonstrates Query Builder with simple queries
// example10QueryBuilderSimple은 간단한 쿼리로 Query Builder를 시연합니다
func example10QueryBuilderSimple(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🏗️  Example 10: Query Builder - Simple query")
	logger.Info("🏗️  예제 10: Query Builder - 간단한 쿼리")
	logger.Info("")

	// Query: SELECT name, email, age FROM users WHERE age > 25 ORDER BY age DESC LIMIT 3
	// 쿼리: 나이가 25 이상인 사용자를 나이 내림차순으로 3명 선택
	users, err := db.Select("name", "email", "age").
		From("users").
		Where("age > ?", 25).
		OrderBy("age DESC").
		Limit(3).
		All(ctx)
	if err != nil {
		return fmt.Errorf("query builder failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Top 3 oldest users (age > 25):"))
	logger.Info(fmt.Sprintf("나이 25세 이상 중 나이가 많은 상위 3명:"))
	for i, user := range users {
		logger.Info(fmt.Sprintf("  %d. %s (age: %v, email: %s)",
			i+1, user["name"], user["age"], user["email"]))
	}
	logger.Info("")
	return nil
}

// example11QueryBuilderGroupBy demonstrates Query Builder with GROUP BY
// example11QueryBuilderGroupBy는 GROUP BY를 사용한 Query Builder를 시연합니다
func example11QueryBuilderGroupBy(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("📊 Example 11: Query Builder - GROUP BY with HAVING")
	logger.Info("📊 예제 11: Query Builder - HAVING을 사용한 GROUP BY")
	logger.Info("")

	// Query: SELECT city, COUNT(*) as count FROM users GROUP BY city HAVING COUNT(*) >= 1 ORDER BY count DESC
	// 쿼리: 도시별 사용자 수를 계산하되, 1명 이상인 도시만 선택하고 사용자 수 내림차순 정렬
	results, err := db.Select("city", "COUNT(*) as count").
		From("users").
		GroupBy("city").
		Having("COUNT(*) >= ?", 1).
		OrderBy("count DESC").
		All(ctx)
	if err != nil {
		return fmt.Errorf("query builder GROUP BY failed: %w", err)
	}

	logger.Info("Cities with 1+ users (sorted by count):")
	logger.Info("1명 이상 거주하는 도시 (사용자 수 내림차순):")
	for _, row := range results {
		logger.Info(fmt.Sprintf("  - %s: %v users", row["city"], row["count"]))
	}
	logger.Info("")
	return nil
}

// example12QueryBuilderComplex demonstrates Query Builder with complex conditions
// example12QueryBuilderComplex는 복잡한 조건을 사용한 Query Builder를 시연합니다
func example12QueryBuilderComplex(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔗 Example 12: Query Builder - Multiple WHERE conditions")
	logger.Info("🔗 예제 12: Query Builder - 다중 WHERE 조건")
	logger.Info("")

	// Query: SELECT * FROM users WHERE age > 25 AND city IN ('Seoul', 'Busan') ORDER BY name
	// 쿼리: 나이가 25 이상이고 서울 또는 부산에 거주하는 사용자를 이름순으로 선택
	users, err := db.Select("*").
		From("users").
		Where("age > ?", 25).
		Where("city IN (?, ?)", "Seoul", "Busan").
		OrderBy("name ASC").
		All(ctx)
	if err != nil {
		return fmt.Errorf("query builder complex query failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Users older than 25 in Seoul or Busan:"))
	logger.Info(fmt.Sprintf("서울 또는 부산에 거주하는 25세 이상 사용자:"))
	if len(users) == 0 {
		logger.Info("  (No users found / 사용자 없음)")
	} else {
		for i, user := range users {
			logger.Info(fmt.Sprintf("  %d. %s (age: %v, city: %s)",
				i+1, user["name"], user["age"], user["city"]))
		}
	}
	logger.Info("")
	return nil
}

// example13SelectWhereSimple demonstrates SelectWhere with simple options
// example13SelectWhereSimple은 간단한 옵션으로 SelectWhere를 시연합니다
func example13SelectWhereSimple(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("✨ Example 13: SelectWhere - Simple query with options")
	logger.Info("✨ 예제 13: SelectWhere - 옵션을 사용한 간단한 쿼리")
	logger.Info("")

	// One-liner query with options / 옵션을 사용한 한 줄 쿼리
	// SELECT name, email, age FROM users WHERE age > 25 ORDER BY age DESC LIMIT 3
	users, err := db.SelectWhere(ctx, "users", "age > ?", 25,
		mysql.WithColumns("name", "email", "age"),
		mysql.WithOrderBy("age DESC"),
		mysql.WithLimit(3))
	if err != nil {
		return fmt.Errorf("selectWhere failed: %w", err)
	}

	logger.Info("Top 3 oldest users (age > 25) - using SelectWhere:")
	logger.Info("나이 25세 이상 중 나이가 많은 상위 3명 - SelectWhere 사용:")
	for i, user := range users {
		logger.Info(fmt.Sprintf("  %d. %s (age: %v, email: %s)",
			i+1, user["name"], user["age"], user["email"]))
	}
	logger.Info("")
	return nil
}

// example14SelectWhereGroupBy demonstrates SelectWhere with GROUP BY
// example14SelectWhereGroupBy는 GROUP BY를 사용한 SelectWhere를 시연합니다
func example14SelectWhereGroupBy(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("📈 Example 14: SelectWhere - GROUP BY with HAVING")
	logger.Info("📈 예제 14: SelectWhere - HAVING을 사용한 GROUP BY")
	logger.Info("")

	// One-liner GROUP BY query / 한 줄 GROUP BY 쿼리
	// SELECT city, COUNT(*) as count FROM users GROUP BY city HAVING COUNT(*) >= 1 ORDER BY count DESC
	results, err := db.SelectWhere(ctx, "users", "",
		mysql.WithColumns("city", "COUNT(*) as count"),
		mysql.WithGroupBy("city"),
		mysql.WithHaving("COUNT(*) >= ?", 1),
		mysql.WithOrderBy("count DESC"))
	if err != nil {
		return fmt.Errorf("selectWhere with GROUP BY failed: %w", err)
	}

	logger.Info("Cities with 1+ users (using SelectWhere):")
	logger.Info("1명 이상 거주하는 도시 (SelectWhere 사용):")
	for _, row := range results {
		logger.Info(fmt.Sprintf("  - %s: %v users", row["city"], row["count"]))
	}
	logger.Info("")
	return nil
}

// example15SelectWhereComplex demonstrates SelectWhere with multiple options
// example15SelectWhereComplex는 여러 옵션을 사용한 SelectWhere를 시연합니다
func example15SelectWhereComplex(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🌟 Example 15: SelectWhere - Multiple conditions and options")
	logger.Info("🌟 예제 15: SelectWhere - 다중 조건과 옵션")
	logger.Info("")

	// Complex query in one line / 한 줄로 복잡한 쿼리
	// SELECT DISTINCT city FROM users WHERE age > 25 ORDER BY city
	cities, err := db.SelectWhere(ctx, "users", "age > ?", 25,
		mysql.WithColumns("city"),
		mysql.WithDistinct(),
		mysql.WithOrderBy("city ASC"))
	if err != nil {
		return fmt.Errorf("selectWhere with DISTINCT failed: %w", err)
	}

	logger.Info("Distinct cities with users older than 25:")
	logger.Info("25세 이상 사용자가 있는 도시 목록:")
	for i, city := range cities {
		logger.Info(fmt.Sprintf("  %d. %s", i+1, city["city"]))
	}
	logger.Info("")
	return nil
}

