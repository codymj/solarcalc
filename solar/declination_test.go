package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDeclination tests Declination()
func TestDeclination(t *testing.T) {
	// stub
	jd := 2453097.0
	p := 2

	// invocation
	d := Declination(jd, p)

	// assertion
	assert.Equal(t, 4.740184662324431, d)
}
