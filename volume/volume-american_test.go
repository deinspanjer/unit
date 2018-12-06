package volume

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestVolume_American(t *testing.T) {

	// SI derived
	AssertFloatEqual(t, 1e3, (1 * Zepoliter).Yoctoliters())
	AssertFloatEqual(t, 1e3, (1 * Attoliter).Zepoliters())
	AssertFloatEqual(t, 1e3, (1 * Femtoliter).Attoliters())
	AssertFloatEqual(t, 1e3, (1 * Picoliter).Femtoliters())
	AssertFloatEqual(t, 1e3, (1 * Nanoliter).Picoliters())
	AssertFloatEqual(t, 1e3, (1 * Microliter).Nanoliters())
	AssertFloatEqual(t, 1e3, (1 * Milliliter).Microliters())

	AssertFloatEqual(t, 1e3, (1 * Liter).Milliliters())
	AssertFloatEqual(t, 1e2, (1 * Liter).Centiliters())
	AssertFloatEqual(t, 1e1, (1 * Liter).Deciliters())
	AssertFloatEqual(t, 1e0, (1 * Liter).Liters())
	AssertFloatEqual(t, 1e-1, (1 * Liter).Decaliters())
	AssertFloatEqual(t, 1e-2, (1 * Liter).Hectoliters())
	AssertFloatEqual(t, 1e-3, (1 * Liter).Kiloliters())

	AssertFloatEqual(t, 1e-3, (1 * Kiloliter).Megaliters())
	AssertFloatEqual(t, 1e-3, (1 * Megaliter).Gigaliters())
	AssertFloatEqual(t, 1e-3, (1 * Gigaliter).Teraliters())
	AssertFloatEqual(t, 1e-3, (1 * Teraliter).Petaliters())
	AssertFloatEqual(t, 1e-3, (1 * Petaliter).Exaliters())
	AssertFloatEqual(t, 1e-3, (1 * Exaliter).Zettaliters())
	AssertFloatEqual(t, 1e-3, (1 * Zettaliter).Yottaliters())

	// SI
	AssertFloatEqual(t, 1e9, (1 * CubicZeptometer).CubicYoctometers())
	AssertFloatEqual(t, 1e9, (1 * CubicAttometer).CubicZeptometers())
	AssertFloatEqual(t, 1e9, (1 * CubicFemtometer).CubicAttometers())
	AssertFloatEqual(t, 1e9, (1 * CubicPicometer).CubicFemtometers())
	AssertFloatEqual(t, 1e9, (1 * CubicNanometer).CubicPicometers())
	AssertFloatEqual(t, 1e9, (1 * CubicMicrometer).CubicNanometers())
	AssertFloatEqual(t, 1e9, (1 * CubicMillimeter).CubicMicrometers())

	AssertFloatEqual(t, 1e9, (1 * CubicMeter).CubicMillimeters())
	AssertFloatEqual(t, 1e6, (1 * CubicMeter).CubicCentimeters())
	AssertFloatEqual(t, 1e3, (1 * CubicMeter).CubicDecimeters())

	AssertFloatEqual(t, 1e-3, (1 * Liter).CubicMeters())
	AssertFloatEqual(t, 1e0, (1 * CubicMeter).CubicMeters())

	AssertFloatEqual(t, 1e-3, (1 * CubicMeter).CubicDecameters())
	AssertFloatEqual(t, 1e-6, (1 * CubicMeter).CubicHectometers())
	AssertFloatEqual(t, 1e-9, (1 * CubicMeter).CubicKilometers())

	AssertFloatEqual(t, 1e-9, (1 * CubicKilometer).CubicMegameters())
	AssertFloatEqual(t, 1e-9, (1 * CubicMegameter).CubicGigameters())
	AssertFloatEqual(t, 1e-9, (1 * CubicGigameter).CubicTerameters())
	AssertFloatEqual(t, 1e-9, (1 * CubicTerameter).CubicPetameters())
	AssertFloatEqual(t, 1e-9, (1 * CubicPetameter).CubicExameters())
	AssertFloatEqual(t, 1e-9, (1 * CubicExameter).CubicZettameters())
	AssertFloatEqual(t, 1e-9, (1 * CubicZettameter).CubicYottameters())

	// US liquid
	AssertFloatEqual(t, 0.26417205235814845, (1 * Liter).USLiquidGallons())
	AssertFloatEqual(t, 1.0566882094325938, (1 * Liter).USLiquidQuarts())
	AssertFloatEqual(t, 2.1133764188651876, (1 * Liter).USLiquidPints())
	AssertFloatEqual(t, 4.226752837730375, (1 * Liter).USCups())
	AssertFloatEqual(t, 4.166666666666667, (1 * Liter).USLegalCups())
	AssertFloatEqual(t, 8.45350567546075, (1 * Liter).USGills())
	AssertFloatEqual(t, 67.628045403686, (1 * Liter).USTableSpoons())
	AssertFloatEqual(t, 202.884136211058, (1 * Liter).USTeaSpoons())
	AssertFloatEqual(t, 270.512181614744, (1 * Liter).USFluidDrams())
	AssertFloatEqual(t, 33.814022701843, (1 * Liter).USFluidOunces())

	// US dry
	AssertFloatEqual(t, 0.9080829842688559, (1 * Liter).USDryQuarts())
	AssertFloatEqual(t, 0.028377593258401747, (1 * Liter).USBushels())
	AssertFloatEqual(t, 0.11351037303360699, (1 * Liter).USPecks())
	AssertFloatEqual(t, 0.22702074606721398, (1 * Liter).USDryGallons())
	AssertFloatEqual(t, 1.8161659685377118, (1 * Liter).USDryPints())
}
