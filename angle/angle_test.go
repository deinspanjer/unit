package angle

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestAngle(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptoradian).Yoctoradians())
	AssertFloatEqual(t, 1e3, (1 * Attoradian).Zeptoradians())
	AssertFloatEqual(t, 1e3, (1 * Femtoradian).Attoradians())
	AssertFloatEqual(t, 1e3, (1 * Picoradian).Femtoradians())
	AssertFloatEqual(t, 1e3, (1 * Nanoradian).Picoradians())
	AssertFloatEqual(t, 1e3, (1 * Microradian).Nanoradians())
	AssertFloatEqual(t, 1e3, (1 * Milliradian).Microradians())

	AssertFloatEqual(t, 1e3, (1 * Radian).Milliradians())
	AssertFloatEqual(t, 1e2, (1 * Radian).Centiradians())
	AssertFloatEqual(t, 1e1, (1 * Radian).Deciradians())
	AssertFloatEqual(t, 1e0, (1 * Radian).Radians())

	// constant ~57.2958 from https://en.wikipedia.org/wiki/Radian#Conversion_between_radians_and_degrees
	AssertFloatEqual(t, 57.295779513082321, (1 * Radian).Degrees())

	// constants from https://en.wikipedia.org/wiki/Minute_and_second_of_arc#Symbols_and_abbreviations
	AssertFloatEqual(t, 17.453292519943295, (1 * Degree).Milliradians())
	AssertFloatEqual(t, 290.8882086657216, (1 * Arcminute).Microradians())
	AssertFloatEqual(t, 4.84813681109536, (1 * Arcsecond).Microradians())

	AssertFloatEqual(t, 60.0, (1 * Degree).Arcminutes())
	AssertFloatEqual(t, 60.0, (1 * Arcminute).Arcseconds())
	AssertFloatEqual(t, 1e3, (1 * Arcsecond).Milliarcseconds())
	AssertFloatEqual(t, 1e6, (1 * Arcsecond).Microarcseconds())
}
