package volume

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestVolume_Imperial(t *testing.T) {

	// imperial volume
	AssertFloatEqual(t, 1728.0, (1 * CubicFoot).CubicInches())
	AssertFloatEqual(t, 1e0, (1 * CubicFoot).CubicFeet())
	AssertFloatEqual(t, 0.037037037037037035, (1 * CubicFoot).CubicYards())
	AssertFloatEqual(t, 1.8342646506386176e-10, (1 * CubicYard).CubicMiles())
	AssertFloatEqual(t, 9.391362885602761e-8, (1 * CubicYard).CubicFurlongs())

	// imperial liquid
	AssertFloatEqual(t, 0.21996924829908776, (1 * Litre).ImperialGallons())
	AssertFloatEqual(t, 0.879876993196351, (1 * Litre).ImperialQuarts())
	AssertFloatEqual(t, 1.759753986392702, (1 * Litre).ImperialPints())
	AssertFloatEqual(t, 7.039015945570808, (1 * Litre).ImperialGills())
	AssertFloatEqual(t, 3.519507972785404, (1 * Litre).ImperialCups())
	AssertFloatEqual(t, 35.19507972785404, (1 * Litre).ImperialFluidOunces())
	AssertFloatEqual(t, 281.56063782283235, (1 * Litre).ImperialFluidDrams())
	AssertFloatEqual(t, 0.10998462414954388, (1 * Litre).ImperialPecks())
	AssertFloatEqual(t, 0.02749615603738597, (1 * Litre).ImperialBushels())

	// alternative teaspoons and tablespoons
	AssertFloatEqual(t, 66.66666666666667, (1 * Litre).ImperialTableSpoons())
	AssertFloatEqual(t, 200.00000000000003, (1 * Litre).ImperialTeaSpoons())
}
