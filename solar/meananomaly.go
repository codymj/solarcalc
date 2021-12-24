package solar

import (
	"math"

	"github.com/codymj/solarcalc/julian"
)

// MeanAnomaly calculates the position that the Earth would have relative to its
// perihelion if the orbit were a circle.
//
// jd - Julian day
//
// M - mean anomaly in degrees
func MeanAnomaly(jd float64) (M float64) {
	M = math.Mod(M0+M1*(jd-julian.J2000), 360.0)
	return
}
