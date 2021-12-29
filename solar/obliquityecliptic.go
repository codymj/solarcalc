package solar

// ObliquityEcliptic (e) is the angle between the ecliptic and the celestial
// equator of the planet.
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
func ObliquityEcliptic(p int) (e float64) {
	switch p {
	case 0:
		e = EMercury
	case 1:
		e = EVenus
	case 2:
		e = EEarth
	case 3:
		e = EMars
	case 4:
		e = EJupiter
	case 5:
		e = ESaturn
	case 6:
		e = EUranus
	case 7:
		e = ENeptune
	case 8:
		e = EPluto
	}

	return
}
