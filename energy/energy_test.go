package energy

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestEnergy(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptojoule).Yoctojoules())
	AssertFloatEqual(t, 1e3, (1 * Attojoule).Zeptojoules())
	AssertFloatEqual(t, 1e3, (1 * Femtojoule).Attojoules())
	AssertFloatEqual(t, 1e3, (1 * Picojoule).Femtojoules())
	AssertFloatEqual(t, 1e3, (1 * Nanojoule).Picojoules())
	AssertFloatEqual(t, 1e3, (1 * Microjoule).Nanojoules())
	AssertFloatEqual(t, 1e3, (1 * Millijoule).Microjoules())

	AssertFloatEqual(t, 1e3, (1 * Joule).Millijoules())
	AssertFloatEqual(t, 1e2, (1 * Joule).Centijoules())
	AssertFloatEqual(t, 1e1, (1 * Joule).Decijoules())
	AssertFloatEqual(t, 1e0, (1 * Joule).Joules())
	AssertFloatEqual(t, 1e-1, (1 * Joule).Decajoules())
	AssertFloatEqual(t, 1e-2, (1 * Joule).Hectojoules())
	AssertFloatEqual(t, 1e-3, (1 * Joule).Kilojoules())

	AssertFloatEqual(t, 1e-3, (1 * Kilojoule).Megajoules())
	AssertFloatEqual(t, 1e-3, (1 * Megajoule).Gigajoules())
	AssertFloatEqual(t, 1e-3, (1 * Gigajoule).Terajoules())
	AssertFloatEqual(t, 1e-3, (1 * Terajoule).Petajoules())
	AssertFloatEqual(t, 1e-3, (1 * Petajoule).Exajoules())
	AssertFloatEqual(t, 1e-3, (1 * Exajoule).Zettajoules())
	AssertFloatEqual(t, 1e-3, (1 * Zettajoule).Yottajoules())

	// watt hours
	AssertFloatEqual(t, 1e3, (1 * ZeptowattHour).YoctowattHours())
	AssertFloatEqual(t, 1e3, (1 * AttowattHour).ZeptowattHours())
	AssertFloatEqual(t, 1e3, (1 * FemtowattHour).AttowattHours())
	AssertFloatEqual(t, 1e3, (1 * PicowattHour).FemtowattHours())
	AssertFloatEqual(t, 1e3, (1 * NanowattHour).PicowattHours())
	AssertFloatEqual(t, 1e3, (1 * MicrowattHour).NanowattHours())
	AssertFloatEqual(t, 1e3, (1 * MilliwattHour).MicrowattHours())

	AssertFloatEqual(t, 1e3, (1 * WattHour).MilliwattHours())
	AssertFloatEqual(t, 1e2, (1 * WattHour).CentiwattHours())
	AssertFloatEqual(t, 1e1, (1 * WattHour).DeciwattHours())

	AssertFloatEqual(t, 3600.0, (1 * WattHour).Joules())
	AssertFloatEqual(t, 1e0, (1 * WattHour).WattHours())

	AssertFloatEqual(t, 1e-1, (1 * WattHour).DecawattHours())
	AssertFloatEqual(t, 1e-2, (1 * WattHour).HectowattHours())
	AssertFloatEqual(t, 1e-3, (1 * WattHour).KilowattHours())

	AssertFloatEqual(t, 1e-3, (1 * KilowattHour).MegawattHours())
	AssertFloatEqual(t, 1e-3, (1 * MegawattHour).GigawattHours())
	AssertFloatEqual(t, 1e-3, (1 * GigawattHour).TerawattHours())
	AssertFloatEqual(t, 1e-3, (1 * TerawattHour).PetawattHours())
	AssertFloatEqual(t, 1e-3, (1 * PetawattHour).ExawattHours())
	AssertFloatEqual(t, 1e-3, (1 * ExawattHour).ZettawattHours())
	AssertFloatEqual(t, 1e-3, (1 * ZettawattHour).YottawattHours())

	// calories
	AssertFloatEqual(t, 4.184, (1 * Gramcalorie).Joules())
	AssertFloatEqual(t, 4184.0, (1 * Kilocalorie).Joules())
	AssertFloatEqual(t, 4.184e+6, (1 * Megacalorie).Joules())
}
