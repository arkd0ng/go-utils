// +build integration

package mysql

import (
	"context"
	"testing"
)

// TestIntegrationSelectAll tests SelectAll with real database
// TestIntegrationSelectAll은 실제 데이터베이스로 SelectAll을 테스트합니다
//
// Run with: go test -tags=integration -v
func TestIntegrationSelectAll(t *testing.T) {
	// Setup Docker MySQL / Docker MySQL 설정
	helper := NewTestHelper(t)
	helper.SetupDocker()
	defer helper.TeardownDocker()

	// Create test client / 테스트 클라이언트 생성
	client, err := helper.CreateTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	defer helper.CleanupTestData(client)

	ctx := context.Background()

	// Test SelectAll / SelectAll 테스트
	users, err := client.SelectAll("users")
	if err != nil {
		t.Fatalf("SelectAll failed: %v", err)
	}

	if len(users) == 0 {
		t.Error("Expected at least one user, got 0")
	}

	t.Logf("Found %d users", len(users))
}

// TestIntegrationInsertAndSelect tests Insert and Select operations
// TestIntegrationInsertAndSelect는 Insert 및 Select 작업을 테스트합니다
//
// Run with: go test -tags=integration -v
func TestIntegrationInsertAndSelect(t *testing.T) {
	// Setup Docker MySQL / Docker MySQL 설정
	helper := NewTestHelper(t)
	helper.SetupDocker()
	defer helper.TeardownDocker()

	// Create test client / 테스트 클라이언트 생성
	client, err := helper.CreateTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	defer helper.CleanupTestData(client)

	ctx := context.Background()

	// Insert test user / 테스트 사용자 삽입
	data := map[string]interface{}{
		"name":  "Test User",
		"email": "test@test.example.com",
		"age":   25,
		"city":  "Seoul",
	}

	result, err := client.Insert(ctx, "users", data)
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	insertedID, _ := result.LastInsertId()
	t.Logf("Inserted user with ID: %d", insertedID)

	// Select the inserted user / 삽입된 사용자 선택
	user, err := client.SelectOne("users", "id = ?", insertedID)
	if err != nil {
		t.Fatalf("SelectOne failed: %v", err)
	}

	if user["email"] != "test@test.example.com" {
		t.Errorf("Expected email 'test@test.example.com', got '%v'", user["email"])
	}

	t.Logf("Successfully retrieved user: %v", user["name"])
}
