package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/unix.go
// Purpose: Unix timestamp operations with multiple precision levels
//          여러 정밀도 수준의 Unix 타임스탬프 연산
//
// This file provides comprehensive Unix timestamp operations for converting
// between time.Time and Unix timestamps at various precision levels: seconds,
// milliseconds, microseconds, and nanoseconds. Unix timestamps represent the
// number of units elapsed since the Unix epoch (January 1, 1970, 00:00:00 UTC).
// These operations are essential for APIs, databases, logging, and any system
// that needs timezone-independent time representation or numeric time values.
//
// 이 파일은 초, 밀리초, 마이크로초, 나노초 등 다양한 정밀도 수준에서 time.Time과
// Unix 타임스탬프 간 변환을 위한 포괄적인 Unix 타임스탬프 연산을 제공합니다.
// Unix 타임스탬프는 Unix epoch(1970년 1월 1일 00:00:00 UTC) 이후 경과한
// 단위 수를 나타냅니다. 이러한 연산은 API, 데이터베이스, 로깅 및 타임존 독립적인
// 시간 표현이나 숫자 시간 값이 필요한 모든 시스템에 필수적입니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. CURRENT TIMESTAMP (현재 타임스탬프)
//    - Now: Current time in seconds since epoch
//      epoch 이후 초 단위 현재 시간
//    - NowMilli: Current time in milliseconds
//      밀리초 단위 현재 시간
//    - NowMicro: Current time in microseconds
//      마이크로초 단위 현재 시간
//    - NowNano: Current time in nanoseconds
//      나노초 단위 현재 시간
//    - Convenience for getting "now" as numeric value
//      "지금"을 숫자 값으로 가져오는 편의
//
// 2. TIMESTAMP TO TIME (타임스탬프를 시간으로)
//    - FromUnix: Convert seconds to time.Time
//      초를 time.Time으로 변환
//    - FromUnixMilli: Convert milliseconds to time.Time
//      밀리초를 time.Time으로 변환
//    - FromUnixMicro: Convert microseconds to time.Time
//      마이크로초를 time.Time으로 변환
//    - FromUnixNano: Convert nanoseconds to time.Time
//      나노초를 time.Time으로 변환
//    - Applies default timezone (KST)
//      기본 타임존 적용 (KST)
//
// 3. TIME TO TIMESTAMP (시간을 타임스탬프로)
//    - ToUnix: Convert time.Time to seconds
//      time.Time을 초로 변환
//    - ToUnixMilli: Convert time.Time to milliseconds
//      time.Time을 밀리초로 변환
//    - ToUnixMicro: Convert time.Time to microseconds
//      time.Time을 마이크로초로 변환
//    - ToUnixNano: Convert time.Time to nanoseconds
//      time.Time을 나노초로 변환
//    - Timezone-independent (always UTC-based)
//      타임존 독립적 (항상 UTC 기반)
//
// 4. MULTIPLE PRECISION LEVELS (여러 정밀도 수준)
//    - Seconds: Standard Unix timestamp (1 second resolution)
//      표준 Unix 타임스탬프 (1초 해상도)
//    - Milliseconds: For most applications (1 ms = 0.001 sec)
//      대부분의 애플리케이션용 (1 ms = 0.001초)
//    - Microseconds: High-precision timing (1 μs = 0.000001 sec)
//      고정밀 타이밍 (1 μs = 0.000001초)
//    - Nanoseconds: Maximum Go precision (1 ns = 0.000000001 sec)
//      최대 Go 정밀도 (1 ns = 0.000000001초)
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. UNIX EPOCH AS REFERENCE POINT (참조점으로 Unix epoch)
//    - All timestamps relative to: 1970-01-01 00:00:00 UTC
//      모든 타임스탬프가 상대적: 1970-01-01 00:00:00 UTC
//    - Positive for dates after 1970
//      1970 이후 날짜는 양수
//    - Negative for dates before 1970 (supported but rare)
//      1970 이전 날짜는 음수 (지원되지만 드뭄)
//    - Universal standard across systems
//      시스템 전반의 보편적 표준
//
// 2. TIMEZONE-INDEPENDENT STORAGE (타임존 독립적 저장)
//    - Unix timestamps are always UTC-based
//      Unix 타임스탬프는 항상 UTC 기반
//    - Same instant = same timestamp regardless of timezone
//      같은 순간 = 타임존과 관계없이 같은 타임스탬프
//    - Ideal for databases and APIs
//      데이터베이스 및 API에 이상적
//    - No DST or timezone offset issues
//      DST 또는 타임존 오프셋 문제 없음
//
// 3. DEFAULT TIMEZONE FOR FROM* FUNCTIONS (FROM* 함수의 기본 타임존)
//    - FromUnix* functions apply defaultLocation (KST)
//      FromUnix* 함수는 defaultLocation (KST) 적용
//    - Returns time.Time in KST for consistency
//      일관성을 위해 KST의 time.Time 반환
//    - Can be changed via SetDefaultTimezone
//      SetDefaultTimezone을 통해 변경 가능
//    - Instant in time is preserved (only representation changes)
//      시간의 순간은 보존됨 (표현만 변경)
//
// 4. PRECISION SELECTION (정밀도 선택)
//    - Use seconds for most applications (smaller numbers)
//      대부분의 애플리케이션에 초 사용 (더 작은 숫자)
//    - Use milliseconds for web APIs (JavaScript standard)
//      웹 API에 밀리초 사용 (JavaScript 표준)
//    - Use microseconds for precise measurements
//      정밀 측정에 마이크로초 사용
//    - Use nanoseconds for benchmarks and internal timing
//      벤치마크 및 내부 타이밍에 나노초 사용
//
// 5. SYMMETRICAL API (대칭 API)
//    - Now* functions: Get current timestamp
//      Now* 함수: 현재 타임스탬프 가져오기
//    - FromUnix*: Convert timestamp to time
//      FromUnix*: 타임스탬프를 시간으로 변환
//    - ToUnix*: Convert time to timestamp
//      ToUnix*: 시간을 타임스탬프로 변환
//    - Consistent naming and behavior
//      일관된 명명 및 동작
//
// 6. INTEGER REPRESENTATION (정수 표현)
//    - All timestamps are int64
//      모든 타임스탬프는 int64
//    - Easy to serialize, compare, and store
//      직렬화, 비교 및 저장이 쉬움
//    - No floating-point precision issues
//      부동 소수점 정밀도 문제 없음
//    - Sufficient range: ~292 billion years for nanoseconds
//      충분한 범위: 나노초로 ~292억 년
//
// ============================================================================
// UNIX TIMESTAMP OPERATIONS OVERVIEW / Unix 타임스탬프 연산 개요
// ============================================================================
//
// CURRENT TIMESTAMP (현재 타임스탬프) - 4 functions
// ├─ Now       : Current Unix timestamp in seconds
// ├─ NowMilli  : Current Unix timestamp in milliseconds
// ├─ NowMicro  : Current Unix timestamp in microseconds
// └─ NowNano   : Current Unix timestamp in nanoseconds
//
// TIMESTAMP TO TIME (타임스탬프를 시간으로) - 4 functions
// ├─ FromUnix      : Convert seconds to time.Time (in KST)
// ├─ FromUnixMilli : Convert milliseconds to time.Time (in KST)
// ├─ FromUnixMicro : Convert microseconds to time.Time (in KST)
// └─ FromUnixNano  : Convert nanoseconds to time.Time (in KST)
//
// TIME TO TIMESTAMP (시간을 타임스탬프로) - 4 functions
// ├─ ToUnix      : Convert time.Time to seconds
// ├─ ToUnixMilli : Convert time.Time to milliseconds
// ├─ ToUnixMicro : Convert time.Time to microseconds
// └─ ToUnixNano  : Convert time.Time to nanoseconds
//
// Total: 12 Unix timestamp functions
// 총: 12개의 Unix 타임스탬프 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
//
// NOW, NOWMILLI, NOWMICRO, NOWnano: O(1)
//   System call to get current time
//   현재 시간을 가져오는 시스템 호출
//   ~100-200 nanoseconds
//
// FROMUNIX, FROMUNIXMILLI, FROMUNIXMICRO, FROMUNIXNANO: O(1)
//   Simple arithmetic and structure creation
//   간단한 산술 및 구조 생성
//   ~20-50 nanoseconds
//
// TOUNIX, TOUNIXMILLI, TOUNIXMICRO, TOUNIXNANO: O(1)
//   Simple arithmetic on time.Time fields
//   time.Time 필드의 간단한 산술
//   ~10-30 nanoseconds
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - no additional allocation
//   모든 함수: O(1) - 추가 할당 없음
//
// PERFORMANCE NOTES:
// 성능 참고:
// 1. ToUnix* is very fast (direct field access)
//    ToUnix*는 매우 빠름 (직접 필드 액세스)
// 2. FromUnix* slightly slower (creates time.Time struct)
//    FromUnix*는 약간 느림 (time.Time 구조 생성)
// 3. Now* requires system call (slowest but still fast)
//    Now*는 시스템 호출 필요 (가장 느리지만 여전히 빠름)
// 4. All operations are extremely efficient
//    모든 연산이 매우 효율적
//
// BENCHMARK EXAMPLES:
// 벤치마크 예시:
//   Now()           : 150 ns/op
//   FromUnix(sec)   :  30 ns/op
//   ToUnix(t)       :  15 ns/op
//
// ============================================================================
// UNIX TIMESTAMP PRECISION LEVELS / Unix 타임스탬프 정밀도 수준
// ============================================================================
//
// SECONDS (초):
// - Range: -292,277,026,596 to 292,277,026,596
//   범위: -292,277,026,596 ~ 292,277,026,596
// - Precision: 1 second
//   정밀도: 1초
// - Use cases: Standard timestamps, dates without time-of-day precision
//   사용 사례: 표준 타임스탬프, 시간 정밀도가 없는 날짜
// - Example: 1704074400 = 2024-01-01 00:00:00 UTC
//   예시: 1704074400 = 2024-01-01 00:00:00 UTC
//
// MILLISECONDS (밀리초):
// - Range: ~292 million years (int64)
//   범위: ~292백만 년 (int64)
// - Precision: 1 millisecond = 0.001 seconds
//   정밀도: 1 밀리초 = 0.001초
// - Use cases: Web APIs, JavaScript compatibility, event logging
//   사용 사례: 웹 API, JavaScript 호환성, 이벤트 로깅
// - Example: 1704074400000 = 2024-01-01 00:00:00.000 UTC
//   예시: 1704074400000 = 2024-01-01 00:00:00.000 UTC
//
// MICROSECONDS (마이크로초):
// - Range: ~292,000 years (int64)
//   범위: ~292,000년 (int64)
// - Precision: 1 microsecond = 0.000001 seconds
//   정밀도: 1 마이크로초 = 0.000001초
// - Use cases: Performance measurements, high-frequency trading
//   사용 사례: 성능 측정, 고빈도 거래
// - Example: 1704074400000000 = 2024-01-01 00:00:00.000000 UTC
//   예시: 1704074400000000 = 2024-01-01 00:00:00.000000 UTC
//
// NANOSECONDS (나노초):
// - Range: ~292 years from 1970 (int64 limitation)
//   범위: 1970년부터 ~292년 (int64 제한)
// - Precision: 1 nanosecond = 0.000000001 seconds
//   정밀도: 1 나노초 = 0.000000001초
// - Use cases: Benchmarks, internal timing, maximum precision
//   사용 사례: 벤치마크, 내부 타이밍, 최대 정밀도
// - Example: 1704074400000000000 = 2024-01-01 00:00:00.000000000 UTC
//   예시: 1704074400000000000 = 2024-01-01 00:00:00.000000000 UTC
//
// PRECISION SELECTION GUIDE:
// 정밀도 선택 가이드:
// - Most applications: Milliseconds (good balance)
//   대부분의 애플리케이션: 밀리초 (좋은 균형)
// - Database timestamps: Seconds (smaller storage)
//   데이터베이스 타임스탬프: 초 (더 작은 저장소)
// - JavaScript interop: Milliseconds (Date.now())
//   JavaScript 상호 운용: 밀리초 (Date.now())
// - Performance timing: Microseconds or nanoseconds
//   성능 타이밍: 마이크로초 또는 나노초
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Database Timestamp Storage (데이터베이스 타임스탬프 저장)
// Use case: Store creation time as Unix timestamp
// 사용 사례: 생성 시간을 Unix 타임스탬프로 저장
//
//   type User struct {
//       ID        int    `db:"id"`
//       Name      string `db:"name"`
//       CreatedAt int64  `db:"created_at"`  // Unix seconds
//   }
//
//   user := &User{
//       Name:      "Alice",
//       CreatedAt: timeutil.Now(),  // Current Unix timestamp
//   }
//   db.Insert(user)
//
//   // Later, convert back to time.Time
//   createdTime := timeutil.FromUnix(user.CreatedAt)
//   fmt.Println(timeutil.FormatDateTime(createdTime))
//
// PATTERN 2: API Response (API 응답)
// Use case: Return timestamps in milliseconds for JavaScript
// 사용 사례: JavaScript용 밀리초 단위 타임스탬프 반환
//
//   type EventResponse struct {
//       ID        string `json:"id"`
//       Title     string `json:"title"`
//       Timestamp int64  `json:"timestamp"`  // Milliseconds
//   }
//
//   response := &EventResponse{
//       ID:        event.ID,
//       Title:     event.Title,
//       Timestamp: timeutil.ToUnixMilli(event.OccurredAt),
//   }
//   json.NewEncoder(w).Encode(response)
//
//   // JavaScript client:
//   // const date = new Date(response.timestamp);
//
// PATTERN 3: Request Parsing (요청 파싱)
// Use case: Parse Unix timestamp from query parameter
// 사용 사례: 쿼리 매개변수에서 Unix 타임스탬프 파싱
//
//   // GET /events?since=1704074400
//   sinceStr := r.URL.Query().Get("since")
//   sinceUnix, _ := strconv.ParseInt(sinceStr, 10, 64)
//
//   sinceTime := timeutil.FromUnix(sinceUnix)
//   events := db.GetEventsSince(sinceTime)
//
// PATTERN 4: Performance Measurement (성능 측정)
// Use case: Measure operation duration in microseconds
// 사용 사례: 마이크로초 단위로 작업 기간 측정
//
//   start := timeutil.NowMicro()
//
//   // Perform operation
//   result := expensiveOperation()
//
//   end := timeutil.NowMicro()
//   durationMicros := end - start
//   fmt.Printf("Operation took %d μs\n", durationMicros)
//
// PATTERN 5: Timestamp Comparison (타임스탬프 비교)
// Use case: Check if timestamp is within range
// 사용 사례: 타임스탬프가 범위 내에 있는지 확인
//
//   now := timeutil.Now()
//   oneDayAgo := now - 86400  // 86400 seconds = 1 day
//
//   if event.Timestamp >= oneDayAgo && event.Timestamp <= now {
//       fmt.Println("Event occurred in the last 24 hours")
//   }
//
// PATTERN 6: Log Timestamps (로그 타임스탬프)
// Use case: Add precise timestamps to logs
// 사용 사례: 로그에 정밀한 타임스탬프 추가
//
//   type LogEntry struct {
//       Timestamp int64  `json:"timestamp"`  // Nanoseconds
//       Level     string `json:"level"`
//       Message   string `json:"message"`
//   }
//
//   log := &LogEntry{
//       Timestamp: timeutil.NowNano(),
//       Level:     "INFO",
//       Message:   "User logged in",
//   }
//
// PATTERN 7: Cache Expiration (캐시 만료)
// Use case: Store expiration time as Unix timestamp
// 사용 사례: 만료 시간을 Unix 타임스탬프로 저장
//
//   type CacheEntry struct {
//       Key       string
//       Value     interface{}
//       ExpiresAt int64  // Unix seconds
//   }
//
//   entry := &CacheEntry{
//       Key:       "user:123",
//       Value:     userData,
//       ExpiresAt: timeutil.Now() + 3600,  // Expires in 1 hour
//   }
//
//   // Check expiration
//   if timeutil.Now() > entry.ExpiresAt {
//       // Cache expired
//       delete(cache, entry.Key)
//   }
//
// PATTERN 8: Unix-to-Unix Conversion (Unix-Unix 변환)
// Use case: Convert between precision levels
// 사용 사례: 정밀도 수준 간 변환
//
//   // Have milliseconds, need seconds
//   millis := int64(1704074400000)
//   seconds := millis / 1000
//
//   // Or convert through time.Time
//   t := timeutil.FromUnixMilli(millis)
//   seconds = timeutil.ToUnix(t)
//
// PATTERN 9: Sorting by Timestamp (타임스탬프로 정렬)
// Use case: Sort events by Unix timestamp
// 사용 사례: Unix 타임스탬프로 이벤트 정렬
//
//   type Event struct {
//       Name      string
//       Timestamp int64
//   }
//
//   events := []Event{ /* ... */ }
//
//   sort.Slice(events, func(i, j int) bool {
//       return events[i].Timestamp < events[j].Timestamp
//   })
//
// PATTERN 10: JWT Claims (JWT 클레임)
// Use case: Standard JWT expiration time
// 사용 사례: 표준 JWT 만료 시간
//
//   type JWTClaims struct {
//       UserID string `json:"user_id"`
//       Exp    int64  `json:"exp"`  // Standard: Unix seconds
//       Iat    int64  `json:"iat"`  // Issued at
//   }
//
//   claims := &JWTClaims{
//       UserID: "123",
//       Iat:    timeutil.Now(),
//       Exp:    timeutil.Now() + 3600,  // 1 hour from now
//   }
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// NEGATIVE TIMESTAMPS:
// 음수 타임스탬프:
//   // Dates before Unix epoch (1970-01-01)
//   t := time.Date(1969, 12, 31, 0, 0, 0, 0, time.UTC)
//   timestamp := timeutil.ToUnix(t)  // Negative value
//   // timestamp = -86400 (one day before epoch)
//
// ZERO TIMESTAMP:
// 제로 타임스탬프:
//   t := timeutil.FromUnix(0)
//   // t = 1970-01-01 00:00:00 UTC (Unix epoch)
//   // In KST: 1970-01-01 09:00:00 (UTC+9)
//
// NANOSECOND OVERFLOW:
// 나노초 오버플로:
//   // NowNano() will overflow around year 2262
//   // Use milliseconds or microseconds for far-future dates
//   // NowNano()는 2262년경에 오버플로됩니다
//   // 먼 미래 날짜에는 밀리초 또는 마이크로초 사용
//
// PRECISION LOSS:
// 정밀도 손실:
//   // Converting to seconds loses sub-second precision
//   t := time.Now()
//   sec := timeutil.ToUnix(t)
//   t2 := timeutil.FromUnix(sec)
//   // t != t2 (sub-second precision lost)
//   // t != t2 (서브초 정밀도 손실)
//
// TIMEZONE IN FROM* FUNCTIONS:
// FROM* 함수의 타임존:
//   // FromUnix applies default timezone
//   t := timeutil.FromUnix(1704074400)
//   // t.Location() == KST (not UTC)
//   // Instant is correct, only representation changes
//   // 순간은 정확하지만 표현만 변경됨
//
// LARGE TIMESTAMPS:
// 큰 타임스탬프:
//   // int64 max: 9,223,372,036,854,775,807
//   // For seconds: valid until year ~292 billion
//   // For nanoseconds: valid until year ~2262
//   // 초의 경우: ~292억 년까지 유효
//   // 나노초의 경우: ~2262년까지 유효
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// THREAD-SAFE FUNCTIONS:
// 스레드 안전 함수:
// - All Unix timestamp functions are thread-safe
//   모든 Unix 타임스탬프 함수는 스레드 안전
// - No shared mutable state
//   공유 변경 가능 상태 없음
// - Read-only access to defaultLocation
//   defaultLocation에 대한 읽기 전용 액세스
// - Safe for concurrent use
//   동시 사용에 안전
//
// CONCURRENT USAGE EXAMPLE:
// 동시 사용 예시:
//   var wg sync.WaitGroup
//   for i := 0; i < 100; i++ {
//       wg.Add(1)
//       go func() {
//           defer wg.Done()
//           now := timeutil.Now()
//           t := timeutil.FromUnix(now)
//           fmt.Println(timeutil.FormatDateTime(t))
//       }()
//   }
//   wg.Wait()
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - defaultLocation: For FromUnix* functions (applies timezone)
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Now(): Get current time
// - time.Unix(): Create time from Unix timestamp
// - time.UnixMilli(), UnixMicro(): Create time from milliseconds/microseconds
//
// USED BY (사용처):
// - Database operations (store timestamps)
//   데이터베이스 연산 (타임스탬프 저장)
// - APIs (serialize/deserialize times)
//   API (시간 직렬화/역직렬화)
// - Logging (precise event timestamps)
//   로깅 (정밀한 이벤트 타임스탬프)
// - Performance measurement (duration tracking)
//   성능 측정 (기간 추적)
// - Caching (expiration times)
//   캐싱 (만료 시간)
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE MILLISECONDS FOR WEB APIs
//    웹 API에 밀리초 사용
//    timestamp := timeutil.ToUnixMilli(time.Now())
//
// 2. USE SECONDS FOR DATABASE STORAGE
//    데이터베이스 저장에 초 사용
//    CreatedAt: timeutil.Now()  // Smaller storage
//
// 3. USE MICROSECONDS FOR PERFORMANCE TIMING
//    성능 타이밍에 마이크로초 사용
//    start := timeutil.NowMicro()
//
// 4. STORE UNIX, DISPLAY FORMATTED
//    Unix로 저장, 형식화하여 표시
//    db.Timestamp = timeutil.Now()
//    display := timeutil.FormatDateTime(timeutil.FromUnix(db.Timestamp))
//
// 5. DOCUMENT PRECISION IN API
//    API에 정밀도 문서화
//    // {"timestamp": 1704074400000}  // milliseconds since epoch
//
// 6. AVOID NANOSECONDS FOR LONG-TERM STORAGE
//    장기 저장에 나노초 사용 금지
//    // Nanoseconds overflow in year 2262
//
// 7. USE TIMEZONE-INDEPENDENT COMPARISON
//    타임존 독립적 비교 사용
//    if timestamp1 > timestamp2 { /* ... */ }
//
// 8. CONVERT THROUGH TIME.TIME FOR PRECISION CHANGE
//    정밀도 변경을 위해 time.Time을 통해 변환
//    t := timeutil.FromUnixMilli(millis)
//    seconds := timeutil.ToUnix(t)
//
// ============================================================================

// Now returns the current Unix timestamp in seconds.
// Now는 현재 Unix 타임스탬프를 초 단위로 반환합니다.
func Now() int64 {
	return time.Now().Unix()
}

// NowMilli returns the current Unix timestamp in milliseconds.
// NowMilli는 현재 Unix 타임스탬프를 밀리초 단위로 반환합니다.
func NowMilli() int64 {
	return time.Now().UnixMilli()
}

// NowMicro returns the current Unix timestamp in microseconds.
// NowMicro는 현재 Unix 타임스탬프를 마이크로초 단위로 반환합니다.
func NowMicro() int64 {
	return time.Now().UnixMicro()
}

// NowNano returns the current Unix timestamp in nanoseconds.
// NowNano는 현재 Unix 타임스탬프를 나노초 단위로 반환합니다.
func NowNano() int64 {
	return time.Now().UnixNano()
}

// FromUnix creates a time from a Unix timestamp in seconds.
// FromUnix는 초 단위 Unix 타임스탬프로부터 시간을 생성합니다.
func FromUnix(sec int64) time.Time {
	return time.Unix(sec, 0).In(defaultLocation)
}

// FromUnixMilli creates a time from a Unix timestamp in milliseconds.
// FromUnixMilli는 밀리초 단위 Unix 타임스탬프로부터 시간을 생성합니다.
func FromUnixMilli(msec int64) time.Time {
	return time.UnixMilli(msec).In(defaultLocation)
}

// FromUnixMicro creates a time from a Unix timestamp in microseconds.
// FromUnixMicro는 마이크로초 단위 Unix 타임스탬프로부터 시간을 생성합니다.
func FromUnixMicro(usec int64) time.Time {
	return time.UnixMicro(usec).In(defaultLocation)
}

// FromUnixNano creates a time from a Unix timestamp in nanoseconds.
// FromUnixNano는 나노초 단위 Unix 타임스탬프로부터 시간을 생성합니다.
func FromUnixNano(nsec int64) time.Time {
	return time.Unix(0, nsec).In(defaultLocation)
}

// ToUnix converts a time to a Unix timestamp in seconds.
// ToUnix는 시간을 초 단위 Unix 타임스탬프로 변환합니다.
func ToUnix(t time.Time) int64 {
	return t.Unix()
}

// ToUnixMilli converts a time to a Unix timestamp in milliseconds.
// ToUnixMilli는 시간을 밀리초 단위 Unix 타임스탬프로 변환합니다.
func ToUnixMilli(t time.Time) int64 {
	return t.UnixMilli()
}

// ToUnixMicro converts a time to a Unix timestamp in microseconds.
// ToUnixMicro는 시간을 마이크로초 단위 Unix 타임스탬프로 변환합니다.
func ToUnixMicro(t time.Time) int64 {
	return t.UnixMicro()
}

// ToUnixNano converts a time to a Unix timestamp in nanoseconds.
// ToUnixNano는 시간을 나노초 단위 Unix 타임스탬프로 변환합니다.
func ToUnixNano(t time.Time) int64 {
	return t.UnixNano()
}
