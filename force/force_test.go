package force

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestForce(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Force#Units_of_measurement
	AssertFloatEqual(t, 1e5, (1 * Newton).Dynes())
	AssertFloatEqual(t, 0.10197162129779283, (1 * Newton).KilogramForce())
	AssertFloatEqual(t, 0.22480892365533914, (1 * Newton).PoundForce())
	AssertFloatEqual(t, 7.233011464323171, (1 * Newton).Poundals())

	AssertFloatEqual(t, 1e-5, (1 * Dyne).Newtons())
	AssertFloatEqual(t, 1, Kilopond.KilogramForce())
	AssertFloatEqual(t, 1.0197162129779284e-6, (1 * Dyne).KilogramForce())
	AssertFloatEqual(t, 2.2480892365533917e-6, (1 * Dyne).PoundForce())
	AssertFloatEqual(t, 7.233011464323172e-05, (1 * Dyne).Poundals())

	AssertFloatEqual(t, 9.80665, (1 * KilogramForce).Newtons())
	AssertFloatEqual(t, 980665, (1 * KilogramForce).Dynes())
	AssertFloatEqual(t, 2.204622431164631, (1 * KilogramForce).PoundForce())
	AssertFloatEqual(t, 70.93161187660482, (1 * KilogramForce).Poundals())

	AssertFloatEqual(t, 4.448222, (1 * PoundForce).Newtons())
	AssertFloatEqual(t, 444822.2, (1 * PoundForce).Dynes())
	AssertFloatEqual(t, 0.45359240923251065, (1 * PoundForce).KilogramForce())
	AssertFloatEqual(t, 32.174040721854546, (1 * PoundForce).Poundals())

	AssertFloatEqual(t, 0.138255, (1 * Poundal).Newtons())
	AssertFloatEqual(t, 13825.5, (1 * Poundal).Dynes())
	AssertFloatEqual(t, 0.014098086502526346, (1 * Poundal).KilogramForce())
	AssertFloatEqual(t, 0.03108095773996891, (1 * Poundal).PoundForce())
}
