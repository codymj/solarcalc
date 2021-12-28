package solar

import (
	"math"
)

const (
	J2000 = 2451545.0

	M0 = 357.5291
	M1 = 0.98560028
)

// MeanAnomaly calculates the position that the Earth would have relative to its
// perihelion if the orbit were a circle.
// 	jd: Julian day
// 	M: mean anomaly in degrees
func MeanAnomaly(jd float64) (M float64) {
	M = math.Mod(M0+M1*(jd-J2000), 360.0)
	return
}
