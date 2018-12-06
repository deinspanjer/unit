package amount_of_substance

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestAmountOfSubstance(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e0, (1 * Mole).Moles())
}
