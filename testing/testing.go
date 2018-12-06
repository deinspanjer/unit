package testing

import (
	"math"
	"testing"

	"github.com/stretchrcom/testify/assert"
)

// AssertFloatEqual asserts that the actual float64 value is within
// two epsilons of the expected float64 value.
func AssertFloatEqual(t *testing.T, expected, actual float64, args ...interface{}) {
	epsilon := math.Abs(math.Nextafter(expected, actual) - expected)
	epsilon += math.Abs(math.Nextafter(actual, expected) - actual)
	assert.InDelta(t, expected, actual, epsilon, args...)
}
