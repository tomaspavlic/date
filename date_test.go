package date

import (
	"reflect"
	"testing"
	"time"
)

func TestYear(t *testing.T) {
	tests := []struct {
		name string
		do   Date
		want int
	}{
		{"1989", Create(1989, 10, 11), 1989},
		{"1434", Create(1434, 1, 1), 1434},
		{"2055", Create(2055, 5, 22), 2055},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.do.Year(); got != tt.want {
				t.Errorf("Date.Year() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonth(t *testing.T) {
	tests := []struct {
		name string
		d    Date
		want time.Month
	}{
		{"October", Create(1989, 10, 11), time.October},
		{"January", Create(1434, 1, 1), time.January},
		{"May", Create(2055, 5, 22), time.May},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Month(); got != tt.want {
				t.Errorf("Date.Month() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay(t *testing.T) {
	tests := []struct {
		name string
		d    Date
		want int
	}{
		{"11", Create(1989, 10, 11), 11},
		{"1", Create(1434, 1, 1), 1},
		{"22", Create(2055, 5, 22), 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Day(); got != tt.want {
				t.Errorf("Date.Day() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddDays(t *testing.T) {
	tests := []struct {
		name  string
		d     Date
		value int
		want  Date
	}{
		{"10", Create(1989, 10, 25), 10, Create(1989, 11, 4)},
		{"222", Create(1434, 1, 1), 222, Create(1434, 8, 11)},
		{"123", Create(2055, 5, 22), 123, Create(2055, 9, 22)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddDays(tt.value); got != tt.want {
				t.Errorf("Date.AddDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekday(t *testing.T) {
	tests := []struct {
		name string
		d    Date
		want time.Weekday
	}{
		{"Sunday", Create(2022, 6, 19), time.Sunday},
		{"Friday", Create(2021, 1, 1), time.Friday},
		{"Wednesday", Create(2021, 2, 10), time.Wednesday},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Weekday(); got != tt.want {
				t.Errorf("Date.Weekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYearDay(t *testing.T) {
	tests := []struct {
		name string
		d    Date
		want int
	}{
		{"Sunday", Create(2022, 12, 31), 365},
		{"1", Create(2021, 1, 1), 1},
		{"Wednesday", Create(2021, 2, 10), 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.YearDay(); got != tt.want {
				t.Errorf("Date.YearDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddDate(t *testing.T) {
	type args struct {
		years  int
		months int
		days   int
	}
	tests := []struct {
		name string
		d    Date
		args args
		want Date
	}{
		{"Years", Create(1989, 2, 19), args{1, 0, 0}, Create(1990, 2, 19)},
		{"Months", Create(2021, 1, 1), args{0, 3, 0}, Create(2021, 4, 1)},
		{"Days", Create(2021, 2, 10), args{0, 0, 10}, Create(2021, 2, 20)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AddDate(tt.args.years, tt.args.months, tt.args.days); got != tt.want {
				t.Errorf("Date.AddDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTime(t *testing.T) {
	type args struct {
		hour int
		min  int
		sec  int
		nsec int
		loc  *time.Location
	}
	tests := []struct {
		name string
		d    Date
		args args
		want time.Time
	}{
		{
			"2022-02-05 02:00:05",
			Create(2022, 2, 5),
			args{2, 0, 5, 0, time.UTC},
			time.Date(2022, 2, 5, 2, 0, 5, 0, time.UTC),
		},
		{
			"1989-04-22 11:10:55",
			Create(1989, 4, 22),
			args{11, 10, 55, 0, time.UTC},
			time.Date(1989, 4, 22, 11, 10, 55, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ToTime(tt.args.hour, tt.args.min, tt.args.sec, tt.args.nsec, tt.args.loc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.ToTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
