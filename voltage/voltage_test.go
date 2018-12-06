package voltage

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestVoltage(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptovolt).Yoctovolts())
	AssertFloatEqual(t, 1e3, (1 * Attovolt).Zeptovolts())
	AssertFloatEqual(t, 1e3, (1 * Femtovolt).Attovolts())
	AssertFloatEqual(t, 1e3, (1 * Picovolt).Femtovolts())
	AssertFloatEqual(t, 1e3, (1 * Nanovolt).Picovolts())
	AssertFloatEqual(t, 1e3, (1 * Microvolt).Nanovolts())
	AssertFloatEqual(t, 1e3, (1 * Millivolt).Microvolts())

	AssertFloatEqual(t, 1e3, (1 * Volt).Millivolts())
	AssertFloatEqual(t, 1e2, (1 * Volt).Centivolts())
	AssertFloatEqual(t, 1e1, (1 * Volt).Decivolts())
	AssertFloatEqual(t, 1e0, (1 * Volt).Volts())
	AssertFloatEqual(t, 1e-1, (1 * Volt).Decavolts())
	AssertFloatEqual(t, 1e-2, (1 * Volt).Hectovolts())
	AssertFloatEqual(t, 1e-3, (1 * Volt).Kilovolts())

	AssertFloatEqual(t, 1e-3, (1 * Kilovolt).Megavolts())
	AssertFloatEqual(t, 1e-3, (1 * Megavolt).Gigavolts())
	AssertFloatEqual(t, 1e-3, (1 * Gigavolt).Teravolts())
	AssertFloatEqual(t, 1e-3, (1 * Teravolt).Petavolts())
	AssertFloatEqual(t, 1e-3, (1 * Petavolt).Exavolts())
	AssertFloatEqual(t, 1e-3, (1 * Exavolt).Zettavolts())
	AssertFloatEqual(t, 1e-3, (1 * Zettavolt).Yottavolts())
}
