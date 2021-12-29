package solar

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
		w = WMercury
	case 1:
		w = WVenus
	case 2:
		w = WEarth
	case 3:
		w = WMars
	case 4:
		w = WJupiter
	case 5:
		w = WSaturn
	case 6:
		w = WUranus
	case 7:
		w = WNeptune
	case 8:
		w = WPluto
	}

	return
}
