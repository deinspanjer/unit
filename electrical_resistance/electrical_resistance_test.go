package electrical_resistance

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestElectricalResistance(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e0, (1 * Ohm).Ohms())
}
