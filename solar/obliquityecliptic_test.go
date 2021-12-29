package solar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestObliquityEcliptic tests ObliquityEcliptic()
func TestObliquityEcliptic(t *testing.T) {
	// stub
	p := 2

	// invocation
	e := ObliquityEcliptic(p)

	// assertion
	assert.Equal(t, EEarth, e)
}
