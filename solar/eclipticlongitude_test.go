package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEclipticLongitude tests EclipticLongitude()
func TestEclipticLongitude(t *testing.T) {
	// stub
	jd := 2453097.0
	p := 2

	// assertion
	lambda := EclipticLongitude(jd, p)

	// assertion
	assert.Equal(t, 12.032185297938668, lambda)
}
