package julian

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

// TestFromCalendar
func TestFromCalendar(t *testing.T) {
	// stub
	timeStr := "2000-01-01T18:11:10Z"
	tm, err := time.Parse(time.RFC3339, timeStr)
	require.NoError(t, err)

	// invocation
	jd := FromCalendar(tm)

	// assertion
	assert.Equal(t, 2.4515452577546295e+06, jd)
}
