package speed

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestSpeed(t *testing.T) {

	// SI
	AssertFloatEqual(t, 10.0, (10 * MetersPerSecond).MetersPerSecond())
	AssertFloatEqual(t, 35.99997120002304, (10 * MetersPerSecond).KilometersPerHour())

	// US
	AssertFloatEqual(t, 32.808398950131235, (10 * MetersPerSecond).FeetPerSecond())
	AssertFloatEqual(t, 22.369362920544024, (10 * MetersPerSecond).MilesPerHour())

	// misc
	AssertFloatEqual(t, 19.438461717893492, (10 * MetersPerSecond).Knots())

	// space
	AssertFloatEqual(t, 0.03335640951981521, (10000000 * MetersPerSecond).SpeedOfLight())
}
