package datarate

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestDatarate(t *testing.T) {

	// base 10 (SI prefixes)
	AssertFloatEqual(t, 1e0, (1 * BitPerSecond).BitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * BitPerSecond).KilobitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * KilobitPerSecond).MegabitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * MegabitPerSecond).GigabitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * GigabitPerSecond).TerabitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * TerabitPerSecond).PetabitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * PetabitPerSecond).ExabitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * ExabitPerSecond).ZettabitsPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * ZettabitPerSecond).YottabitsPerSecond())

	AssertFloatEqual(t, 1e0, (1 * BytePerSecond).BytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * BytePerSecond).KilobytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * KilobytePerSecond).MegabytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * MegabytePerSecond).GigabytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * GigabytePerSecond).TerabytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * TerabytePerSecond).PetabytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * PetabytePerSecond).ExabytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * ExabytePerSecond).ZettabytesPerSecond())
	AssertFloatEqual(t, 1e-3, (1 * ZettabytePerSecond).YottabytesPerSecond())

	// base 2 (IEC prefixes)
	AssertFloatEqual(t, 0.0009765625, (1 * BitPerSecond).KibibitsPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * KibibitPerSecond).MebibitsPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * MebibitPerSecond).GibibitsPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * GibibitPerSecond).TebibitsPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * TebibitPerSecond).PebibitsPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * PebibitPerSecond).ExbibitsPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * ExbibitPerSecond).ZebibitsPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * ZebibitPerSecond).YobibitsPerSecond())

	AssertFloatEqual(t, 0.0009765625, (1 * BytePerSecond).KibibytesPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * KibibytePerSecond).MebibytesPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * MebibytePerSecond).GibibytesPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * GibibytePerSecond).TebibytesPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * TebibytePerSecond).PebibytesPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * PebibytePerSecond).ExbibytesPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * ExbibytePerSecond).ZebibytesPerSecond())
	AssertFloatEqual(t, 0.0009765625, (1 * ZebibytePerSecond).YobibytesPerSecond())
}
