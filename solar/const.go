package solar

import "math"

const (
	RAD = math.Pi / 180
	DEG = 180 / math.Pi
)

// mean anomaly
const (
	M0Mercury = 174.7948
	M1Mercury = 4.09233445

	M0Venus = 50.4161
	M1Venus = 1.60213034

	M0Earth = 357.5291
	M1Earth = 0.98560028

	M0Mars = 19.3730
	M1Mars = 0.52402068

	M0Jupiter = 20.0202
	M1Jupiter = 0.08308529

	M0Saturn = 317.0207
	M1Saturn = 0.03344414

	M0Uranus = 141.0498
	M1Uranus = 0.01172834

	M0Neptune = 256.2250
	M1Neptune = 0.00598103

	M0Pluto = 14.882
	M1Pluto = 0.00396
)

// true anomaly
const (
	C1Mercury = 23.4400
	C2Mercury = 2.9818
	C3Mercury = 0.5255
	C4Mercury = 0.1058
	C5Mercury = 0.0241
	C6Mercury = 0.0055

	C1Venus = 0.7758
	C2Venus = 0.0033
	C3Venus = 0.0000
	C4Venus = 0.0000
	C5Venus = 0.0000
	C6Venus = 0.0000

	C1Earth = 1.9148
	C2Earth = 0.0200
	C3Earth = 0.0003
	C4Earth = 0.0000
	C5Earth = 0.0000
	C6Earth = 0.0000

	C1Mars = 10.6912
	C2Mars = 0.6228
	C3Mars = 0.0503
	C4Mars = 0.0046
	C5Mars = 0.0005
	C6Mars = 0.0000

	C1Jupiter = 5.5549
	C2Jupiter = 0.1683
	C3Jupiter = 0.0071
	C4Jupiter = 0.0003
	C5Jupiter = 0.0000
	C6Jupiter = 0.0000

	C1Saturn = 6.3585
	C2Saturn = 0.2204
	C3Saturn = 0.0106
	C4Saturn = 0.0006
	C5Saturn = 0.0000
	C6Saturn = 0.0000

	C1Uranus = 5.3042
	C2Uranus = 0.1534
	C3Uranus = 0.0062
	C4Uranus = 0.0003
	C5Uranus = 0.0000
	C6Uranus = 0.0000

	C1Neptune = 1.0302
	C2Neptune = 0.0058
	C3Neptune = 0.0000
	C4Neptune = 0.0000
	C5Neptune = 0.0000
	C6Neptune = 0.0000

	C1Pluto = 28.3150
	C2Pluto = 4.3408
	C3Pluto = 0.9214
	C4Pluto = 0.2235
	C5Pluto = 0.0627
	C6Pluto = 0.0174
)

// obliquity of the ecliptic
const (
	EMercury = 0.0351
	EVenus   = 2.6376
	EEarth   = 23.4393
	EMars    = 25.1918
	EJupiter = 3.1189
	ESaturn  = 26.7285
	EUranus  = 82.2298
	ENeptune = 27.8477
	EPluto   = 119.6075
)

// perihelion
const (
	WMercury = 230.3265
	WVenus   = 73.7576
	WEarth   = 102.9373
	WMars    = 71.0041
	WJupiter = 237.1015
	WSaturn  = 99.4587
	WUranus  = 5.4634
	WNeptune = 182.2100
	WPluto   = 184.5484
)
