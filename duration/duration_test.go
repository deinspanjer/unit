package duration

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestDuration(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptosecond).Yoctoseconds())
	AssertFloatEqual(t, 1e3, (1 * Attosecond).Zeptoseconds())
	AssertFloatEqual(t, 1e3, (1 * Femtosecond).Attoseconds())
	AssertFloatEqual(t, 1e3, (1 * Picosecond).Femtoseconds())
	AssertFloatEqual(t, 1e3, (1 * Nanosecond).Picoseconds())
	AssertFloatEqual(t, 1e3, (1 * Microsecond).Nanoseconds())
	AssertFloatEqual(t, 1e3, (1 * Millisecond).Microseconds())

	AssertFloatEqual(t, 1e3, (1 * Second).Milliseconds())
	AssertFloatEqual(t, 1e2, (1 * Second).Centiseconds())
	AssertFloatEqual(t, 1e1, (1 * Second).Deciseconds())
	AssertFloatEqual(t, 1e0, (1 * Second).Seconds())
	AssertFloatEqual(t, 1e-1, (1 * Second).Decaseconds())
	AssertFloatEqual(t, 1e-2, (1 * Second).Hectoseconds())
	AssertFloatEqual(t, 1e-3, (1 * Second).Kiloseconds())

	AssertFloatEqual(t, 1e-3, (1 * Kilosecond).Megaseconds())
	AssertFloatEqual(t, 1e-3, (1 * Megasecond).Gigaseconds())
	AssertFloatEqual(t, 1e-3, (1 * Gigasecond).Teraseconds())
	AssertFloatEqual(t, 1e-3, (1 * Terasecond).Petaseconds())
	AssertFloatEqual(t, 1e-3, (1 * Petasecond).Exaseconds())
	AssertFloatEqual(t, 1e-3, (1 * Exasecond).Zettaseconds())
	AssertFloatEqual(t, 1e-3, (1 * Zettasecond).Yottaseconds())

	// non-SI
	AssertFloatEqual(t, 0.016666666666666666, (1 * Second).Minutes())
	AssertFloatEqual(t, 0.0002777777777777778, (1 * Second).Hours())
	AssertFloatEqual(t, 1.1574074074074073e-5, (1 * Second).Days())
	AssertFloatEqual(t, 1.6534391534391535e-6, (1 * Second).Weeks())
	AssertFloatEqual(t, 3.8580246913580245e-7, (1 * Second).ThirtyDayMonths())
	AssertFloatEqual(t, 3.168808781402895e-8, (1 * Second).JulianYears())
}
