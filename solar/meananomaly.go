package solar

import (
	"math"

	"github.com/codymj/solarcalc/julian"
)

const ()

// MeanAnomaly (M) calculates the position that the Earth would have relative to
// its perihelion if the orbit were a circle.
// 	jd: Julian day
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
func MeanAnomaly(jd float64, p int) (M float64) {
	switch p {
	case 0:
		M = math.Mod(M0Mercury+M1Mercury*(jd-julian.J2000), 360.0)
	case 1:
		M = math.Mod(M0Venus+M1Venus*(jd-julian.J2000), 360.0)
	case 2:
		M = math.Mod(M0Earth+M1Earth*(jd-julian.J2000), 360.0)
	case 3:
		M = math.Mod(M0Mars+M1Mars*(jd-julian.J2000), 360.0)
	case 4:
		M = math.Mod(M0Jupiter+M1Jupiter*(jd-julian.J2000), 360.0)
	case 5:
		M = math.Mod(M0Saturn+M1Saturn*(jd-julian.J2000), 360.0)
	case 6:
		M = math.Mod(M0Uranus+M1Uranus*(jd-julian.J2000), 360.0)
	case 7:
		M = math.Mod(M0Neptune+M1Neptune*(jd-julian.J2000), 360.0)
	case 8:
		M = math.Mod(M0Pluto+M1Pluto*(jd-julian.J2000), 360.0)
	}

	return
}
