package length

import . "github.com/deinspanjer/units/unit"

// Length represents a SI unit of length (in metres, m)
type Length Unit

// ...
const (
	// SI
	Yoctometre        = Metre * 1e-24
	Zeptometre        = Metre * 1e-21
	Attometre         = Metre * 1e-18
	Femtometre        = Metre * 1e-15
	Picometre         = Metre * 1e-12
	Nanometre         = Metre * 1e-9
	Micrometre        = Metre * 1e-6
	Millimetre        = Metre * 1e-3
	Centimetre        = Metre * 1e-2
	Decimetre         = Metre * 1e-1
	Metre      Length = 1e0
	Decametre         = Metre * 1e1
	Hectometre        = Metre * 1e2
	Kilometre         = Metre * 1e3
	Megametre         = Metre * 1e6
	Gigametre         = Metre * 1e9
	Terametre         = Metre * 1e12
	Petametre         = Metre * 1e15
	Exametre          = Metre * 1e18
	Zettametre        = Metre * 1e21
	Yottametre        = Metre * 1e24

	// space
	LunarDistance    = Kilometre * 384400
	AstronomicalUnit = Metre * 149597870700
	LightYear        = Metre * 9460730472580800
)

// Yoctometres returns the length in ym
func (l Length) Yoctometres() float64 {
	return float64(l / Yoctometre)
}

// Zeptometres returns the length in zm
func (l Length) Zeptometres() float64 {
	return float64(l / Zeptometre)
}

// Attometres returns the length in am
func (l Length) Attometres() float64 {
	return float64(l / Attometre)
}

// Femtometres returns the length in fm
func (l Length) Femtometres() float64 {
	return float64(l / Femtometre)
}

// Picometres returns the length in pm
func (l Length) Picometres() float64 {
	return float64(l / Picometre)
}

// Nanometres returns the length in nm
func (l Length) Nanometres() float64 {
	return float64(l / Nanometre)
}

// Micrometres returns the length in microms
func (l Length) Micrometres() float64 {
	return float64(l / Micrometre)
}

// Millimetres returns the length in mm
func (l Length) Millimetres() float64 {
	return float64(l / Millimetre)
}

// Centimetres returns the length in cm
func (l Length) Centimetres() float64 {
	return float64(l / Centimetre)
}

// Decimetres returns the length in dm
func (l Length) Decimetres() float64 {
	return float64(l / Decimetre)
}

// Metres returns the length in m
func (l Length) Metres() float64 {
	return float64(l)
}

// Decametres returns the length in dam
func (l Length) Decametres() float64 {
	return float64(l / Decametre)
}

// Hectometres returns the length in hm
func (l Length) Hectometres() float64 {
	return float64(l / Hectometre)
}

// Kilometres returns the length in km
func (l Length) Kilometres() float64 {
	return float64(l / Kilometre)
}

// Megametres returns the length in in Mm
func (l Length) Megametres() float64 {
	return float64(l / Megametre)
}

// Gigametres returns the length in in Gm
func (l Length) Gigametres() float64 {
	return float64(l / Gigametre)
}

// Terametres returns the length in in Tm
func (l Length) Terametres() float64 {
	return float64(l / Terametre)
}

// Petametres returns the length in in Pm
func (l Length) Petametres() float64 {
	return float64(l / Petametre)
}

// Exametres returns the length in in Em
func (l Length) Exametres() float64 {
	return float64(l / Exametre)
}

// Zettametres returns the length in in Zm
func (l Length) Zettametres() float64 {
	return float64(l / Zettametre)
}

// Yottametres returns the length in in Ym
func (l Length) Yottametres() float64 {
	return float64(l / Yottametre)
}

// LunarDistances returns the length in ld
func (l Length) LunarDistances() float64 {
	return float64(l / LunarDistance)
}

// AstronomicalUnits returns the length in au
func (l Length) AstronomicalUnits() float64 {
	return float64(l / AstronomicalUnit)
}

// LightYears returns the length in ly
func (l Length) LightYears() float64 {
	return float64(l / LightYear)
}
