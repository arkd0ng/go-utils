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
	// Create results directories if they don't exist / 결과 디렉토리가 없다면 새롭게 생성
	if err := os.MkdirAll("logs/", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create logs directory: %v\n", err)
		os.Exit(1)
	}
	if err := os.MkdirAll("logs/mysql_export", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create mysql_export directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(fmt.Sprintf("logs/mysql-example-%s.log", time.Now().Format("20060102-150405"))),
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
	logger.Info("Loading database configuration from cfg/database-mysql.yaml")
	logger.Info("cfg/database-mysql.yaml에서 데이터베이스 설정 로드 중")
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

	// Check if Docker MySQL is running / Docker MySQL 실행 여부 확인
	wasRunning := isDockerMySQLRunning()

	if !wasRunning {
		logger.Info("Docker MySQL is not running, starting container...")
		logger.Info("Docker MySQL이 실행 중이 아닙니다. 컨테이너를 시작합니다...")
		if err := startDockerMySQL(); err != nil {
			logger.Error("Failed to start Docker MySQL", "error", err)
			logger.Info("")
			logger.Info("Please ensure Docker is installed and running:")
			logger.Info("Docker가 설치되어 실행 중인지 확인하세요:")
			logger.Info("  1. Install Docker Desktop: https://www.docker.com/products/docker-desktop")
			logger.Info("  2. Start Docker Desktop")
			logger.Info("  3. Run: docker compose up -d")
			os.Exit(1)
		}

		// Wait for MySQL to be ready / MySQL 준비 대기
		logger.Info("Waiting for Docker MySQL to be ready...")
		logger.Info("Docker MySQL 준비 중...")
		if err := waitForDockerMySQL(config.MySQL, 30*time.Second); err != nil {
			logger.Error("Docker MySQL failed to become ready", "error", err)
			// Clean up - stop MySQL if we started it / 정리 - 시작한 경우 MySQL 중지
			stopDockerMySQL()
			os.Exit(1)
		}
		logger.Info("Docker MySQL is ready!")
		logger.Info("Docker MySQL 준비 완료!")

		// Ensure MySQL stops when program exits / 프로그램 종료 시 MySQL 중지 보장
		defer func() {
			logger.Info("")
			logger.Info("Stopping Docker MySQL container...")
			logger.Info("Docker MySQL 컨테이너 중지 중...")
			if err := stopDockerMySQL(); err != nil {
				logger.Warn("Failed to stop Docker MySQL", "error", err)
			} else {
				logger.Info("✅ Docker MySQL stopped successfully")
				logger.Info("✅ Docker MySQL이 성공적으로 중지되었습니다")
			}
		}()
	} else {
		logger.Info("Docker MySQL is already running")
		logger.Info("Docker MySQL이 이미 실행 중입니다")
		logger.Info("(Will not stop MySQL as it was already running)")
		logger.Info("(이미 실행 중이었으므로 MySQL을 중지하지 않습니다)")
	}

	logger.Info("")
	logger.Info(strings.Repeat("=", 70))
	logger.Info("Running Examples / 예제 실행 중")
	logger.Info(strings.Repeat("=", 70))
	logger.Info("")

	// Initialize database with sample data if needed / 필요한 경우 샘플 데이터로 데이터베이스 초기화
	if err := initializeDatabaseIfNeeded(dsn, logger); err != nil {
		logger.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}

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
	configPath := filepath.Join(projectRoot, "cfg", "database-mysql.yaml")

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

// isDockerMySQLRunning checks if Docker MySQL container is running
// isDockerMySQLRunning은 Docker MySQL 컨테이너가 실행 중인지 확인합니다
func isDockerMySQLRunning() bool {
	cmd := exec.Command("docker", "ps", "--filter", "name=go-utils-mysql", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "go-utils-mysql"
}

// startDockerMySQL starts the Docker MySQL container using docker compose
// startDockerMySQL은 docker compose를 사용하여 Docker MySQL 컨테이너를 시작합니다
func startDockerMySQL() error {
	// Get project root directory / 프로젝트 루트 디렉토리 가져오기
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}
	projectRoot := filepath.Join(wd, "..", "..", ".docker")

	// Start Docker Compose / Docker Compose 시작
	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to start docker compose: %w (output: %s)", err, string(output))
	}
	return nil
}

// waitForDockerMySQL waits for Docker MySQL to be ready
// waitForDockerMySQL은 Docker MySQL이 준비될 때까지 대기합니다
func waitForDockerMySQL(cfg MySQLConfig, timeout time.Duration) error {
	start := time.Now()
	for {
		// Try to connect to MySQL / MySQL 연결 시도
		cmd := exec.Command("docker", "exec", "go-utils-mysql",
			"mysqladmin", "ping", "-h", "localhost", "-u", cfg.User,
			fmt.Sprintf("-p%s", cfg.Password))
		err := cmd.Run()
		if err == nil {
			return nil
		}

		// Check timeout / 타임아웃 확인
		if time.Since(start) > timeout {
			return fmt.Errorf("timeout waiting for MySQL to be ready")
		}

		// Wait before retry / 재시도 전 대기
		time.Sleep(1 * time.Second)
	}
}

// stopDockerMySQL stops the Docker MySQL container
// stopDockerMySQL은 Docker MySQL 컨테이너를 중지합니다
func stopDockerMySQL() error {
	// Get project root directory / 프로젝트 루트 디렉토리 가져오기
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}
	projectRoot := filepath.Join(wd, "..", "..")

	// Stop Docker Compose / Docker Compose 중지
	cmd := exec.Command("docker", "compose", "down")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stop docker compose: %w (output: %s)", err, string(output))
	}
	return nil
}

// initializeDatabaseIfNeeded checks if the database has sample data and initializes it if needed
// initializeDatabaseIfNeeded는 데이터베이스에 샘플 데이터가 있는지 확인하고 필요한 경우 초기화합니다
func initializeDatabaseIfNeeded(dsn string, logger *logging.Logger) error {
	// Create a temporary client to check and initialize database
	// 데이터베이스 확인 및 초기화를 위한 임시 클라이언트 생성
	db, err := mysql.New(mysql.WithDSN(dsn))
	if err != nil {
		return fmt.Errorf("failed to create temporary client: %w", err)
	}
	defer db.Close()

	ctx := context.Background()

	// Check if users table exists / users 테이블 존재 확인
	tableExists, err := db.TableExists(ctx, "users")
	if err != nil {
		return fmt.Errorf("failed to check if users table exists: %w", err)
	}

	if !tableExists {
		logger.Info("users table does not exist, creating...")
		logger.Info("users 테이블이 존재하지 않습니다. 생성 중...")

		// Create users table / users 테이블 생성
		createTableSQL := `
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			age INT NOT NULL,
			city VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP NULL DEFAULT NULL,
			INDEX idx_email (email),
			INDEX idx_city (city),
			INDEX idx_age (age)
		`
		if err := db.CreateTable(ctx, "users", createTableSQL); err != nil {
			return fmt.Errorf("failed to create users table: %w", err)
		}
		logger.Info("✅ users table created successfully")
		logger.Info("✅ users 테이블이 성공적으로 생성되었습니다")
	}

	// Check if sample data exists by looking for a known sample user / 알려진 샘플 사용자를 찾아 샘플 데이터 존재 확인
	exists, err := db.Exists("users", "email = ?", "john@example.com")
	if err != nil {
		return fmt.Errorf("failed to check for sample data: %w", err)
	}

	if !exists {
		// Clear any existing data and insert fresh sample data / 기존 데이터를 지우고 새 샘플 데이터 삽입
		logger.Info("Sample data not found or incomplete, resetting users table...")
		logger.Info("샘플 데이터가 없거나 불완전합니다. users 테이블을 재설정합니다...")

		// Truncate table to remove any leftover data / 테이블을 비워 남은 데이터 제거
		if err := db.TruncateTable(ctx, "users"); err != nil {
			return fmt.Errorf("failed to truncate users table: %w", err)
		}

		// Insert sample data / 샘플 데이터 삽입
		sampleUsers := []map[string]interface{}{
			{"name": "John Doe", "email": "john@example.com", "age": 30, "city": "Seoul"},
			{"name": "Jane Smith", "email": "jane@example.com", "age": 25, "city": "Seoul"},
			{"name": "Bob Johnson", "email": "bob@example.com", "age": 35, "city": "Seoul"},
			{"name": "Alice Williams", "email": "alice@example.com", "age": 28, "city": "Incheon"},
			{"name": "Emily Park", "email": "emily@example.com", "age": 27, "city": "Gwangju"},
			{"name": "Frank Lee", "email": "frank@example.com", "age": 32, "city": "Daegu"},
			{"name": "Grace Kim", "email": "grace@example.com", "age": 29, "city": "Busan"},
			{"name": "Henry Choi", "email": "henry@example.com", "age": 31, "city": "Ulsan"},
			{"name": "Iris Jung", "email": "iris@example.com", "age": 26, "city": "Daejeon"},
			{"name": "Jack Yoon", "email": "jack@example.com", "age": 33, "city": "Gwangju"},
			{"name": "Charlie Brown", "email": "charlie@example.com", "age": 40, "city": "Seoul"},
		}

		_, err := db.BatchInsert(ctx, "users", sampleUsers)
		if err != nil {
			return fmt.Errorf("failed to insert sample data: %w", err)
		}

		logger.Info(fmt.Sprintf("✅ Inserted %d sample users", len(sampleUsers)))
		logger.Info(fmt.Sprintf("✅ %d명의 샘플 사용자를 삽입했습니다", len(sampleUsers)))
	} else {
		count, _ := db.Count("users")
		logger.Info(fmt.Sprintf("Sample data already exists (%d users), skipping initialization", count))
		logger.Info(fmt.Sprintf("샘플 데이터가 이미 존재합니다 (%d명의 사용자). 초기화를 건너뜁니다", count))
	}

	logger.Info("")
	return nil
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

	// Example 16: SelectColumn - Select single column / 단일 컬럼 선택
	if err := example16SelectColumn(ctx, db, logger); err != nil {
		return err
	}

	// Example 17: SelectColumns - Select multiple columns / 여러 컬럼 선택
	if err := example17SelectColumns(ctx, db, logger); err != nil {
		return err
	}

	// Example 18: BatchInsert - Batch insert operations / 배치 삽입 작업
	if err := example18BatchInsert(ctx, db, logger); err != nil {
		return err
	}

	// Example 19: BatchUpdate - Batch update operations / 배치 업데이트 작업
	if err := example19BatchUpdate(ctx, db, logger); err != nil {
		return err
	}

	// Example 20: BatchDelete - Batch delete operations / 배치 삭제 작업
	if err := example20BatchDelete(ctx, db, logger); err != nil {
		return err
	}

	// Example 21: Upsert - Insert or update on duplicate / 중복 시 삽입 또는 업데이트
	if err := example21Upsert(ctx, db, logger); err != nil {
		return err
	}

	// Example 22: UpsertBatch - Batch upsert operations / 배치 upsert 작업
	if err := example22UpsertBatch(ctx, db, logger); err != nil {
		return err
	}

	// Example 23: Pagination - Basic pagination / 기본 페이지네이션
	if err := example23Pagination(ctx, db, logger); err != nil {
		return err
	}

	// Example 24: PaginationWithOptions - Pagination with options / 옵션이 있는 페이지네이션
	if err := example24PaginationWithOptions(ctx, db, logger); err != nil {
		return err
	}

	// Example 24.5: Prepare table for soft delete - Add deleted_at column / 소프트 삭제를 위한 테이블 준비 - deleted_at 컬럼 추가
	if err := example24_5PrepareForSoftDelete(ctx, db, logger); err != nil {
		return err
	}

	// Example 25: SoftDelete - Soft delete user / 사용자 소프트 삭제
	if err := example25SoftDelete(ctx, db, logger); err != nil {
		return err
	}

	// Example 26: RestoreSoftDeleted - Restore soft-deleted user / 소프트 삭제된 사용자 복구
	if err := example26RestoreSoftDeleted(ctx, db, logger); err != nil {
		return err
	}

	// Example 27: SelectTrashed - Query trashed users / 삭제된 사용자 조회
	if err := example27SelectTrashed(ctx, db, logger); err != nil {
		return err
	}

	// Example 28: QueryStats - Query execution statistics / 쿼리 실행 통계
	if err := example28QueryStats(ctx, db, logger); err != nil {
		return err
	}

	// Example 29: SlowQueryLog - Slow query logging / 느린 쿼리 로깅
	if err := example29SlowQueryLog(ctx, db, logger); err != nil {
		return err
	}

	// Example 30: PoolMetrics - Connection pool metrics / 연결 풀 메트릭
	if err := example30PoolMetrics(ctx, db, logger); err != nil {
		return err
	}

	// Example 31: GetTables - List all tables / 모든 테이블 나열
	if err := example31GetTables(ctx, db, logger); err != nil {
		return err
	}

	// Example 32: InspectTable - Inspect table structure / 테이블 구조 검사
	if err := example32InspectTable(ctx, db, logger); err != nil {
		return err
	}

	// Example 33: CreateTable - Create new table / 새 테이블 생성
	if err := example33CreateTable(ctx, db, logger); err != nil {
		return err
	}

	// Example 34: AddColumn - Migration operations / 마이그레이션 작업
	if err := example34AddColumn(ctx, db, logger); err != nil {
		return err
	}

	// Example 35: ExportCSV - Export table to CSV / 테이블을 CSV로 내보내기
	if err := example35ExportCSV(ctx, db, logger); err != nil {
		return err
	}

	logger.Info("========================================")
	logger.Info("All examples completed successfully!")
	logger.Info("모든 예제가 성공적으로 완료되었습니다!")
	logger.Info("========================================")
	logger.Info("")

	return nil
}

// example1SelectAll demonstrates SelectAll method
// example1SelectAll은 SelectAll 메서드를 시연합니다
func example1SelectAll(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("📋 Example 1: SelectAll - Select all users")
	logger.Info("📋 예제 1: SelectAll - 모든 사용자 선택")
	logger.Info("")

	// Select all users from Seoul / 서울의 모든 사용자 선택
	// Using non-context version for simplicity / 간단함을 위해 non-context 버전 사용
	users, err := db.SelectAll("users", "city = ?", "Seoul")
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
func example2SelectOne(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("👤 Example 2: SelectOne - Select single user")
	logger.Info("👤 예제 2: SelectOne - 단일 사용자 선택")
	logger.Info("")

	// Using non-context version / non-context 버전 사용
	user, err := db.SelectOne("users", "email = ?", "john@example.com")
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
func example3Insert(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("➕ Example 3: Insert - Insert new user")
	logger.Info("➕예제 3: Insert - 새 사용자 삽입")
	logger.Info("")

	// Generate unique email with timestamp / 타임스탬프로 유니크한 이메일 생성
	timestamp := time.Now().Unix()
	email := fmt.Sprintf("david.kim.%d@example.com", timestamp)

	// Using non-context version / non-context 버전 사용
	result, err := db.Insert("users", map[string]any{
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
func example4Update(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔄 Example 4: Update - Update user")
	logger.Info("🔄 예제 4: Update - 사용자 업데이트")
	logger.Info("")

	// Update Jane Smith's age / Jane Smith의 나이 업데이트
	result, err := db.Update("users",
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
func example5Count(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔢 Example 5: Count - Count users")
	logger.Info("🔢 예제 5: Count - 사용자 수 계산")
	logger.Info("")

	// Count all users / 모든 사용자 수
	totalCount, err := db.Count("users")
	if err != nil {
		return fmt.Errorf("count failed: %w", err)
	}
	logger.Info(fmt.Sprintf("Total users: %d", totalCount))
	logger.Info(fmt.Sprintf("전체 사용자: %d명", totalCount))

	// Count users older than 25 / 25세 이상 사용자 수
	adultCount, err := db.Count("users", "age > ?", 25)
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
func example6Exists(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🔍 Example 6: Exists - Check if user exists")
	logger.Info("🔍 예제 6: Exists - 사용자 존재 확인")
	logger.Info("")

	// Check if John Doe exists / John Doe 존재 확인
	exists, err := db.Exists("users", "email = ?", "john@example.com")
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
		result1, err := tx.Insert("users", map[string]any{
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
		result2, err := tx.Insert("users", map[string]any{
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
func example8Delete(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("🗑️  Example 8: Delete - Delete user")
	logger.Info("🗑️  예제 8: Delete - 사용자 삭제")
	logger.Info("")

	// Delete Charlie Brown (one of the sample users) / 샘플 사용자 중 한 명인 Charlie Brown 삭제
	result, err := db.Delete("users", "email = ?", "charlie@example.com")
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

	logger.Info("Top 3 oldest users (age > 25):")
	logger.Info("나이 25세 이상 중 나이가 많은 상위 3명:")
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

	logger.Info("Users older than 25 in Seoul or Busan:")
	logger.Info("서울 또는 부산에 거주하는 25세 이상 사용자:")
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

// example16SelectColumn demonstrates SelectColumn - single column selection
// example16SelectColumn은 SelectColumn을 시연합니다 - 단일 컬럼 선택
func example16SelectColumn(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 16: SelectColumn - Single Column Selection")
	logger.Info("예제 16: SelectColumn - 단일 컬럼 선택")
	logger.Info("========================================")

	// SELECT email FROM users
	logger.Info("Selecting all email addresses...")
	logger.Info("모든 이메일 주소 선택 중...")
	// Using non-context version / non-context 버전 사용
	emails, err := db.SelectColumn("users", "email")
	if err != nil {
		return fmt.Errorf("SelectColumn failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d email addresses:", len(emails)))
	logger.Info(fmt.Sprintf("%d개의 이메일 주소를 찾았습니다:", len(emails)))
	for i, row := range emails {
		if i >= 5 {
			logger.Info("  ... (truncated)")
			break
		}
		logger.Info(fmt.Sprintf("  %d. %s", i+1, row["email"]))
	}

	// SELECT name FROM users WHERE age > 25
	logger.Info("")
	logger.Info("Selecting names of users older than 25...")
	logger.Info("25세 이상 사용자의 이름 선택 중...")
	names, err := db.SelectColumn("users", "name", "age > ?", 25)
	if err != nil {
		return fmt.Errorf("SelectColumn with condition failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d names:", len(names)))
	logger.Info(fmt.Sprintf("%d개의 이름을 찾았습니다:", len(names)))
	for i, row := range names {
		if i >= 5 {
			logger.Info("  ... (truncated)")
			break
		}
		logger.Info(fmt.Sprintf("  %d. %s", i+1, row["name"]))
	}

	logger.Info("")
	return nil
}

// example17SelectColumns demonstrates SelectColumns - multiple columns selection
// example17SelectColumns는 SelectColumns를 시연합니다 - 여러 컬럼 선택
func example17SelectColumns(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 17: SelectColumns - Multiple Columns Selection")
	logger.Info("예제 17: SelectColumns - 여러 컬럼 선택")
	logger.Info("========================================")

	// SELECT name, email FROM users
	logger.Info("Selecting name and email of all users...")
	logger.Info("모든 사용자의 이름과 이메일 선택 중...")
	// Using non-context version / non-context 버전 사용
	users, err := db.SelectColumns("users", []string{"name", "email"})
	if err != nil {
		return fmt.Errorf("SelectColumns failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d users:", len(users)))
	logger.Info(fmt.Sprintf("%d명의 사용자를 찾았습니다:", len(users)))
	for i, user := range users {
		if i >= 5 {
			logger.Info("  ... (truncated)")
			break
		}
		logger.Info(fmt.Sprintf("  %d. %s <%s>", i+1, user["name"], user["email"]))
	}

	// SELECT name, age, city FROM users WHERE age > 25
	logger.Info("")
	logger.Info("Selecting name, age, and city of users older than 25...")
	logger.Info("25세 이상 사용자의 이름, 나이, 도시 선택 중...")
	usersWithAge, err := db.SelectColumns("users", []string{"name", "age", "city"}, "age > ?", 25)
	if err != nil {
		return fmt.Errorf("SelectColumns with condition failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d users:", len(usersWithAge)))
	logger.Info(fmt.Sprintf("%d명의 사용자를 찾았습니다:", len(usersWithAge)))
	for i, user := range usersWithAge {
		if i >= 5 {
			logger.Info("  ... (truncated)")
			break
		}
		logger.Info(fmt.Sprintf("  %d. %s (age: %v, city: %s)",
			i+1, user["name"], user["age"], user["city"]))
	}

	logger.Info("")
	return nil
}

// example18BatchInsert demonstrates BatchInsert - multiple row insertion
// example18BatchInsert는 BatchInsert를 시연합니다 - 여러 행 삽입
func example18BatchInsert(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 18: BatchInsert - Insert multiple users in single query")
	logger.Info("예제 18: BatchInsert - 단일 쿼리로 여러 사용자 삽입")
	logger.Info("========================================")

	// Generate unique emails with timestamp / 타임스탬프로 유니크한 이메일 생성
	timestamp := time.Now().Unix()

	// Prepare batch data / 배치 데이터 준비
	data := []map[string]interface{}{
		{
			"name":  "Alice Johnson",
			"email": fmt.Sprintf("alice.%d@example.com", timestamp),
			"age":   28,
			"city":  "Incheon",
		},
		{
			"name":  "Bob Smith",
			"email": fmt.Sprintf("bob.%d@example.com", timestamp),
			"age":   35,
			"city":  "Daejeon",
		},
		{
			"name":  "Carol White",
			"email": fmt.Sprintf("carol.%d@example.com", timestamp),
			"age":   29,
			"city":  "Busan",
		},
		{
			"name":  "Dave Lee",
			"email": fmt.Sprintf("dave.%d@example.com", timestamp),
			"age":   42,
			"city":  "Seoul",
		},
	}

	logger.Info(fmt.Sprintf("Inserting %d users in a single batch operation...", len(data)))
	logger.Info(fmt.Sprintf("배치 작업으로 %d명의 사용자를 삽입합니다...", len(data)))

	result, err := db.BatchInsert(ctx, "users", data)
	if err != nil {
		return fmt.Errorf("BatchInsert failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Successfully inserted %d users", rows))
	logger.Info(fmt.Sprintf("✅ %d명의 사용자를 성공적으로 삽입했습니다", rows))

	logger.Info("")
	return nil
}

// example19BatchUpdate demonstrates BatchUpdate - multiple row updates in transaction
// example19BatchUpdate는 BatchUpdate를 시연합니다 - 트랜잭션에서 여러 행 업데이트
func example19BatchUpdate(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 19: BatchUpdate - Update multiple users in transaction")
	logger.Info("예제 19: BatchUpdate - 트랜잭션에서 여러 사용자 업데이트")
	logger.Info("========================================")

	// Prepare batch updates / 배치 업데이트 준비
	updates := []mysql.BatchUpdateItem{
		{
			Data:             map[string]interface{}{"age": 31},
			ConditionAndArgs: []interface{}{"email = ?", "john@example.com"},
		},
		{
			Data:             map[string]interface{}{"age": 27, "city": "Seoul"},
			ConditionAndArgs: []interface{}{"email = ?", "jane@example.com"},
		},
	}

	logger.Info(fmt.Sprintf("Updating %d users in a transaction...", len(updates)))
	logger.Info(fmt.Sprintf("트랜잭션에서 %d명의 사용자를 업데이트합니다...", len(updates)))

	err := db.BatchUpdate(ctx, "users", updates)
	if err != nil {
		return fmt.Errorf("BatchUpdate failed: %w", err)
	}

	logger.Info("✅ Batch update completed successfully")
	logger.Info("✅ 배치 업데이트가 성공적으로 완료되었습니다")

	logger.Info("")
	return nil
}

// example20BatchDelete demonstrates BatchDelete - delete multiple rows by IDs
// example20BatchDelete는 BatchDelete를 시연합니다 - ID로 여러 행 삭제
func example20BatchDelete(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 20: BatchDelete - Delete multiple users by IDs")
	logger.Info("예제 20: BatchDelete - ID로 여러 사용자 삭제")
	logger.Info("========================================")

	// First, get some user IDs to delete / 먼저 삭제할 사용자 ID 가져오기
	users, err := db.SelectWhere(ctx, "users", "city = ?", "Daejeon",
		mysql.WithColumns("id"),
		mysql.WithLimit(2))
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	if len(users) == 0 {
		logger.Info("No users found to delete")
		logger.Info("삭제할 사용자가 없습니다")
		logger.Info("")
		return nil
	}

	// Collect IDs / ID 수집
	ids := make([]interface{}, len(users))
	for i, user := range users {
		ids[i] = user["id"]
	}

	logger.Info(fmt.Sprintf("Deleting %d users with IDs: %v", len(ids), ids))
	logger.Info(fmt.Sprintf("ID가 %v인 %d명의 사용자를 삭제합니다", ids, len(ids)))

	result, err := db.BatchDelete(ctx, "users", "id", ids)
	if err != nil {
		return fmt.Errorf("BatchDelete failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Deleted %d users", rows))
	logger.Info(fmt.Sprintf("✅ %d명의 사용자를 삭제했습니다", rows))

	logger.Info("")
	return nil
}

// example21Upsert demonstrates Upsert - insert or update on duplicate key
// example21Upsert는 Upsert를 시연합니다 - 중복 키에서 삽입 또는 업데이트
func example21Upsert(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 21: Upsert - Insert or update on duplicate key")
	logger.Info("예제 21: Upsert - 중복 키에서 삽입 또는 업데이트")
	logger.Info("========================================")

	// First upsert - will insert / 첫 번째 upsert - 삽입됨
	data := map[string]interface{}{
		"email": "upsert.test@example.com",
		"name":  "Upsert Test User",
		"age":   30,
		"city":  "Seoul",
	}

	logger.Info("First upsert (will insert new record)...")
	logger.Info("첫 번째 upsert (새 레코드 삽입)...")

	result, err := db.Upsert(ctx, "users", data, []string{"name", "age", "city"})
	if err != nil {
		return fmt.Errorf("[ERROR] Upsert failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Rows affected: %d (1 = insert, 2 = update)", rows))
	logger.Info(fmt.Sprintf("✅ 영향받은 행: %d (1 = 삽입, 2 = 업데이트)", rows))

	// Second upsert - will update / 두 번째 upsert - 업데이트됨
	logger.Info("")
	logger.Info("Second upsert with same email (will update)...")
	logger.Info("같은 이메일로 두 번째 upsert (업데이트됨)...")

	data["age"] = 31
	data["city"] = "Busan"

	result, err = db.Upsert(ctx, "users", data, []string{"name", "age", "city"})
	if err != nil {
		return fmt.Errorf("[ERROR] Upsert failed: %w", err)
	}

	rows, _ = result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Rows affected: %d", rows))
	logger.Info(fmt.Sprintf("✅ 영향받은 행: %d", rows))

	logger.Info("")
	return nil
}

// example22UpsertBatch demonstrates UpsertBatch - batch upsert operations
// example22UpsertBatch는 UpsertBatch를 시연합니다 - 배치 upsert 작업
func example22UpsertBatch(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 22: UpsertBatch - Batch upsert operations")
	logger.Info("예제 22: UpsertBatch - 배치 upsert 작업")
	logger.Info("========================================")

	timestamp := time.Now().Unix()

	data := []map[string]interface{}{
		{
			"email": fmt.Sprintf("batch.upsert1.%d@example.com", timestamp),
			"name":  "Batch User 1",
			"age":   25,
			"city":  "Seoul",
		},
		{
			"email": fmt.Sprintf("batch.upsert2.%d@example.com", timestamp),
			"name":  "Batch User 2",
			"age":   30,
			"city":  "Busan",
		},
	}

	logger.Info(fmt.Sprintf("Performing batch upsert for %d users...", len(data)))
	logger.Info(fmt.Sprintf("%d명의 사용자에 대해 배치 upsert를 수행합니다...", len(data)))

	result, err := db.UpsertBatch(ctx, "users", data, []string{"name", "age", "city"})
	if err != nil {
		return fmt.Errorf("UpsertBatch failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Rows affected: %d", rows))
	logger.Info(fmt.Sprintf("✅ 영향받은 행: %d", rows))

	logger.Info("")
	return nil
}

// example23Pagination demonstrates Paginate - basic pagination
// example23Pagination은 Paginate를 시연합니다 - 기본 페이지네이션
func example23Pagination(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 23: Paginate - Basic pagination with metadata")
	logger.Info("예제 23: Paginate - 메타데이터가 있는 기본 페이지네이션")
	logger.Info("========================================")

	// Get first page with 5 items / 5개 항목의 첫 페이지 가져오기
	page := 1
	pageSize := 5

	logger.Info(fmt.Sprintf("Fetching page %d (size: %d)...", page, pageSize))
	logger.Info(fmt.Sprintf("페이지 %d (크기: %d) 가져오는 중...", page, pageSize))

	result, err := db.Paginate(ctx, "users", page, pageSize)
	if err != nil {
		return fmt.Errorf("[ERROR] Paginate failed: %w", err)
	}

	logger.Info(fmt.Sprintf("✅ Page: %d/%d", result.Page, result.TotalPages))
	logger.Info(fmt.Sprintf("✅ Total rows: %d", result.TotalRows))
	logger.Info(fmt.Sprintf("✅ Has next: %v, Has prev: %v", result.HasNext, result.HasPrev))
	logger.Info("")
	logger.Info("Page data:")
	logger.Info("페이지 데이터:")
	for i, user := range result.Data {
		logger.Info(fmt.Sprintf("  %d. %s (%s)", i+1, user["name"], user["email"]))
	}

	logger.Info("")
	return nil
}

// example24PaginationWithOptions demonstrates Paginate with WHERE and ORDER BY
// example24PaginationWithOptions는 WHERE 및 ORDER BY를 사용한 Paginate를 시연합니다
func example24PaginationWithOptions(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 24: Paginate - With WHERE and ORDER BY")
	logger.Info("예제 24: Paginate - WHERE 및 ORDER BY 사용")
	logger.Info("========================================")

	// Get page 1 of users older than 25, ordered by age / 25세 이상 사용자의 1페이지, 나이순 정렬
	result, err := db.Paginate(ctx, "users", 1, 3, "age > ?", 25,
		mysql.WithOrderBy("age DESC"),
		mysql.WithColumns("name", "email", "age"))
	if err != nil {
		return fmt.Errorf("[ERROR] Paginate with options failed: %w", err)
	}

	logger.Info("Users older than 25, ordered by age (descending):")
	logger.Info("25세 이상 사용자, 나이 내림차순 정렬:")
	logger.Info(fmt.Sprintf("Page %d/%d (Total: %d users)", result.Page, result.TotalPages, result.TotalRows))
	logger.Info(fmt.Sprintf("페이지 %d/%d (전체: %d명)", result.Page, result.TotalPages, result.TotalRows))
	for i, user := range result.Data {
		logger.Info(fmt.Sprintf("  %d. %s (age: %v)", i+1, user["name"], user["age"]))
	}

	logger.Info("")
	return nil
}

// example24_5PrepareForSoftDelete prepares the users table for soft delete by adding deleted_at column
// example24_5PrepareForSoftDelete는 deleted_at 컬럼을 추가하여 users 테이블을 소프트 삭제를 위해 준비합니다
func example24_5PrepareForSoftDelete(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 24.5: Prepare for SoftDelete - Add deleted_at column")
	logger.Info("예제 24.5: SoftDelete 준비 - deleted_at 컬럼 추가")
	logger.Info("========================================")

	// Check if deleted_at column already exists / deleted_at 컬럼이 이미 존재하는지 확인
	columns, err := db.GetColumns(ctx, "users")
	if err != nil {
		return fmt.Errorf("GetColumns failed: %w", err)
	}

	hasDeletedAt := false
	for _, col := range columns {
		if col.Name == "deleted_at" {
			hasDeletedAt = true
			break
		}
	}

	if hasDeletedAt {
		logger.Info("✅ deleted_at column already exists, skipping migration")
		logger.Info("✅ deleted_at 컬럼이 이미 존재하므로 마이그레이션을 건너뜁니다")
		logger.Info("")
		return nil
	}

	// Add deleted_at column for soft delete / 소프트 삭제를 위한 deleted_at 컬럼 추가
	logger.Info("Adding deleted_at column to users table...")
	logger.Info("users 테이블에 deleted_at 컬럼 추가 중...")

	err = db.AddColumn(ctx, "users", "deleted_at", "TIMESTAMP NULL DEFAULT NULL")
	if err != nil {
		return fmt.Errorf("AddColumn failed: %w", err)
	}

	logger.Info("✅ Successfully added deleted_at column")
	logger.Info("✅ deleted_at 컬럼을 성공적으로 추가했습니다")

	// Verify the column was added / 컬럼이 추가되었는지 확인
	columns, err = db.GetColumns(ctx, "users")
	if err != nil {
		return fmt.Errorf("GetColumns verification failed: %w", err)
	}

	logger.Info("")
	logger.Info("Current table structure:")
	logger.Info("현재 테이블 구조:")
	for _, col := range columns {
		logger.Info(fmt.Sprintf("  - %s (%s)", col.Name, col.Type))
	}

	logger.Info("")
	return nil
}

// example25SoftDelete demonstrates SoftDelete - mark row as deleted
// example25SoftDelete는 SoftDelete를 시연합니다 - 행을 삭제로 표시
func example25SoftDelete(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 25: SoftDelete - Mark user as deleted")
	logger.Info("예제 25: SoftDelete - 사용자를 삭제로 표시")
	logger.Info("========================================")

	// Find a user to soft delete / 소프트 삭제할 사용자 찾기
	users, err := db.SelectWhere(ctx, "users", "city = ?", "Seoul",
		mysql.WithLimit(1))
	if err != nil || len(users) == 0 {
		logger.Info("No users found to soft delete")
		logger.Info("소프트 삭제할 사용자가 없습니다")
		logger.Info("")
		return nil
	}

	userID := users[0]["id"]
	userName := users[0]["name"]

	logger.Info(fmt.Sprintf("Soft deleting user: %s (ID: %v)", userName, userID))
	logger.Info(fmt.Sprintf("사용자 소프트 삭제 중: %s (ID: %v)", userName, userID))

	result, err := db.SoftDelete(ctx, "users", "id = ?", userID)
	if err != nil {
		return fmt.Errorf("SoftDelete failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Soft deleted %d user(s)", rows))
	logger.Info(fmt.Sprintf("✅ %d명의 사용자를 소프트 삭제했습니다", rows))

	logger.Info("")
	return nil
}

// example26RestoreSoftDeleted demonstrates Restore - restore soft-deleted rows
// example26RestoreSoftDeleted는 Restore를 시연합니다 - 소프트 삭제된 행 복구
func example26RestoreSoftDeleted(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 26: Restore - Restore soft-deleted user")
	logger.Info("예제 26: Restore - 소프트 삭제된 사용자 복구")
	logger.Info("========================================")

	// Find a soft-deleted user / 소프트 삭제된 사용자 찾기
	trashedUsers, err := db.SelectAllOnlyTrashed(ctx, "users", "", mysql.WithLimit(1))
	if err != nil || len(trashedUsers) == 0 {
		logger.Info("No soft-deleted users found to restore")
		logger.Info("복구할 소프트 삭제된 사용자가 없습니다")
		logger.Info("")
		return nil
	}

	userID := trashedUsers[0]["id"]
	userName := trashedUsers[0]["name"]

	logger.Info(fmt.Sprintf("Restoring user: %s (ID: %v)", userName, userID))
	logger.Info(fmt.Sprintf("사용자 복구 중: %s (ID: %v)", userName, userID))

	result, err := db.Restore(ctx, "users", "id = ?", userID)
	if err != nil {
		return fmt.Errorf("[ERROR] Restore failed: %w", err)
	}

	rows, _ := result.RowsAffected()
	logger.Info(fmt.Sprintf("✅ Restored %d user(s)", rows))
	logger.Info(fmt.Sprintf("✅ %d명의 사용자를 복구했습니다", rows))

	logger.Info("")
	return nil
}

// example27SelectTrashed demonstrates SelectAllWithTrashed and SelectAllOnlyTrashed
// example27SelectTrashed는 SelectAllWithTrashed 및 SelectAllOnlyTrashed를 시연합니다
func example27SelectTrashed(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 27: SelectTrashed - Query trashed and all users")
	logger.Info("예제 27: SelectTrashed - 삭제된 사용자와 전체 사용자 조회")
	logger.Info("========================================")

	// Count all users including trashed / 삭제된 사용자를 포함한 전체 사용자 수
	totalCount, err := db.CountWithTrashed(ctx, "users")
	if err != nil {
		return fmt.Errorf("CountWithTrashed failed: %w", err)
	}

	// Count only trashed users / 삭제된 사용자만 계산
	trashedCount, err := db.CountOnlyTrashed(ctx, "users")
	if err != nil {
		return fmt.Errorf("CountOnlyTrashed failed: %w", err)
	}

	activeCount := totalCount - trashedCount

	logger.Info(fmt.Sprintf("Total users: %d", totalCount))
	logger.Info(fmt.Sprintf("Active users: %d", activeCount))
	logger.Info(fmt.Sprintf("Trashed users: %d", trashedCount))
	logger.Info(fmt.Sprintf("전체 사용자: %d명, 활성: %d명, 삭제됨: %d명", totalCount, activeCount, trashedCount))

	logger.Info("")
	return nil
}

// example28QueryStats demonstrates GetQueryStats - query execution statistics
// example28QueryStats는 GetQueryStats를 시연합니다 - 쿼리 실행 통계
func example28QueryStats(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 28: QueryStats - Query execution statistics")
	logger.Info("예제 28: QueryStats - 쿼리 실행 통계")
	logger.Info("========================================")

	// Enable query stats / 쿼리 통계 활성화
	db.EnableQueryStats()
	logger.Info("Query statistics enabled")
	logger.Info("쿼리 통계가 활성화되었습니다")

	// Perform some queries / 몇 가지 쿼리 수행
	db.Count("users")
	db.SelectAll("users", "city = ?", "Seoul")
	db.SelectOne("users", "email = ?", "john@example.com")

	// Get statistics / 통계 가져오기
	stats := db.GetQueryStats()

	logger.Info("Query Statistics:")
	logger.Info("쿼리 통계:")
	logger.Info(fmt.Sprintf("  Total queries: %d", stats.TotalQueries))
	logger.Info(fmt.Sprintf("  Successful: %d", stats.SuccessQueries))
	logger.Info(fmt.Sprintf("  Failed: %d", stats.FailedQueries))
	logger.Info(fmt.Sprintf("  Average duration: %v", stats.AvgDuration))
	logger.Info(fmt.Sprintf("  Slow queries: %d", stats.SlowQueries))

	logger.Info("")
	return nil
}

// example29SlowQueryLog demonstrates EnableSlowQueryLog - slow query logging
// example29SlowQueryLog는 EnableSlowQueryLog를 시연합니다 - 느린 쿼리 로깅
func example29SlowQueryLog(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 29: SlowQueryLog - Slow query detection")
	logger.Info("예제 29: SlowQueryLog - 느린 쿼리 감지")
	logger.Info("========================================")

	// Enable slow query log (threshold: 100ms) / 느린 쿼리 로그 활성화 (임계값: 100ms)
	db.EnableSlowQueryLog(100*time.Millisecond, func(info mysql.SlowQueryInfo) {
		logger.Warn("Slow query detected",
			"duration", info.Duration,
			"query", info.Query[:min(50, len(info.Query))]+"...")
	})

	logger.Info("Slow query logging enabled (threshold: 100ms)")
	logger.Info("느린 쿼리 로깅이 활성화되었습니다 (임계값: 100ms)")

	// Perform some queries / 일부 쿼리 수행
	db.SelectAll("users")
	time.Sleep(150 * time.Millisecond) // Simulate slow query / 느린 쿼리 시뮬레이션

	// Get slow queries / 느린 쿼리 가져오기
	slowQueries := db.GetSlowQueries(5)
	if len(slowQueries) > 0 {
		logger.Info(fmt.Sprintf("Found %d slow queries:", len(slowQueries)))
		logger.Info(fmt.Sprintf("%d개의 느린 쿼리를 찾았습니다:", len(slowQueries)))
		for i, sq := range slowQueries {
			logger.Info(fmt.Sprintf("  %d. Duration: %v", i+1, sq.Duration))
		}
	} else {
		logger.Info("No slow queries detected")
		logger.Info("느린 쿼리가 감지되지 않았습니다")
	}

	logger.Info("")
	return nil
}

// example30PoolMetrics demonstrates GetPoolMetrics - connection pool metrics
// example30PoolMetrics는 GetPoolMetrics를 시연합니다 - 연결 풀 메트릭
func example30PoolMetrics(_ context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 30: PoolMetrics - Connection pool metrics")
	logger.Info("예제 30: PoolMetrics - 연결 풀 메트릭")
	logger.Info("========================================")

	// Get pool metrics / 풀 메트릭 가져오기
	metrics := db.GetPoolMetrics()

	logger.Info(fmt.Sprintf("Total connection pools: %d", metrics.PoolCount))
	logger.Info(fmt.Sprintf("Total connections: %d", metrics.TotalConnections))
	logger.Info(fmt.Sprintf("전체 연결 풀: %d개, 전체 연결: %d개", metrics.PoolCount, metrics.TotalConnections))

	logger.Info("")
	logger.Info("Pool statistics:")
	logger.Info("풀 통계:")
	for _, pool := range metrics.PoolStats {
		logger.Info(fmt.Sprintf("  Pool %d:", pool.Index))
		logger.Info(fmt.Sprintf("    Max open: %d", pool.MaxOpenConns))
		logger.Info(fmt.Sprintf("    Open: %d", pool.OpenConnections))
		logger.Info(fmt.Sprintf("    In use: %d", pool.InUse))
		logger.Info(fmt.Sprintf("    Idle: %d", pool.Idle))
		logger.Info(fmt.Sprintf("    Wait count: %d", pool.WaitCount))
	}

	// Get connection utilization / 연결 사용률 가져오기
	utilization := db.GetConnectionUtilization()
	logger.Info("")
	logger.Info("Connection utilization:")
	logger.Info("연결 사용률:")
	for i, util := range utilization {
		logger.Info(fmt.Sprintf("  Pool %d: %.2f%%", i, util))
	}

	logger.Info("")
	return nil
}

// example31GetTables demonstrates GetTables - list all tables
// example31GetTables는 GetTables를 시연합니다 - 모든 테이블 나열
func example31GetTables(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 31: GetTables - List all tables in database")
	logger.Info("예제 31: GetTables - 데이터베이스의 모든 테이블 나열")
	logger.Info("========================================")

	// Get all tables / 모든 테이블 가져오기
	tables, err := db.GetTables(ctx)
	if err != nil {
		return fmt.Errorf("GetTables failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d tables:", len(tables)))
	logger.Info(fmt.Sprintf("%d개의 테이블을 찾았습니다:", len(tables)))
	for i, table := range tables {
		logger.Info(fmt.Sprintf("  %d. %s (Engine: %s, Rows: %d)",
			i+1, table.Name, table.Engine, table.Rows))
	}

	logger.Info("")
	return nil
}

// example32InspectTable demonstrates InspectTable - comprehensive table inspection
// example32InspectTable은 InspectTable을 시연합니다 - 포괄적인 테이블 검사
func example32InspectTable(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 32: InspectTable - Comprehensive table inspection")
	logger.Info("예제 32: InspectTable - 포괄적인 테이블 검사")
	logger.Info("========================================")

	// Inspect users table / users 테이블 검사
	inspection, err := db.InspectTable(ctx, "users")
	if err != nil {
		return fmt.Errorf("InspectTable failed: %w", err)
	}

	logger.Info(fmt.Sprintf("Table: %s", inspection.Info.Name))
	logger.Info(fmt.Sprintf("Engine: %s", inspection.Info.Engine))
	logger.Info(fmt.Sprintf("Rows: %d", inspection.Info.Rows))
	logger.Info(fmt.Sprintf("Size: %.2f KB", float64(inspection.Size)/1024))
	logger.Info("")

	logger.Info(fmt.Sprintf("Columns (%d):", len(inspection.Columns)))
	logger.Info(fmt.Sprintf("컬럼 (%d개):", len(inspection.Columns)))
	for i, col := range inspection.Columns {
		if i >= 5 {
			logger.Info(fmt.Sprintf("  ... and %d more columns", len(inspection.Columns)-5))
			break
		}
		nullable := "NOT NULL"
		if col.Nullable {
			nullable = "NULL"
		}
		logger.Info(fmt.Sprintf("  - %s %s %s", col.Name, col.Type, nullable))
	}

	logger.Info("")
	logger.Info(fmt.Sprintf("Indexes (%d):", len(inspection.Indexes)))
	logger.Info(fmt.Sprintf("인덱스 (%d개):", len(inspection.Indexes)))
	for _, idx := range inspection.Indexes {
		uniqueStr := ""
		if idx.Unique {
			uniqueStr = "UNIQUE "
		}
		logger.Info(fmt.Sprintf("  - %s%s (%s)", uniqueStr, idx.Name, strings.Join(idx.Columns, ", ")))
	}

	logger.Info("")
	return nil
}

// example33CreateTable demonstrates CreateTable - create new table
// example33CreateTable은 CreateTable을 시연합니다 - 새 테이블 생성
func example33CreateTable(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 33: CreateTable - Create new test table")
	logger.Info("예제 33: CreateTable - 새 테스트 테이블 생성")
	logger.Info("========================================")

	// Drop table if exists / 테이블이 존재하면 삭제
	db.DropTable(ctx, "test_migration", true)

	// Create test table / 테스트 테이블 생성
	schema := `
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		status ENUM('active', 'inactive') DEFAULT 'active',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	`

	logger.Info("Creating test_migration table...")
	logger.Info("test_migration 테이블 생성 중...")

	err := db.CreateTable(ctx, "test_migration", schema)
	if err != nil {
		return fmt.Errorf("CreateTable failed: %w", err)
	}

	logger.Info("✅ Table created successfully")
	logger.Info("✅ 테이블이 성공적으로 생성되었습니다")

	// Verify table exists / 테이블 존재 확인
	exists, _ := db.TableExists(ctx, "test_migration")
	logger.Info(fmt.Sprintf("Table exists: %v", exists))
	logger.Info(fmt.Sprintf("테이블 존재: %v", exists))

	logger.Info("")
	return nil
}

// example34AddColumn demonstrates AddColumn and other migration operations
// example34AddColumn은 AddColumn 및 기타 마이그레이션 작업을 시연합니다
func example34AddColumn(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 34: Migration Operations - Add/Modify/Drop columns")
	logger.Info("예제 34: 마이그레이션 작업 - 컬럼 추가/수정/삭제")
	logger.Info("========================================")

	// Add column / 컬럼 추가
	logger.Info("Adding 'email' column...")
	logger.Info("'email' 컬럼 추가 중...")

	err := db.AddColumn(ctx, "test_migration", "email", "VARCHAR(255)")
	if err != nil {
		return fmt.Errorf("AddColumn failed: %w", err)
	}

	logger.Info("✅ Column added")
	logger.Info("✅ 컬럼이 추가되었습니다")

	// Modify column / 컬럼 수정
	logger.Info("")
	logger.Info("Modifying 'email' column to add UNIQUE constraint...")
	logger.Info("'email' 컬럼에 UNIQUE 제약 조건 추가 중...")

	err = db.ModifyColumn(ctx, "test_migration", "email", "VARCHAR(255) UNIQUE")
	if err != nil {
		return fmt.Errorf("ModifyColumn failed: %w", err)
	}

	logger.Info("✅ Column modified")
	logger.Info("✅ 컬럼이 수정되었습니다")

	// Add index / 인덱스 추가
	logger.Info("")
	logger.Info("Adding index on 'name' column...")
	logger.Info("'name' 컬럼에 인덱스 추가 중...")

	err = db.AddIndex(ctx, "test_migration", "idx_name", []string{"name"}, false)
	if err != nil {
		return fmt.Errorf("AddIndex failed: %w", err)
	}

	logger.Info("✅ Index added")
	logger.Info("✅ 인덱스가 추가되었습니다")

	logger.Info("")
	return nil
}

// example35ExportCSV demonstrates ExportTableToCSV - export table to CSV
// example35ExportCSV는 ExportTableToCSV를 시연합니다 - 테이블을 CSV로 내보내기
func example35ExportCSV(ctx context.Context, db *mysql.Client, logger *logging.Logger) error {
	logger.Info("========================================")
	logger.Info("Example 35: ExportCSV - Export table to CSV file")
	logger.Info("예제 35: ExportCSV - 테이블을 CSV 파일로 내보내기")
	logger.Info("========================================")

	// Create CSV file path / CSV 파일 경로 생성
	csvPath := fmt.Sprintf("logs/mysql_export/users_export_%s.csv", time.Now().Format("20060102_150405"))

	// Configure export options / 내보내기 옵션 설정
	opts := mysql.DefaultCSVExportOptions()
	opts.Columns = []string{"id", "name", "email", "age", "city"}
	opts.Where = "age > ?"
	opts.WhereArgs = []interface{}{25}
	opts.OrderBy = "age DESC"
	opts.Limit = 10

	logger.Info(fmt.Sprintf("Exporting users (age > 25) to: %s", csvPath))
	logger.Info(fmt.Sprintf("사용자 (나이 > 25)를 다음으로 내보내는 중: %s", csvPath))

	err := db.ExportTableToCSV(ctx, "users", csvPath, opts)
	if err != nil {
		return fmt.Errorf("ExportTableToCSV failed: %w", err)
	}

	logger.Info("✅ Export completed successfully")
	logger.Info("✅ 내보내기가 성공적으로 완료되었습니다")

	// Read and display first few lines / 처음 몇 줄 읽어서 표시
	content, err := os.ReadFile(csvPath)
	if err == nil {
		lines := strings.Split(string(content), "\n")
		logger.Info("")
		logger.Info("First 3 lines of CSV:")
		logger.Info("CSV의 처음 3줄:")
		for i := 0; i < min(3, len(lines)); i++ {
			if lines[i] != "" {
				logger.Info(fmt.Sprintf("  %s", lines[i]))
			}
		}
	}

	logger.Info("")
	logger.Info(fmt.Sprintf("Note: CSV file saved to %s", csvPath))
	logger.Info(fmt.Sprintf("참고: CSV 파일이 %s에 저장되었습니다", csvPath))

	logger.Info("")
	return nil
}

// Helper function for min / min을 위한 헬퍼 함수
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
