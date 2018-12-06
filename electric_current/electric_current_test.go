package electric_current

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestElectricCurrent(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptoampere).Yoctoamperes())
	AssertFloatEqual(t, 1e3, (1 * Attoampere).Zeptoamperes())
	AssertFloatEqual(t, 1e3, (1 * Femtoampere).Attoamperes())
	AssertFloatEqual(t, 1e3, (1 * Picoampere).Femtoamperes())
	AssertFloatEqual(t, 1e3, (1 * Nanoampere).Picoamperes())
	AssertFloatEqual(t, 1e3, (1 * Microampere).Nanoamperes())
	AssertFloatEqual(t, 1e3, (1 * Milliampere).Microamperes())

	AssertFloatEqual(t, 1e3, (1 * Ampere).Milliamperes())
	AssertFloatEqual(t, 1e2, (1 * Ampere).Deciamperes())
	AssertFloatEqual(t, 1e1, (1 * Ampere).Centiamperes())
	AssertFloatEqual(t, 1e0, (1 * Ampere).Amperes())
	AssertFloatEqual(t, 1e-1, (1 * Ampere).Decaamperes())
	AssertFloatEqual(t, 1e-2, (1 * Ampere).Hectoamperes())
	AssertFloatEqual(t, 1e-3, (1 * Ampere).Kiloamperes())

	AssertFloatEqual(t, 1e-3, (1 * Kiloampere).Megaamperes())
	AssertFloatEqual(t, 1e-3, (1 * Megaampere).Gigaamperes())
	AssertFloatEqual(t, 1e-3, (1 * Gigaampere).Teraamperes())
	AssertFloatEqual(t, 1e-3, (1 * Teraampere).Petaamperes())
	AssertFloatEqual(t, 1e-3, (1 * Petaampere).Exaamperes())
	AssertFloatEqual(t, 1e-3, (1 * Exaampere).Zettaamperes())
	AssertFloatEqual(t, 1e-3, (1 * Zettaampere).Yottaamperes())
}
