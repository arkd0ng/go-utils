package mysql

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestHelper provides helper functions for MySQL tests
// TestHelper는 MySQL 테스트를 위한 헬퍼 함수를 제공합니다
type TestHelper struct {
	startedDocker bool
	t             *testing.T
}

// NewTestHelper creates a new test helper
// NewTestHelper는 새로운 테스트 헬퍼를 생성합니다
func NewTestHelper(t *testing.T) *TestHelper {
	return &TestHelper{
		t:             t,
		startedDocker: false,
	}
}

// SetupDocker ensures Docker MySQL is running for tests
// SetupDocker는 테스트를 위해 Docker MySQL이 실행 중인지 확인합니다
func (h *TestHelper) SetupDocker() {
	// Check if Docker MySQL is already running / Docker MySQL이 이미 실행 중인지 확인
	if h.isDockerMySQLRunning() {
		h.t.Log("Docker MySQL is already running (will not stop after tests)")
		return
	}

	h.t.Log("Starting Docker MySQL for tests...")
	if err := h.startDockerMySQL(); err != nil {
		h.t.Skipf("Skipping test: Docker MySQL not available: %v", err)
		return
	}

	h.startedDocker = true
	h.t.Log("Waiting for Docker MySQL to be ready...")

	// Wait for MySQL to be ready / MySQL 준비 대기
	if err := h.waitForDockerMySQL(30 * time.Second); err != nil {
		h.stopDockerMySQL() // Clean up / 정리
		h.t.Skipf("Skipping test: Docker MySQL failed to become ready: %v", err)
		return
	}

	h.t.Log("Docker MySQL is ready!")
}

// TeardownDocker stops Docker MySQL if it was started by this helper
// TeardownDocker는 이 헬퍼에서 시작한 경우 Docker MySQL을 중지합니다
func (h *TestHelper) TeardownDocker() {
	if !h.startedDocker {
		return
	}

	h.t.Log("Stopping Docker MySQL...")
	if err := h.stopDockerMySQL(); err != nil {
		h.t.Logf("Warning: Failed to stop Docker MySQL: %v", err)
	} else {
		h.t.Log("Docker MySQL stopped successfully")
	}
}

// isDockerMySQLRunning checks if Docker MySQL container is running
// isDockerMySQLRunning은 Docker MySQL 컨테이너가 실행 중인지 확인합니다
func (h *TestHelper) isDockerMySQLRunning() bool {
	cmd := exec.Command("docker", "ps", "--filter", "name=go-utils-mysql", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(output)) == "go-utils-mysql"
}

// startDockerMySQL starts the Docker MySQL container
// startDockerMySQL은 Docker MySQL 컨테이너를 시작합니다
func (h *TestHelper) startDockerMySQL() error {
	// Get project root directory / 프로젝트 루트 디렉토리 가져오기
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	// Navigate to project root (database/mysql -> go-utils)
	// 프로젝트 루트로 이동
	projectRoot := filepath.Join(wd, "..", "..")

	// Start Docker Compose / Docker Compose 시작
	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = projectRoot
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to start docker compose: %w (output: %s)", err, string(output))
	}
	return nil
}

// stopDockerMySQL stops the Docker MySQL container
// stopDockerMySQL은 Docker MySQL 컨테이너를 중지합니다
func (h *TestHelper) stopDockerMySQL() error {
	// Get project root directory / 프로젝트 루트 디렉토리 가져오기
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	// Navigate to project root
	// 프로젝트 루트로 이동
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

// waitForDockerMySQL waits for Docker MySQL to be ready
// waitForDockerMySQL은 Docker MySQL이 준비될 때까지 대기합니다
func (h *TestHelper) waitForDockerMySQL(timeout time.Duration) error {
	start := time.Now()
	for {
		// Try to connect to MySQL / MySQL 연결 시도
		cmd := exec.Command("docker", "exec", "go-utils-mysql",
			"mysqladmin", "ping", "-h", "localhost", "-u", "root", "-prootpassword")
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

// GetTestDSN returns a DSN for testing
// GetTestDSN은 테스트용 DSN을 반환합니다
func (h *TestHelper) GetTestDSN() string {
	return "root:rootpassword@tcp(localhost:3306)/testdb?parseTime=true&charset=utf8mb4&loc=Local"
}

// CreateTestClient creates a MySQL client for testing
// CreateTestClient는 테스트용 MySQL 클라이언트를 생성합니다
func (h *TestHelper) CreateTestClient() (*Client, error) {
	return New(
		WithDSN(h.GetTestDSN()),
		WithMaxOpenConns(10),
		WithMaxIdleConns(5),
		WithConnMaxLifetime(5*time.Minute),
	)
}

// CleanupTestData cleans up test data from the database
// CleanupTestData는 데이터베이스에서 테스트 데이터를 정리합니다
func (h *TestHelper) CleanupTestData(client *Client) {
	// Delete test users created during tests / 테스트 중 생성된 테스트 사용자 삭제
	client.Delete("users", "email LIKE ?", "%@test.example.com")
	client.Delete("users", "email LIKE ?", "test_%@example.com")
}
