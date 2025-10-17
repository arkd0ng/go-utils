package timeutil

import "time"

// MonthKorean returns the Korean name of the month.
// MonthKorean은 월의 한글 이름을 반환합니다.
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
// month := timeutil.MonthKorean(t) / fmt.Println(month) // "10월"
func MonthKorean(t time.Time) string {
	return t.Format("1월")
}

// MonthName returns the English name of the month.
// MonthName은 월의 영문 이름을 반환합니다.
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	month := timeutil.MonthName(t)
//	fmt.Println(month) // "October"
func MonthName(t time.Time) string {
	return t.Month().String()
}

// MonthNameShort returns the short English name of the month (3 letters).
// MonthNameShort는 월의 짧은 영문 이름을 반환합니다 (3글자).
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	month := timeutil.MonthNameShort(t)
//	fmt.Println(month) // "Oct"
func MonthNameShort(t time.Time) string {
	return t.Format("Jan")
}

// Quarter returns the quarter of the year (1-4).
// Quarter는 년의 분기를 반환합니다 (1-4).
//
// Q1: January-March, Q2: April-June, Q3: July-September, Q4: October-December
// Q1: 1-3월, Q2: 4-6월, Q3: 7-9월, Q4: 10-12월
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	quarter := timeutil.Quarter(t)
//	fmt.Println(quarter) // 4
func Quarter(t time.Time) int {
	month := int(t.Month())
	return (month-1)/MonthsPerQuarter + 1
}
