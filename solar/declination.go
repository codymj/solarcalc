package solar

import "math"

// Declination (d) determines from which parts of the planet the object can be
// visible.
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
func Declination(jd float64, p int) (d float64) {
	lambda := EclipticLongitude(jd, p)
	e := ObliquityEcliptic(p)

	d = math.Atan(math.Sin(lambda*RAD) * math.Sin(e*RAD))
	return d * DEG
}
