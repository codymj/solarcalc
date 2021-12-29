package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPerihelionLongitude tests PerihelionLongitude()
func TestPerihelionLongitude(t *testing.T) {
	// stub
	p := 2

	// invocation
	w := PerihelionLongitude(p)

	// assertion
	assert.Equal(t, WEarth, w)
}
