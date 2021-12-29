package solar

import "math"

// RightAscension (a) determines when an object is visible (in degrees).
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
func RightAscension(jd float64, p int) (a float64) {
	lambda := EclipticLongitude(jd, p)
	e := ObliquityEcliptic(p)

	a = math.Atan2(math.Sin(lambda*RAD)*math.Cos(e*RAD), math.Cos(lambda*RAD))
	return a * DEG
}
