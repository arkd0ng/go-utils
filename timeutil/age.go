package timeutil

import "time"

// AgeInYears calculates age in years from birth date.
// AgeInYears는 생년월일로부터 나이를 년 단위로 계산합니다.
//
// Example / 예제:
//
//	birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
//	age := timeutil.AgeInYears(birthDate) // 35
func AgeInYears(birthDate time.Time) int {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	years := now.Year() - birthDate.Year()

	// Adjust if birthday hasn't occurred this year yet
	// 올해 생일이 아직 지나지 않았으면 조정
	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		years--
	}

	return years
}

// AgeInMonths calculates age in months from birth date.
// AgeInMonths는 생년월일로부터 나이를 월 단위로 계산합니다.
//
// Example / 예제:
//
//	months := timeutil.AgeInMonths(birthDate)
func AgeInMonths(birthDate time.Time) int {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	years := now.Year() - birthDate.Year()
	months := int(now.Month()) - int(birthDate.Month())

	// Adjust if day hasn't occurred this month yet
	// 이번 달 일이 아직 지나지 않았으면 조정
	if now.Day() < birthDate.Day() {
		months--
	}

	return years*MonthsPerYear + months
}

// AgeInDays calculates age in days from birth date.
// AgeInDays는 생년월일로부터 나이를 일 단위로 계산합니다.
//
// Example / 예제:
//
//	days := timeutil.AgeInDays(birthDate)
func AgeInDays(birthDate time.Time) int {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	// Truncate to start of day for accurate day count
	// 정확한 일 수 계산을 위해 하루의 시작으로 절단
	now = StartOfDay(now)
	birthDate = StartOfDay(birthDate)

	return int(now.Sub(birthDate).Hours() / 24)
}

// Age calculates detailed age (years, months, days) from birth date.
// Age는 생년월일로부터 상세 나이 (년, 월, 일)를 계산합니다.
//
// Example / 예제:
//
//	age := timeutil.Age(birthDate)
//	fmt.Printf("%d years %d months %d days\n", age.Years, age.Months, age.Days)
func Age(birthDate time.Time) *AgeDetail {
	now := time.Now().In(defaultLocation)
	birthDate = birthDate.In(defaultLocation)

	years := now.Year() - birthDate.Year()
	months := int(now.Month()) - int(birthDate.Month())
	days := now.Day() - birthDate.Day()

	// Adjust days / 일 조정
	if days < 0 {
		months--
		// Get days in previous month / 이전 달의 일 수 가져오기
		prevMonth := now.AddDate(0, -1, 0)
		daysInPrevMonth := time.Date(prevMonth.Year(), prevMonth.Month()+1, 0, 0, 0, 0, 0, defaultLocation).Day()
		days += daysInPrevMonth
	}

	// Adjust months / 월 조정
	if months < 0 {
		years--
		months += MonthsPerYear
	}

	return &AgeDetail{
		Years:  years,
		Months: months,
		Days:   days,
	}
}
