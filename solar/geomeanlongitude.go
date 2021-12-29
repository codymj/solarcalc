package solar

import "github.com/codymj/solarcalc/julian"

// GeoMeanLongitude (L0) calculates the geometric mean longitude of the Sun in
// decimal degrees.
// 	jd: Julian day
func GeoMeanLongitude(jd float64) (L0 float64) {
	L0 = 280.46646 + 0.9856474*(jd-julian.J2000)
	return
}
