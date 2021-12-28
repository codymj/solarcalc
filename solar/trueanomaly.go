package solar

import (
	"math"
)

const (
	DegreeToRadian = math.Pi / 180

	C1Earth = 1.9148
	C2Earth = 0.0200
	C3Earth = 0.0003
	C4Earth = 0.0000
	C5Earth = 0.0000
	C6Earth = 0.0000
)

// TrueAnomaly is the sum of the mean anomaly (M) and equation of center (C).
//	v: true anomaly (v = C + M)
func TrueAnomaly(jd float64) (v float64) {
	M := MeanAnomaly(jd)
	MRad := M * DegreeToRadian

	C := C1Earth*math.Sin(MRad) + C2Earth*math.Sin(2*MRad) +
		C3Earth*math.Sin(3*MRad) + C4Earth*math.Sin(4*MRad) +
		C5Earth*math.Sin(5*MRad) + C6Earth*math.Sin(6*MRad)

	v = M + C
	return
}
