package illuminance

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestIlluminance(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e0, (1 * Lux).Lux())
}
