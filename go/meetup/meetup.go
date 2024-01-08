package meetup

import "time"

type WeekSchedule string

var (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Fifth  WeekSchedule = "fifth"
	Teenth WeekSchedule = "teenth"
	Last   WeekSchedule = "last"
)
var (
	daysPerWeek int = 7
	modRef          = map[WeekSchedule]int{Second: daysPerWeek, Third: daysPerWeek * 2, Fourth: daysPerWeek * 3, Fifth: daysPerWeek * 4}
	monthDays       = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	weekDays        = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	start := 1
	if wSched == Teenth {
		start = 13
	}
	base := time.Date(year, month, start, 0, 0, 0, 0, time.UTC)
	if (base.Weekday() == wDay) && (wSched == First || wSched == Teenth) {
		return base.Day()
	}
	diff := int(wDay - base.Weekday())
	if diff < 0 {
		diff += daysPerWeek
	}
	base = base.AddDate(0, 0, start+diff+modRef[wSched]-base.Day())
	if wSched != Last {
		return base.Day()
	}
	if (year%400 == 0) || ((year%4 == 0) && (year%100 != 0)) {
		monthDays[2] = 29
	}
	for base.Day() <= monthDays[month]-daysPerWeek {
		base = base.AddDate(0, 0, daysPerWeek)
	}
	monthDays[2] = 28
	return base.Day()
}
