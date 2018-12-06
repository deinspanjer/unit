package acceleration

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestAcceleration(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Standard_gravity#Conversions
	AssertFloatEqual(t, 0.03280839895013123, (1 * Gal).FeetPerSecondSquared())
	AssertFloatEqual(t, 0.01, (1 * Gal).MetersPerSecondSquared())
	AssertFloatEqual(t, 1e0, (1 * Gal).CentimetersPerSecondSquared()) // alias
	AssertFloatEqual(t, 0.0010197162129779284, (1 * Gal).StandardGravity())

	AssertFloatEqual(t, 30.4800, (1 * FootPerSecondSquared).Gals())
	AssertFloatEqual(t, 0.304800, (1 * FootPerSecondSquared).MetersPerSecondSquared())
	AssertFloatEqual(t, 0.031080950171567256, (1 * FootPerSecondSquared).StandardGravity())

	AssertFloatEqual(t, 100.0, (1 * MeterPerSecondSquared).Gals())
	AssertFloatEqual(t, 3.280839895013123, (1 * MeterPerSecondSquared).FeetPerSecondSquared())
	AssertFloatEqual(t, 0.10197162129779283, (1 * MeterPerSecondSquared).StandardGravity())

	AssertFloatEqual(t, 980.665, (1 * StandardGravity).Gals())
	AssertFloatEqual(t, 32.17404855643044, (1 * StandardGravity).FeetPerSecondSquared())
	AssertFloatEqual(t, 9.80665, (1 * StandardGravity).MetersPerSecondSquared())
}
