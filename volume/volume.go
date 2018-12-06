package volume

import . "github.com/deinspanjer/units/unit"

// Volume represents a volume in cubic metres
type Volume Unit

// ...
const (
	// SI
	CubicYoctometre        = CubicMetre * 1e-72
	CubicZeptometre        = CubicMetre * 1e-63
	CubicAttometre         = CubicMetre * 1e-54
	CubicFemtometre        = CubicMetre * 1e-45
	CubicPicometre         = CubicMetre * 1e-36
	CubicNanometre         = CubicMetre * 1e-27
	CubicMicrometre        = CubicMetre * 1e-18
	CubicMillimetre        = CubicMetre * 1e-9
	CubicCentimetre        = CubicMetre * 1e-6
	CubicDecimetre         = CubicMetre * 1e-3
	CubicMetre      Volume = 1e0
	CubicDecametre         = CubicMetre * 1e3
	CubicHectometre        = CubicMetre * 1e6
	CubicKilometre         = CubicMetre * 1e9
	CubicMegametre         = CubicMetre * 1e18
	CubicGigametre         = CubicMetre * 1e27
	CubicTerametre         = CubicMetre * 1e36
	CubicPetametre         = CubicMetre * 1e45
	CubicExametre          = CubicMetre * 1e54
	CubicZettametre        = CubicMetre * 1e63
	CubicYottametre        = CubicMetre * 1e72

	// SI derived
	Yoctolitre = Litre * 1e-24
	Zepolitre  = Litre * 1e-21
	Attolitre  = Litre * 1e-18
	Femtolitre = Litre * 1e-15
	Picolitre  = Litre * 1e-12
	Nanolitre  = Litre * 1e-9
	Microlitre = Litre * 1e-6
	Millilitre = Litre * 1e-3
	Centilitre = Litre * 1e-2
	Decilitre  = Litre * 1e-1
	Litre      = CubicMetre * 1e-3
	Decalitre  = Litre * 1e1
	Hectolitre = Litre * 1e2
	Kilolitre  = Litre * 1e3
	Megalitre  = Litre * 1e6
	Gigalitre  = Litre * 1e9
	Teralitre  = Litre * 1e12
	Petalitre  = Litre * 1e15
	Exalitre   = Litre * 1e18
	Zettalitre = Litre * 1e21
	Yottalitre = Litre * 1e24

	// metric cooking
	MetricTableSpoon = Millilitre * 15
	MetricTeaSpoon   = Millilitre * 5
)

// Yoctolitres returns the volume in yl
func (v Volume) Yoctolitres() float64 {
	return float64(v / Yoctolitre)
}

// Zepolitres returns the volume in zl
func (v Volume) Zepolitres() float64 {
	return float64(v / Zepolitre)
}

// Attolitres returns the volume in al
func (v Volume) Attolitres() float64 {
	return float64(v / Attolitre)
}

// Femtolitres returns the volume in fl
func (v Volume) Femtolitres() float64 {
	return float64(v / Femtolitre)
}

// Picolitres returns the volume in pl
func (v Volume) Picolitres() float64 {
	return float64(v / Picolitre)
}

// Nanolitres returns the volume in nl
func (v Volume) Nanolitres() float64 {
	return float64(v / Nanolitre)
}

// Microlitres returns the volume in µl
func (v Volume) Microlitres() float64 {
	return float64(v / Microlitre)
}

// Millilitres returns the volume in ml
func (v Volume) Millilitres() float64 {
	return float64(v / Millilitre)
}

// Centilitres returns the volume in cl
func (v Volume) Centilitres() float64 {
	return float64(v / Centilitre)
}

// Decilitres returns the volume in dl
func (v Volume) Decilitres() float64 {
	return float64(v / Decilitre)
}

// Litres returns the volume in l
func (v Volume) Litres() float64 {
	return float64(v / Litre)
}

// Decalitres returns the volume in Dl
func (v Volume) Decalitres() float64 {
	return float64(v / Decalitre)
}

// Hectolitres returns the volume in Hl
func (v Volume) Hectolitres() float64 {
	return float64(v / Hectolitre)
}

// Kilolitres returns the volume in Kl
func (v Volume) Kilolitres() float64 {
	return float64(v / Kilolitre)
}

// Megalitres returns the volume in Ml
func (v Volume) Megalitres() float64 {
	return float64(v / Megalitre)
}

// Gigalitres returns the volume in Gl
func (v Volume) Gigalitres() float64 {
	return float64(v / Gigalitre)
}

// Teralitres returns the volume in Tl
func (v Volume) Teralitres() float64 {
	return float64(v / Teralitre)
}

// Petalitres returns the volume in Pl
func (v Volume) Petalitres() float64 {
	return float64(v / Petalitre)
}

// Exalitres returns the volume in El
func (v Volume) Exalitres() float64 {
	return float64(v / Exalitre)
}

// Zettalitres returns the volume in Zl
func (v Volume) Zettalitres() float64 {
	return float64(v / Zettalitre)
}

// Yottalitres returns the volume in Yl
func (v Volume) Yottalitres() float64 {
	return float64(v / Yottalitre)
}

// CubicYoctometres returns the volume in ym³
func (v Volume) CubicYoctometres() float64 {
	return float64(v / CubicYoctometre)
}

// CubicZeptometres returns the volume in zm³
func (v Volume) CubicZeptometres() float64 {
	return float64(v / CubicZeptometre)
}

// CubicAttometres returns the volume in am³
func (v Volume) CubicAttometres() float64 {
	return float64(v / CubicAttometre)
}

// CubicFemtometres returns the volume in fm³
func (v Volume) CubicFemtometres() float64 {
	return float64(v / CubicFemtometre)
}

// CubicPicometres returns the volume in pm³
func (v Volume) CubicPicometres() float64 {
	return float64(v / CubicPicometre)
}

// CubicNanometres returns the volume in nm³
func (v Volume) CubicNanometres() float64 {
	return float64(v / CubicNanometre)
}

// CubicMicrometres returns the volume in µm³
func (v Volume) CubicMicrometres() float64 {
	return float64(v / CubicMicrometre)
}

// CubicMillimetres returns the volume in mm³
func (v Volume) CubicMillimetres() float64 {
	return float64(v / CubicMillimetre)
}

// CubicCentimetres returns the volume in cm³
func (v Volume) CubicCentimetres() float64 {
	return float64(v / CubicCentimetre)
}

// CubicDecimetres returns the volume in dm³
func (v Volume) CubicDecimetres() float64 {
	return float64(v / CubicDecimetre)
}

// CubicMetres returns the volume in m³
func (v Volume) CubicMetres() float64 {
	return float64(v / CubicMetre)
}

// CubicDecametres returns the volume in dam³
func (v Volume) CubicDecametres() float64 {
	return float64(v / CubicDecametre)
}

// CubicHectometres returns the volume in hm³
func (v Volume) CubicHectometres() float64 {
	return float64(v / CubicHectometre)
}

// CubicKilometres returns the volume in km³
func (v Volume) CubicKilometres() float64 {
	return float64(v / CubicKilometre)
}

// CubicMegametres returns the volume in Mm³
func (v Volume) CubicMegametres() float64 {
	return float64(v / CubicMegametre)
}

// CubicGigametres returns the volume in Gm³
func (v Volume) CubicGigametres() float64 {
	return float64(v / CubicGigametre)
}

// CubicTerametres returns the volume in Tm³
func (v Volume) CubicTerametres() float64 {
	return float64(v / CubicTerametre)
}

// CubicPetametres returns the volume in Pm³
func (v Volume) CubicPetametres() float64 {
	return float64(v / CubicPetametre)
}

// CubicExametres returns the volume in Em³
func (v Volume) CubicExametres() float64 {
	return float64(v / CubicExametre)
}

// CubicZettametres returns the volume in Zm³
func (v Volume) CubicZettametres() float64 {
	return float64(v / CubicZettametre)
}

// CubicYottametres returns the volume in Ym³
func (v Volume) CubicYottametres() float64 {
	return float64(v / CubicYottametre)
}

// MetricTableSpoons returns the volume in metric/imperial tablespoons
func (v Volume) MetricTableSpoons() float64 {
	return float64(v / MetricTableSpoon)
}

// MetricTeaSpoons returns the volume in metric/imperial teaspoons
func (v Volume) MetricTeaSpoons() float64 {
	return float64(v / MetricTeaSpoon)
}
