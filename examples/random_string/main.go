package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/random"
)

func main() {
	// Setup log file with backup management / 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/random-example.log"

	// Check if previous log file exists / 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/random-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file / 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
			}
		}

		// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/random-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time / 수정 시간으로 정렬
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			var files []fileInfo
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			// Sort oldest first / 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5 / 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename / 고정 파일명으로 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Banner("Random String Package Examples", "go-utils/random")

	logger.Info("=== Random String Generation Examples ===")
	logger.Info("=== 랜덤 문자열 생성 예제 ===")
	logger.Info("")

	// Example 1: Letters only
	// 예제 1: 알파벳만
	logger.Info("1. Letters only (8-12 characters) / 알파벳만 (8-12자):")
	str1, err := random.GenString.Letters(8, 12)
	if err != nil {
		logger.Fatal("Failed to generate letters string", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result / 결과: %s (length / 길이: %d)", str1, len(str1)))
	logger.Info("")

	// Example 2: Alphanumeric
	// 예제 2: 영숫자
	logger.Info("2. Alphanumeric (32-128 characters) / 영숫자 (32-128자):")
	str2, err := random.GenString.Alnum(32, 128)
	if err != nil {
		logger.Fatal("Failed to generate alphanumeric string", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result / 결과: %s (length / 길이: %d)", str2, len(str2)))
	logger.Info("")

	// Example 3: Fixed length
	// 예제 3: 고정 길이
	logger.Info("3. Fixed length alphanumeric (exactly 32 characters) / 고정 길이 영숫자 (정확히 32자):")
	str3, err := random.GenString.Alnum(32)
	if err != nil {
		logger.Fatal("Failed to generate fixed length string", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result / 결과: %s (length / 길이: %d)", str3, len(str3)))
	logger.Info("")

	// Example 4: Complex with all special characters
	// 예제 4: 모든 특수 문자 포함
	logger.Info("4. Complex with all special characters (16-24 characters) / 모든 특수 문자 포함 (16-24자):")
	str4, err := random.GenString.Complex(16, 24)
	if err != nil {
		logger.Fatal("Failed to generate complex string", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result / 결과: %s (length / 길이: %d)", str4, len(str4)))
	logger.Info("")

	// Example 5: Standard with safe special characters
	// 예제 5: 안전한 특수 문자 포함
	logger.Info("5. Standard with safe special characters (20-30 characters) / 안전한 특수 문자 포함 (20-30자):")
	str5, err := random.GenString.Standard(20, 30)
	if err != nil {
		logger.Fatal("Failed to generate standard string", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result / 결과: %s (length / 길이: %d)", str5, len(str5)))
	logger.Info("")

	// Example 6: Custom charset - numbers only
	// 예제 6: 사용자 정의 문자 집합 - 숫자만
	logger.Info("6. Custom charset - Numbers only (6 digits) / 사용자 정의 - 숫자만 (6자리):")
	str6, err := random.GenString.Custom("0123456789", 6)
	if err != nil {
		logger.Fatal("Failed to generate custom numeric string", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result / 결과: %s (length / 길이: %d)", str6, len(str6)))
	logger.Info("")

	// Example 7: Custom charset - hexadecimal
	// 예제 7: 사용자 정의 문자 집합 - 16진수
	logger.Info("7. Custom charset - Hexadecimal (16 characters) / 사용자 정의 - 16진수 (16자):")
	str7, err := random.GenString.Custom("0123456789ABCDEF", 16)
	if err != nil {
		logger.Fatal("Failed to generate custom hexadecimal string", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result / 결과: %s (length / 길이: %d)", str7, len(str7)))
	logger.Info("")

	// Common use cases
	// 일반적인 사용 사례
	logger.Info("=== Common Use Cases ===")
	logger.Info("=== 일반적인 사용 사례 ===")
	logger.Info("")

	// Password / 비밀번호
	password, err := random.GenString.Complex(16, 24)
	if err != nil {
		logger.Fatal("Failed to generate password", "error", err)
	}
	logger.Info(fmt.Sprintf("Secure Password / 안전한 비밀번호:  %s", password))

	// API Key / API 키
	apiKey, err := random.GenString.Alnum(40)
	if err != nil {
		logger.Fatal("Failed to generate API key", "error", err)
	}
	logger.Info(fmt.Sprintf("API Key / API 키:                   %s", apiKey))

	// Username / 사용자명
	username, err := random.GenString.Letters(8, 12)
	if err != nil {
		logger.Fatal("Failed to generate username", "error", err)
	}
	logger.Info(fmt.Sprintf("Username / 사용자명:                 %s", username))

	// Verification Code / 인증 코드
	verificationCode, err := random.GenString.Custom("0123456789", 6)
	if err != nil {
		logger.Fatal("Failed to generate verification code", "error", err)
	}
	logger.Info(fmt.Sprintf("Verification / 인증 코드:            %s", verificationCode))

	// Session Token / 세션 토큰
	sessionToken, err := random.GenString.Alnum(64)
	if err != nil {
		logger.Fatal("Failed to generate session token", "error", err)
	}
	logger.Info(fmt.Sprintf("Session Token / 세션 토큰:           %s", sessionToken))

	logger.Info("")
	logger.Info("=== Additional Method Examples ===")
	logger.Info("=== 추가 메서드 예제 ===")
	logger.Info("")

	// PIN Code / PIN 코드
	pinCode, err := random.GenString.Digits(6)
	if err != nil {
		logger.Fatal("Failed to generate PIN code", "error", err)
	}
	logger.Info(fmt.Sprintf("PIN Code / PIN 코드:                 %s", pinCode))

	// Hex Color Code / 16진수 색상 코드
	colorCode, err := random.GenString.Hex(6)
	if err != nil {
		logger.Fatal("Failed to generate hex color code", "error", err)
	}
	logger.Info(fmt.Sprintf("Hex Color / 16진수 색상:             #%s", colorCode))

	// UUID-like (lowercase hex) / UUID 형태 (소문자 16진수)
	uuidLike, err := random.GenString.HexLower(32)
	if err != nil {
		logger.Fatal("Failed to generate UUID-like string", "error", err)
	}
	logger.Info(fmt.Sprintf("UUID-like / UUID 형태:               %s", uuidLike))

	// Coupon Code (uppercase) / 쿠폰 코드 (대문자)
	couponCode, err := random.GenString.AlphaUpper(10)
	if err != nil {
		logger.Fatal("Failed to generate coupon code", "error", err)
	}
	logger.Info(fmt.Sprintf("Coupon Code / 쿠폰 코드:             %s", couponCode))

	// Subdomain (lowercase) / 서브도메인 (소문자)
	subdomain, err := random.GenString.AlphaLower(8, 12)
	if err != nil {
		logger.Fatal("Failed to generate subdomain", "error", err)
	}
	logger.Info(fmt.Sprintf("Subdomain / 서브도메인:              %s", subdomain))

	// License Key (uppercase + digits) / 라이선스 키 (대문자 + 숫자)
	licenseKey, err := random.GenString.AlnumUpper(16)
	if err != nil {
		logger.Fatal("Failed to generate license key", "error", err)
	}
	logger.Info(fmt.Sprintf("License Key / 라이선스 키:           %s", licenseKey))

	// Token (lowercase + digits) / 토큰 (소문자 + 숫자)
	token, err := random.GenString.AlnumLower(20)
	if err != nil {
		logger.Fatal("Failed to generate token", "error", err)
	}
	logger.Info(fmt.Sprintf("Token / 토큰:                        %s", token))

	// URL-safe Token / URL-safe 토큰
	urlSafeToken, err := random.GenString.Base64URL(32)
	if err != nil {
		logger.Fatal("Failed to generate URL-safe token", "error", err)
	}
	logger.Info(fmt.Sprintf("URL-safe Token / URL-safe 토큰:      %s", urlSafeToken))

	// Base64-like / Base64 형태
	base64Like, err := random.GenString.Base64(24)
	if err != nil {
		logger.Fatal("Failed to generate Base64-like string", "error", err)
	}
	logger.Info(fmt.Sprintf("Base64-like / Base64 형태:           %s", base64Like))

	logger.Info("")
	logger.Info("=== All examples completed successfully! ===")
	logger.Info("=== 모든 예제가 성공적으로 완료되었습니다! ===")
	logger.Info(fmt.Sprintf("Log saved to: ./results/logs/random_example_%s.log", time.Now().Format("20060102_150405")))
	logger.Info(fmt.Sprintf("로그 저장 위치: ./results/logs/random_example_%s.log", time.Now().Format("20060102_150405")))
}
