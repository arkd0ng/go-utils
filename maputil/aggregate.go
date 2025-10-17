package maputil

// Reduce reduces a map to a single value using an accumulator function.
// Reduce는 누산기 함수를 사용하여 맵을 단일 값으로 축소합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	sum := maputil.Reduce(m, 0, func(acc int, k string, v int) int {
//	    return acc + v
//	}) // sum = 6
func Reduce[K comparable, V any, R any](m map[K]V, initial R, fn func(R, K, V) R) R {
	result := initial
	for k, v := range m {
		result = fn(result, k, v)
	}
	return result
}

// Sum calculates the sum of all numeric values in the map.
// Sum은 맵의 모든 숫자 값의 합을 계산합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	total := maputil.Sum(m) // total = 6
func Sum[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

// Min finds the key-value pair with the minimum value.
// Min은 최소값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	key, value, ok := maputil.Min(m) // key = "b", value = 1, ok = true
func Min[K comparable, V Ordered](m map[K]V) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var minKey K
	var minValue V
	first := true

	for k, v := range m {
		if first || v < minValue {
			minKey = k
			minValue = v
			first = false
		}
	}

	return minKey, minValue, true
}

// Max finds the key-value pair with the maximum value.
// Max는 최대값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 3, "b": 1, "c": 2}
//	key, value, ok := maputil.Max(m) // key = "a", value = 3, ok = true
func Max[K comparable, V Ordered](m map[K]V) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var maxKey K
	var maxValue V
	first := true

	for k, v := range m {
		if first || v > maxValue {
			maxKey = k
			maxValue = v
			first = false
		}
	}

	return maxKey, maxValue, true
}

// MinBy finds the key-value pair with the minimum value according to a custom function.
// MinBy는 사용자 정의 함수에 따라 최소값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; Age int }
//	m := map[string]User{
//	    "alice": {Name: "Alice", Age: 30},
//	    "bob":   {Name: "Bob", Age: 25},
//	}
//	key, user, ok := maputil.MinBy(m, func(u User) float64 {
//	    return float64(u.Age)
//	}) // key = "bob", user.Age = 25, ok = true
func MinBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var minKey K
	var minValue V
	var minScore float64
	first := true

	for k, v := range m {
		score := fn(v)
		if first || score < minScore {
			minKey = k
			minValue = v
			minScore = score
			first = false
		}
	}

	return minKey, minValue, true
}

// MaxBy finds the key-value pair with the maximum value according to a custom function.
// MaxBy는 사용자 정의 함수에 따라 최대값을 가진 키-값 쌍을 찾습니다.
//
// Returns the key, value, and true if the map is not empty.
// Returns zero values and false if the map is empty.
//
// 맵이 비어 있지 않으면 키, 값, true를 반환합니다.
// 맵이 비어 있으면 zero 값과 false를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; Score int }
//	m := map[string]User{
//	    "alice": {Name: "Alice", Score: 95},
//	    "bob":   {Name: "Bob", Score: 88},
//	}
//	key, user, ok := maputil.MaxBy(m, func(u User) float64 {
//	    return float64(u.Score)
//	}) // key = "alice", user.Score = 95, ok = true
func MaxBy[K comparable, V any](m map[K]V, fn func(V) float64) (K, V, bool) {
	if len(m) == 0 {
		var zeroK K
		var zeroV V
		return zeroK, zeroV, false
	}

	var maxKey K
	var maxValue V
	var maxScore float64
	first := true

	for k, v := range m {
		score := fn(v)
		if first || score > maxScore {
			maxKey = k
			maxValue = v
			maxScore = score
			first = false
		}
	}

	return maxKey, maxValue, true
}

// Average calculates the average of all numeric values in the map.
// Average는 맵의 모든 숫자 값의 평균을 계산합니다.
//
// Returns 0 if the map is empty.
// 맵이 비어 있으면 0을 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	avg := maputil.Average(m) // avg = 2.5
func Average[K comparable, V Number](m map[K]V) float64 {
	if len(m) == 0 {
		return 0
	}

	var sum V
	for _, v := range m {
		sum += v
	}

	return float64(sum) / float64(len(m))
}

// GroupBy groups a slice of elements by a key extracted from each element.
// GroupBy는 각 요소에서 추출한 키로 요소 슬라이스를 그룹화합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; City string }
//	users := []User{
//	    {Name: "Alice", City: "Seoul"},
//	    {Name: "Bob", City: "Seoul"},
//	    {Name: "Charlie", City: "Busan"},
//	}
//	byCity := maputil.GroupBy(users, func(u User) string {
//	    return u.City
//	})
//	// byCity = map[string][]User{
//	//     "Seoul": []User{{Name: "Alice", ...}, {Name: "Bob", ...}},
//	//     "Busan": []User{{Name: "Charlie", ...}},
//	// }
func GroupBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G][]V {
	result := make(map[G][]V)

	for _, item := range slice {
		key := fn(item)
		result[key] = append(result[key], item)
	}

	return result
}

// CountBy counts the number of elements in a slice for each key extracted from each element.
// CountBy는 각 요소에서 추출한 키별로 슬라이스의 요소 수를 계산합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { Name string; City string }
//	users := []User{
//	    {Name: "Alice", City: "Seoul"},
//	    {Name: "Bob", City: "Seoul"},
//	    {Name: "Charlie", City: "Busan"},
//	}
//	counts := maputil.CountBy(users, func(u User) string {
//	    return u.City
//	}) // counts = map[string]int{"Seoul": 2, "Busan": 1}
func CountBy[K comparable, V any, G comparable](slice []V, fn func(V) G) map[G]int {
	result := make(map[G]int)

	for _, item := range slice {
		key := fn(item)
		result[key]++
	}

	return result
}

// Median calculates the median of all numeric values in the map.
// Median는 맵의 모든 숫자 값의 중앙값을 계산합니다.
//
// This function collects all values, sorts them, and returns the middle value.
// For even-length maps, it returns the average of the two middle values.
// Returns (0, false) for empty maps.
//
// 이 함수는 모든 값을 수집하고 정렬한 후 중간 값을 반환합니다.
// 짝수 길이 맵의 경우 두 중간 값의 평균을 반환합니다.
// 빈 맵의 경우 (0, false)를 반환합니다.
//
// Time Complexity
// 시간 복잡도: O(n log n) due to sorting
// Space Complexity
// 공간 복잡도: O(n) for collecting values
//
// Parameters
// 매개변수:
// - m: The input map with numeric values
// 숫자 값이 있는 입력 맵
//
// Returns
// 반환값:
// - float64: The median value
// 중앙값
// - bool: false if map is empty, true otherwise
// 맵이 비어있으면 false, 그렇지 않으면 true
//
// Example
// 예제:
//
//	scores := map[string]int{
//		"Alice":   85,
//		"Bob":     90,
//		"Charlie": 75,
//		"Diana":   95,
//		"Eve":     80,
//	}
//	median, ok := maputil.Median(scores) // median = 85.0 (middle value when sorted)
//
//	evenScores := map[string]int{
//		"Alice": 80,
//		"Bob":   90,
//		"Charlie": 70,
//		"Diana": 100,
//	}
//	median2, ok2 := maputil.Median(evenScores) // median2 = 85.0 (average of 80 and 90)
//
// Use Case
// 사용 사례:
// - Statistical analysis
// 통계 분석
// - Grade distribution analysis
// 성적 분포 분석
// - Performance metrics
// 성능 메트릭
// - Finding typical values
// 대표값 찾기
func Median[K comparable, V Number](m map[K]V) (float64, bool) {
	if len(m) == 0 {
		return 0, false
	}

	// Collect all values into a slice
	values := make([]float64, 0, len(m))
	for _, v := range m {
		values = append(values, float64(v))
	}

	// Sort values
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}

	n := len(values)
	if n%2 == 1 {
		// Odd length: return middle value
		return values[n/2], true
	}

	// Even length: return average of two middle values
	mid1 := values[n/2-1]
	mid2 := values[n/2]
	return (mid1 + mid2) / 2, true
}

// Frequencies counts the occurrence of each unique value in the map.
// Frequencies는 맵의 각 고유 값의 출현 빈도를 계산합니다.
//
// This function inverts the map structure, creating a map where keys are the
// original values and values are their occurrence counts. Useful for finding
// duplicate values or analyzing value distributions.
//
// 이 함수는 맵 구조를 반전하여 키가 원래 값이고 값이 출현 횟수인 맵을 생성합니다.
// 중복 값을 찾거나 값 분포를 분석하는 데 유용합니다.
//
// Time Complexity
// 시간 복잡도: O(n)
// Space Complexity
// 공간 복잡도: O(u) where u is unique values
//
// Parameters
// 매개변수:
// - m: The input map
// 입력 맵
//
// Returns
// 반환값:
// - map[V]int: Map of value → count
// 값 → 개수 맵
//
// Example
// 예제:
//
//	grades := map[string]string{
//		"Alice":   "A",
//		"Bob":     "B",
//		"Charlie": "A",
//		"Diana":   "C",
//		"Eve":     "B",
//		"Frank":   "A",
//	}
//	freq := maputil.Frequencies(grades)
//	// freq = map[string]int{"A": 3, "B": 2, "C": 1}
//
//	scores := map[string]int{
//		"test1": 85,
//		"test2": 90,
//		"test3": 85,
//		"test4": 90,
//		"test5": 75,
//	}
//	scoreFreq := maputil.Frequencies(scores)
//	// scoreFreq = map[int]int{85: 2, 90: 2, 75: 1}
//
// Use Case
// 사용 사례:
// - Finding duplicate values
// 중복 값 찾기
// - Value distribution analysis
// 값 분포 분석
// - Histogram generation
// 히스토그램 생성
// - Data quality checks
// 데이터 품질 확인
func Frequencies[K comparable, V comparable](m map[K]V) map[V]int {
	result := make(map[V]int)

	for _, value := range m {
		result[value]++
	}

	return result
}
