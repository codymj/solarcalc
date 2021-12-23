package julian

import (
	"time"
)

// FromCalendar converts a calendar datetime into a Julian day
func FromCalendar(t time.Time) float64 {
	A := (1461 * (t.Year() + 4800 + (int(t.Month())-14)/12)) / 4
	B := (367 * (int(t.Month()) - 2 - 12*((int(t.Month())-14)/12))) / 12
	C := (3 * ((t.Year() + 4900 + (int(t.Month())-14)/12) / 100)) / 4
	D := t.Day() - 32075
	H := float64(t.Hour()-12) / 24
	M := float64(t.Minute()) / 1440
	S := float64(t.Second()) / 86400
	return float64(A+B-C+D) + H + M + S
}
