package solar

import "github.com/codymj/solarcalc/julian"

// GeoMeanLongitude calculates the geometric mean longitude of the sun.
//
// jd - Julian day
//
// L0 - decimal degrees
func GeoMeanLongitude(jd float64) (L0 float64) {
	L0 = 280.46646 + 0.9856474*(jd-julian.J2000)
	return
}
