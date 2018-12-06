package electrical_conductance

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestElectricalConductance(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e0, (1 * Siemens).Siemens())
}
