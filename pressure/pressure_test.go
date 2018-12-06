package pressure

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestPressure(t *testing.T) {

	// SI derived
	AssertFloatEqual(t, 1e3, (1 * Zeptopascal).Yoctopascals())
	AssertFloatEqual(t, 1e3, (1 * Attopascal).Zeptopascals())
	AssertFloatEqual(t, 1e3, (1 * Femtopascal).Attopascals())
	AssertFloatEqual(t, 1e3, (1 * Picopascal).Femtopascals())
	AssertFloatEqual(t, 1e3, (1 * Nanopascal).Picopascals())
	AssertFloatEqual(t, 1e3, (1 * Micropascal).Nanopascals())
	AssertFloatEqual(t, 1e3, (1 * Millipascal).Micropascals())

	AssertFloatEqual(t, 1e3, (1 * Pascal).Millipascals())
	AssertFloatEqual(t, 1e2, (1 * Pascal).Centipascals())
	AssertFloatEqual(t, 1e1, (1 * Pascal).Decipascals())
	AssertFloatEqual(t, 1e0, (1 * Pascal).Pascals())
	AssertFloatEqual(t, 1e-1, (1 * Pascal).Decapascals())
	AssertFloatEqual(t, 1e-2, (1 * Pascal).Hectopascals())
	AssertFloatEqual(t, 1e-3, (1 * Pascal).Kilopascals())

	AssertFloatEqual(t, 1e-3, (1 * Kilopascal).Megapascals())
	AssertFloatEqual(t, 1e-3, (1 * Megapascal).Gigapascals())
	AssertFloatEqual(t, 1e-3, (1 * Gigapascal).Terapascals())
	AssertFloatEqual(t, 1e-3, (1 * Terapascal).Petapascals())
	AssertFloatEqual(t, 1e-3, (1 * Petapascal).Exapascals())
	AssertFloatEqual(t, 1e-3, (1 * Exapascal).Zettapascals())
	AssertFloatEqual(t, 1e-3, (1 * Zettapascal).Yottapascals())

	// non-SI
	AssertFloatEqual(t, 1e3, (1 * Zeptobar).Yoctobars())
	AssertFloatEqual(t, 1e3, (1 * Attobar).Zeptobars())
	AssertFloatEqual(t, 1e3, (1 * Femtobar).Attobars())
	AssertFloatEqual(t, 1e3, (1 * Picobar).Femtobars())
	AssertFloatEqual(t, 1e3, (1 * Nanobar).Picobars())
	AssertFloatEqual(t, 1e3, (1 * Microbar).Nanobars())
	AssertFloatEqual(t, 1e3, (1 * Millibar).Microbars())

	AssertFloatEqual(t, 1e3, (1 * Bar).Millibars())
	AssertFloatEqual(t, 1e2, (1 * Bar).Centibars())
	AssertFloatEqual(t, 1e1, (1 * Bar).Decibars())
	AssertFloatEqual(t, 1e0, (1 * Bar).Bars())
	AssertFloatEqual(t, 1e-1, (1 * Bar).Decabars())
	AssertFloatEqual(t, 1e-2, (1 * Bar).Hectobars())
	AssertFloatEqual(t, 1e-3, (1 * Bar).Kilobars())

	AssertFloatEqual(t, 1e-3, (1 * Kilobar).Megabars())
	AssertFloatEqual(t, 1e-3, (1 * Megabar).Gigabars())
	AssertFloatEqual(t, 1e-3, (1 * Gigabar).Terabars())
	AssertFloatEqual(t, 1e-3, (1 * Terabar).Petabars())
	AssertFloatEqual(t, 1e-3, (1 * Petabar).Exabars())
	AssertFloatEqual(t, 1e-3, (1 * Exabar).Zettabars())
	AssertFloatEqual(t, 1e-3, (1 * Zettabar).Yottabars())

	// misc
	AssertFloatEqual(t, 0.09869232667160129, (10000 * Pascal).Atmospheres())
	AssertFloatEqual(t, 0.10197162129779283, (10000 * Pascal).TechAtmospheres())
	AssertFloatEqual(t, 75.00615050434136, (10000 * Pascal).Torrs())
	AssertFloatEqual(t, 1.4503683935719673, (10000 * Pascal).PoundsPerSquareInch())
}
