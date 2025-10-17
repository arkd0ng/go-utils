package sliceutil

import (
	"errors"
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

// statistics.go provides statistical analysis operations for slices.
//
// This file implements statistical measures for analyzing data distributions,
// central tendencies, variability, and frequency patterns:
//
// Central Tendency Measures:
//
// Median (Middle Value):
//   - Median(slice): Middle value in sorted data
//     Time: O(n log n), Space: O(n) (requires sorting)
//     Odd length: Returns middle element
//     Even length: Returns average of two middle elements
//     Robust to outliers (better than mean for skewed data)
//     Example: Median([1,2,3,4,5]) = 3, Median([1,2,3,4]) = 2.5
//     Use cases: Income analysis, real estate prices, robust statistics
//
// Mode (Most Frequent):
//   - Mode(slice): Most frequently occurring element
//     Time: O(n), Space: O(n) (frequency map)
//     Returns first mode if multiple exist (ties)
//     Useful for categorical data
//     Empty slice returns error
//     Example: Mode([1,2,2,3,3,3,4]) = 3
//     Use cases: Survey responses, categorical data, popular choices
//
// Variability Measures:
//
// Variance (Dispersion):
//   - Variance(slice): Average squared deviation from mean
//     Time: O(n), Space: O(1)
//     Population variance: Σ(x - μ)² / N
//     Measures spread of data
//     Squared units (e.g., if data in meters, variance in meters²)
//     Example: Variance([2,4,4,4,5,5,7,9]) = 4.0
//     Use cases: Risk analysis, quality control, consistency measurement
//
// StandardDeviation (Spread):
//   - StandardDeviation(slice): Square root of variance
//     Time: O(n), Space: O(1)
//     Same units as original data
//     More interpretable than variance
//     σ = √variance
//     ~68% of data within 1σ of mean (normal distribution)
//     Example: StandardDeviation([2,4,4,4,5,5,7,9]) ≈ 2.0
//     Use cases: Volatility, measurement precision, outlier detection
//
// Percentile Analysis:
//
// Percentile (Rank Position):
//   - Percentile(slice, p): Value below which p% of data falls
//     Time: O(n log n), Space: O(n) (requires sorting)
//     p must be 0-100 inclusive
//     Uses linear interpolation for smooth results
//     50th percentile = Median
//     25th/75th percentiles = Quartiles
//     Example: Percentile([1..10], 75) = 7.75
//     Use cases: Performance benchmarks, grade cutoffs, SLA monitoring
//
// Frequency Analysis:
//
// Frequencies (Count Map):
//   - Frequencies(slice): Map of element → count
//     Time: O(n), Space: O(k) where k = unique elements
//     Builds frequency distribution
//     Foundation for Mode, MostCommon, LeastCommon
//     Example: Frequencies([1,2,2,3,3,3]) = {1:1, 2:2, 3:3}
//     Use cases: Histogram data, distribution analysis, word counts
//
// MostCommon (Top N Frequent):
//   - MostCommon(slice, n): N most frequent elements
//     Time: O(n + k log k), Space: O(k) where k = unique elements
//     Descending frequency order
//     Ties broken by first occurrence
//     Caps at number of unique elements
//     Example: MostCommon([1,2,2,3,3,3,4,4,4,4], 2) = [4,3]
//     Use cases: Top products, trending topics, frequent patterns
//
// LeastCommon (Bottom N Frequent):
//   - LeastCommon(slice, n): N least frequent elements
//     Time: O(n + k log k), Space: O(k) where k = unique elements
//     Ascending frequency order
//     Ties broken by first occurrence
//     Caps at number of unique elements
//     Example: LeastCommon([1,2,2,3,3,3,4,4,4,4], 2) = [1,2]
//     Use cases: Rare items, edge cases, anomaly detection
//
// Statistical Concepts:
//
// Population vs Sample:
//   - This implementation uses population formulas (divide by N)
//   - For sample statistics, divide by (N-1) - Bessel's correction
//   - Population: Analyzing entire dataset
//   - Sample: Estimating from subset of population
//
// Variance and Standard Deviation Relationship:
//   - StandardDeviation = √Variance
//   - Variance emphasizes larger deviations (squared)
//   - StdDev in same units as data (more interpretable)
//
// Median vs Mean:
//   - Median: Robust to outliers, better for skewed data
//   - Mean (Average): Affected by outliers, better for normal data
//   - Median: Use for income, house prices, skewed distributions
//   - Mean: Use for normally distributed continuous data
//
// Mode Characteristics:
//   - Can have multiple modes (bimodal, multimodal)
//   - This implementation returns first encountered mode
//   - Most useful for discrete/categorical data
//   - Less meaningful for continuous data with unique values
//
// Design Principles:
//   - Robust: Empty slice handling with errors
//   - Efficient: Single-pass algorithms where possible
//   - Standard: Uses common statistical formulas
//   - Type-safe: Number constraint for numeric operations
//   - Comparable: Frequency operations work with any comparable type
//
// Performance Characteristics:
//
// Time Complexity:
//   - Median: O(n log n) - requires sorting
//   - Mode: O(n) - frequency counting
//   - Variance/StdDev: O(n) - single pass after mean
//   - Percentile: O(n log n) - requires sorting
//   - Frequencies: O(n) - hash map building
//   - MostCommon/LeastCommon: O(n + k log k) - count + sort
//
// Space Complexity:
//   - Median/Percentile: O(n) - sorted copy
//   - Mode: O(k) - frequency map (k = unique elements)
//   - Variance/StdDev: O(1) - just accumulation
//   - Frequencies: O(k) - map storage
//   - MostCommon/LeastCommon: O(k) - unique elements + sort
//
// Common Usage Patterns:
//
//	// Analyze test scores
//	scores := []float64{65, 70, 75, 80, 85, 90, 95, 100}
//	median, _ := sliceutil.Median(scores)
//	stddev, _ := sliceutil.StandardDeviation(scores)
//	p90, _ := sliceutil.Percentile(scores, 90)
//	fmt.Printf("Median: %.1f, StdDev: %.1f, 90th%%: %.1f\n", median, stddev, p90)
//
//	// Find popular items
//	purchases := []string{"apple", "banana", "apple", "cherry", "apple", "banana"}
//	topSellers := sliceutil.MostCommon(purchases, 3)
//	fmt.Printf("Top sellers: %v\n", topSellers)
//
//	// Analyze variability
//	measurements := []float64{10.1, 10.2, 9.9, 10.0, 10.3}
//	variance, _ := sliceutil.Variance(measurements)
//	if variance < 0.1 {
//	    fmt.Println("Low variability - consistent measurements")
//	}
//
//	// Frequency distribution
//	votes := []int{1, 2, 1, 3, 2, 1, 4, 2, 1}
//	freqs := sliceutil.Frequencies(votes)
//	for option, count := range freqs {
//	    fmt.Printf("Option %d: %d votes\n", option, count)
//	}
//
// Statistical Analysis Patterns:
//
//	// Five-number summary (Tukey's summary)
//	min, _ := sliceutil.Min(data)
//	q1, _ := sliceutil.Percentile(data, 25)
//	median, _ := sliceutil.Median(data)
//	q3, _ := sliceutil.Percentile(data, 75)
//	max, _ := sliceutil.Max(data)
//	fmt.Printf("Min: %.1f, Q1: %.1f, Median: %.1f, Q3: %.1f, Max: %.1f\n",
//	    min, q1, median, q3, max)
//
//	// Detect outliers using IQR method
//	q1, _ := sliceutil.Percentile(data, 25)
//	q3, _ := sliceutil.Percentile(data, 75)
//	iqr := q3 - q1
//	lowerBound := q1 - 1.5*iqr
//	upperBound := q3 + 1.5*iqr
//	outliers := sliceutil.Filter(data, func(x float64) bool {
//	    return x < lowerBound || x > upperBound
//	})
//
// Error Handling:
//   - Empty slice: Most functions return errors
//   - Invalid parameters: Percentile validates p range (0-100)
//   - Frequencies: Returns empty map for empty slice (no error)
//   - MostCommon/LeastCommon: Returns empty slice for invalid n
//
// Comparison with Standard Library:
//   - Go standard library has limited statistics support
//   - math package provides basic functions (sqrt, etc.)
//   - This package provides higher-level statistical measures
//   - For advanced statistics, consider external libraries
//
// statistics.go는 슬라이스에 대한 통계 분석 작업을 제공합니다.
//
// 이 파일은 데이터 분포, 중심 경향, 변동성 및 빈도 패턴을 분석하기 위한
// 통계적 측정을 구현합니다:
//
// 중심 경향 측정:
//
// Median (중간값):
//   - Median(slice): 정렬된 데이터의 중간값
//     시간: O(n log n), 공간: O(n) (정렬 필요)
//     홀수 길이: 중간 요소 반환
//     짝수 길이: 두 중간 요소의 평균 반환
//     이상값에 강건 (치우친 데이터에 평균보다 나음)
//     예: Median([1,2,3,4,5]) = 3, Median([1,2,3,4]) = 2.5
//     사용 사례: 소득 분석, 부동산 가격, 강건한 통계
//
// Mode (최빈값):
//   - Mode(slice): 가장 자주 나타나는 요소
//     시간: O(n), 공간: O(n) (빈도 맵)
//     여러 개 존재 시 첫 번째 반환 (동점)
//     범주형 데이터에 유용
//     빈 슬라이스는 에러 반환
//     예: Mode([1,2,2,3,3,3,4]) = 3
//     사용 사례: 설문 응답, 범주형 데이터, 인기 선택
//
// 변동성 측정:
//
// Variance (분산):
//   - Variance(slice): 평균으로부터 제곱 편차의 평균
//     시간: O(n), 공간: O(1)
//     모집단 분산: Σ(x - μ)² / N
//     데이터 분산 측정
//     제곱 단위 (예: 데이터가 미터면 분산은 미터²)
//     예: Variance([2,4,4,4,5,5,7,9]) = 4.0
//     사용 사례: 위험 분석, 품질 관리, 일관성 측정
//
// StandardDeviation (표준편차):
//   - StandardDeviation(slice): 분산의 제곱근
//     시간: O(n), 공간: O(1)
//     원래 데이터와 같은 단위
//     분산보다 해석 용이
//     σ = √분산
//     정규 분포에서 데이터의 ~68%가 평균의 1σ 내
//     예: StandardDeviation([2,4,4,4,5,5,7,9]) ≈ 2.0
//     사용 사례: 변동성, 측정 정밀도, 이상값 감지
//
// 백분위 분석:
//
// Percentile (순위 위치):
//   - Percentile(slice, p): 데이터의 p%가 그 아래에 있는 값
//     시간: O(n log n), 공간: O(n) (정렬 필요)
//     p는 0-100 사이여야 함
//     부드러운 결과를 위해 선형 보간 사용
//     50번째 백분위 = 중간값
//     25/75 백분위 = 사분위수
//     예: Percentile([1..10], 75) = 7.75
//     사용 사례: 성능 벤치마크, 등급 컷오프, SLA 모니터링
//
// 빈도 분석:
//
// Frequencies (개수 맵):
//   - Frequencies(slice): 요소 → 개수 맵
//     시간: O(n), 공간: O(k) (k = 고유 요소)
//     빈도 분포 구축
//     Mode, MostCommon, LeastCommon의 기반
//     예: Frequencies([1,2,2,3,3,3]) = {1:1, 2:2, 3:3}
//     사용 사례: 히스토그램 데이터, 분포 분석, 단어 개수
//
// MostCommon (상위 N 빈도):
//   - MostCommon(slice, n): 가장 빈번한 N개 요소
//     시간: O(n + k log k), 공간: O(k) (k = 고유 요소)
//     빈도 내림차순
//     동점은 첫 발생으로 결정
//     고유 요소 수로 제한
//     예: MostCommon([1,2,2,3,3,3,4,4,4,4], 2) = [4,3]
//     사용 사례: 인기 제품, 트렌딩 주제, 빈번한 패턴
//
// LeastCommon (하위 N 빈도):
//   - LeastCommon(slice, n): 가장 드문 N개 요소
//     시간: O(n + k log k), 공간: O(k) (k = 고유 요소)
//     빈도 오름차순
//     동점은 첫 발생으로 결정
//     고유 요소 수로 제한
//     예: LeastCommon([1,2,2,3,3,3,4,4,4,4], 2) = [1,2]
//     사용 사례: 희귀 항목, 엣지 케이스, 이상 감지
//
// 통계 개념:
//
// 모집단 vs 표본:
//   - 이 구현은 모집단 공식 사용 (N으로 나눔)
//   - 표본 통계는 (N-1)로 나눔 - 베셀 보정
//   - 모집단: 전체 데이터셋 분석
//   - 표본: 모집단의 부분집합에서 추정
//
// 분산과 표준편차 관계:
//   - 표준편차 = √분산
//   - 분산은 큰 편차 강조 (제곱)
//   - 표준편차는 데이터와 같은 단위 (더 해석 용이)
//
// 중간값 vs 평균:
//   - 중간값: 이상값에 강건, 치우친 데이터에 나음
//   - 평균: 이상값 영향, 정규 데이터에 나음
//   - 중간값: 소득, 집값, 치우친 분포에 사용
//   - 평균: 정규 분포 연속 데이터에 사용
//
// 최빈값 특성:
//   - 여러 최빈값 가능 (이봉, 다봉)
//   - 이 구현은 첫 번째 최빈값 반환
//   - 이산/범주형 데이터에 가장 유용
//   - 고유 값의 연속 데이터엔 덜 의미 있음
//
// 설계 원칙:
//   - 강건함: 에러로 빈 슬라이스 처리
//   - 효율적: 가능한 곳에서 단일 패스 알고리즘
//   - 표준: 일반적인 통계 공식 사용
//   - 타입 안전: 숫자 작업을 위한 Number 제약
//   - Comparable: 빈도 작업은 모든 comparable 타입 작동
//
// 성능 특성:
//
// 시간 복잡도:
//   - Median: O(n log n) - 정렬 필요
//   - Mode: O(n) - 빈도 계산
//   - Variance/StdDev: O(n) - 평균 후 단일 패스
//   - Percentile: O(n log n) - 정렬 필요
//   - Frequencies: O(n) - 해시 맵 구축
//   - MostCommon/LeastCommon: O(n + k log k) - 개수 + 정렬
//
// 공간 복잡도:
//   - Median/Percentile: O(n) - 정렬된 복사본
//   - Mode: O(k) - 빈도 맵 (k = 고유 요소)
//   - Variance/StdDev: O(1) - 누적만
//   - Frequencies: O(k) - 맵 저장
//   - MostCommon/LeastCommon: O(k) - 고유 요소 + 정렬
//
// 일반적인 사용 패턴:
//
//	// 시험 점수 분석
//	scores := []float64{65, 70, 75, 80, 85, 90, 95, 100}
//	median, _ := sliceutil.Median(scores)
//	stddev, _ := sliceutil.StandardDeviation(scores)
//	p90, _ := sliceutil.Percentile(scores, 90)
//	fmt.Printf("중간값: %.1f, 표준편차: %.1f, 90%%: %.1f\n", median, stddev, p90)
//
//	// 인기 항목 찾기
//	purchases := []string{"사과", "바나나", "사과", "체리", "사과", "바나나"}
//	topSellers := sliceutil.MostCommon(purchases, 3)
//	fmt.Printf("인기 상품: %v\n", topSellers)
//
//	// 변동성 분석
//	measurements := []float64{10.1, 10.2, 9.9, 10.0, 10.3}
//	variance, _ := sliceutil.Variance(measurements)
//	if variance < 0.1 {
//	    fmt.Println("낮은 변동성 - 일관된 측정")
//	}
//
//	// 빈도 분포
//	votes := []int{1, 2, 1, 3, 2, 1, 4, 2, 1}
//	freqs := sliceutil.Frequencies(votes)
//	for option, count := range freqs {
//	    fmt.Printf("옵션 %d: %d표\n", option, count)
//	}
//
// 통계 분석 패턴:
//
//	// 5개 수 요약 (Tukey 요약)
//	min, _ := sliceutil.Min(data)
//	q1, _ := sliceutil.Percentile(data, 25)
//	median, _ := sliceutil.Median(data)
//	q3, _ := sliceutil.Percentile(data, 75)
//	max, _ := sliceutil.Max(data)
//	fmt.Printf("최소: %.1f, Q1: %.1f, 중간: %.1f, Q3: %.1f, 최대: %.1f\n",
//	    min, q1, median, q3, max)
//
//	// IQR 방법으로 이상값 감지
//	q1, _ := sliceutil.Percentile(data, 25)
//	q3, _ := sliceutil.Percentile(data, 75)
//	iqr := q3 - q1
//	lowerBound := q1 - 1.5*iqr
//	upperBound := q3 + 1.5*iqr
//	outliers := sliceutil.Filter(data, func(x float64) bool {
//	    return x < lowerBound || x > upperBound
//	})
//
// 에러 처리:
//   - 빈 슬라이스: 대부분 함수가 에러 반환
//   - 잘못된 매개변수: Percentile은 p 범위 (0-100) 검증
//   - Frequencies: 빈 슬라이스에 빈 맵 반환 (에러 없음)
//   - MostCommon/LeastCommon: 잘못된 n에 빈 슬라이스 반환
//
// 표준 라이브러리와 비교:
//   - Go 표준 라이브러리는 제한적인 통계 지원
//   - math 패키지는 기본 함수 제공 (sqrt 등)
//   - 이 패키지는 더 높은 수준의 통계 측정 제공
//   - 고급 통계는 외부 라이브러리 고려

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
