package julian

import "math"

// DayFromJulianDay calculates the day number of a solar year from Julian day.
// 	jd: Julian day
// 	dn: day number
func DayFromJulianDay(jd float64) (dn uint) {
	z := math.Floor(jd + 0.5)
	f := jd + 0.5 - z

	var A, alpha float64
	if z < 2299161 {
		A = z
	} else {
		alpha = math.Floor((z - 1867216.25) / 36524.25)
		A = z + 1 + alpha - math.Floor(alpha/4.0)
	}

	B := A + 1524.0
	C := math.Floor((B - 122.1) / 365.25)
	D := math.Floor(365.25 * C)
	E := math.Floor((B - D) / 30.6001)

	day := uint(B - D - math.Floor(30.6001*E) + f)
	month := func(E float64) uint {
		if E < 14 {
			return uint(E - 1)
		} else {
			return uint(E - 13)
		}
	}(E)
	year := func(C float64, month uint) uint {
		if month > 2 {
			return uint(C - 4716)
		} else {
			return uint(C - 4715)
		}
	}(C, month)
	k := func(year uint) uint {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 1
		} else {
			return 2
		}
	}(year)

	dn = uint(math.Floor(float64(
		275*month/9 - k*uint(math.Floor(float64((month+9)/12))) + day - 30)),
	)
	return
}
