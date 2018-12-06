package length

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestLength_Imperial(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptometer).Yoctometers())
	AssertFloatEqual(t, 1e3, (1 * Attometer).Zeptometers())
	AssertFloatEqual(t, 1e3, (1 * Femtometer).Attometers())
	AssertFloatEqual(t, 1e3, (1 * Picometer).Femtometers())
	AssertFloatEqual(t, 1e3, (1 * Nanometer).Picometers())
	AssertFloatEqual(t, 1e3, (1 * Micrometer).Nanometers())
	AssertFloatEqual(t, 1e3, (1 * Millimeter).Micrometers())

	AssertFloatEqual(t, 1e3, (1 * Meter).Millimeters())
	AssertFloatEqual(t, 1e2, (1 * Meter).Centimeters())
	AssertFloatEqual(t, 1e1, (1 * Meter).Decimeters())
	AssertFloatEqual(t, 1e0, (1 * Meter).Meters())
	AssertFloatEqual(t, 1e-1, (1 * Meter).Decameters())
	AssertFloatEqual(t, 1e-2, (1 * Meter).Hectometers())
	AssertFloatEqual(t, 1e-3, (1 * Meter).Kilometers())
	AssertFloatEqual(t, 1e-4, (1 * Meter).ScandinavianMiles())

	AssertFloatEqual(t, 1e-3, (1 * Kilometer).Megameters())
	AssertFloatEqual(t, 1e-3, (1 * Megameter).Gigameters())
	AssertFloatEqual(t, 1e-3, (1 * Gigameter).Terameters())
	AssertFloatEqual(t, 1e-3, (1 * Terameter).Petameters())
	AssertFloatEqual(t, 1e-3, (1 * Petameter).Exameters())
	AssertFloatEqual(t, 1e-3, (1 * Exameter).Zettameters())
	AssertFloatEqual(t, 1e-3, (1 * Zettameter).Yottameters())

	// imperial
	AssertFloatEqual(t, 0.3048, (1 * Foot).Meters())
	AssertFloatEqual(t, 12, (1 * Foot).Inches())
	AssertFloatEqual(t, 3, (1 * Foot).Hands())
	AssertFloatEqual(t, 0.08333333333333334, (1 * Inch).Feet())
	AssertFloatEqual(t, 0.3333333333333333, (1 * Foot).Yards())
	AssertFloatEqual(t, 7.92, (1 * Link).Inches())
	AssertFloatEqual(t, 101.6, (1 * Hand).Millimeters())
	AssertFloatEqual(t, 0.18181818181818182, (1 * Yard).Rods())
	AssertFloatEqual(t, 0.25, (1 * Rod).Chains())
	AssertFloatEqual(t, 25., (1 * Rod).Links())
	AssertFloatEqual(t, 0.1, (1 * Chain).Furlongs())
	AssertFloatEqual(t, 12.5, (100 * Furlong).Miles())
	AssertFloatEqual(t, 54.6806649168854, (100 * Meter).Fathoms())
	AssertFloatEqual(t, 0.5399568034557236, (100 * Meter).Cables())

	// imperial maritime
	AssertFloatEqual(t, 1.8288, (1 * Fathom).Meters())
	AssertFloatEqual(t, 185.2, (1 * Cable).Meters())
	AssertFloatEqual(t, 0.8689762419006479, (1 * Mile).NauticalMiles())
}
