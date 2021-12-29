package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRightAscension tests RightAscension()
func TestRightAscension(t *testing.T) {
	// stub
	jd := 2453097.0
	p := 2

	// invocation
	a := RightAscension(jd, p)

	// assertion
	assert.Equal(t, 11.064870715700355, a)
}
