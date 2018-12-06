package volume

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestVolume(t *testing.T) {

	// SI derived
	AssertFloatEqual(t, 1e3, (1 * Zepolitre).Yoctolitres())
	AssertFloatEqual(t, 1e3, (1 * Attolitre).Zepolitres())
	AssertFloatEqual(t, 1e3, (1 * Femtolitre).Attolitres())
	AssertFloatEqual(t, 1e3, (1 * Picolitre).Femtolitres())
	AssertFloatEqual(t, 1e3, (1 * Nanolitre).Picolitres())
	AssertFloatEqual(t, 1e3, (1 * Microlitre).Nanolitres())
	AssertFloatEqual(t, 1e3, (1 * Millilitre).Microlitres())

	AssertFloatEqual(t, 1e3, (1 * Litre).Millilitres())
	AssertFloatEqual(t, 1e2, (1 * Litre).Centilitres())
	AssertFloatEqual(t, 1e1, (1 * Litre).Decilitres())
	AssertFloatEqual(t, 1e0, (1 * Litre).Litres())
	AssertFloatEqual(t, 1e-1, (1 * Litre).Decalitres())
	AssertFloatEqual(t, 1e-2, (1 * Litre).Hectolitres())
	AssertFloatEqual(t, 1e-3, (1 * Litre).Kilolitres())

	AssertFloatEqual(t, 1e-3, (1 * Kilolitre).Megalitres())
	AssertFloatEqual(t, 1e-3, (1 * Megalitre).Gigalitres())
	AssertFloatEqual(t, 1e-3, (1 * Gigalitre).Teralitres())
	AssertFloatEqual(t, 1e-3, (1 * Teralitre).Petalitres())
	AssertFloatEqual(t, 1e-3, (1 * Petalitre).Exalitres())
	AssertFloatEqual(t, 1e-3, (1 * Exalitre).Zettalitres())
	AssertFloatEqual(t, 1e-3, (1 * Zettalitre).Yottalitres())

	// SI
	AssertFloatEqual(t, 1e9, (1 * CubicZeptometre).CubicYoctometres())
	AssertFloatEqual(t, 1e9, (1 * CubicAttometre).CubicZeptometres())
	AssertFloatEqual(t, 1e9, (1 * CubicFemtometre).CubicAttometres())
	AssertFloatEqual(t, 1e9, (1 * CubicPicometre).CubicFemtometres())
	AssertFloatEqual(t, 1e9, (1 * CubicNanometre).CubicPicometres())
	AssertFloatEqual(t, 1e9, (1 * CubicMicrometre).CubicNanometres())
	AssertFloatEqual(t, 1e9, (1 * CubicMillimetre).CubicMicrometres())

	AssertFloatEqual(t, 1e9, (1 * CubicMetre).CubicMillimetres())
	AssertFloatEqual(t, 1e6, (1 * CubicMetre).CubicCentimetres())
	AssertFloatEqual(t, 1e3, (1 * CubicMetre).CubicDecimetres())

	AssertFloatEqual(t, 1e-3, (1 * Litre).CubicMetres())
	AssertFloatEqual(t, 1e0, (1 * CubicMetre).CubicMetres())

	AssertFloatEqual(t, 1e-3, (1 * CubicMetre).CubicDecametres())
	AssertFloatEqual(t, 1e-6, (1 * CubicMetre).CubicHectometres())
	AssertFloatEqual(t, 1e-9, (1 * CubicMetre).CubicKilometres())

	AssertFloatEqual(t, 1e-9, (1 * CubicKilometre).CubicMegametres())
	AssertFloatEqual(t, 1e-9, (1 * CubicMegametre).CubicGigametres())
	AssertFloatEqual(t, 1e-9, (1 * CubicGigametre).CubicTerametres())
	AssertFloatEqual(t, 1e-9, (1 * CubicTerametre).CubicPetametres())
	AssertFloatEqual(t, 1e-9, (1 * CubicPetametre).CubicExametres())
	AssertFloatEqual(t, 1e-9, (1 * CubicExametre).CubicZettametres())
	AssertFloatEqual(t, 1e-9, (1 * CubicZettametre).CubicYottametres())

	AssertFloatEqual(t, 66.66666666666667, (1 * Litre).MetricTableSpoons())
	AssertFloatEqual(t, 200.00000000000003, (1 * Litre).MetricTeaSpoons())
}
