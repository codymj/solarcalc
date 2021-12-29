package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTrueAnomaly tests TrueAnomaly()
func TestTrueAnomaly(t *testing.T) {
	// stub
	jd := 2453097.0
	p := 2

	// invocation
	v := TrueAnomaly(jd, p)

	// assertion
	assert.Equal(t, 89.09488529793867, v)
}
