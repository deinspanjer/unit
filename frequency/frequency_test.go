package frequency

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestFrequency(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptohertz).Yoctohertz())
	AssertFloatEqual(t, 1e3, (1 * Attohertz).Zeptohertz())
	AssertFloatEqual(t, 1e3, (1 * Femtohertz).Attohertz())
	AssertFloatEqual(t, 1e3, (1 * Picohertz).Femtohertz())
	AssertFloatEqual(t, 1e3, (1 * Nanohertz).Picohertz())
	AssertFloatEqual(t, 1e3, (1 * Microhertz).Nanohertz())
	AssertFloatEqual(t, 1e6, (1 * Hertz).Microhertz())

	AssertFloatEqual(t, 1e3, (1 * Hertz).Millihertz())
	AssertFloatEqual(t, 1e2, (1 * Hertz).Centihertz())
	AssertFloatEqual(t, 1e1, (1 * Hertz).Decihertz())
	AssertFloatEqual(t, 1e0, (1 * Hertz).Hertz())
	AssertFloatEqual(t, 1e-1, (1 * Hertz).Decahertz())
	AssertFloatEqual(t, 1e-2, (1 * Hertz).Hectohertz())
	AssertFloatEqual(t, 1e-3, (1 * Hertz).Kilohertz())

	AssertFloatEqual(t, 1e-6, (1 * Hertz).Megahertz())
	AssertFloatEqual(t, 1e-3, (1 * Megahertz).Gigahertz())
	AssertFloatEqual(t, 1e-3, (1 * Gigahertz).Terahertz())
	AssertFloatEqual(t, 1e-3, (1 * Terahertz).Petahertz())
	AssertFloatEqual(t, 1e-3, (1 * Petahertz).Exahertz())
	AssertFloatEqual(t, 1e-3, (1 * Exahertz).Zettahertz())
	AssertFloatEqual(t, 1e-3, (1 * Zettahertz).Yottahertz())
}
