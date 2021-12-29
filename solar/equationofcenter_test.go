package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEquationOfCenter tests EquationOfCenter()
func TestEquationOfCenter(t *testing.T) {
	// stub
	jd := 2453097.0
	p := 2

	// invocation
	C := EquationOfCenter(jd, p)

	// assertion
	assert.Equal(t, 1.9141507379386618, C)
}
