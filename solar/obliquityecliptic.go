package solar

const (
	eMercury = 0.0351
	eVenus   = 2.6376
	eEarth   = 23.4393
	eMars    = 25.1918
	eJupiter = 3.1189
	eSaturn  = 26.7285
	eUranus  = 82.2298
	eNeptune = 27.8477
	ePluto   = 119.6075
)

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
		e = eMercury
	case 1:
		e = eVenus
	case 2:
		e = eEarth
	case 3:
		e = eMars
	case 4:
		e = eJupiter
	case 5:
		e = eSaturn
	case 6:
		e = eUranus
	case 7:
		e = eNeptune
	case 8:
		e = ePluto
	}

	return
}
