package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMeanAnomaly tests MeanAnomaly()
func TestMeanAnomaly(t *testing.T) {
	// stub
	jd := 2453097.0
	p := 2

	// invocation
	M := MeanAnomaly(jd, p)

	// assertion
	assert.Equal(t, 87.18073456000002, M)
}
