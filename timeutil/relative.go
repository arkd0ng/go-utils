package timeutil

import (
	"fmt"
	"math"
	"time"
)

// RelativeTime returns a human-readable relative time string.
// RelativeTime은 사람이 읽기 쉬운 상대 시간 문자열을 반환합니다.
//
// Examples:
//   - "2 hours ago"
//   - "in 3 days"
//   - "just now"
func RelativeTime(t time.Time) string {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	diff := now.Sub(t)

	if diff < 0 {
		// Future / 미래
		return relativeTimeFuture(-diff)
	}

	// Past / 과거
	return relativeTimePast(diff)
}

// RelativeTimeShort returns a short human-readable relative time string.
// RelativeTimeShort는 짧은 사람이 읽기 쉬운 상대 시간 문자열을 반환합니다.
//
// Examples:
//   - "2h ago"
//   - "in 3d"
//   - "now"
func RelativeTimeShort(t time.Time) string {
	now := time.Now().In(defaultLocation)
	t = t.In(defaultLocation)
	diff := now.Sub(t)

	if diff < 0 {
		// Future / 미래
		return relativeTimeFutureShort(-diff)
	}

	// Past / 과거
	return relativeTimePastShort(diff)
}

// TimeAgo is an alias for RelativeTime.
// TimeAgo는 RelativeTime의 별칭입니다.
func TimeAgo(t time.Time) string {
	return RelativeTime(t)
}

// relativeTimePast returns a relative time string for past times.
// relativeTimePast는 과거 시간에 대한 상대 시간 문자열을 반환합니다.
func relativeTimePast(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "just now"
	case seconds < 60:
		return fmt.Sprintf("%d seconds ago", seconds)
	case minutes == 1:
		return "1 minute ago"
	case minutes < 60:
		return fmt.Sprintf("%d minutes ago", minutes)
	case hours == 1:
		return "1 hour ago"
	case hours < 24:
		return fmt.Sprintf("%d hours ago", hours)
	case days == 1:
		return "1 day ago"
	case days < 7:
		return fmt.Sprintf("%d days ago", days)
	case weeks == 1:
		return "1 week ago"
	case weeks < 4:
		return fmt.Sprintf("%d weeks ago", weeks)
	case months == 1:
		return "1 month ago"
	case months < 12:
		return fmt.Sprintf("%d months ago", months)
	case years == 1:
		return "1 year ago"
	default:
		return fmt.Sprintf("%d years ago", years)
	}
}

// relativeTimeFuture returns a relative time string for future times.
// relativeTimeFuture는 미래 시간에 대한 상대 시간 문자열을 반환합니다.
func relativeTimeFuture(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "just now"
	case seconds < 60:
		return fmt.Sprintf("in %d seconds", seconds)
	case minutes == 1:
		return "in 1 minute"
	case minutes < 60:
		return fmt.Sprintf("in %d minutes", minutes)
	case hours == 1:
		return "in 1 hour"
	case hours < 24:
		return fmt.Sprintf("in %d hours", hours)
	case days == 1:
		return "in 1 day"
	case days < 7:
		return fmt.Sprintf("in %d days", days)
	case weeks == 1:
		return "in 1 week"
	case weeks < 4:
		return fmt.Sprintf("in %d weeks", weeks)
	case months == 1:
		return "in 1 month"
	case months < 12:
		return fmt.Sprintf("in %d months", months)
	case years == 1:
		return "in 1 year"
	default:
		return fmt.Sprintf("in %d years", years)
	}
}

// relativeTimePastShort returns a short relative time string for past times.
// relativeTimePastShort는 과거 시간에 대한 짧은 상대 시간 문자열을 반환합니다.
func relativeTimePastShort(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "now"
	case seconds < 60:
		return fmt.Sprintf("%ds ago", seconds)
	case minutes < 60:
		return fmt.Sprintf("%dm ago", minutes)
	case hours < 24:
		return fmt.Sprintf("%dh ago", hours)
	case days < 7:
		return fmt.Sprintf("%dd ago", days)
	case weeks < 4:
		return fmt.Sprintf("%dw ago", weeks)
	case months < 12:
		return fmt.Sprintf("%dmo ago", months)
	default:
		return fmt.Sprintf("%dy ago", years)
	}
}

// relativeTimeFutureShort returns a short relative time string for future times.
// relativeTimeFutureShort는 미래 시간에 대한 짧은 상대 시간 문자열을 반환합니다.
func relativeTimeFutureShort(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := hours / 24
	weeks := days / 7
	months := days / 30
	years := days / 365

	switch {
	case seconds < 10:
		return "now"
	case seconds < 60:
		return fmt.Sprintf("in %ds", seconds)
	case minutes < 60:
		return fmt.Sprintf("in %dm", minutes)
	case hours < 24:
		return fmt.Sprintf("in %dh", hours)
	case days < 7:
		return fmt.Sprintf("in %dd", days)
	case weeks < 4:
		return fmt.Sprintf("in %dw", weeks)
	case months < 12:
		return fmt.Sprintf("in %dmo", months)
	default:
		return fmt.Sprintf("in %dy", years)
	}
}

// HumanizeDuration converts a duration to a human-readable string.
// HumanizeDuration은 duration을 사람이 읽기 쉬운 문자열로 변환합니다.
//
// Example / 예제:
//
//	d := 2*time.Hour + 30*time.Minute
//	str := timeutil.HumanizeDuration(d) // "2 hours 30 minutes"
func HumanizeDuration(d time.Duration) string {
	if d < 0 {
		d = -d
	}

	seconds := int(math.Abs(d.Seconds()))
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24

	seconds = seconds % 60
	minutes = minutes % 60
	hours = hours % 24

	var result string
	if days > 0 {
		result = fmt.Sprintf("%d days", days)
		if hours > 0 {
			result += fmt.Sprintf(" %d hours", hours)
		}
	} else if hours > 0 {
		result = fmt.Sprintf("%d hours", hours)
		if minutes > 0 {
			result += fmt.Sprintf(" %d minutes", minutes)
		}
	} else if minutes > 0 {
		result = fmt.Sprintf("%d minutes", minutes)
		if seconds > 0 {
			result += fmt.Sprintf(" %d seconds", seconds)
		}
	} else {
		result = fmt.Sprintf("%d seconds", seconds)
	}

	return result
}
