package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/arkd0ng/go-utils/database/redis"
	"github.com/arkd0ng/go-utils/logging"
	"gopkg.in/yaml.v3"
)

// RedisConfig represents Redis configuration from YAML
// RedisConfig는 YAML에서 읽은 Redis 설정을 나타냅니다
type RedisConfig struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	DB           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
}

// DatabaseConfig represents database configuration from YAML
// DatabaseConfig는 YAML에서 읽은 데이터베이스 설정을 나타냅니다
type DatabaseConfig struct {
	Redis RedisConfig `yaml:"redis"`
}

func main() {
	// Create results directories if they don't exist / 결과 디렉토리가 없다면 새롭게 생성
	if err := os.MkdirAll("./results/logs", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create logs directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(fmt.Sprintf("./results/logs/redis_example_%s.log", time.Now().Format("20060102_150405"))),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Banner("Redis Package Examples", "go-utils/database/redis")

	// Load database configuration / 데이터베이스 설정 로드
	logger.Info("Loading database configuration from cfg/database-redis.yaml")
	logger.Info("cfg/database-redis.yaml에서 데이터베이스 설정 로드 중")
	config, err := loadRedisConfig()
	if err != nil {
		logger.Error("Failed to load Redis config", "error", err)
		os.Exit(1)
	}

	logger.Info("Configuration loaded successfully",
		"addr", config.Redis.Addr,
		"db", config.Redis.DB)

	// Check if Docker Redis is running / Docker Redis 실행 여부 확인
	wasRunning := isDockerRedisRunning()

	if !wasRunning {
		logger.Info("Docker Redis is not running, starting container...")
		logger.Info("Docker Redis가 실행 중이 아닙니다. 컨테이너를 시작합니다...")
		if err := startDockerRedis(); err != nil {
			logger.Error("Failed to start Docker Redis", "error", err)
			logger.Info("")
			logger.Info("Please ensure Docker is installed and running:")
			logger.Info("Docker가 설치되어 실행 중인지 확인하세요:")
			logger.Info("  1. Install Docker Desktop: https://www.docker.com/products/docker-desktop")
			logger.Info("  2. Start Docker Desktop")
			logger.Info("  3. Run: ./scripts/docker-redis-start.sh")
			os.Exit(1)
		}

		// Wait for Redis to be ready / Redis 준비 대기
		logger.Info("Waiting for Docker Redis to be ready...")
		logger.Info("Docker Redis 준비 중...")
		if err := waitForDockerRedis(30 * time.Second); err != nil {
			logger.Error("Docker Redis failed to become ready", "error", err)
			// Clean up - stop Redis if we started it / 정리 - 시작한 경우 Redis 중지
			stopDockerRedis()
			os.Exit(1)
		}
		logger.Info("Docker Redis is ready!")
		logger.Info("Docker Redis 준비 완료!")
	} else {
		logger.Info("Docker Redis is already running")
		logger.Info("Docker Redis가 이미 실행 중입니다")
	}

	// Create Redis client / Redis 클라이언트 생성
	logger.Info("Connecting to Redis...")
	logger.Info("Redis에 연결 중...")

	rdb, err := redis.New(
		redis.WithAddr(config.Redis.Addr),
		redis.WithPassword(config.Redis.Password),
		redis.WithDB(config.Redis.DB),
		redis.WithPoolSize(config.Redis.PoolSize),
		redis.WithMinIdleConns(config.Redis.MinIdleConns),
	)
	if err != nil {
		logger.Error("Failed to connect to Redis", "error", err)
		// Stop Redis if we started it / 시작한 경우 Redis 중지
		if !wasRunning {
			stopDockerRedis()
		}
		os.Exit(1)
	}
	defer rdb.Close()

	logger.Info("Connected to Redis successfully")
	logger.Info("Redis 연결 성공")
	logger.Info("")

	ctx := context.Background()

	// Run examples / 예제 실행
	logger.Info("=== Running Redis Examples ===")
	logger.Info("=== Redis 예제 실행 중 ===")
	logger.Info("")

	// 1. String Operations / 문자열 작업
	logger.Info("--- Example 1: String Operations ---")
	logger.Info("--- 예제 1: 문자열 작업 ---")
	stringOperations(ctx, rdb, logger)

	// 2. Hash Operations / 해시 작업
	logger.Info("")
	logger.Info("--- Example 2: Hash Operations ---")
	logger.Info("--- 예제 2: 해시 작업 ---")
	hashOperations(ctx, rdb, logger)

	// 3. List Operations / 리스트 작업
	logger.Info("")
	logger.Info("--- Example 3: List Operations ---")
	logger.Info("--- 예제 3: 리스트 작업 ---")
	listOperations(ctx, rdb, logger)

	// 4. Set Operations / 집합 작업
	logger.Info("")
	logger.Info("--- Example 4: Set Operations ---")
	logger.Info("--- 예제 4: 집합 작업 ---")
	setOperations(ctx, rdb, logger)

	// 5. Sorted Set Operations / 정렬 집합 작업
	logger.Info("")
	logger.Info("--- Example 5: Sorted Set Operations ---")
	logger.Info("--- 예제 5: 정렬 집합 작업 ---")
	sortedSetOperations(ctx, rdb, logger)

	// 6. Key Operations / 키 작업
	logger.Info("")
	logger.Info("--- Example 6: Key Operations ---")
	logger.Info("--- 예제 6: 키 작업 ---")
	keyOperations(ctx, rdb, logger)

	// 7. Pipeline Operations / 파이프라인 작업
	logger.Info("")
	logger.Info("--- Example 7: Pipeline Operations ---")
	logger.Info("--- 예제 7: 파이프라인 작업 ---")
	pipelineOperations(ctx, rdb, logger)

	// 8. Transaction Operations / 트랜잭션 작업
	logger.Info("")
	logger.Info("--- Example 8: Transaction Operations ---")
	logger.Info("--- 예제 8: 트랜잭션 작업 ---")
	transactionOperations(ctx, rdb, logger)

	// Summary / 요약
	logger.Info("")
	logger.Info("=== All examples completed successfully! ===")
	logger.Info("=== 모든 예제가 성공적으로 완료되었습니다! ===")

	// Cleanup / 정리
	if !wasRunning {
		logger.Info("")
		logger.Info("Stopping Docker Redis (was started by examples)...")
		logger.Info("Docker Redis 중지 중 (예제에서 시작함)...")
		stopDockerRedis()
		logger.Info("Docker Redis stopped")
		logger.Info("Docker Redis 중지됨")
	}

	logger.Info("")
	logger.Info("Logs saved to: %s", fmt.Sprintf("./results/logs/redis_example_%s.log", time.Now().Format("20060102")))
	logger.Info("로그 저장 위치: %s", fmt.Sprintf("./results/logs/redis_example_%s.log", time.Now().Format("20060102")))
}

// stringOperations demonstrates string operations
// stringOperations는 문자열 작업을 시연합니다
func stringOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	// Set / 설정
	err := rdb.Set(ctx, "name", "John Doe")
	if err != nil {
		logger.Error("SET failed", "error", err)
		return
	}
	logger.Info("SET name = 'John Doe'")

	// Get / 가져오기
	name, err := rdb.Get(ctx, "name")
	if err != nil {
		logger.Error("GET failed", "error", err)
		return
	}
	logger.Info("GET name", "value", name)

	// Set with expiration / 만료와 함께 설정
	err = rdb.Set(ctx, "session", "abc123", 10*time.Second)
	if err != nil {
		logger.Error("SET with expiration failed", "error", err)
		return
	}
	logger.Info("SET session = 'abc123' with 10s expiration")

	// Increment / 증가
	err = rdb.Set(ctx, "counter", "0")
	if err != nil {
		logger.Error("SET counter failed", "error", err)
		return
	}
	count, err := rdb.Incr(ctx, "counter")
	if err != nil {
		logger.Error("INCR failed", "error", err)
		return
	}
	logger.Info("INCR counter", "value", count)

	// Multiple set / 다중 설정
	err = rdb.MSet(ctx, map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	if err != nil {
		logger.Error("MSET failed", "error", err)
		return
	}
	logger.Info("MSET key1, key2, key3")

	// Multiple get / 다중 가져오기
	values, err := rdb.MGet(ctx, "key1", "key2", "key3")
	if err != nil {
		logger.Error("MGET failed", "error", err)
		return
	}
	logger.Info("MGET", "values", values)

	// Cleanup / 정리
	rdb.Del(ctx, "name", "session", "counter", "key1", "key2", "key3")
	logger.Debug("Cleaned up string operation keys")
}

// hashOperations demonstrates hash operations
// hashOperations는 해시 작업을 시연합니다
func hashOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	// Set hash field / 해시 필드 설정
	err := rdb.HSet(ctx, "user:1001", "name", "Alice")
	if err != nil {
		logger.Error("HSET failed", "error", err)
		return
	}
	logger.Info("HSET user:1001 name = 'Alice'")

	// Set multiple fields / 여러 필드 설정
	err = rdb.HSetMap(ctx, "user:1001", map[string]interface{}{
		"email": "alice@example.com",
		"age":   "25",
		"city":  "Seoul",
	})
	if err != nil {
		logger.Error("HMSET failed", "error", err)
		return
	}
	logger.Info("HMSET user:1001 email, age, city")

	// Get hash field / 해시 필드 가져오기
	name, err := rdb.HGet(ctx, "user:1001", "name")
	if err != nil {
		logger.Error("HGET failed", "error", err)
		return
	}
	logger.Info("HGET user:1001 name", "value", name)

	// Get all fields / 모든 필드 가져오기
	fields, err := rdb.HGetAll(ctx, "user:1001")
	if err != nil {
		logger.Error("HGETALL failed", "error", err)
		return
	}
	logger.Info("HGETALL user:1001", "fields", fields)

	// Increment hash field / 해시 필드 증가
	newAge, err := rdb.HIncrBy(ctx, "user:1001", "age", 1)
	if err != nil {
		logger.Error("HINCRBY failed", "error", err)
		return
	}
	logger.Info("HINCRBY user:1001 age 1", "new_value", newAge)

	// Cleanup / 정리
	rdb.Del(ctx, "user:1001")
	logger.Debug("Cleaned up hash operation keys")
}

// listOperations demonstrates list operations
// listOperations는 리스트 작업을 시연합니다
func listOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	// Push to list / 리스트에 추가
	err := rdb.RPush(ctx, "queue", "task1", "task2", "task3")
	if err != nil {
		logger.Error("RPUSH failed", "error", err)
		return
	}
	logger.Info("RPUSH queue = ['task1', 'task2', 'task3']")

	// Get list length / 리스트 길이 가져오기
	length, err := rdb.LLen(ctx, "queue")
	if err != nil {
		logger.Error("LLEN failed", "error", err)
		return
	}
	logger.Info("LLEN queue", "length", length)

	// Get range / 범위 가져오기
	items, err := rdb.LRange(ctx, "queue", 0, -1)
	if err != nil {
		logger.Error("LRANGE failed", "error", err)
		return
	}
	logger.Info("LRANGE queue 0 -1", "items", items)

	// Pop from list / 리스트에서 제거
	item, err := rdb.LPop(ctx, "queue")
	if err != nil {
		logger.Error("LPOP failed", "error", err)
		return
	}
	logger.Info("LPOP queue", "item", item)

	// Cleanup / 정리
	rdb.Del(ctx, "queue")
	logger.Debug("Cleaned up list operation keys")
}

// setOperations demonstrates set operations
// setOperations는 집합 작업을 시연합니다
func setOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	// Add to set / 집합에 추가
	err := rdb.SAdd(ctx, "languages", "Go", "Python", "JavaScript", "Rust")
	if err != nil {
		logger.Error("SADD failed", "error", err)
		return
	}
	logger.Info("SADD languages = ['Go', 'Python', 'JavaScript', 'Rust']")

	// Check membership / 멤버 확인
	exists, err := rdb.SIsMember(ctx, "languages", "Go")
	if err != nil {
		logger.Error("SISMEMBER failed", "error", err)
		return
	}
	logger.Info("SISMEMBER languages 'Go'", "exists", exists)

	// Get all members / 모든 멤버 가져오기
	members, err := rdb.SMembers(ctx, "languages")
	if err != nil {
		logger.Error("SMEMBERS failed", "error", err)
		return
	}
	logger.Info("SMEMBERS languages", "members", members)

	// Get cardinality / 크기 가져오기
	size, err := rdb.SCard(ctx, "languages")
	if err != nil {
		logger.Error("SCARD failed", "error", err)
		return
	}
	logger.Info("SCARD languages", "size", size)

	// Set operations / 집합 작업
	err = rdb.SAdd(ctx, "backend", "Go", "Python", "Java")
	if err != nil {
		logger.Error("SADD backend failed", "error", err)
		return
	}
	union, err := rdb.SUnion(ctx, "languages", "backend")
	if err != nil {
		logger.Error("SUNION failed", "error", err)
		return
	}
	logger.Info("SUNION languages backend", "union", union)

	// Cleanup / 정리
	rdb.Del(ctx, "languages", "backend")
	logger.Debug("Cleaned up set operation keys")
}

// sortedSetOperations demonstrates sorted set operations
// sortedSetOperations는 정렬 집합 작업을 시연합니다
func sortedSetOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	// Add members with scores / 점수와 함께 멤버 추가
	err := rdb.ZAddMultiple(ctx, "leaderboard", map[string]float64{
		"Alice":   100,
		"Bob":     85,
		"Charlie": 95,
		"David":   90,
	})
	if err != nil {
		logger.Error("ZADD failed", "error", err)
		return
	}
	logger.Info("ZADD leaderboard with scores")

	// Get range (ascending) / 범위 가져오기 (오름차순)
	players, err := rdb.ZRange(ctx, "leaderboard", 0, -1)
	if err != nil {
		logger.Error("ZRANGE failed", "error", err)
		return
	}
	logger.Info("ZRANGE leaderboard 0 -1", "players", players)

	// Get range (descending) / 범위 가져오기 (내림차순)
	topPlayers, err := rdb.ZRevRange(ctx, "leaderboard", 0, 2)
	if err != nil {
		logger.Error("ZREVRANGE failed", "error", err)
		return
	}
	logger.Info("Top 3 players", "players", topPlayers)

	// Get score / 점수 가져오기
	score, err := rdb.ZScore(ctx, "leaderboard", "Alice")
	if err != nil {
		logger.Error("ZSCORE failed", "error", err)
		return
	}
	logger.Info("ZSCORE leaderboard 'Alice'", "score", score)

	// Increment score / 점수 증가
	newScore, err := rdb.ZIncrBy(ctx, "leaderboard", 5, "Bob")
	if err != nil {
		logger.Error("ZINCRBY failed", "error", err)
		return
	}
	logger.Info("ZINCRBY leaderboard 'Bob' 5", "new_score", newScore)

	// Cleanup / 정리
	rdb.Del(ctx, "leaderboard")
	logger.Debug("Cleaned up sorted set operation keys")
}

// keyOperations demonstrates key operations
// keyOperations는 키 작업을 시연합니다
func keyOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	// Set keys / 키 설정
	err := rdb.Set(ctx, "app:config:debug", "true")
	if err != nil {
		logger.Error("SET failed", "error", err)
		return
	}
	err = rdb.Set(ctx, "app:config:timeout", "30")
	if err != nil {
		logger.Error("SET failed", "error", err)
		return
	}
	err = rdb.Set(ctx, "app:user:1001", "Alice")
	if err != nil {
		logger.Error("SET failed", "error", err)
		return
	}
	logger.Info("SET multiple app:* keys")

	// Find keys by pattern / 패턴으로 키 찾기
	keys, err := rdb.Keys(ctx, "app:config:*")
	if err != nil {
		logger.Error("KEYS failed", "error", err)
		return
	}
	logger.Info("KEYS app:config:*", "keys", keys)

	// Check existence / 존재 확인
	count, err := rdb.Exists(ctx, "app:config:debug", "app:config:timeout")
	if err != nil {
		logger.Error("EXISTS failed", "error", err)
		return
	}
	logger.Info("EXISTS app:config:*", "count", count)

	// Set expiration / 만료 설정
	err = rdb.Expire(ctx, "app:user:1001", 60*time.Second)
	if err != nil {
		logger.Error("EXPIRE failed", "error", err)
		return
	}
	ttl, err := rdb.TTL(ctx, "app:user:1001")
	if err != nil {
		logger.Error("TTL failed", "error", err)
		return
	}
	logger.Info("TTL app:user:1001", "ttl", ttl)

	// Delete keys / 키 삭제
	rdb.Del(ctx, "app:config:debug", "app:config:timeout", "app:user:1001")
	logger.Debug("Cleaned up key operation keys")
}

// pipelineOperations demonstrates pipeline operations
// pipelineOperations는 파이프라인 작업을 시연합니다
func pipelineOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	logger.Info("Executing multiple commands in pipeline...")
	logger.Info("파이프라인에서 여러 명령 실행 중...")

	// Execute pipeline / 파이프라인 실행
	err := rdb.Pipeline(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, "batch:1", "value1", 0)
		pipe.Set(ctx, "batch:2", "value2", 0)
		pipe.Set(ctx, "batch:3", "value3", 0)
		pipe.Incr(ctx, "batch:counter")
		pipe.SAdd(ctx, "batch:set", "member1", "member2")
		return nil
	})

	if err != nil {
		logger.Error("Pipeline execution failed", "error", err)
		return
	}

	logger.Info("Pipeline executed successfully")
	logger.Info("파이프라인 성공적으로 실행됨")

	// Verify results / 결과 확인
	val1, _ := rdb.Get(ctx, "batch:1")
	counter, _ := rdb.Get(ctx, "batch:counter")
	logger.Info("Pipeline results", "batch:1", val1, "batch:counter", counter)

	// Cleanup / 정리
	rdb.Del(ctx, "batch:1", "batch:2", "batch:3", "batch:counter", "batch:set")
	logger.Debug("Cleaned up pipeline operation keys")
}

// transactionOperations demonstrates transaction operations
// transactionOperations는 트랜잭션 작업을 시연합니다
func transactionOperations(ctx context.Context, rdb *redis.Client, logger *logging.Logger) {
	logger.Info("Executing transaction with optimistic locking...")
	logger.Info("낙관적 잠금을 사용한 트랜잭션 실행 중...")

	// Initialize counter / 카운터 초기화
	err := rdb.Set(ctx, "tx:counter", "10")
	if err != nil {
		logger.Error("SET tx:counter failed", "error", err)
		return
	}

	// Execute transaction / 트랜잭션 실행
	err = rdb.Transaction(ctx, func(tx *redis.Tx) error {
		// Get current value / 현재 값 가져오기
		val, err := tx.Get(ctx, "tx:counter")
		if err != nil {
			return err
		}

		logger.Info("Current counter value", "value", val)

		// Execute commands atomically / 명령을 원자적으로 실행
		return tx.Exec(ctx, func(pipe redis.Pipeliner) error {
			pipe.Incr(ctx, "tx:counter")
			pipe.Set(ctx, "tx:last_update", time.Now().Unix(), 0)
			return nil
		})
	}, "tx:counter") // Watch the counter key / counter 키 감시

	if err != nil {
		logger.Error("Transaction failed", "error", err)
		return
	}

	// Verify result / 결과 확인
	newVal, _ := rdb.Get(ctx, "tx:counter")
	logger.Info("New counter value", "value", newVal)
	logger.Info("Transaction completed successfully")
	logger.Info("트랜잭션 성공적으로 완료됨")

	// Cleanup / 정리
	rdb.Del(ctx, "tx:counter", "tx:last_update")
	logger.Debug("Cleaned up transaction operation keys")
}

// loadRedisConfig loads Redis configuration from YAML file
// loadRedisConfig는 YAML 파일에서 Redis 설정을 로드합니다
func loadRedisConfig() (*DatabaseConfig, error) {
	data, err := os.ReadFile("../../cfg/database-redis.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config DatabaseConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// isDockerRedisRunning checks if Docker Redis container is running
// isDockerRedisRunning은 Docker Redis 컨테이너가 실행 중인지 확인합니다
func isDockerRedisRunning() bool {
	cmd := exec.Command("docker", "ps", "--filter", "name=go-utils-redis", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "go-utils-redis"
}

// startDockerRedis starts Docker Redis container
// startDockerRedis는 Docker Redis 컨테이너를 시작합니다
func startDockerRedis() error {
	// Try using the start script first / 먼저 시작 스크립트 사용 시도
	cmd := exec.Command("../../scripts/docker-redis-start.sh")
	if err := cmd.Run(); err == nil {
		return nil
	}

	// Fallback to docker compose / docker compose로 폴백
	cmd = exec.Command("docker", "compose", "up", "-d", "redis")
	cmd.Dir = "../.."
	return cmd.Run()
}

// stopDockerRedis stops Docker Redis container
// stopDockerRedis는 Docker Redis 컨테이너를 중지합니다
func stopDockerRedis() {
	// Try using the stop script first / 먼저 중지 스크립트 사용 시도
	cmd := exec.Command("../../scripts/docker-redis-stop.sh")
	if err := cmd.Run(); err == nil {
		return
	}

	// Fallback to docker compose / docker compose로 폴백
	cmd = exec.Command("docker", "compose", "down", "redis")
	cmd.Dir = "../.."
	cmd.Run()
}

// waitForDockerRedis waits for Docker Redis to be ready
// waitForDockerRedis는 Docker Redis가 준비될 때까지 대기합니다
func waitForDockerRedis(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		// Try to connect / 연결 시도
		cmd := exec.Command("docker", "exec", "go-utils-redis", "redis-cli", "ping")
		if output, err := cmd.Output(); err == nil {
			if strings.TrimSpace(string(output)) == "PONG" {
				return nil
			}
		}

		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("timeout waiting for Redis to be ready")
}
