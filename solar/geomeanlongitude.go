package solar

// GeoMeanLongitude (L0) calculates the geometric mean longitude of the Sun in
// decimal degrees.
// 	jd: Julian day
func GeoMeanLongitude(jd float64) (L0 float64) {
	L0 = 280.46646 + 0.9856474*(jd-J2000)
	return
}
