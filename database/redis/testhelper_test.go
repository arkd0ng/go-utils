package redis

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

const testRedisAddr = "localhost:6379"

var (
	redisSetupOnce      sync.Once
	redisSetupErr       error
	redisStartedByTests bool
	repoRootOnce        sync.Once
	repoRootDir         string
	repoRootErr         error
	testKeyCounter      uint64
)

func TestMain(m *testing.M) {
	code := m.Run()
	if redisStartedByTests {
		if err := stopDockerRedis(); err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to stop docker redis: %v\n", err)
		}
	}
	os.Exit(code)
}

func ensureRedisRunning(t *testing.T) {
	t.Helper()

	redisSetupOnce.Do(func() {
		if err := pingRedis(); err == nil {
			return
		}

		if err := startDockerRedis(); err != nil {
			redisSetupErr = fmt.Errorf("failed to start docker redis: %w", err)
			return
		}
		redisStartedByTests = true

		if err := waitForRedis(30 * time.Second); err != nil {
			redisSetupErr = fmt.Errorf("redis did not become ready: %w", err)
		}
	})

	if redisSetupErr != nil {
		t.Skipf("Redis not available for tests: %v", redisSetupErr)
	}
}

func newTestClient(t *testing.T, opts ...Option) *Client {
	t.Helper()

	ensureRedisRunning(t)

	options := append([]Option{WithAddr(testRedisAddr)}, opts...)
	client, err := New(options...)
	if err != nil {
		t.Fatalf("failed to create redis test client: %v", err)
	}

	flushDB(t, client)

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := client.rdb.FlushDB(ctx).Err(); err != nil {
			t.Logf("warning: failed to flush redis db during cleanup: %v", err)
		}
		if err := client.Close(); err != nil {
			t.Logf("warning: failed to close redis client: %v", err)
		}
	})

	return client
}

func flushDB(t *testing.T, client *Client) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := client.rdb.FlushDB(ctx).Err(); err != nil {
		t.Fatalf("failed to flush redis db: %v", err)
	}
}

func testKey(prefix string) string {
	counter := atomic.AddUint64(&testKeyCounter, 1)
	return fmt.Sprintf("test:%s:%d", prefix, counter)
}

func pingRedis() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:        testRedisAddr,
		DialTimeout: 500 * time.Millisecond,
	})
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	return rdb.Ping(ctx).Err()
}

func startDockerRedis() error {
	root, err := getRepoRoot()
	if err != nil {
		return err
	}

	// docker-compose.yml is now in .docker/ directory
	dockerDir := filepath.Join(root, ".docker")
	cmd := exec.Command("docker", "compose", "up", "-d", "redis")
	cmd.Dir = dockerDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w (output: %s)", err, string(output))
	}

	return nil
}

func waitForRedis(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if err := pingRedis(); err == nil {
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("timeout waiting for redis")
}

func stopDockerRedis() error {
	root, err := getRepoRoot()
	if err != nil {
		return err
	}

	// docker-compose.yml is now in .docker/ directory
	dockerDir := filepath.Join(root, ".docker")

	cmd := exec.Command("docker", "compose", "stop", "redis")
	cmd.Dir = dockerDir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w (output: %s)", err, string(output))
	}

	cmd = exec.Command("docker", "compose", "rm", "-f", "redis")
	cmd.Dir = dockerDir
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
