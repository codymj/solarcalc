package julian

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestToJulianCentury tests ToJulianCentury()
func TestToJulianCentury(t *testing.T) {
	// stub
	timeStr := "2000-01-01T18:11:10Z"
	tm, err := time.Parse(time.RFC3339, timeStr)
	require.NoError(t, err)

	jd := FromCalendar(tm)

	// invocation
	jc := ToJulianCentury(jd)

	assert.Equal(t, 24515.955985266733, jc)
}
