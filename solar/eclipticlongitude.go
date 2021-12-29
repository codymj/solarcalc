package solar

import "math"

// EclipticLongitude (lambda) is the position along the ecliptic relative to the
// vernal equinox (in degrees).
//	jd: Julian day
//	p: enum of the planet:
// 		0 = Mercury
//		1 = Venus
//		2 = Earth
//		3 = Mars
//		4 = Jupiter
//		5 = Saturn
// 		6 = Uranus
//		7 = Neptune
//		8 = Pluto
func EclipticLongitude(jd float64, p int) (lambda float64) {
	M := MeanAnomaly(jd, p)
	w := PerihelionLongitude(p)
	C := EquationOfCenter(jd, p)
	L := M + w + 180

	lambda = L + C
	for lambda > 360.0 {
		lambda = math.Mod(lambda, 360.0)
	}

	return lambda
}
