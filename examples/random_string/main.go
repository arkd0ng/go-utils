package main

import (
	"fmt"
	"github.com/arkd0ng/go-utils/random"
)

func main() {
	fmt.Println("=== Random String Generation Examples ===\n")

	// Example 1: Alphabetic only
	fmt.Println("1. Alphabetic only (8-12 characters):")
	str1 := random.GenString.Alpha(8, 12)
	fmt.Printf("   Result: %s (length: %d)\n\n", str1, len(str1))

	// Example 2: Alphanumeric
	fmt.Println("2. Alphanumeric (32-128 characters):")
	str2 := random.GenString.AlphaNum(32, 128)
	fmt.Printf("   Result: %s (length: %d)\n\n", str2, len(str2))

	// Example 3: Fixed length
	fmt.Println("3. Fixed length alphanumeric (exactly 32 characters):")
	str3 := random.GenString.AlphaNum(32, 32)
	fmt.Printf("   Result: %s (length: %d)\n\n", str3, len(str3))

	// Example 4: With special characters
	fmt.Println("4. With all special characters (16-24 characters):")
	str4 := random.GenString.AlphaNumSpecial(16, 24)
	fmt.Printf("   Result: %s (length: %d)\n\n", str4, len(str4))

	// Example 5: With limited special characters
	fmt.Println("5. With limited special characters (20-30 characters):")
	str5 := random.GenString.AlphaNumSpecialLimited(20, 30)
	fmt.Printf("   Result: %s (length: %d)\n\n", str5, len(str5))

	// Example 6: Custom charset - numbers only
	fmt.Println("6. Custom charset - Numbers only (6 digits):")
	str6 := random.GenString.Custom("0123456789", 6, 6)
	fmt.Printf("   Result: %s (length: %d)\n\n", str6, len(str6))

	// Example 7: Custom charset - hexadecimal
	fmt.Println("7. Custom charset - Hexadecimal (16 characters):")
	str7 := random.GenString.Custom("0123456789ABCDEF", 16, 16)
	fmt.Printf("   Result: %s (length: %d)\n\n", str7, len(str7))

	// Common use cases
	fmt.Println("=== Common Use Cases ===\n")

	// Password
	password := random.GenString.AlphaNumSpecial(16, 24)
	fmt.Printf("Secure Password: %s\n", password)

	// API Key
	apiKey := random.GenString.AlphaNum(40, 40)
	fmt.Printf("API Key:         %s\n", apiKey)

	// Username
	username := random.GenString.Alpha(8, 12)
	fmt.Printf("Username:        %s\n", username)

	// Verification Code
	verificationCode := random.GenString.Custom("0123456789", 6, 6)
	fmt.Printf("Verification:    %s\n", verificationCode)

	// Session Token
	sessionToken := random.GenString.AlphaNum(64, 64)
	fmt.Printf("Session Token:   %s\n", sessionToken)
}
