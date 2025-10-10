package main

import (
	"fmt"
	"github.com/arkd0ng/go-utils/random"
)

func main() {
	fmt.Println("=== Random String Generation Examples ===")
	fmt.Println("=== 랜덤 문자열 생성 예제 ===")
	fmt.Println()

	// Example 1: Letters only
	// 예제 1: 알파벳만
	fmt.Println("1. Letters only (8-12 characters) / 알파벳만 (8-12자):")
	str1 := random.GenString.Letters(8, 12)
	fmt.Printf("   Result / 결과: %s (length / 길이: %d)\n\n", str1, len(str1))

	// Example 2: Alphanumeric
	// 예제 2: 영숫자
	fmt.Println("2. Alphanumeric (32-128 characters) / 영숫자 (32-128자):")
	str2 := random.GenString.Alnum(32, 128)
	fmt.Printf("   Result / 결과: %s (length / 길이: %d)\n\n", str2, len(str2))

	// Example 3: Fixed length
	// 예제 3: 고정 길이
	fmt.Println("3. Fixed length alphanumeric (exactly 32 characters) / 고정 길이 영숫자 (정확히 32자):")
	str3 := random.GenString.Alnum(32, 32)
	fmt.Printf("   Result / 결과: %s (length / 길이: %d)\n\n", str3, len(str3))

	// Example 4: Complex with all special characters
	// 예제 4: 모든 특수 문자 포함
	fmt.Println("4. Complex with all special characters (16-24 characters) / 모든 특수 문자 포함 (16-24자):")
	str4 := random.GenString.Complex(16, 24)
	fmt.Printf("   Result / 결과: %s (length / 길이: %d)\n\n", str4, len(str4))

	// Example 5: Standard with safe special characters
	// 예제 5: 안전한 특수 문자 포함
	fmt.Println("5. Standard with safe special characters (20-30 characters) / 안전한 특수 문자 포함 (20-30자):")
	str5 := random.GenString.Standard(20, 30)
	fmt.Printf("   Result / 결과: %s (length / 길이: %d)\n\n", str5, len(str5))

	// Example 6: Custom charset - numbers only
	// 예제 6: 사용자 정의 문자 집합 - 숫자만
	fmt.Println("6. Custom charset - Numbers only (6 digits) / 사용자 정의 - 숫자만 (6자리):")
	str6 := random.GenString.Custom("0123456789", 6, 6)
	fmt.Printf("   Result / 결과: %s (length / 길이: %d)\n\n", str6, len(str6))

	// Example 7: Custom charset - hexadecimal
	// 예제 7: 사용자 정의 문자 집합 - 16진수
	fmt.Println("7. Custom charset - Hexadecimal (16 characters) / 사용자 정의 - 16진수 (16자):")
	str7 := random.GenString.Custom("0123456789ABCDEF", 16, 16)
	fmt.Printf("   Result / 결과: %s (length / 길이: %d)\n\n", str7, len(str7))

	// Common use cases
	// 일반적인 사용 사례
	fmt.Println("=== Common Use Cases ===")
	fmt.Println("=== 일반적인 사용 사례 ===")
	fmt.Println()

	// Password / 비밀번호
	password := random.GenString.Complex(16, 24)
	fmt.Printf("Secure Password / 안전한 비밀번호:  %s\n", password)

	// API Key / API 키
	apiKey := random.GenString.Alnum(40, 40)
	fmt.Printf("API Key / API 키:                   %s\n", apiKey)

	// Username / 사용자명
	username := random.GenString.Letters(8, 12)
	fmt.Printf("Username / 사용자명:                 %s\n", username)

	// Verification Code / 인증 코드
	verificationCode := random.GenString.Custom("0123456789", 6, 6)
	fmt.Printf("Verification / 인증 코드:            %s\n", verificationCode)

	// Session Token / 세션 토큰
	sessionToken := random.GenString.Alnum(64, 64)
	fmt.Printf("Session Token / 세션 토큰:           %s\n", sessionToken)
}
