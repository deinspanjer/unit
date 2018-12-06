package length

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestLength_ScandinavianMilesLength(t *testing.T) {
	AssertFloatEqual(t, 1e-4, (1 * Metre).ScandinavianMiles())
}
