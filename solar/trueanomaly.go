package solar

// TrueAnomaly (v) is the sum of the mean anomaly (M) and the result from the
// equation of center (C).
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
func TrueAnomaly(jd float64, p int) (v float64) {
	M := MeanAnomaly(jd, p)
	C := EquationOfCenter(jd, p)

	return M + C
}
