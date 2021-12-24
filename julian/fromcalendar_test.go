package julian

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

// TestFromCalendar tests FromCalendar()
func TestFromCalendar(t *testing.T) {
	// stub
	timeStr := "2000-01-01T18:11:10-01:00"
	tm, err := time.Parse(time.RFC3339, timeStr)
	require.NoError(t, err)

	// invocation
	jd := FromCalendar(tm)

	// assertion
	assert.Equal(t, 2.451545216087963e+06, jd)
}
