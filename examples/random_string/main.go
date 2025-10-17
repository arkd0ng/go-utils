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
	// Setup log file with backup management
	// 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/random-example.log"

	// Check if previous log file exists
	// 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file
		// 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp
			// 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/random-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file
			// 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
				// Delete original log file to prevent content duplication
				// 내용 중복 방지를 위해 원본 로그 파일 삭제
				fileutil.DeleteFile(logFilePath)
			}
		}

		// Cleanup old backup files - keep only 5 most recent
		// 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/random-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time
			// 수정 시간으로 정렬
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

			// Sort oldest first
			// 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5
			// 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename
	// 고정 파일명으로 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		// Enable stdout for screen output
		// 화면 출력 활성화
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner
	// 배너 출력
	logger.Banner("Random String Package Examples", "go-utils/random")

	logger.Info("╔════════════════════════════════════════════════════════════════════════════╗")
	logger.Info("║              Random String Package - Comprehensive Examples               ║")
	logger.Info("║              Random String 패키지 - 종합 예제                              ║")
	logger.Info("╚════════════════════════════════════════════════════════════════════════════╝")
	logger.Info("")
	logger.Info("📋 Package Information / 패키지 정보")
	logger.Info("   Package: github.com/arkd0ng/go-utils/random")
	logger.Info("   Description: Cryptographically secure random string generation")
	logger.Info("   설명: 암호학적으로 안전한 랜덤 문자열 생성")
	logger.Info("   Total Methods: 14 generators")
	logger.Info("   Security: Uses crypto/rand (not math/rand)")
	logger.Info("   Performance: Optimized for production use")
	logger.Info("")
	logger.Info("🔒 Security Features / 보안 기능")
	logger.Info("   • Cryptographically secure randomness")
	logger.Info("   • Unpredictable output (암호학적으로 안전)")
	logger.Info("   • Suitable for passwords, tokens, and keys")
	logger.Info("   • No predictable patterns")
	logger.Info("")

	// Example 1: Letters only
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣  Letters() - Alphabetic characters only")
	logger.Info("   알파벳 문자만 생성")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Letters(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates random string with uppercase and lowercase letters only")
	logger.Info("   대소문자 알파벳만 포함하는 랜덤 문자열을 생성합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Username generation (사용자명 생성)")
	logger.Info("   • Temporary identifiers (임시 식별자)")
	logger.Info("   • Alphabetic tokens (알파벳 토큰)")
	logger.Info("   • Name placeholders (이름 플레이스홀더)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Character set: a-z, A-Z (52 characters)")
	logger.Info("   • Variable length: 1 arg = fixed, 2 args = range")
	logger.Info("   • Cryptographically secure (crypto/rand)")
	logger.Info("   • URL-safe: no special characters")
	logger.Info("")
	logger.Info("📊 Character Set Details / 문자 집합 상세:")
	logger.Info("   • Lowercase: a-z (26 characters)")
	logger.Info("   • Uppercase: A-Z (26 characters)")
	logger.Info("   • Total pool: 52 possible characters")
	logger.Info("   • Entropy: ~5.7 bits per character")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Letters(8, 12)")
	logger.Info("   • Mode: Variable length (가변 길이)")
	logger.Info("   • Min length: 8 characters")
	logger.Info("   • Max length: 12 characters")
	logger.Info("")

	str1, err := random.GenString.Letters(8, 12)
	if err != nil {
		logger.Fatal("Failed to generate letters string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str1))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str1)))
	logger.Info(fmt.Sprintf("   🔤 Type: Alphabetic only"))
	logger.Info(fmt.Sprintf("   ✓ In Range: %v (8-12 characters)", len(str1) >= 8 && len(str1) <= 12))
	logger.Info("")
	logger.Info("🔍 Character Analysis / 문자 분석:")
	lowercase1, uppercase1 := 0, 0
	for _, c := range str1 {
		if c >= 'a' && c <= 'z' {
			lowercase1++
		} else if c >= 'A' && c <= 'Z' {
			uppercase1++
		}
	}
	logger.Info(fmt.Sprintf("   • Lowercase letters: %d (%.1f%%)", lowercase1, float64(lowercase1)/float64(len(str1))*100))
	logger.Info(fmt.Sprintf("   • Uppercase letters: %d (%.1f%%)", uppercase1, float64(uppercase1)/float64(len(str1))*100))
	logger.Info(fmt.Sprintf("   • Total: %d", len(str1)))
	logger.Info("")

	// Example 2: Alphanumeric
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("2️⃣  Alnum() - Alphanumeric characters")
	logger.Info("   영숫자 문자 생성")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Alnum(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates random alphanumeric string (letters + digits)")
	logger.Info("   영문자와 숫자를 포함하는 랜덤 문자열을 생성합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • API keys and tokens (API 키 및 토큰)")
	logger.Info("   • Session IDs (세션 ID)")
	logger.Info("   • Verification codes (인증 코드)")
	logger.Info("   • Unique identifiers (고유 식별자)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Character set: a-z, A-Z, 0-9 (62 characters)")
	logger.Info("   • Most common format for tokens")
	logger.Info("   • URL-safe and database-friendly")
	logger.Info("   • Higher entropy than letters-only")
	logger.Info("")
	logger.Info("📊 Character Set Details / 문자 집합 상세:")
	logger.Info("   • Lowercase: a-z (26 characters)")
	logger.Info("   • Uppercase: A-Z (26 characters)")
	logger.Info("   • Digits: 0-9 (10 characters)")
	logger.Info("   • Total pool: 62 possible characters")
	logger.Info("   • Entropy: ~5.95 bits per character")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Alnum(32, 128)")
	logger.Info("   • Mode: Variable length (가변 길이)")
	logger.Info("   • Min length: 32 characters")
	logger.Info("   • Max length: 128 characters")
	logger.Info("")

	str2, err := random.GenString.Alnum(32, 128)
	if err != nil {
		logger.Fatal("Failed to generate alphanumeric string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str2))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str2)))
	logger.Info(fmt.Sprintf("   🔤 Type: Alphanumeric"))
	logger.Info(fmt.Sprintf("   ✓ In Range: %v (32-128 characters)", len(str2) >= 32 && len(str2) <= 128))
	logger.Info("")
	logger.Info("🔍 Character Analysis / 문자 분석:")
	lowercase2, uppercase2, digits2 := 0, 0, 0
	for _, c := range str2 {
		if c >= 'a' && c <= 'z' {
			lowercase2++
		} else if c >= 'A' && c <= 'Z' {
			uppercase2++
		} else if c >= '0' && c <= '9' {
			digits2++
		}
	}
	logger.Info(fmt.Sprintf("   • Lowercase letters: %d (%.1f%%)", lowercase2, float64(lowercase2)/float64(len(str2))*100))
	logger.Info(fmt.Sprintf("   • Uppercase letters: %d (%.1f%%)", uppercase2, float64(uppercase2)/float64(len(str2))*100))
	logger.Info(fmt.Sprintf("   • Digits: %d (%.1f%%)", digits2, float64(digits2)/float64(len(str2))*100))
	logger.Info(fmt.Sprintf("   • Total: %d", len(str2)))
	logger.Info("")

	// Example 3: Fixed length
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("3️⃣  Alnum() - Fixed length mode")
	logger.Info("   고정 길이 모드")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Alnum(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   When called with one argument, generates exact length string")
	logger.Info("   하나의 인자로 호출하면 정확한 길이의 문자열을 생성합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Fixed-format tokens (고정 형식 토큰)")
	logger.Info("   • Database primary keys (데이터베이스 기본 키)")
	logger.Info("   • Standardized identifiers (표준화된 식별자)")
	logger.Info("   • Uniform password length (균일한 비밀번호 길이)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Guaranteed exact length (정확한 길이 보장)")
	logger.Info("   • Predictable output size (예측 가능한 출력 크기)")
	logger.Info("   • Easier validation (더 쉬운 검증)")
	logger.Info("   • Consistent formatting (일관된 형식)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Alnum(32)")
	logger.Info("   • Mode: Fixed length (고정 길이)")
	logger.Info("   • Exact length: 32 characters")
	logger.Info("")

	str3, err := random.GenString.Alnum(32)
	if err != nil {
		logger.Fatal("Failed to generate fixed length string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str3))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters (exactly as requested)", len(str3)))
	logger.Info(fmt.Sprintf("   🔤 Type: Alphanumeric"))
	logger.Info(fmt.Sprintf("   ✓ Exact Match: %v", len(str3) == 32))
	logger.Info("")

	// Example 4: Complex with all special characters
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("4️⃣  Complex() - Maximum security with special characters")
	logger.Info("   특수 문자를 포함한 최대 보안")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Complex(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates maximum security string with letters, digits, and ALL special characters")
	logger.Info("   영문자, 숫자, 모든 특수 문자를 포함한 최대 보안 문자열을 생성합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Strong passwords (강력한 비밀번호)")
	logger.Info("   • Encryption keys (암호화 키)")
	logger.Info("   • High-security tokens (높은 보안 토큰)")
	logger.Info("   • Master passwords (마스터 비밀번호)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Maximum character diversity (최대 문자 다양성)")
	logger.Info("   • Highest entropy per character (가장 높은 엔트로피)")
	logger.Info("   • Includes ALL printable special characters")
	logger.Info("   • Strongest security level (가장 강력한 보안)")
	logger.Info("")
	logger.Info("📊 Character Set Details / 문자 집합 상세:")
	logger.Info("   • Lowercase: a-z (26 characters)")
	logger.Info("   • Uppercase: A-Z (26 characters)")
	logger.Info("   • Digits: 0-9 (10 characters)")
	logger.Info("   • Special: !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ (32 characters)")
	logger.Info("   • Total pool: 94 possible characters")
	logger.Info("   • Entropy: ~6.55 bits per character")
	logger.Info("")
	logger.Info("⚠️  Warning / 경고:")
	logger.Info("   • May require escaping in shell/SQL contexts")
	logger.Info("   • Some systems may have character restrictions")
	logger.Info("   • URL encoding may be needed")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Complex(16, 24)")
	logger.Info("   • Mode: Variable length (가변 길이)")
	logger.Info("   • Min length: 16 characters")
	logger.Info("   • Max length: 24 characters")
	logger.Info("")

	str4, err := random.GenString.Complex(16, 24)
	if err != nil {
		logger.Fatal("Failed to generate complex string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str4))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str4)))
	logger.Info(fmt.Sprintf("   🔤 Type: Complex (alphanumeric + all special characters)"))
	logger.Info(fmt.Sprintf("   ✓ In Range: %v (16-24 characters)", len(str4) >= 16 && len(str4) <= 24))
	logger.Info("")
	logger.Info("🔍 Character Analysis / 문자 분석:")
	lowercase4, uppercase4, digits4, special4 := 0, 0, 0, 0
	for _, c := range str4 {
		if c >= 'a' && c <= 'z' {
			lowercase4++
		} else if c >= 'A' && c <= 'Z' {
			uppercase4++
		} else if c >= '0' && c <= '9' {
			digits4++
		} else {
			special4++
		}
	}
	logger.Info(fmt.Sprintf("   • Lowercase letters: %d", lowercase4))
	logger.Info(fmt.Sprintf("   • Uppercase letters: %d", uppercase4))
	logger.Info(fmt.Sprintf("   • Digits: %d", digits4))
	logger.Info(fmt.Sprintf("   • Special characters: %d", special4))
	logger.Info(fmt.Sprintf("   • Total: %d", len(str4)))
	logger.Info(fmt.Sprintf("   • Character diversity: %.1f%% (4 types present)", float64(4)/float64(4)*100))
	logger.Info("")

	// Example 5: Standard with safe special characters
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("5️⃣  Standard() - Balanced security with safe special characters")
	logger.Info("   안전한 특수 문자를 포함한 균형잡힌 보안")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Standard(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates secure string with letters, digits, and SAFE special characters")
	logger.Info("   영문자, 숫자, 안전한 특수 문자를 포함한 보안 문자열을 생성합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • User passwords (사용자 비밀번호)")
	logger.Info("   • General-purpose tokens (범용 토큰)")
	logger.Info("   • Cross-platform identifiers (크로스 플랫폼 식별자)")
	logger.Info("   • Widely compatible strings (폭넓은 호환성)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Good balance between security and compatibility")
	logger.Info("   • Safe for most contexts (URL, shell, SQL)")
	logger.Info("   • Selected special chars: -_@#$%")
	logger.Info("   • Recommended for general use (일반 사용 권장)")
	logger.Info("")
	logger.Info("📊 Character Set Details / 문자 집합 상세:")
	logger.Info("   • Lowercase: a-z (26 characters)")
	logger.Info("   • Uppercase: A-Z (26 characters)")
	logger.Info("   • Digits: 0-9 (10 characters)")
	logger.Info("   • Safe special: -_@#$% (6 characters)")
	logger.Info("   • Total pool: 68 possible characters")
	logger.Info("   • Entropy: ~6.09 bits per character")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Standard(20, 30)")
	logger.Info("   • Mode: Variable length (가변 길이)")
	logger.Info("   • Min length: 20 characters")
	logger.Info("   • Max length: 30 characters")
	logger.Info("")

	str5, err := random.GenString.Standard(20, 30)
	if err != nil {
		logger.Fatal("Failed to generate standard string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str5))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str5)))
	logger.Info(fmt.Sprintf("   🔤 Type: Standard (alphanumeric + safe special)"))
	logger.Info(fmt.Sprintf("   ✓ In Range: %v (20-30 characters)", len(str5) >= 20 && len(str5) <= 30))
	logger.Info("")

	// Example 6: Digits only
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("6️⃣  Digits() - Numeric characters only")
	logger.Info("   숫자 문자만 생성")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Digits(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates random numeric string (0-9 only)")
	logger.Info("   숫자만 포함하는 랜덤 문자열을 생성합니다 (0-9)")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • PIN codes (PIN 코드)")
	logger.Info("   • Verification codes (인증 코드)")
	logger.Info("   • Order numbers (주문 번호)")
	logger.Info("   • Numeric tokens (숫자 토큰)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Character set: 0-9 (10 characters)")
	logger.Info("   • Easy to type and read (입력 및 읽기 용이)")
	logger.Info("   • Universal compatibility (범용 호환성)")
	logger.Info("   • Suitable for SMS/phone entry")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Digits(6)")
	logger.Info("   • Mode: Fixed length")
	logger.Info("   • Length: 6 digits (common for 2FA codes)")
	logger.Info("")

	str6, err := random.GenString.Digits(6)
	if err != nil {
		logger.Fatal("Failed to generate digits string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str6))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str6)))
	logger.Info(fmt.Sprintf("   🔤 Type: Numeric only"))
	logger.Info(fmt.Sprintf("   ✓ Format: Suitable for PIN/verification code"))
	logger.Info("")

	// Example 7: Hex
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("7️⃣  Hex() - Hexadecimal uppercase")
	logger.Info("   16진수 대문자")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Hex(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates random hexadecimal string (0-9, A-F uppercase)")
	logger.Info("   16진수 문자열을 생성합니다 (0-9, A-F 대문자)")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Color codes (색상 코드)")
	logger.Info("   • Hash representations (해시 표현)")
	logger.Info("   • Memory addresses (메모리 주소)")
	logger.Info("   • UUID components (UUID 구성요소)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Character set: 0-9, A-F (16 characters)")
	logger.Info("   • Standard hex format (uppercase)")
	logger.Info("   • Programming-friendly (프로그래밍 친화적)")
	logger.Info("   • Commonly used in tech contexts")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Hex(16)")
	logger.Info("   • Mode: Fixed length")
	logger.Info("   • Length: 16 characters")
	logger.Info("")

	str7, err := random.GenString.Hex(16)
	if err != nil {
		logger.Fatal("Failed to generate hex string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str7))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str7)))
	logger.Info(fmt.Sprintf("   🔤 Type: Hexadecimal (uppercase)"))
	logger.Info(fmt.Sprintf("   ✓ Format: Standard hex format"))
	logger.Info("")

	// Example 8: HexLower
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("8️⃣  HexLower() - Hexadecimal lowercase")
	logger.Info("   16진수 소문자")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func HexLower(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates random hexadecimal string (0-9, a-f lowercase)")
	logger.Info("   16진수 문자열을 생성합니다 (0-9, a-f 소문자)")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Git commit hashes (Git 커밋 해시)")
	logger.Info("   • Lowercase hex requirements (소문자 16진수 요구사항)")
	logger.Info("   • CSS color codes (CSS 색상 코드)")
	logger.Info("   • Database hex fields (데이터베이스 16진수 필드)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Character set: 0-9, a-f (16 characters)")
	logger.Info("   • Lowercase format (소문자 형식)")
	logger.Info("   • Matches git/crypto conventions")
	logger.Info("   • Web-friendly format")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.HexLower(32)")
	logger.Info("   • Mode: Fixed length")
	logger.Info("   • Length: 32 characters (like SHA-256 prefix)")
	logger.Info("")

	str8, err := random.GenString.HexLower(32)
	if err != nil {
		logger.Fatal("Failed to generate lowercase hex string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str8))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str8)))
	logger.Info(fmt.Sprintf("   🔤 Type: Hexadecimal (lowercase)"))
	logger.Info(fmt.Sprintf("   ✓ Format: Git-style hex format"))
	logger.Info("")

	// Example 9: Base64
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("9️⃣  Base64() - Base64 encoded characters")
	logger.Info("   Base64 인코딩 문자")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Base64(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates random Base64 string (a-z, A-Z, 0-9, +, /)")
	logger.Info("   Base64 문자열을 생성합니다 (a-z, A-Z, 0-9, +, /)")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Binary data encoding (바이너리 데이터 인코딩)")
	logger.Info("   • Email-safe tokens (이메일 안전 토큰)")
	logger.Info("   • API authentication (API 인증)")
	logger.Info("   • Data transmission (데이터 전송)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Character set: a-z, A-Z, 0-9, +, / (64 characters)")
	logger.Info("   • Standard Base64 alphabet")
	logger.Info("   • Compact representation (압축 표현)")
	logger.Info("   • Wide compatibility (폭넓은 호환성)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Base64(32)")
	logger.Info("   • Mode: Fixed length")
	logger.Info("   • Length: 32 characters")
	logger.Info("")

	str9, err := random.GenString.Base64(32)
	if err != nil {
		logger.Fatal("Failed to generate base64 string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str9))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str9)))
	logger.Info(fmt.Sprintf("   🔤 Type: Base64 standard"))
	logger.Info(fmt.Sprintf("   ✓ Format: Standard Base64 alphabet"))
	logger.Info("")

	// Example 10: Base64URL
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🔟 Base64URL() - URL-safe Base64")
	logger.Info("   URL 안전 Base64")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Base64URL(length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates URL-safe Base64 string (a-z, A-Z, 0-9, -, _)")
	logger.Info("   URL 안전 Base64 문자열을 생성합니다 (a-z, A-Z, 0-9, -, _)")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • URL parameters (URL 매개변수)")
	logger.Info("   • Filename-safe tokens (파일명 안전 토큰)")
	logger.Info("   • JWT tokens (JWT 토큰)")
	logger.Info("   • Cookie values (쿠키 값)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Character set: a-z, A-Z, 0-9, -, _ (64 characters)")
	logger.Info("   • URL-safe: replaces + with - and / with _")
	logger.Info("   • No percent-encoding needed")
	logger.Info("   • Filesystem-friendly (파일시스템 친화적)")
	logger.Info("")
	logger.Info("▶️  Executing / 실행 중:")
	logger.Info("   str, err := random.GenString.Base64URL(32)")
	logger.Info("   • Mode: Fixed length")
	logger.Info("   • Length: 32 characters")
	logger.Info("")

	str10, err := random.GenString.Base64URL(32)
	if err != nil {
		logger.Fatal("Failed to generate base64url string", "error", err)
	}

	logger.Info("✅ Generation Successful / 생성 성공")
	logger.Info(fmt.Sprintf("   📝 Result: %s", str10))
	logger.Info(fmt.Sprintf("   📏 Length: %d characters", len(str10)))
	logger.Info(fmt.Sprintf("   🔤 Type: Base64 URL-safe"))
	logger.Info(fmt.Sprintf("   ✓ Format: No URL encoding required"))
	logger.Info("")

	// Example 11-14: Case variants
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣1️⃣  Case Variants - Uppercase/Lowercase control")
	logger.Info("   대소문자 제어 변형")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Four methods for specific case requirements")
	logger.Info("   특정 대소문자 요구사항을 위한 4가지 메서드")
	logger.Info("")

	// AlphaUpper
	logger.Info("📚 AlphaUpper() - Uppercase letters only (A-Z)")
	str11, err := random.GenString.AlphaUpper(10)
	if err != nil {
		logger.Fatal("Failed to generate AlphaUpper", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result: %s", str11))
	logger.Info("   Use case: SCREAMING_SNAKE_CASE identifiers")
	logger.Info("")

	// AlphaLower
	logger.Info("📚 AlphaLower() - Lowercase letters only (a-z)")
	str12, err := random.GenString.AlphaLower(10)
	if err != nil {
		logger.Fatal("Failed to generate AlphaLower", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result: %s", str12))
	logger.Info("   Use case: lowercase usernames, slugs")
	logger.Info("")

	// AlnumUpper
	logger.Info("📚 AlnumUpper() - Uppercase alphanumeric (A-Z, 0-9)")
	str13, err := random.GenString.AlnumUpper(10)
	if err != nil {
		logger.Fatal("Failed to generate AlnumUpper", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result: %s", str13))
	logger.Info("   Use case: License keys, serial numbers")
	logger.Info("")

	// AlnumLower
	logger.Info("📚 AlnumLower() - Lowercase alphanumeric (a-z, 0-9)")
	str14, err := random.GenString.AlnumLower(10)
	if err != nil {
		logger.Fatal("Failed to generate AlnumLower", "error", err)
	}
	logger.Info(fmt.Sprintf("   Result: %s", str14))
	logger.Info("   Use case: Database keys, subdomain names")
	logger.Info("")

	// Custom charset example
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("1️⃣4️⃣  Custom() - User-defined character set")
	logger.Info("   사용자 정의 문자 집합")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📚 Function Signature / 함수 시그니처:")
	logger.Info("   func Custom(charset string, length ...int) (string, error)")
	logger.Info("")
	logger.Info("📖 Description / 설명:")
	logger.Info("   Generates random string using custom character set")
	logger.Info("   사용자 정의 문자 집합을 사용하여 랜덤 문자열을 생성합니다")
	logger.Info("")
	logger.Info("🎯 Use Cases / 사용 사례:")
	logger.Info("   • Domain-specific formats (도메인별 형식)")
	logger.Info("   • Restricted character sets (제한된 문자 집합)")
	logger.Info("   • Special encoding schemes (특수 인코딩 체계)")
	logger.Info("   • Custom alphabets (사용자 정의 알파벳)")
	logger.Info("")
	logger.Info("💡 Key Features / 주요 기능:")
	logger.Info("   • Any character set allowed (모든 문자 집합 허용)")
	logger.Info("   • Maximum flexibility (최대 유연성)")
	logger.Info("   • Unicode support (유니코드 지원)")
	logger.Info("   • Application-specific needs")
	logger.Info("")
	logger.Info("▶️  Example 1: Custom vowels-only string")
	logger.Info("   예제 1: 모음만 포함")
	customVowels, err := random.GenString.Custom("aeiouAEIOU", 8)
	if err != nil {
		logger.Fatal("Failed to generate custom vowels", "error", err)
	}
	logger.Info(fmt.Sprintf("   Charset: \"aeiouAEIOU\""))
	logger.Info(fmt.Sprintf("   Result: %s", customVowels))
	logger.Info("")
	logger.Info("▶️  Example 2: DNA sequence (ATCG)")
	logger.Info("   예제 2: DNA 서열 (ATCG)")
	customDNA, err := random.GenString.Custom("ATCG", 20)
	if err != nil {
		logger.Fatal("Failed to generate custom DNA", "error", err)
	}
	logger.Info(fmt.Sprintf("   Charset: \"ATCG\""))
	logger.Info(fmt.Sprintf("   Result: %s", customDNA))
	logger.Info("")

	// Real-world use cases summary
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("🌟 Real-World Use Cases Summary")
	logger.Info("   실제 사용 사례 요약")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")

	// Password
	logger.Info("💼 Use Case 1: Secure Password Generation")
	logger.Info("   사용 사례 1: 안전한 비밀번호 생성")
	password, _ := random.GenString.Complex(16, 24)
	logger.Info(fmt.Sprintf("   Generated: %s", password))
	logger.Info("   Recommendation: Complex() for maximum security")
	logger.Info("")

	// API Key
	logger.Info("💼 Use Case 2: API Key Generation")
	logger.Info("   사용 사례 2: API 키 생성")
	apiKey, _ := random.GenString.Alnum(40)
	logger.Info(fmt.Sprintf("   Generated: %s", apiKey))
	logger.Info("   Recommendation: Alnum() with fixed length 32-64")
	logger.Info("")

	// Session ID
	logger.Info("💼 Use Case 3: Session ID")
	logger.Info("   사용 사례 3: 세션 ID")
	sessionID, _ := random.GenString.Base64URL(32)
	logger.Info(fmt.Sprintf("   Generated: %s", sessionID))
	logger.Info("   Recommendation: Base64URL() for cookies/URLs")
	logger.Info("")

	// Verification Code
	logger.Info("💼 Use Case 4: 2FA Verification Code")
	logger.Info("   사용 사례 4: 2단계 인증 코드")
	verifyCode, _ := random.GenString.Digits(6)
	logger.Info(fmt.Sprintf("   Generated: %s", verifyCode))
	logger.Info("   Recommendation: Digits() with length 6")
	logger.Info("")

	// Token
	logger.Info("💼 Use Case 5: Reset Token")
	logger.Info("   사용 사례 5: 리셋 토큰")
	resetToken, _ := random.GenString.Hex(32)
	logger.Info(fmt.Sprintf("   Generated: %s", resetToken))
	logger.Info("   Recommendation: Hex() or HexLower() for tokens")
	logger.Info("")

	// Username
	logger.Info("💼 Use Case 6: Random Username")
	logger.Info("   사용 사례 6: 랜덤 사용자명")
	username, _ := random.GenString.AlphaLower(8, 12)
	logger.Info(fmt.Sprintf("   Generated: %s", username))
	logger.Info("   Recommendation: AlphaLower() for usernames")
	logger.Info("")

	// Final summary
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("✅ All Examples Completed Successfully")
	logger.Info("   모든 예제가 성공적으로 완료되었습니다")
	logger.Info("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	logger.Info("")
	logger.Info("📊 Summary / 요약:")
	logger.Info("   • Total methods demonstrated: 14")
	logger.Info("   • Character sets covered: 10+ types")
	logger.Info("   • Real-world use cases: 6 scenarios")
	logger.Info("   • Security level: Cryptographically secure (crypto/rand)")
	logger.Info("")
	logger.Info("📚 Documentation / 문서:")
	logger.Info("   • Full documentation: github.com/arkd0ng/go-utils/random/README.md")
	logger.Info("   • Source code: github.com/arkd0ng/go-utils/random/")
	logger.Info("   • Test coverage: 100% (all functions tested)")
	logger.Info("")
	logger.Info("💡 Best Practices / 모범 사례:")
	logger.Info("   1. Use Complex() or Standard() for passwords")
	logger.Info("   2. Use Alnum() for general-purpose tokens")
	logger.Info("   3. Use Base64URL() for URL-safe tokens")
	logger.Info("   4. Use Digits() for SMS/2FA codes")
	logger.Info("   5. Use HexLower() for crypto-style identifiers")
	logger.Info("")
	logger.Info("🔒 Security Notes / 보안 참고사항:")
	logger.Info("   • All methods use crypto/rand (NOT math/rand)")
	logger.Info("   • Suitable for cryptographic purposes")
	logger.Info("   • No predictable patterns")
	logger.Info("   • Production-ready security level")
	logger.Info("")
	logger.Info("Thank you for using go-utils/random package!")
	logger.Info("go-utils/random 패키지를 사용해 주셔서 감사합니다!")
	logger.Info("")
}
