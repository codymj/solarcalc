package solar

const (
	PiMercury = 230.3265
	PiVenus   = 73.7576
	PiEarth   = 102.9373
	PiMars    = 71.0041
	PiJupiter = 237.1015
	PiSaturn  = 99.4587
	PiUranus  = 5.4634
	PiNeptune = 182.2100
	PiPluto   = 184.5484
)

// PerihelionLongitude (Pi) is the plane of the orbit of the planet relative to
// the angle of its orbit.
// 	p: enum of the planet:
// 		0 = Mercury
//		1 = Venus
//		2 = Earth
//		3 = Mars
//		4 = Jupiter
//		5 = Saturn
// 		6 = Uranus
//		7 = Neptune
//		8 = Pluto
func PerihelionLongitude(p int) (Pi float64) {
	switch p {
	case 0:
		Pi = PiMercury
	case 1:
		Pi = PiVenus
	case 2:
		Pi = PiEarth
	case 3:
		Pi = PiMars
	case 4:
		Pi = PiJupiter
	case 5:
		Pi = PiSaturn
	case 6:
		Pi = PiUranus
	case 7:
		Pi = PiNeptune
	case 8:
		Pi = PiPluto
	}

	return
}
