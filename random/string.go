package random

import (
	"crypto/rand"
	"math/big"
)

// Character sets for random string generation
const (
	charsetAlpha          = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	charsetDigits         = "0123456789"
	charsetSpecial        = "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	charsetSpecialLimited = "!@#$%^&*-_"
)

// stringGenerator provides methods for generating random strings
type stringGenerator struct{}

// GenString is a global instance for generating random strings
var GenString = stringGenerator{}

// Alpha generates a random string containing only alphabetic characters (a-z, A-Z)
// min: minimum length of the generated string
// max: maximum length of the generated string
// Returns a random string with length between min and max (inclusive)
func (stringGenerator) Alpha(min, max int) string {
	return generateRandomString(charsetAlpha, min, max)
}

// AlphaNum generates a random string containing alphabetic and numeric characters (a-z, A-Z, 0-9)
// min: minimum length of the generated string
// max: maximum length of the generated string
// Returns a random string with length between min and max (inclusive)
func (stringGenerator) AlphaNum(min, max int) string {
	return generateRandomString(charsetAlpha+charsetDigits, min, max)
}

// AlphaNumSpecial generates a random string containing alphabetic, numeric, and special characters
// Special characters include: !@#$%^&*()-_=+[]{}|;:,.<>?/
// min: minimum length of the generated string
// max: maximum length of the generated string
// Returns a random string with length between min and max (inclusive)
func (stringGenerator) AlphaNumSpecial(min, max int) string {
	return generateRandomString(charsetAlpha+charsetDigits+charsetSpecial, min, max)
}

// AlphaNumSpecialLimited generates a random string with limited special characters
// Special characters include only: !@#$%^&*-_
// min: minimum length of the generated string
// max: maximum length of the generated string
// Returns a random string with length between min and max (inclusive)
func (stringGenerator) AlphaNumSpecialLimited(min, max int) string {
	return generateRandomString(charsetAlpha+charsetDigits+charsetSpecialLimited, min, max)
}

// Custom generates a random string using a custom character set
// charset: custom set of characters to use for generation
// min: minimum length of the generated string
// max: maximum length of the generated string
// Returns a random string with length between min and max (inclusive)
func (stringGenerator) Custom(charset string, min, max int) string {
	return generateRandomString(charset, min, max)
}

// generateRandomString is a helper function that generates a random string
// from the given charset with a length between min and max
func generateRandomString(charset string, min, max int) string {
	if min < 0 {
		min = 0
	}
	if max < min {
		max = min
	}
	if len(charset) == 0 {
		return ""
	}

	// Determine the actual length of the string to generate
	length := min
	if max > min {
		// Generate random length between min and max
		lengthRange := max - min + 1
		randomLength, err := rand.Int(rand.Reader, big.NewInt(int64(lengthRange)))
		if err == nil {
			length = min + int(randomLength.Int64())
		}
	}

	// Generate the random string
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			// Fallback to first character if random generation fails
			result[i] = charset[0]
			continue
		}
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result)
}
