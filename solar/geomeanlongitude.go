package solar

// GeoMeanLongitude calculates the geometric mean longitude of the Sun.
// 	jd: Julian day
// 	L0: decimal degrees
func GeoMeanLongitude(jd float64) (L0 float64) {
	L0 = 280.46646 + 0.9856474*(jd-J2000)
	return
}
