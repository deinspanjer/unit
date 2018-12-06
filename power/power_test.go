package power

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestPower(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptowatt).Yoctowatts())
	AssertFloatEqual(t, 1e3, (1 * Attowatt).Zeptowatts())
	AssertFloatEqual(t, 1e3, (1 * Femtowatt).Attowatts())
	AssertFloatEqual(t, 1e3, (1 * Picowatt).Femtowatts())
	AssertFloatEqual(t, 1e3, (1 * Nanowatt).Picowatts())
	AssertFloatEqual(t, 1e3, (1 * Microwatt).Nanowatts())
	AssertFloatEqual(t, 1e3, (1 * Milliwatt).Microwatts())

	AssertFloatEqual(t, 1e3, (1 * Watt).Milliwatts())
	AssertFloatEqual(t, 1e2, (1 * Watt).Centiwatts())
	AssertFloatEqual(t, 1e1, (1 * Watt).Deciwatts())
	AssertFloatEqual(t, 1e0, (1 * Watt).Watts())
	AssertFloatEqual(t, 1e-1, (1 * Watt).Decawatts())
	AssertFloatEqual(t, 1e-2, (1 * Watt).Hectowatts())
	AssertFloatEqual(t, 1e-3, (1 * Watt).Kilowatts())

	AssertFloatEqual(t, 1e-3, (1 * Kilowatt).Megawatts())
	AssertFloatEqual(t, 1e-3, (1 * Megawatt).Gigawatts())
	AssertFloatEqual(t, 1e-3, (1 * Gigawatt).Terawatts())
	AssertFloatEqual(t, 1e-3, (1 * Terawatt).Petawatts())
	AssertFloatEqual(t, 1e-3, (1 * Petawatt).Exawatts())
	AssertFloatEqual(t, 1e-3, (1 * Exawatt).Zettawatts())
	AssertFloatEqual(t, 1e-3, (1 * Zettawatt).Yottawatts())

	// non-SI
	AssertFloatEqual(t, 13.596216173039045, (10000 * Watt).Pferdestarke())
}
