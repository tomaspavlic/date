package date

import (
	"time"
	"unsafe"
)

const (
	hasMonotonic = 1 << 63
	minDayNumber = 0
	maxDayNumber = 3_652_058 // Maps to December 31, 9999.
	nsecShift    = 30
	secondsInDay = 86_400

	wallToInternal int64 = (1884*365 + 1884/4 - 1884/100 + 1884/400) * secondsInDay
)

// Represents dates with values from January 1, 0001 Anno Domini (Common Era)
// through December 31, 9999 A.D (C.E) in Gregorian calendar.
type Date int64

// ext returns a ext value from time.Time
func ext(t *time.Time) int64 {
	return *(*int64)(unsafe.Add(unsafe.Pointer(t), 8))
}

// wall returns a wall value from time.Time
func wall(t *time.Time) uint64 {
	return *(*uint64)(unsafe.Pointer(t))
}

// sec returns the time's seconds since Jan 1 year 1.
func sec(t *time.Time) int64 {
	if wall(t)&hasMonotonic != 0 {
		return wallToInternal + int64(wall(t)<<1>>(nsecShift+1))
	}

	return ext(t)
}

// dayNumber extracts ext field from time struct which contains
// number of seconds since January 1, 0001.
func dayNumber(t *time.Time) int64 {
	return sec(t) / secondsInDay
}

// unsafeCreateTime creates a time.Time from date.Date by initializing
// time.Time with ext.
func unsafeCreateTime(d Date) time.Time {
	t := time.Time{}
	ext := (*int64)(unsafe.Add(unsafe.Pointer(&t), 8))
	*ext = int64(d) * secondsInDay
	return t
}

// Create creates a date containing only day number.
func Create(year int, month time.Month, day int) Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return Date(dayNumber(&t))
}

// Today returns current date only.
func Today() Date {
	n := time.Now()
	return Date(dayNumber(&n))
}

// FromTime converts time.Time to Date.
func FromTime(t *time.Time) Date {
	return Date(dayNumber(t))
}

// Parse parses a formatted string and returns the date value it represents.
// See the documentation for the constant called Layout to see how to
// represent the format. The second argument must be parseable using
// the format string (layout) provided as the first argument.
func Parse(layout, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return -1, err
	}

	return Date(dayNumber(&t)), nil
}

// Since returns number of days since d.
func Since(d Date) int64 {
	return int64(Today() - d)
}

// Weekday returns the day of the week specified by t.
func (d Date) Weekday() time.Weekday {
	return time.Weekday((int(d) + 1) % 7)
}

// AddDays the specified number of days to the value of this instance.
func (d Date) AddDays(value int) Date {
	newDateNumber := int64(d) + int64(value)
	if newDateNumber > maxDayNumber {
		panic("Calculated date is out of range")
	}

	return Date(newDateNumber)
}

// Year returns the year in which t occurs.
func (d Date) Year() int {
	return unsafeCreateTime(d).Year()
}

// Month returns the month of the year specified by t.
func (d Date) Month() time.Month {
	return unsafeCreateTime(d).Month()
}

// Day returns the day of specified by d.
func (d Date) Day() int {
	return unsafeCreateTime(d).Day()
}

// YearDay returns the day of the year specified by t, in the
// range [1,365] for non-leap years, and [1,366] in leap years.
func (d Date) YearDay() int {
	return unsafeCreateTime(d).YearDay()
}

// AddDate returns the date corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
func (d Date) AddDate(years int, months int, days int) Date {
	t := unsafeCreateTime(d).AddDate(years, months, days)
	return Date(dayNumber(&t))
}

// Before reports whether the date instant d is before u.
func (d Date) Before(u Date) bool {
	return d < u
}

// Equal reports whether d and u represent the same date instant.
func (d Date) Equal(u Date) bool {
	return d == u
}

// After reports whether the date instant d is after u.
func (d Date) After(u Date) bool {
	return d > u
}

// Sub returns number of days between d and u.
func (d Date) Sub(u Date) int {
	return int(d - u)
}

// ToTime creates time.Time from Date by adding time.
func (d Date) ToTime(hour, min, sec, nsec int, loc *time.Location) time.Time {
	t := unsafeCreateTime(d)
	newTime := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		hour,
		min,
		sec,
		nsec,
		loc,
	)

	return newTime
}

// String returns the date formatted using the format string
//	"2006-01-02"
func (d Date) String() string {
	t := unsafeCreateTime(d)
	return t.Format("2006-01-02")
}
