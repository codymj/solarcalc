# solarcalc

A Go library for calculating certain phenomena in relation to the solar system.

## Data

### Planet Enumeration

These enumerations are used throughout the module:

| enum | planet  |
|------|---------|
| 0    | Mercury |
| 1    | Venus   |
| 2    | Earth   |
| 3    | Mars    |
| 4    | Jupiter |
| 5    | Saturn  |
| 6    | Uranus  |
| 7    | Neptune |
| 8    | Pluto   |

## Functions

### `func MeanAnomaly(jd float64, p int) (M float64)`

The position that the planet would have relative to its perihelion if the orbit 
of the planet were a circle.

| parameter | description            |
|-----------|------------------------|
| jd        | julian day             |
| p         | planet enum            |
 | M         | mean anomaly (degrees) |

### `func EquationOfCenter(jd float64, p int) (C float64)`

The angular difference between the actual position of a body in its elliptical 
orbit and the position it would occupy if its motion were uniform, in a circular
orbit of the same period.

| parameter | description                  |
|-----------|------------------------------|
| jd        | julian day                   |
| p         | planet enum                  |
| C         | equation of center (degrees) |

### `func PerihelionLongitude(p int) (w float64)`

The sum of the longitude of ascending node (measured on the ecliptic plane) and
the argument of periapsis (measured on the orbital plane).

| parameter | description                  |
|-----------|------------------------------|
| p         | planet enum                  |
| w         | equation of center (degrees) |

### `func ObliquityEcliptic(p int) (w float64)`

The angle between the ecliptic and the celestial equator of the planet.

| parameter | description                         |
|-----------|-------------------------------------|
| p         | planet enum                         |
| e         | obliquity of the ecliptic (degrees) |

### `func EclipticLongitude(jd float64, p int) (lambda float64)`

The position along the ecliptic relative to the vernal equinox.

| parameter | description                  |
|-----------|------------------------------|
| jd        | julian day                   |
| p         | planet enum                  |
| lambda    | ecliptic longitude (degrees) |

### `func RightAscension(jd float64, p int) (a float64)`

Right ascension and declination define the position of a celestial object.

| parameter | description               |
|-----------|---------------------------|
| jd        | julian day                |
| p         | planet enum               |
| a         | right ascension (degrees) |

### `func Declination(jd float64, p int) (d float64)`

Right ascension and declination define the position of a celestial object.

| parameter | description           |
|-----------|-----------------------|
| jd        | julian day            |
| p         | planet enum           |
| d         | declination (degrees) |
