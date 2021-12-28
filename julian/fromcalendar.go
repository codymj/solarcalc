package julian

import (
	"time"
)

const (
	SecondsPerDay = 86400.0
	MinutesPerDay = 1440.0
	HoursPerDay   = 24.0
)

// FromCalendar converts a calendar datetime into a Julian day
//
// t - time
//
// jd - Julian day
func FromCalendar(t time.Time) (jd float64) {
	A := (1461 * (t.Year() + 4800 + (int(t.Month())-14)/12)) / 4
	B := (367 * (int(t.Month()) - 2 - 12*((int(t.Month())-14)/12))) / 12
	C := (3 * ((t.Year() + 4900 + (int(t.Month())-14)/12) / 100)) / 4
	D := t.Day() - 32075
	H := float64(t.Hour()-12) / HoursPerDay
	M := float64(t.Minute()) / MinutesPerDay
	S := float64(t.Second()) / SecondsPerDay
	_, offset := t.Zone()
	Z := float64(offset) / SecondsPerDay

	jd = float64(A+B-C+D) + H + M + S + Z
	return
}
