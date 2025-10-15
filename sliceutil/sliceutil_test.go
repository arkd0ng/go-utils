package sliceutil

import (
	"testing"
)

// TestPackageVersion tests the package version constant.
// TestPackageVersion은 패키지 버전 상수를 테스트합니다.
func TestPackageVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version constant should not be empty")
	}

	expected := "1.7.024"
	if Version != expected {
		t.Errorf("Version = %v, want %v", Version, expected)
	}
}

// TestPackageImport ensures the package can be imported without errors.
// TestPackageImport는 패키지를 오류 없이 가져올 수 있는지 확인합니다.
func TestPackageImport(t *testing.T) {
	// If we can run this test, the package was imported successfully
	// 이 테스트를 실행할 수 있다면 패키지를 성공적으로 가져온 것입니다
	t.Log("Package sliceutil imported successfully")
}
