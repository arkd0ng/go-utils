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
	logger := logging.Default()
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

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("Running Examples / 예제 실행 중")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	// Run all examples / 모든 예제 실행
	if err := runExamples(dsn, config.MySQL, logger); err != nil {
		logger.Error("Examples failed", "error", err)
		os.Exit(1)
	}

	fmt.Println("\n" + strings.Repeat("=", 70))
	logger.Info("All examples completed successfully!")
	logger.Info("모든 예제가 성공적으로 완료되었습니다!")
	fmt.Println(strings.Repeat("=", 70) + "\n")
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
	if err := example1SelectAll(ctx, db); err != nil {
		return err
	}

	// Example 2: SelectOne - Select single user / 단일 사용자 선택
	if err := example2SelectOne(ctx, db); err != nil {
		return err
	}

	// Example 3: Insert - Insert new user / 새 사용자 삽입
	if err := example3Insert(ctx, db); err != nil {
		return err
	}

	// Example 4: Update - Update user / 사용자 업데이트
	if err := example4Update(ctx, db); err != nil {
		return err
	}

	// Example 5: Count - Count users / 사용자 수 계산
	if err := example5Count(ctx, db); err != nil {
		return err
	}

	// Example 6: Exists - Check if user exists / 사용자 존재 확인
	if err := example6Exists(ctx, db); err != nil {
		return err
	}

	// Example 7: Transaction - Insert with transaction / 트랜잭션으로 삽입
	if err := example7Transaction(ctx, db); err != nil {
		return err
	}

	// Example 8: Delete - Delete user / 사용자 삭제
	if err := example8Delete(ctx, db); err != nil {
		return err
	}

	// Example 9: Raw SQL - Use raw SQL queries / Raw SQL 쿼리 사용
	if err := example9RawSQL(ctx, db); err != nil {
		return err
	}

	return nil
}

// example1SelectAll demonstrates SelectAll method
// example1SelectAll은 SelectAll 메서드를 시연합니다
func example1SelectAll(ctx context.Context, db *mysql.Client) error {
	fmt.Println("📋 Example 1: SelectAll - Select all users")
	fmt.Println("📋 예제 1: SelectAll - 모든 사용자 선택\n")

	// Select all users from Seoul / 서울의 모든 사용자 선택
	users, err := db.SelectAll(ctx, "users", "city = ?", "Seoul")
	if err != nil {
		return fmt.Errorf("SelectAll failed: %w", err)
	}

	fmt.Printf("Found %d users from Seoul:\n", len(users))
	fmt.Printf("서울에서 %d명의 사용자를 찾았습니다:\n", len(users))
	for i, user := range users {
		fmt.Printf("  %d. %s (age: %v, email: %s)\n",
			i+1, user["name"], user["age"], user["email"])
	}
	fmt.Println()
	return nil
}

// example2SelectOne demonstrates SelectOne method
// example2SelectOne은 SelectOne 메서드를 시연합니다
func example2SelectOne(ctx context.Context, db *mysql.Client) error {
	fmt.Println("👤 Example 2: SelectOne - Select single user")
	fmt.Println("👤 예제 2: SelectOne - 단일 사용자 선택\n")

	user, err := db.SelectOne(ctx, "users", "email = ?", "john@example.com")
	if err != nil {
		return fmt.Errorf("SelectOne failed: %w", err)
	}

	fmt.Printf("Found user: %s\n", user["name"])
	fmt.Printf("  - Email: %s\n", user["email"])
	fmt.Printf("  - Age: %v\n", user["age"])
	fmt.Printf("  - City: %s\n", user["city"])
	fmt.Println()
	return nil
}

// example3Insert demonstrates Insert method
// example3Insert는 Insert 메서드를 시연합니다
func example3Insert(ctx context.Context, db *mysql.Client) error {
	fmt.Println("➕ Example 3: Insert - Insert new user")
	fmt.Println("➕예제 3: Insert - 새 사용자 삽입\n")

	result, err := db.Insert(ctx, "users", map[string]interface{}{
		"name":  "David Kim",
		"email": "david@example.com",
		"age":   32,
		"city":  "Daegu",
	})
	if err != nil {
		return fmt.Errorf("Insert failed: %w", err)
	}

	id, _ := result.LastInsertId()
	fmt.Printf("✅ Inserted user 'David Kim' with ID: %d\n", id)
	fmt.Printf("✅ 'David Kim' 사용자를 ID %d로 삽입했습니다\n", id)
	fmt.Println()
	return nil
}

// example4Update demonstrates Update method
// example4Update는 Update 메서드를 시연합니다
func example4Update(ctx context.Context, db *mysql.Client) error {
	fmt.Println("🔄 Example 4: Update - Update user")
	fmt.Println("🔄 예제 4: Update - 사용자 업데이트\n")

	result, err := db.Update(ctx, "users",
		map[string]interface{}{
			"age":  33,
			"city": "Daejeon",
		},
		"email = ?", "david@example.com")
	if err != nil {
		return fmt.Errorf("Update failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	fmt.Printf("✅ Updated %d user(s)\n", rows)
	fmt.Printf("✅ %d명의 사용자를 업데이트했습니다\n", rows)
	fmt.Println()
	return nil
}

// example5Count demonstrates Count method
// example5Count는 Count 메서드를 시연합니다
func example5Count(ctx context.Context, db *mysql.Client) error {
	fmt.Println("🔢 Example 5: Count - Count users")
	fmt.Println("🔢 예제 5: Count - 사용자 수 계산\n")

	// Count all users / 모든 사용자 수
	totalCount, err := db.Count(ctx, "users")
	if err != nil {
		return fmt.Errorf("Count failed: %w", err)
	}
	fmt.Printf("Total users: %d\n", totalCount)
	fmt.Printf("전체 사용자: %d명\n", totalCount)

	// Count users older than 25 / 25세 이상 사용자 수
	adultCount, err := db.Count(ctx, "users", "age > ?", 25)
	if err != nil {
		return fmt.Errorf("Count with condition failed: %w", err)
	}
	fmt.Printf("Users older than 25: %d\n", adultCount)
	fmt.Printf("25세 이상 사용자: %d명\n", adultCount)
	fmt.Println()
	return nil
}

// example6Exists demonstrates Exists method
// example6Exists는 Exists 메서드를 시연합니다
func example6Exists(ctx context.Context, db *mysql.Client) error {
	fmt.Println("🔍 Example 6: Exists - Check if user exists")
	fmt.Println("🔍 예제 6: Exists - 사용자 존재 확인\n")

	exists, err := db.Exists(ctx, "users", "email = ?", "david@example.com")
	if err != nil {
		return fmt.Errorf("Exists failed: %w", err)
	}

	if exists {
		fmt.Println("✅ User 'david@example.com' exists")
		fmt.Println("✅ 사용자 'david@example.com'이 존재합니다")
	} else {
		fmt.Println("❌ User 'david@example.com' does not exist")
		fmt.Println("❌ 사용자 'david@example.com'이 존재하지 않습니다")
	}
	fmt.Println()
	return nil
}

// example7Transaction demonstrates Transaction method
// example7Transaction은 Transaction 메서드를 시연합니다
func example7Transaction(ctx context.Context, db *mysql.Client) error {
	fmt.Println("💳 Example 7: Transaction - Insert with transaction")
	fmt.Println("💳 예제 7: Transaction - 트랜잭션으로 삽입\n")

	err := db.Transaction(ctx, func(tx *mysql.Tx) error {
		// Insert first user / 첫 번째 사용자 삽입
		result1, err := tx.Insert(ctx, "users", map[string]interface{}{
			"name":  "Emily Park",
			"email": "emily@example.com",
			"age":   27,
			"city":  "Gwangju",
		})
		if err != nil {
			return err // Auto rollback / 자동 롤백
		}
		id1, _ := result1.LastInsertId()
		fmt.Printf("  - Inserted Emily Park (ID: %d)\n", id1)

		// Insert second user / 두 번째 사용자 삽입
		result2, err := tx.Insert(ctx, "users", map[string]interface{}{
			"name":  "Frank Lee",
			"email": "frank@example.com",
			"age":   29,
			"city":  "Ulsan",
		})
		if err != nil {
			return err // Auto rollback / 자동 롤백
		}
		id2, _ := result2.LastInsertId()
		fmt.Printf("  - Inserted Frank Lee (ID: %d)\n", id2)

		return nil // Auto commit / 자동 커밋
	})

	if err != nil {
		return fmt.Errorf("Transaction failed: %w", err)
	}

	fmt.Println("✅ Transaction completed successfully")
	fmt.Println("✅ 트랜잭션이 성공적으로 완료되었습니다")
	fmt.Println()
	return nil
}

// example8Delete demonstrates Delete method
// example8Delete는 Delete 메서드를 시연합니다
func example8Delete(ctx context.Context, db *mysql.Client) error {
	fmt.Println("🗑️  Example 8: Delete - Delete user")
	fmt.Println("🗑️  예제 8: Delete - 사용자 삭제\n")

	result, err := db.Delete(ctx, "users", "email = ?", "david@example.com")
	if err != nil {
		return fmt.Errorf("Delete failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	fmt.Printf("✅ Deleted %d user(s)\n", rows)
	fmt.Printf("✅ %d명의 사용자를 삭제했습니다\n", rows)
	fmt.Println()
	return nil
}

// example9RawSQL demonstrates raw SQL queries
// example9RawSQL은 raw SQL 쿼리를 시연합니다
func example9RawSQL(ctx context.Context, db *mysql.Client) error {
	fmt.Println("🔧 Example 9: Raw SQL - Use raw SQL queries")
	fmt.Println("🔧 예제 9: Raw SQL - Raw SQL 쿼리 사용\n")

	// Execute raw query / Raw 쿼리 실행
	rows, err := db.Query(ctx, "SELECT city, COUNT(*) as count FROM users GROUP BY city ORDER BY count DESC")
	if err != nil {
		return fmt.Errorf("Raw query failed: %w", err)
	}
	defer rows.Close()

	fmt.Println("Users per city:")
	fmt.Println("도시별 사용자 수:")

	for rows.Next() {
		var city string
		var count int
		if err := rows.Scan(&city, &count); err != nil {
			return err
		}
		fmt.Printf("  - %s: %d users\n", city, count)
	}

	fmt.Println()
	return nil
}
