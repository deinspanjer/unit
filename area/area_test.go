package area

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestArea(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e6, (1 * SquareZeptometer).SquareYoctometers())
	AssertFloatEqual(t, 1e6, (1 * SquareAttometer).SquareZeptometers())
	AssertFloatEqual(t, 1e6, (1 * SquareFemtometer).SquareAttometers())
	AssertFloatEqual(t, 1e6, (1 * SquarePicometer).SquareFemtometers())
	AssertFloatEqual(t, 1e6, (1 * SquareNanometer).SquarePicometers())
	AssertFloatEqual(t, 1e6, (1 * SquareMicrometer).SquareNanometers())
	AssertFloatEqual(t, 1e6, (1 * SquareMillimeter).SquareMicrometers())
	AssertFloatEqual(t, 1e6, (1 * SquareMeter).SquareMillimeters())

	AssertFloatEqual(t, 1e4, (1 * SquareMeter).SquareCentimeters())
	AssertFloatEqual(t, 1e2, (1 * SquareMeter).SquareDecimeters())
	AssertFloatEqual(t, 1e0, (1 * SquareMeter).SquareMeters())
	AssertFloatEqual(t, 1e-2, (1 * SquareMeter).SquareDecameter())
	AssertFloatEqual(t, 1e-4, (1 * SquareMeter).SquareHectometer())

	AssertFloatEqual(t, 1e-6, (1 * SquareMeter).SquareKilometers())
	AssertFloatEqual(t, 1e-6, (1 * SquareKilometer).SquareMegameters())
	AssertFloatEqual(t, 1e-6, (1 * SquareMegameter).SquareGigameters())
	AssertFloatEqual(t, 1e-6, (1 * SquareGigameter).SquareTerameters())
	AssertFloatEqual(t, 1e-6, (1 * SquareTerameter).SquarePetameters())
	AssertFloatEqual(t, 1e-6, (1 * SquarePetameter).SquareExameters())
	AssertFloatEqual(t, 1e-6, (1 * SquareExameter).SquareZettameters())
	AssertFloatEqual(t, 1e-6, (1 * SquareZettameter).SquareYottameters())

	// US
	AssertFloatEqual(t, 10.763910416709724, (1 * SquareMeter).SquareFeet())
	AssertFloatEqual(t, 144.0, (1 * SquareFoot).SquareInches())
	AssertFloatEqual(t, 9.0, (1 * SquareYard).SquareFeet())
	AssertFloatEqual(t, 4840.0, (1 * Acre).SquareYards())
	AssertFloatEqual(t, 640.0, (1 * SquareMile).Acres())
	AssertFloatEqual(t, 0.0015625, (1 * Acre).SquareMiles())

	// imperial
	AssertFloatEqual(t, 25.29285264, (1 * SquareRod).SquareMeters())
	AssertFloatEqual(t, 1011.7141055999998, (1 * Rood).SquareMeters())
	AssertFloatEqual(t, 0.025, (1 * SquareRod).Roods())

	// aliases
	AssertFloatEqual(t, 1.0, (1 * SquareMeter).Centiares())
	AssertFloatEqual(t, 1.0, (1 * SquareDecameter).Ares())
	AssertFloatEqual(t, 1.0, (1 * SquareHectometer).Hectares())
	AssertFloatEqual(t, 1.0, (1 * SquarePerch).SquareRods())
}
