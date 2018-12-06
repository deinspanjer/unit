package luminous_flux

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestLuminousFlux(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e0, (1 * Lumen).Lumen())
}
