package stringutil

import "math"

// LevenshteinDistance calculates the Levenshtein distance between two strings.
// LevenshteinDistance는 두 문자열 사이의 Levenshtein 거리를 계산합니다.
//
// The Levenshtein distance is the minimum number of single-character edits
// (insertions, deletions, or substitutions) required to change one string into another.
//
// Levenshtein 거리는 한 문자열을 다른 문자열로 변경하는 데 필요한
// 최소 단일 문자 편집(삽입, 삭제 또는 치환) 횟수입니다.
//
// Example / 예제:
//
//	LevenshteinDistance("kitten", "sitting")  // 3
//	LevenshteinDistance("hello", "hello")     // 0
//	LevenshteinDistance("", "hello")          // 5
func LevenshteinDistance(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)

	lenA := len(runesA)
	lenB := len(runesB)

	// Early exit for empty strings / 빈 문자열에 대한 조기 종료
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}

	// Create matrix / 행렬 생성
	matrix := make([][]int, lenA+1)
	for i := range matrix {
		matrix[i] = make([]int, lenB+1)
	}

	// Initialize first row and column / 첫 행과 열 초기화
	for i := 0; i <= lenA; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		matrix[0][j] = j
	}

	// Fill matrix / 행렬 채우기
	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			cost := 0
			if runesA[i-1] != runesB[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion / 삭제
				matrix[i][j-1]+1,      // insertion / 삽입
				matrix[i-1][j-1]+cost, // substitution / 치환
			)
		}
	}

	return matrix[lenA][lenB]
}

// Similarity calculates the similarity ratio between two strings (0.0 to 1.0).
// Similarity는 두 문자열 사이의 유사도 비율을 계산합니다 (0.0 ~ 1.0).
//
// Returns 1.0 for identical strings, 0.0 for completely different strings.
// 동일한 문자열은 1.0, 완전히 다른 문자열은 0.0을 반환합니다.
//
// Formula: 1 - (distance / max(len(a), len(b)))
// 공식: 1 - (거리 / max(len(a), len(b)))
//
// Example / 예제:
//
//	Similarity("hello", "hello")   // 1.0
//	Similarity("hello", "hallo")   // 0.8
//	Similarity("hello", "world")   // 0.2
func Similarity(a, b string) float64 {
	// Early exit for identical strings / 동일한 문자열에 대한 조기 종료
	if a == b {
		return 1.0
	}

	runesA := []rune(a)
	runesB := []rune(b)

	maxLen := len(runesA)
	if len(runesB) > maxLen {
		maxLen = len(runesB)
	}

	// Early exit for empty strings / 빈 문자열에 대한 조기 종료
	if maxLen == 0 {
		return 1.0
	}

	distance := LevenshteinDistance(a, b)
	return 1.0 - (float64(distance) / float64(maxLen))
}

// HammingDistance calculates the Hamming distance between two strings of equal length.
// HammingDistance는 같은 길이의 두 문자열 사이의 Hamming 거리를 계산합니다.
//
// The Hamming distance is the number of positions at which the corresponding characters are different.
// Hamming 거리는 해당 문자가 다른 위치의 수입니다.
//
// Returns -1 if strings have different lengths.
// 문자열 길이가 다르면 -1을 반환합니다.
//
// Example / 예제:
//
//	HammingDistance("karolin", "kathrin")  // 3
//	HammingDistance("hello", "hello")      // 0
//	HammingDistance("hello", "world")      // 4
//	HammingDistance("hello", "hi")         // -1 (different lengths)
func HammingDistance(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)

	if len(runesA) != len(runesB) {
		return -1
	}

	distance := 0
	for i := 0; i < len(runesA); i++ {
		if runesA[i] != runesB[i] {
			distance++
		}
	}

	return distance
}

// JaroWinklerSimilarity calculates the Jaro-Winkler similarity between two strings (0.0 to 1.0).
// JaroWinklerSimilarity는 두 문자열 사이의 Jaro-Winkler 유사도를 계산합니다 (0.0 ~ 1.0).
//
// The Jaro-Winkler similarity is a variant of the Jaro similarity metric
// with extra weight given to common prefixes.
//
// Jaro-Winkler 유사도는 공통 접두사에 추가 가중치를 부여하는
// Jaro 유사도 메트릭의 변형입니다.
//
// Example / 예제:
//
//	JaroWinklerSimilarity("martha", "marhta")  // 0.961
//	JaroWinklerSimilarity("hello", "hello")    // 1.0
//	JaroWinklerSimilarity("hello", "world")    // 0.466
func JaroWinklerSimilarity(a, b string) float64 {
	jaroSim := jaroSimilarity(a, b)

	// Calculate common prefix length (up to 4 characters)
	// 공통 접두사 길이 계산 (최대 4문자)
	runesA := []rune(a)
	runesB := []rune(b)
	prefixLen := 0
	maxPrefix := 4
	if len(runesA) < maxPrefix {
		maxPrefix = len(runesA)
	}
	if len(runesB) < maxPrefix {
		maxPrefix = len(runesB)
	}

	for i := 0; i < maxPrefix; i++ {
		if runesA[i] == runesB[i] {
			prefixLen++
		} else {
			break
		}
	}

	// Jaro-Winkler formula / Jaro-Winkler 공식
	return jaroSim + (float64(prefixLen) * 0.1 * (1.0 - jaroSim))
}

// jaroSimilarity calculates the Jaro similarity between two strings.
// jaroSimilarity는 두 문자열 사이의 Jaro 유사도를 계산합니다.
func jaroSimilarity(a, b string) float64 {
	if a == b {
		return 1.0
	}

	runesA := []rune(a)
	runesB := []rune(b)
	lenA := len(runesA)
	lenB := len(runesB)

	if lenA == 0 && lenB == 0 {
		return 1.0
	}
	if lenA == 0 || lenB == 0 {
		return 0.0
	}

	// Calculate match window / 매치 윈도우 계산
	matchWindow := int(math.Max(float64(lenA), float64(lenB))/2) - 1
	if matchWindow < 1 {
		matchWindow = 1
	}

	matchesA := make([]bool, lenA)
	matchesB := make([]bool, lenB)

	matches := 0
	transpositions := 0

	// Find matches / 매치 찾기
	for i := 0; i < lenA; i++ {
		start := int(math.Max(0, float64(i-matchWindow)))
		end := int(math.Min(float64(lenB), float64(i+matchWindow+1)))

		for j := start; j < end; j++ {
			if matchesB[j] || runesA[i] != runesB[j] {
				continue
			}
			matchesA[i] = true
			matchesB[j] = true
			matches++
			break
		}
	}

	if matches == 0 {
		return 0.0
	}

	// Count transpositions / 전치 횟수 계산
	k := 0
	for i := 0; i < lenA; i++ {
		if !matchesA[i] {
			continue
		}
		for !matchesB[k] {
			k++
		}
		if runesA[i] != runesB[k] {
			transpositions++
		}
		k++
	}

	// Jaro formula / Jaro 공식
	return (float64(matches)/float64(lenA) +
		float64(matches)/float64(lenB) +
		float64(matches-transpositions/2)/float64(matches)) / 3.0
}

// min returns the minimum of three integers.
// min은 세 정수 중 최소값을 반환합니다.
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
