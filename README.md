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
