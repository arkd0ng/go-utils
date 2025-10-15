package sliceutil

import (
	"errors"
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

// Number is a constraint for numeric types.
// Number는 숫자 타입에 대한 제약입니다.
type Number interface {
	constraints.Integer | constraints.Float
}

// Median calculates the median value of a slice of numbers.
// Returns the middle value for odd-length slices, or the average of the two middle values for even-length slices.
// Returns an error if the slice is empty.
//
// Median은 숫자 슬라이스의 중앙값을 계산합니다.
// 홀수 길이 슬라이스의 경우 중간 값을, 짝수 길이 슬라이스의 경우 두 중간 값의 평균을 반환합니다.
// 슬라이스가 비어 있으면 에러를 반환합니다.
//
// Example:
//
//	numbers := []int{3, 1, 4, 1, 5, 9, 2}
//	median, err := sliceutil.Median(numbers) // 3
//
//	evens := []int{1, 2, 3, 4}
//	median, err := sliceutil.Median(evens) // 2.5
func Median[T Number](slice []T) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot calculate median of empty slice")
	}

	// Create a sorted copy
	sorted := Sort(slice)
	length := len(sorted)

	if length%2 == 0 {
		// Even length: average of two middle values
		mid1 := sorted[length/2-1]
		mid2 := sorted[length/2]
		return (float64(mid1) + float64(mid2)) / 2.0, nil
	}

	// Odd length: middle value
	return float64(sorted[length/2]), nil
}

// Mode returns the most frequently occurring element in the slice.
// If there are multiple modes, returns the first one encountered.
// Returns an error if the slice is empty.
//
// Mode는 슬라이스에서 가장 자주 나타나는 요소를 반환합니다.
// 여러 최빈값이 있는 경우 처음 발견된 것을 반환합니다.
// 슬라이스가 비어 있으면 에러를 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4}
//	mode, err := sliceutil.Mode(numbers) // 3
func Mode[T comparable](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("cannot calculate mode of empty slice")
	}

	frequencies := make(map[T]int)
	for _, v := range slice {
		frequencies[v]++
	}

	var mode T
	maxCount := 0
	for value, count := range frequencies {
		if count > maxCount {
			maxCount = count
			mode = value
		}
	}

	return mode, nil
}

// Frequencies returns a map of each element to its frequency count in the slice.
//
// Frequencies는 슬라이스의 각 요소를 빈도 수에 매핑한 맵을 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4}
//	freq := sliceutil.Frequencies(numbers) // map[1:1 2:2 3:3 4:1]
func Frequencies[T comparable](slice []T) map[T]int {
	frequencies := make(map[T]int)
	for _, v := range slice {
		frequencies[v]++
	}
	return frequencies
}

// Percentile calculates the p-th percentile of a slice of numbers.
// p should be between 0 and 100 (inclusive).
// Uses the linear interpolation method.
// Returns an error if the slice is empty or p is out of range.
//
// Percentile은 숫자 슬라이스의 p번째 백분위수를 계산합니다.
// p는 0과 100 사이여야 합니다 (포함).
// 선형 보간법을 사용합니다.
// 슬라이스가 비어 있거나 p가 범위를 벗어나면 에러를 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	p50, err := sliceutil.Percentile(numbers, 50) // 5.5 (median)
//	p75, err := sliceutil.Percentile(numbers, 75) // 7.75
//	p90, err := sliceutil.Percentile(numbers, 90) // 9.1
func Percentile[T Number](slice []T, p float64) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot calculate percentile of empty slice")
	}
	if p < 0 || p > 100 {
		return 0, errors.New("percentile must be between 0 and 100")
	}

	// Create a sorted copy
	sorted := Sort(slice)
	length := len(sorted)

	// Calculate position (using linear interpolation)
	position := (p / 100.0) * float64(length-1)
	lower := int(math.Floor(position))
	upper := int(math.Ceil(position))

	if lower == upper {
		return float64(sorted[lower]), nil
	}

	// Linear interpolation between lower and upper
	lowerValue := float64(sorted[lower])
	upperValue := float64(sorted[upper])
	fraction := position - float64(lower)

	return lowerValue + (upperValue-lowerValue)*fraction, nil
}

// StandardDeviation calculates the standard deviation of a slice of numbers.
// Uses the population standard deviation formula (division by N).
// Returns an error if the slice is empty.
//
// StandardDeviation은 숫자 슬라이스의 표준 편차를 계산합니다.
// 모집단 표준 편차 공식을 사용합니다 (N으로 나눔).
// 슬라이스가 비어 있으면 에러를 반환합니다.
//
// Example:
//
//	numbers := []float64{2, 4, 4, 4, 5, 5, 7, 9}
//	stddev, err := sliceutil.StandardDeviation(numbers) // ~2.0
func StandardDeviation[T Number](slice []T) (float64, error) {
	variance, err := Variance(slice)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(variance), nil
}

// Variance calculates the variance of a slice of numbers.
// Uses the population variance formula (division by N).
// Returns an error if the slice is empty.
//
// Variance는 숫자 슬라이스의 분산을 계산합니다.
// 모집단 분산 공식을 사용합니다 (N으로 나눔).
// 슬라이스가 비어 있으면 에러를 반환합니다.
//
// Example:
//
//	numbers := []float64{2, 4, 4, 4, 5, 5, 7, 9}
//	variance, err := sliceutil.Variance(numbers) // 4.0
func Variance[T Number](slice []T) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot calculate variance of empty slice")
	}

	// Calculate mean
	mean := Average(slice)

	// Calculate sum of squared differences
	var sumSquaredDiff float64
	for _, v := range slice {
		diff := float64(v) - mean
		sumSquaredDiff += diff * diff
	}

	// Population variance (divide by N)
	return sumSquaredDiff / float64(len(slice)), nil
}

// MostCommon returns the n most frequently occurring elements in the slice, in descending order of frequency.
// If there are ties, elements are ordered by their first occurrence in the slice.
// If n is greater than the number of unique elements, returns all unique elements.
//
// MostCommon은 슬라이스에서 가장 자주 나타나는 n개의 요소를 빈도의 내림차순으로 반환합니다.
// 동점이 있는 경우 슬라이스에서 처음 나타난 순서대로 정렬됩니다.
// n이 고유 요소의 수보다 크면 모든 고유 요소를 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
//	top2 := sliceutil.MostCommon(numbers, 2) // [4, 3]
func MostCommon[T comparable](slice []T, n int) []T {
	if len(slice) == 0 || n <= 0 {
		return []T{}
	}

	// Get frequencies
	frequencies := Frequencies(slice)

	// Create a slice of unique elements with their frequencies
	type elementFreq struct {
		element   T
		frequency int
		firstSeen int
	}

	// Track first occurrence
	firstOccurrence := make(map[T]int)
	for i, v := range slice {
		if _, exists := firstOccurrence[v]; !exists {
			firstOccurrence[v] = i
		}
	}

	elements := make([]elementFreq, 0, len(frequencies))
	for element, freq := range frequencies {
		elements = append(elements, elementFreq{
			element:   element,
			frequency: freq,
			firstSeen: firstOccurrence[element],
		})
	}

	// Sort by frequency (descending), then by first occurrence (ascending)
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].frequency != elements[j].frequency {
			return elements[i].frequency > elements[j].frequency
		}
		return elements[i].firstSeen < elements[j].firstSeen
	})

	// Take top n elements
	limit := n
	if limit > len(elements) {
		limit = len(elements)
	}

	result := make([]T, limit)
	for i := 0; i < limit; i++ {
		result[i] = elements[i].element
	}

	return result
}

// LeastCommon returns the n least frequently occurring elements in the slice, in ascending order of frequency.
// If there are ties, elements are ordered by their first occurrence in the slice.
// If n is greater than the number of unique elements, returns all unique elements.
//
// LeastCommon은 슬라이스에서 가장 적게 나타나는 n개의 요소를 빈도의 오름차순으로 반환합니다.
// 동점이 있는 경우 슬라이스에서 처음 나타난 순서대로 정렬됩니다.
// n이 고유 요소의 수보다 크면 모든 고유 요소를 반환합니다.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
//	bottom2 := sliceutil.LeastCommon(numbers, 2) // [1, 2]
func LeastCommon[T comparable](slice []T, n int) []T {
	if len(slice) == 0 || n <= 0 {
		return []T{}
	}

	// Get frequencies
	frequencies := Frequencies(slice)

	// Create a slice of unique elements with their frequencies
	type elementFreq struct {
		element   T
		frequency int
		firstSeen int
	}

	// Track first occurrence
	firstOccurrence := make(map[T]int)
	for i, v := range slice {
		if _, exists := firstOccurrence[v]; !exists {
			firstOccurrence[v] = i
		}
	}

	elements := make([]elementFreq, 0, len(frequencies))
	for element, freq := range frequencies {
		elements = append(elements, elementFreq{
			element:   element,
			frequency: freq,
			firstSeen: firstOccurrence[element],
		})
	}

	// Sort by frequency (ascending), then by first occurrence (ascending)
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].frequency != elements[j].frequency {
			return elements[i].frequency < elements[j].frequency
		}
		return elements[i].firstSeen < elements[j].firstSeen
	})

	// Take bottom n elements
	limit := n
	if limit > len(elements) {
		limit = len(elements)
	}

	result := make([]T, limit)
	for i := 0; i < limit; i++ {
		result[i] = elements[i].element
	}

	return result
}
