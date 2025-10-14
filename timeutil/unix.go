package timeutil

import "time"

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
