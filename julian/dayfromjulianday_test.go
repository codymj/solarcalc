package julian

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDayFromJulianDay tests DayFromJulianDay()
func TestDayFromJulianDay(t *testing.T) {
	// stub
	jd := 2451545.2577546295

	// invocation
	d := DayFromJulianDay(jd)

	// assertion
	assert.Equal(t, uint(1), d)
}
