package solar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/codymj/solarcalc/julian"
	"github.com/stretchr/testify/require"
)

// TestGeoMeanLongitude tests GeoMeanLongitude()
func TestGeoMeanLongitude(t *testing.T) {
	// stub
	timeStr := "2000-01-01T12:00:00+00:00"
	tm, err := time.Parse(time.RFC3339, timeStr)
	require.NoError(t, err)

	jd := julian.FromCalendar(tm)

	// invocation
	L0 := GeoMeanLongitude(jd)

	// assertion
	assert.Equal(t, 280.46646, L0)
}
