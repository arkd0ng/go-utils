package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	testMySQLDSN          = "root:rootpassword@tcp(localhost:3306)/testdb?parseTime=true&charset=utf8mb4&loc=Local"
	testMySQLDialTimeout  = 2 * time.Second
	testMySQLQueryTimeout = 5 * time.Second
)

var (
	mysqlSetupOnce      sync.Once
	mysqlSetupErr       error
	mysqlStartedByTests bool
	repoRootOnce        sync.Once
	repoRootDir         string
	repoRootErr         error
)

func TestMain(m *testing.M) {
	ensureMySQLRunning()

	code := m.Run()

	if mysqlStartedByTests {
		if err := stopDockerMySQL(); err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to stop docker mysql: %v\n", err)
		}
	}

	os.Exit(code)
}

func ensureMySQLRunning() {
	mysqlSetupOnce.Do(func() {
		if err := pingMySQL(testMySQLDialTimeout); err == nil {
			return
		}

		if err := startDockerMySQL(); err != nil {
			mysqlSetupErr = fmt.Errorf("failed to start docker mysql: %w", err)
			return
		}
		mysqlStartedByTests = true

		if err := waitForMySQL(60 * time.Second); err != nil {
			mysqlSetupErr = fmt.Errorf("mysql did not become ready: %w", err)
		}
	})
}

func newTestClient(t *testing.T, opts ...Option) *Client {
	t.Helper()

	if mysqlSetupErr != nil {
		t.Skipf("MySQL not available for tests: %v", mysqlSetupErr)
	}

	options := append([]Option{WithDSN(testMySQLDSN)}, opts...)
	client, err := New(options...)
	if err != nil {
		t.Fatalf("failed to create mysql test client: %v", err)
	}

	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Logf("warning: failed to close mysql client: %v", err)
		}
	})

	return client
}

func openRawTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("mysql", testMySQLDSN)
	if err != nil {
		t.Fatalf("failed to open raw mysql connection: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), testMySQLDialTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("failed to ping raw mysql connection: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Logf("warning: failed to close raw mysql connection: %v", err)
		}
	})

	return db
}

func resetTable(t *testing.T, table string) {
	t.Helper()

	db := openRawTestDB(t)
	ctx, cancel := context.WithTimeout(context.Background(), testMySQLQueryTimeout)
	defer cancel()

	if _, err := db.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS = 0"); err != nil {
		t.Fatalf("failed to disable foreign key checks: %v", err)
	}

	if _, err := db.ExecContext(ctx, fmt.Sprintf("TRUNCATE TABLE %s", table)); err != nil {
		t.Fatalf("failed to truncate table %s: %v", table, err)
	}

	if _, err := db.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS = 1"); err != nil {
		t.Fatalf("failed to enable foreign key checks: %v", err)
	}
}

func pingMySQL(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db, err := sql.Open("mysql", testMySQLDSN)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.PingContext(ctx)
}

func startDockerMySQL() error {
	root, err := getRepoRoot()
	if err != nil {
		return err
	}

	cmd := exec.Command("docker", "compose", "up", "-d", "mysql")
	cmd.Dir = root
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w (output: %s)", err, string(output))
	}

	return nil
}

func waitForMySQL(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if err := pingMySQL(2 * time.Second); err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("timeout waiting for mysql")
}

func stopDockerMySQL() error {
	root, err := getRepoRoot()
	if err != nil {
		return err
	}

	cmd := exec.Command("docker", "compose", "stop", "mysql")
	cmd.Dir = root
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w (output: %s)", err, string(output))
	}

	cmd = exec.Command("docker", "compose", "rm", "-f", "mysql")
	cmd.Dir = root
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w (output: %s)", err, string(output))
	}

	return nil
}

func getRepoRoot() (string, error) {
	repoRootOnce.Do(func() {
		wd, err := os.Getwd()
		if err != nil {
			repoRootErr = err
			return
		}
		repoRootDir = filepath.Clean(filepath.Join(wd, "..", ".."))
	})
	return repoRootDir, repoRootErr
}
