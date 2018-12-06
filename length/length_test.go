package length

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestLength(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptometre).Yoctometres())
	AssertFloatEqual(t, 1e3, (1 * Attometre).Zeptometres())
	AssertFloatEqual(t, 1e3, (1 * Femtometre).Attometres())
	AssertFloatEqual(t, 1e3, (1 * Picometre).Femtometres())
	AssertFloatEqual(t, 1e3, (1 * Nanometre).Picometres())
	AssertFloatEqual(t, 1e3, (1 * Micrometre).Nanometres())
	AssertFloatEqual(t, 1e3, (1 * Millimetre).Micrometres())

	AssertFloatEqual(t, 1e3, (1 * Metre).Millimetres())
	AssertFloatEqual(t, 1e2, (1 * Metre).Centimetres())
	AssertFloatEqual(t, 1e1, (1 * Metre).Decimetres())
	AssertFloatEqual(t, 1e0, (1 * Metre).Metres())
	AssertFloatEqual(t, 1e-1, (1 * Metre).Decametres())
	AssertFloatEqual(t, 1e-2, (1 * Metre).Hectometres())
	AssertFloatEqual(t, 1e-3, (1 * Metre).Kilometres())
	AssertFloatEqual(t, 1e-4, (1 * Metre).ScandinavianMiles())

	AssertFloatEqual(t, 1e-3, (1 * Kilometre).Megametres())
	AssertFloatEqual(t, 1e-3, (1 * Megametre).Gigametres())
	AssertFloatEqual(t, 1e-3, (1 * Gigametre).Terametres())
	AssertFloatEqual(t, 1e-3, (1 * Terametre).Petametres())
	AssertFloatEqual(t, 1e-3, (1 * Petametre).Exametres())
	AssertFloatEqual(t, 1e-3, (1 * Exametre).Zettametres())
	AssertFloatEqual(t, 1e-3, (1 * Zettametre).Yottametres())

	// US
	AssertFloatEqual(t, 0.3048, (1 * Foot).Metres())
	AssertFloatEqual(t, 12, (1 * Foot).Inches())
	AssertFloatEqual(t, 3, (1 * Foot).Hands())
	AssertFloatEqual(t, 0.08333333333333334, (1 * Inch).Feet())
	AssertFloatEqual(t, 0.3333333333333333, (1 * Foot).Yards())
	AssertFloatEqual(t, 7.92, (1 * Link).Inches())
	AssertFloatEqual(t, 101.6, (1 * Hand).Millimetres())
	AssertFloatEqual(t, 0.18181818181818182, (1 * Yard).Rods())
	AssertFloatEqual(t, 0.25, (1 * Rod).Chains())
	AssertFloatEqual(t, 25., (1 * Rod).Links())
	AssertFloatEqual(t, 0.1, (1 * Chain).Furlongs())
	AssertFloatEqual(t, 12.5, (100 * Furlong).Miles())
	AssertFloatEqual(t, 54.6806649168854, (100 * Metre).Fathoms())
	AssertFloatEqual(t, 0.5399568034557236, (100 * Metre).Cables())

	// US maritime
	AssertFloatEqual(t, 1.8288, (1 * Fathom).Metres())
	AssertFloatEqual(t, 185.2, (1 * Cable).Metres())
	AssertFloatEqual(t, 0.8689762419006479, (1 * Mile).NauticalMiles())

	// space
	AssertFloatEqual(t, 389.17240036420395, (1 * AstronomicalUnit).LunarDistances())
	AssertFloatEqual(t, 0.0025695552897999903, (1 * LunarDistance).AstronomicalUnits())
	AssertFloatEqual(t, 63241.07708426628, (1 * LightYear).AstronomicalUnits())
	AssertFloatEqual(t, 1.581250740982066e-05, (1 * AstronomicalUnit).LightYears())
}
