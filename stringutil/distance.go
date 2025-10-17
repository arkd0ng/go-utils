package stringutil

import "math"

// =============================================================================
// File: distance.go
// Purpose: String Distance and Similarity Calculations
// 파일: distance.go
// 목적: 문자열 거리 및 유사도 계산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The distance.go file provides algorithms for measuring string similarity
// and dissimilarity using various distance metrics. These algorithms are essential
// for fuzzy string matching, spell checking, duplicate detection, and approximate
// string search. Each metric has different characteristics and use cases, from
// the general-purpose Levenshtein distance to the prefix-sensitive Jaro-Winkler
// similarity.
//
// distance.go 파일은 다양한 거리 메트릭을 사용하여 문자열 유사도와 비유사도를
// 측정하는 알고리즘을 제공합니다. 이러한 알고리즘은 퍼지 문자열 매칭, 맞춤법
// 검사, 중복 감지 및 근사 문자열 검색에 필수적입니다. 각 메트릭은 범용
// Levenshtein 거리부터 접두사 민감 Jaro-Winkler 유사도까지 다른 특성과 사용
// 사례를 가지고 있습니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Multiple Metrics**: Provide various distance/similarity algorithms
//    **다중 메트릭**: 다양한 거리/유사도 알고리즘 제공
//
// 2. **Unicode-Safe**: All algorithms use rune-based indexing
//    **유니코드 안전**: 모든 알고리즘은 룬 기반 인덱싱 사용
//
// 3. **Normalized Values**: Similarity functions return 0.0-1.0 range
//    **정규화된 값**: 유사도 함수는 0.0-1.0 범위 반환
//
// 4. **Optimized Implementations**: Use dynamic programming and early exits
//    **최적화된 구현**: 동적 프로그래밍 및 조기 종료 사용
//
// 5. **Clear Semantics**: Distance (higher = more different) vs Similarity (higher = more alike)
//    **명확한 의미**: 거리 (높을수록 더 다름) vs 유사도 (높을수록 더 유사)
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. EDIT DISTANCE (편집 거리)
//    - LevenshteinDistance: Minimum edit operations (insert, delete, substitute)
//      LevenshteinDistance: 최소 편집 연산 (삽입, 삭제, 치환)
//    - HammingDistance: Character position differences (equal-length only)
//      HammingDistance: 문자 위치 차이 (동일 길이만)
//
// 2. SIMILARITY RATIOS (유사도 비율)
//    - Similarity: Normalized Levenshtein similarity (0.0-1.0)
//      Similarity: 정규화된 Levenshtein 유사도 (0.0-1.0)
//    - JaroWinklerSimilarity: Prefix-weighted similarity (0.0-1.0)
//      JaroWinklerSimilarity: 접두사 가중 유사도 (0.0-1.0)
//
// 3. INTERNAL HELPERS (내부 헬퍼)
//    - jaroSimilarity: Base Jaro similarity calculation
//      jaroSimilarity: 기본 Jaro 유사도 계산
//    - min: Minimum of three integers
//      min: 세 정수의 최소값
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// LevenshteinDistance(a, b string) int
// - Purpose: Calculate minimum edit operations to transform a into b
// - 목적: a를 b로 변환하는 최소 편집 연산 계산
// - Algorithm: Dynamic programming with 2D matrix
// - 알고리즘: 2D 행렬을 사용한 동적 프로그래밍
// - Time Complexity: O(n * m) where n, m are lengths
// - 시간 복잡도: O(n * m), n, m은 길이
// - Space Complexity: O(n * m) for matrix
// - 공간 복잡도: O(n * m), 행렬용
// - Edit Operations: Insertion, deletion, substitution (each cost = 1)
// - 편집 연산: 삽입, 삭제, 치환 (각 비용 = 1)
// - Use Cases: Spell checking, DNA sequence analysis, diff algorithms, fuzzy search
// - 사용 사례: 맞춤법 검사, DNA 서열 분석, diff 알고리즘, 퍼지 검색
//
// Similarity(a, b string) float64
// - Purpose: Calculate normalized similarity ratio (0.0-1.0)
// - 목적: 정규화된 유사도 비율 계산 (0.0-1.0)
// - Formula: 1 - (distance / max(len(a), len(b)))
// - 공식: 1 - (거리 / max(len(a), len(b)))
// - Time Complexity: O(n * m) - uses LevenshteinDistance
// - 시간 복잡도: O(n * m) - LevenshteinDistance 사용
// - Space Complexity: O(n * m)
// - 공간 복잡도: O(n * m)
// - Return Values: 1.0 = identical, 0.0 = completely different
// - 반환 값: 1.0 = 동일, 0.0 = 완전히 다름
// - Use Cases: Duplicate detection, ranking search results, similarity thresholds
// - 사용 사례: 중복 감지, 검색 결과 순위, 유사도 임계값
//
// HammingDistance(a, b string) int
// - Purpose: Count differing positions in equal-length strings
// - 목적: 동일 길이 문자열에서 다른 위치 개수 계산
// - Time Complexity: O(n) - single pass
// - 시간 복잡도: O(n) - 단일 패스
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Constraint: Strings must have equal length (returns -1 otherwise)
// - 제약: 문자열이 같은 길이여야 함 (그렇지 않으면 -1 반환)
// - Use Cases: Error detection codes, cryptography, bioinformatics (fixed-length sequences)
// - 사용 사례: 오류 감지 코드, 암호학, 생물정보학 (고정 길이 서열)
//
// JaroWinklerSimilarity(a, b string) float64
// - Purpose: Calculate prefix-sensitive similarity (0.0-1.0)
// - 목적: 접두사 민감 유사도 계산 (0.0-1.0)
// - Algorithm: Jaro similarity + prefix bonus (up to 4 chars)
// - 알고리즘: Jaro 유사도 + 접두사 보너스 (최대 4문자)
// - Time Complexity: O(n * m) - match window search
// - 시간 복잡도: O(n * m) - 매치 윈도우 검색
// - Space Complexity: O(n + m) - match arrays
// - 공간 복잡도: O(n + m) - 매치 배열
// - Prefix Weight: 0.1 multiplier for common prefix
// - 접두사 가중치: 공통 접두사에 0.1 곱셈
// - Use Cases: Name matching, record linkage, typo tolerance with common prefixes
// - 사용 사례: 이름 매칭, 레코드 연결, 공통 접두사를 가진 타이포 허용
//
// jaroSimilarity(a, b string) float64
// - Purpose: Calculate base Jaro similarity without prefix bonus
// - 목적: 접두사 보너스 없이 기본 Jaro 유사도 계산
// - Algorithm: Matching characters within window, transpositions
// - 알고리즘: 윈도우 내 일치 문자, 전치
// - Match Window: max(len(a), len(b))/2 - 1
// - 매치 윈도우: max(len(a), len(b))/2 - 1
// - Formula: (matches/len(a) + matches/len(b) + (matches-transpositions/2)/matches) / 3
// - 공식: (matches/len(a) + matches/len(b) + (matches-transpositions/2)/matches) / 3
// - Internal: Used by JaroWinklerSimilarity
// - 내부: JaroWinklerSimilarity에서 사용
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - LevenshteinDistance: O(n * m) - full matrix computation
//   LevenshteinDistance: O(n * m) - 전체 행렬 계산
// - Similarity: O(n * m) - wraps LevenshteinDistance
//   Similarity: O(n * m) - LevenshteinDistance 래핑
// - HammingDistance: O(n) - single pass, linear time
//   HammingDistance: O(n) - 단일 패스, 선형 시간
// - JaroWinklerSimilarity: O(n * m) - match window iteration
//   JaroWinklerSimilarity: O(n * m) - 매치 윈도우 반복
// - jaroSimilarity: O(n * m) - match finding
//   jaroSimilarity: O(n * m) - 매치 찾기
//
// Space Complexities:
// 공간 복잡도:
// - LevenshteinDistance: O(n * m) - dynamic programming matrix
//   LevenshteinDistance: O(n * m) - 동적 프로그래밍 행렬
// - Similarity: O(n * m) - same as Levenshtein
//   Similarity: O(n * m) - Levenshtein과 동일
// - HammingDistance: O(1) - no extra space
//   HammingDistance: O(1) - 추가 공간 없음
// - JaroWinklerSimilarity: O(n + m) - match arrays
//   JaroWinklerSimilarity: O(n + m) - 매치 배열
// - jaroSimilarity: O(n + m) - boolean arrays
//   jaroSimilarity: O(n + m) - 불린 배열
//
// Optimization Opportunities:
// 최적화 기회:
// 1. LevenshteinDistance: Use O(min(n, m)) space with single-row optimization
//    LevenshteinDistance: 단일 행 최적화로 O(min(n, m)) 공간 사용
// 2. Early exit on identical strings (already implemented for some)
//    동일 문자열 시 조기 종료 (일부 이미 구현됨)
// 3. For repeated calculations, consider memoization
//    반복 계산의 경우 메모이제이션 고려
// 4. HammingDistance is fastest - use when strings are equal length
//    HammingDistance가 가장 빠름 - 문자열이 같은 길이일 때 사용
// 5. For very long strings, consider approximate algorithms
//    매우 긴 문자열의 경우 근사 알고리즘 고려
//
// Performance Tips:
// 성능 팁:
// 1. Check equality first (== operator) before computing distance
//    거리 계산 전에 먼저 동등성 확인 (== 연산자)
// 2. Use HammingDistance for fixed-length strings (fastest)
//    고정 길이 문자열에는 HammingDistance 사용 (가장 빠름)
// 3. Use JaroWinkler for name matching (prefix-sensitive)
//    이름 매칭에는 JaroWinkler 사용 (접두사 민감)
// 4. Use Levenshtein/Similarity for general fuzzy matching
//    일반 퍼지 매칭에는 Levenshtein/Similarity 사용
// 5. For large datasets, prefilter by length before computing distance
//    큰 데이터셋의 경우 거리 계산 전에 길이로 사전 필터링
//
// ALGORITHM CHARACTERISTICS
// 알고리즘 특성
// -------------------------
//
// Levenshtein Distance:
// - **Symmetric**: distance(a, b) == distance(b, a)
//   **대칭**: distance(a, b) == distance(b, a)
// - **Triangle Inequality**: distance(a, c) <= distance(a, b) + distance(b, c)
//   **삼각 부등식**: distance(a, c) <= distance(a, b) + distance(b, c)
// - **Zero Distance**: distance(a, a) == 0
//   **제로 거리**: distance(a, a) == 0
// - **Length Bound**: distance(a, b) <= max(len(a), len(b))
//   **길이 경계**: distance(a, b) <= max(len(a), len(b))
// - **Use Case**: General-purpose edit distance
//   **사용 사례**: 범용 편집 거리
//
// Hamming Distance:
// - **Constraint**: Only for equal-length strings
//   **제약**: 동일 길이 문자열만
// - **Faster**: O(n) vs O(n²) for Levenshtein
//   **더 빠름**: Levenshtein의 O(n²) vs O(n)
// - **Position-Sensitive**: Order matters
//   **위치 민감**: 순서 중요
// - **Use Case**: Error detection, fixed-length codes
//   **사용 사례**: 오류 감지, 고정 길이 코드
//
// Jaro-Winkler Similarity:
// - **Prefix Bonus**: Common prefixes increase similarity
//   **접두사 보너스**: 공통 접두사가 유사도 증가
// - **Match Window**: Allows character transpositions within range
//   **매치 윈도우**: 범위 내 문자 전치 허용
// - **Name-Friendly**: Designed for personal names
//   **이름 친화적**: 개인 이름에 맞게 설계됨
// - **Asymmetric**: Can favor shorter strings
//   **비대칭**: 짧은 문자열을 선호할 수 있음
// - **Use Case**: Record linkage, name matching
//   **사용 사례**: 레코드 연결, 이름 매칭
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - LevenshteinDistance("", "") = 0
//   LevenshteinDistance("", "") = 0
// - LevenshteinDistance("hello", "") = 5 (length of non-empty)
//   LevenshteinDistance("hello", "") = 5 (비어있지 않은 것의 길이)
// - Similarity("", "") = 1.0 (identical)
//   Similarity("", "") = 1.0 (동일)
// - HammingDistance("", "") = 0
//   HammingDistance("", "") = 0
// - JaroWinklerSimilarity("", "") = 1.0
//   JaroWinklerSimilarity("", "") = 1.0
//
// Equal Strings:
// 동일 문자열:
// - LevenshteinDistance(s, s) = 0
//   LevenshteinDistance(s, s) = 0
// - Similarity(s, s) = 1.0
//   Similarity(s, s) = 1.0
// - HammingDistance(s, s) = 0
//   HammingDistance(s, s) = 0
// - JaroWinklerSimilarity(s, s) = 1.0
//   JaroWinklerSimilarity(s, s) = 1.0
//
// Length Mismatches:
// 길이 불일치:
// - HammingDistance returns -1 for different lengths
//   HammingDistance는 다른 길이에 대해 -1 반환
// - Levenshtein/Similarity/JaroWinkler handle any lengths
//   Levenshtein/Similarity/JaroWinkler는 모든 길이 처리
//
// Unicode Handling:
// 유니코드 처리:
// - All algorithms use []rune conversion for proper Unicode support
//   모든 알고리즘은 적절한 유니코드 지원을 위해 []rune 변환 사용
// - Multi-byte characters counted as single rune
//   다중 바이트 문자는 단일 룬으로 계산
// - Example: "你好" has 2 runes, not 6 bytes
//   예: "你好"는 6바이트가 아닌 2개의 룬
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Spell Checking with Levenshtein
//    Levenshtein으로 맞춤법 검사:
//
//    userInput := "helllo"
//    dictionary := []string{"hello", "help", "hill"}
//    for _, word := range dictionary {
//        distance := stringutil.LevenshteinDistance(userInput, word)
//        if distance <= 2 {
//            fmt.Printf("Did you mean '%s'? (distance: %d)\n", word, distance)
//            // "Did you mean 'hello'? (distance: 1)"
//        }
//    }
//    // Find words within edit distance threshold
//    // 편집 거리 임계값 내의 단어 찾기
//
// 2. Duplicate Detection with Similarity
//    Similarity로 중복 감지:
//
//    entries := []string{"John Smith", "Jon Smith", "Jane Doe"}
//    threshold := 0.8
//    for i := 0; i < len(entries); i++ {
//        for j := i+1; j < len(entries); j++ {
//            sim := stringutil.Similarity(entries[i], entries[j])
//            if sim >= threshold {
//                fmt.Printf("Possible duplicate: '%s' ~ '%s' (%.2f)\n",
//                    entries[i], entries[j], sim)
//                // "Possible duplicate: 'John Smith' ~ 'Jon Smith' (0.90)"
//            }
//        }
//    }
//    // Detect near-duplicate entries
//    // 거의 중복된 항목 감지
//
// 3. Fixed-Length Code Validation with Hamming
//    Hamming으로 고정 길이 코드 검증:
//
//    code1 := "ABC123"
//    code2 := "ABC124"
//    distance := stringutil.HammingDistance(code1, code2)
//    if distance == 1 {
//        fmt.Println("Single bit error detected")
//    }
//    // Error detection in fixed-length codes
//    // 고정 길이 코드의 오류 감지
//
// 4. Name Matching with Jaro-Winkler
//    Jaro-Winkler로 이름 매칭:
//
//    name1 := "Martha"
//    name2 := "Marhta"  // Typo with transposition
//    similarity := stringutil.JaroWinklerSimilarity(name1, name2)
//    // similarity ≈ 0.961 (very high despite typo)
//    // 유사도 ≈ 0.961 (오타에도 불구하고 매우 높음)
//    if similarity > 0.9 {
//        fmt.Println("Likely same person")
//    }
//    // Robust to typos and transpositions
//    // 오타 및 전치에 강함
//
// 5. Fuzzy Search Ranking
//    퍼지 검색 순위:
//
//    query := "python"
//    candidates := []string{"python", "jython", "cython", "java", "ruby"}
//    type Result struct {
//        Word string
//        Score float64
//    }
//    results := []Result{}
//    for _, candidate := range candidates {
//        score := stringutil.Similarity(query, candidate)
//        results = append(results, Result{candidate, score})
//    }
//    // Sort by score descending
//    // 점수 내림차순 정렬
//    // Results: python (1.0), jython (0.83), cython (0.83), ...
//    // 결과: python (1.0), jython (0.83), cython (0.83), ...
//
// 6. Threshold-Based Matching
//    임계값 기반 매칭:
//
//    input := "color"
//    reference := "colour"  // British vs American spelling
//    distance := stringutil.LevenshteinDistance(input, reference)
//    // distance = 1 (one insertion)
//    if distance <= 1 {
//        fmt.Println("Accept as equivalent")
//    }
//    // Allow minor variations
//    // 사소한 변형 허용
//
// 7. DNA Sequence Comparison
//    DNA 서열 비교:
//
//    seq1 := "AGCTTAGC"
//    seq2 := "AGCTTGGC"
//    distance := stringutil.HammingDistance(seq1, seq2)
//    // distance = 1 (one position different)
//    // 거리 = 1 (한 위치 다름)
//    if distance >= 0 {
//        fmt.Printf("%d mutations\n", distance)
//    }
//    // Count point mutations in equal-length sequences
//    // 동일 길이 서열의 점 돌연변이 개수
//
// 8. Autocomplete Suggestions
//    자동완성 제안:
//
//    partial := "hel"
//    words := []string{"hello", "help", "helm", "helicopter"}
//    suggestions := []string{}
//    for _, word := range words {
//        if strings.HasPrefix(word, partial) {
//            distance := stringutil.LevenshteinDistance(partial, word)
//            suggestions = append(suggestions, word)
//        }
//    }
//    // Combine prefix matching with distance for ranking
//    // 순위를 위해 접두사 매칭과 거리 결합
//
// 9. Record Linkage Threshold
//    레코드 연결 임계값:
//
//    record1 := "John A. Smith"
//    record2 := "Jon A Smith"
//    jwScore := stringutil.JaroWinklerSimilarity(record1, record2)
//    // jwScore ≈ 0.93 (high due to common prefix)
//    if jwScore > 0.85 {
//        fmt.Println("Likely same person")
//    }
//    // Jaro-Winkler good for name variations
//    // Jaro-Winkler는 이름 변형에 적합
//
// 10. Multi-Metric Comparison
//     다중 메트릭 비교:
//
//     s1 := "kitten"
//     s2 := "sitting"
//     levDist := stringutil.LevenshteinDistance(s1, s2)  // 3
//     sim := stringutil.Similarity(s1, s2)                // 0.571
//     jwSim := stringutil.JaroWinklerSimilarity(s1, s2)   // 0.746
//     // Different metrics give different perspectives
//     // 다른 메트릭은 다른 관점 제공
//     // Use appropriate metric for your use case
//     // 사용 사례에 적합한 메트릭 사용
//
// COMPARISON BETWEEN METRICS
// 메트릭 간 비교
// --------------------------
//
// Levenshtein vs Hamming:
// - Levenshtein: Any length, allows insertion/deletion
//   Levenshtein: 모든 길이, 삽입/삭제 허용
// - Hamming: Equal length only, position-based
//   Hamming: 동일 길이만, 위치 기반
// - Use Levenshtein for: Variable-length strings
//   Levenshtein 사용: 가변 길이 문자열
// - Use Hamming for: Fixed-length codes, faster
//   Hamming 사용: 고정 길이 코드, 더 빠름
//
// Levenshtein vs Jaro-Winkler:
// - Levenshtein: Symmetric, all edits equal cost
//   Levenshtein: 대칭, 모든 편집 비용 동일
// - Jaro-Winkler: Prefix-weighted, transposition-friendly
//   Jaro-Winkler: 접두사 가중, 전치 친화적
// - Use Levenshtein for: General fuzzy matching
//   Levenshtein 사용: 일반 퍼지 매칭
// - Use Jaro-Winkler for: Name matching, record linkage
//   Jaro-Winkler 사용: 이름 매칭, 레코드 연결
//
// Similarity vs JaroWinklerSimilarity:
// - Similarity: Based on edit distance
//   Similarity: 편집 거리 기반
// - JaroWinklerSimilarity: Based on character matches and transpositions
//   JaroWinklerSimilarity: 문자 일치 및 전치 기반
// - Similarity: Better for general text
//   Similarity: 일반 텍스트에 더 좋음
// - JaroWinklerSimilarity: Better for names with typos
//   JaroWinklerSimilarity: 오타가 있는 이름에 더 좋음
//
// Distance vs Similarity Functions:
// - Distance: Higher value = more different (0 = identical)
//   Distance: 높은 값 = 더 다름 (0 = 동일)
// - Similarity: Higher value = more alike (1.0 = identical, 0.0 = completely different)
//   Similarity: 높은 값 = 더 유사 (1.0 = 동일, 0.0 = 완전히 다름)
// - Use Distance for: Absolute edit count
//   Distance 사용: 절대 편집 개수
// - Use Similarity for: Normalized comparison (0.0-1.0)
//   Similarity 사용: 정규화된 비교 (0.0-1.0)
//
// TYPICAL THRESHOLDS
// 일반적인 임계값
// ------------------
// Different applications require different similarity/distance thresholds:
// 다른 애플리케이션은 다른 유사도/거리 임계값을 필요로 합니다:
//
// Levenshtein Distance:
// - 0-1: Very similar (typo tolerance)
//   0-1: 매우 유사 (오타 허용)
// - 2-3: Similar (fuzzy matching)
//   2-3: 유사 (퍼지 매칭)
// - 4+: Different
//   4+: 다름
//
// Similarity Ratio:
// - 0.9-1.0: Very similar (near-duplicates)
//   0.9-1.0: 매우 유사 (거의 중복)
// - 0.7-0.9: Similar (potential matches)
//   0.7-0.9: 유사 (잠재적 일치)
// - 0.5-0.7: Somewhat similar
//   0.5-0.7: 다소 유사
// - 0.0-0.5: Different
//   0.0-0.5: 다름
//
// Jaro-Winkler Similarity:
// - 0.9-1.0: Very likely same entity (name matching)
//   0.9-1.0: 동일 개체일 가능성 매우 높음 (이름 매칭)
// - 0.8-0.9: Likely same entity
//   0.8-0.9: 동일 개체일 가능성 높음
// - 0.7-0.8: Possibly same entity
//   0.7-0.8: 동일 개체 가능성 있음
// - 0.0-0.7: Different entities
//   0.0-0.7: 다른 개체
//
// Note: Thresholds should be tuned based on your specific data and requirements.
// 참고: 임계값은 특정 데이터 및 요구사항에 따라 조정되어야 합니다.
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe as they operate on immutable strings
// and don't use shared mutable state. Each function call allocates its own
// temporary data structures (matrices, arrays).
//
// 이 파일의 모든 함수는 불변 문자열에서 작동하고 공유 가변 상태를 사용하지
// 않으므로 스레드 안전합니다. 각 함수 호출은 자체 임시 데이터 구조
// (행렬, 배열)를 할당합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         distance := stringutil.LevenshteinDistance(s1, s2)
//     }()
//
//     go func() {
//         similarity := stringutil.JaroWinklerSimilarity(s1, s2)
//     }()
//
//     // All distance functions safe for concurrent use
//     // 모든 거리 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - comparison.go: Basic string comparison (EqualFold, HasPrefix, HasSuffix)
//   comparison.go: 기본 문자열 비교 (EqualFold, HasPrefix, HasSuffix)
// - search.go: String search operations (ContainsAny, ReplaceAll, etc.)
//   search.go: 문자열 검색 연산 (ContainsAny, ReplaceAll 등)
// - validation.go: String validation (IsEmail, IsURL, etc.)
//   validation.go: 문자열 검증 (IsEmail, IsURL 등)
//
// =============================================================================

// LevenshteinDistance calculates the Levenshtein distance between two strings.
// LevenshteinDistance는 두 문자열 사이의 Levenshtein 거리를 계산합니다.
//
// The Levenshtein distance is the minimum number of single-character edits
// (insertions, deletions, or substitutions) required to change one string into another.
//
// Levenshtein 거리는 한 문자열을 다른 문자열로 변경하는 데 필요한
// 최소 단일 문자 편집(삽입, 삭제 또는 치환) 횟수입니다.
//
// Example
// 예제:
//
//	LevenshteinDistance("kitten", "sitting")  // 3
//	LevenshteinDistance("hello", "hello")     // 0
//	LevenshteinDistance("", "hello")          // 5
func LevenshteinDistance(a, b string) int {
	runesA := []rune(a)
	runesB := []rune(b)

	lenA := len(runesA)
	lenB := len(runesB)

	// Early exit for empty strings
	// 빈 문자열에 대한 조기 종료
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}

	// Create matrix
	// 행렬 생성
	matrix := make([][]int, lenA+1)
	for i := range matrix {
		matrix[i] = make([]int, lenB+1)
	}

	// Initialize first row and column
	// 첫 행과 열 초기화
	for i := 0; i <= lenA; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		matrix[0][j] = j
	}

	// Fill matrix
	// 행렬 채우기
	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			cost := 0
			if runesA[i-1] != runesB[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				// deletion
				// 삭제
				matrix[i-1][j]+1,
				// insertion
				// 삽입
				matrix[i][j-1]+1,
				// substitution
				// 치환
				matrix[i-1][j-1]+cost,
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
// Formula: 1 - (distance
// max(len(a), len(b)))
// 공식: 1 - (거리
// max(len(a), len(b)))
//
// Example
// 예제:
//
//	Similarity("hello", "hello")   // 1.0
//	Similarity("hello", "hallo")   // 0.8
//	Similarity("hello", "world")   // 0.2
func Similarity(a, b string) float64 {
	// Early exit for identical strings
	// 동일한 문자열에 대한 조기 종료
	if a == b {
		return 1.0
	}

	runesA := []rune(a)
	runesB := []rune(b)

	maxLen := len(runesA)
	if len(runesB) > maxLen {
		maxLen = len(runesB)
	}

	// Early exit for empty strings
	// 빈 문자열에 대한 조기 종료
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
// Example
// 예제:
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
// Example
// 예제:
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

	// Jaro-Winkler formula
	// Jaro-Winkler 공식
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

	// Calculate match window
	// 매치 윈도우 계산
	matchWindow := int(math.Max(float64(lenA), float64(lenB))/2) - 1
	if matchWindow < 1 {
		matchWindow = 1
	}

	matchesA := make([]bool, lenA)
	matchesB := make([]bool, lenB)

	matches := 0
	transpositions := 0

	// Find matches
	// 매치 찾기
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

	// Count transpositions
	// 전치 횟수 계산
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

	// Jaro formula
	// Jaro 공식
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
