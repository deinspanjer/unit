package datasize

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestDatasize(t *testing.T) {

	// base 10 (SI prefixes)
	AssertFloatEqual(t, 1e0, (1 * Bit).Bits())
	AssertFloatEqual(t, 1e-3, (1 * Bit).Kilobits())
	AssertFloatEqual(t, 1e-3, (1 * Kilobit).Megabits())
	AssertFloatEqual(t, 1e-3, (1 * Megabit).Gigabits())
	AssertFloatEqual(t, 1e-3, (1 * Gigabit).Terabits())
	AssertFloatEqual(t, 1e-3, (1 * Terabit).Petabits())
	AssertFloatEqual(t, 1e-3, (1 * Petabit).Exabits())
	AssertFloatEqual(t, 1e-3, (1 * Exabit).Zettabits())
	AssertFloatEqual(t, 1e-3, (1 * Zettabit).Yottabits())

	AssertFloatEqual(t, 0.125, (1 * Bit).Bytes())
	AssertFloatEqual(t, 1e0, (1 * Byte).Bytes())

	AssertFloatEqual(t, 1e-3, (1 * Byte).Kilobytes())
	AssertFloatEqual(t, 1e-3, (1 * Kilobyte).Megabytes())
	AssertFloatEqual(t, 1e-3, (1 * Megabyte).Gigabytes())
	AssertFloatEqual(t, 1e-3, (1 * Gigabyte).Terabytes())
	AssertFloatEqual(t, 1e-3, (1 * Terabyte).Petabytes())
	AssertFloatEqual(t, 1e-3, (1 * Petabyte).Exabytes())
	AssertFloatEqual(t, 1e-3, (1 * Exabyte).Zettabytes())
	AssertFloatEqual(t, 1e-3, (1 * Zettabyte).Yottabytes())

	// base 2 (IEC prefixes)
	AssertFloatEqual(t, 0.0009765625, (1 * Bit).Kibibits())
	AssertFloatEqual(t, 0.0009765625, (1 * Kibibit).Mebibits())
	AssertFloatEqual(t, 0.0009765625, (1 * Mebibit).Gibibits())
	AssertFloatEqual(t, 0.0009765625, (1 * Gibibit).Tebibits())
	AssertFloatEqual(t, 0.0009765625, (1 * Tebibit).Pebibits())
	AssertFloatEqual(t, 0.0009765625, (1 * Pebibit).Exbibits())
	AssertFloatEqual(t, 0.0009765625, (1 * Exbibit).Zebibits())
	AssertFloatEqual(t, 0.0009765625, (1 * Zebibit).Yobibits())

	AssertFloatEqual(t, 0.0009765625, (1 * Byte).Kibibytes())
	AssertFloatEqual(t, 0.0009765625, (1 * Kibibyte).Mebibytes())
	AssertFloatEqual(t, 0.0009765625, (1 * Mebibyte).Gibibytes())
	AssertFloatEqual(t, 0.0009765625, (1 * Gibibyte).Tebibytes())
	AssertFloatEqual(t, 0.0009765625, (1 * Tebibyte).Pebibytes())
	AssertFloatEqual(t, 0.0009765625, (1 * Pebibyte).Exbibytes())
	AssertFloatEqual(t, 0.0009765625, (1 * Exbibyte).Zebibytes())
	AssertFloatEqual(t, 0.0009765625, (1 * Zebibyte).Yobibytes())
}
