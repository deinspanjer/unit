package luminous_intensity

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestLuminousIntensity(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e0, (1 * Candela).Candela())
}
