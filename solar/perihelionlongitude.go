package solar

const (
	wMercury = 230.3265
	wVenus   = 73.7576
	wEarth   = 102.9373
	wMars    = 71.0041
	wJupiter = 237.1015
	wSaturn  = 99.4587
	wUranus  = 5.4634
	wNeptune = 182.2100
	wPluto   = 184.5484
)

// PerihelionLongitude (w) is the sum of the longitude of ascending node
// (measured on the ecliptic plane) and the argument of periapsis (measured on
// the orbital plane).
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
func PerihelionLongitude(p int) (w float64) {
	switch p {
	case 0:
		w = wMercury
	case 1:
		w = wVenus
	case 2:
		w = wEarth
	case 3:
		w = wMars
	case 4:
		w = wJupiter
	case 5:
		w = wSaturn
	case 6:
		w = wUranus
	case 7:
		w = wNeptune
	case 8:
		w = wPluto
	}

	return
}
