package solar

import "math"

// EquationOfCenter (C) is the angular difference between the actual position of
// a body in its elliptical orbit and the position it would occupy if its motion
// were uniform.
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
func EquationOfCenter(jd float64, p int) (C float64) {
	M := MeanAnomaly(jd, p)
	MRad := M * DegreeToRadian

	switch p {
	case 0:
		C = generalCalculation(
			C1Mercury, C2Mercury, C3Mercury,
			C4Mercury, C5Mercury, C6Mercury,
			MRad,
		)
	case 1:
		C = generalCalculation(
			C1Venus, C2Venus, C3Venus,
			C4Venus, C5Venus, C6Venus,
			MRad,
		)
	case 2:
		C = generalCalculation(
			C1Earth, C2Earth, C3Earth,
			C4Earth, C5Earth, C6Earth,
			MRad,
		)
	case 3:
		C = generalCalculation(
			C1Mars, C2Mars, C3Mars,
			C4Mars, C5Mars, C6Mars,
			MRad,
		)
	case 4:
		C = generalCalculation(
			C1Jupiter, C2Jupiter, C3Jupiter,
			C4Jupiter, C5Jupiter, C6Jupiter,
			MRad,
		)
	case 5:
		C = generalCalculation(
			C1Saturn, C2Saturn, C3Saturn,
			C4Saturn, C5Saturn, C6Saturn,
			MRad,
		)
	case 6:
		C = generalCalculation(
			C1Uranus, C2Uranus, C3Uranus,
			C4Uranus, C5Uranus, C6Uranus,
			MRad,
		)
	case 7:
		C = generalCalculation(
			C1Neptune, C2Neptune, C3Neptune,
			C4Neptune, C5Neptune, C6Neptune,
			MRad,
		)
	case 8:
		C = generalCalculation(
			C1Pluto, C2Pluto, C3Pluto,
			C4Pluto, C5Pluto, C6Pluto,
			MRad,
		)
	}

	return
}

func generalCalculation(c1, c2, c3, c4, c5, c6, m float64) float64 {
	return c1*math.Sin(m) + c2*math.Sin(2*m) + c3*math.Sin(3*m) +
		c4*math.Sin(4*m) + c5*math.Sin(5*m) + c6*math.Sin(6*m)
}
