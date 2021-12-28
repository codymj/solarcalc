package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEclipticLongitude tests EclipticLongitude()
func TestEclipticLongitude(t *testing.T) {
	// stub
	p := 2

	// invocation
	Pi := EclipticLongitude(p)

	// assertion
	assert.Equal(t, PiEarth, Pi)
}
