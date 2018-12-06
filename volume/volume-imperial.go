package volume

// ...
const (
	// imperial liquid
	ImperialGallon     = Litre * 4.54609
	ImperialQuart      = ImperialGallon / 4
	ImperialPint       = ImperialQuart / 2
	ImperialCup        = ImperialPint / 2
	ImperialGill       = ImperialPint / 4
	ImperialFluidOunce = ImperialGill / 5
	ImperialFluidDram  = ImperialFluidOunce / 8
	ImperialPeck       = ImperialGallon * 2
	ImperialBushel     = ImperialPeck * 4

	// imperial volume
	CubicInch    = Litre * 0.016387064
	CubicFoot    = CubicInch * 1728
	CubicYard    = CubicFoot * 27
	CubicMile    = CubicYard * 5451776000
	CubicFurlong = CubicMile * 0.00195314

	// aliases
	ImperialTableSpoon = MetricTableSpoon
	ImperialTeaSpoon   = MetricTeaSpoon
)

// CubicInches returns the volume in in³
func (v Volume) CubicInches() float64 {
	return float64(v / CubicInch)
}

// CubicFeet returns the volume in ft³
func (v Volume) CubicFeet() float64 {
	return float64(v / CubicFoot)
}

// CubicYards returns the volume in yd³
func (v Volume) CubicYards() float64 {
	return float64(v / CubicYard)
}

// CubicMiles returns the volume in mi³
func (v Volume) CubicMiles() float64 {
	return float64(v / CubicMile)
}

// CubicFurlongs returns the volume in furlong³
func (v Volume) CubicFurlongs() float64 {
	return float64(v / CubicFurlong)
}

// ImperialGallons returns the volume in imperial gallons
func (v Volume) ImperialGallons() float64 {
	return float64(v / ImperialGallon)
}

// ImperialQuarts returns the volume in imperial quarts
func (v Volume) ImperialQuarts() float64 {
	return float64(v / ImperialQuart)
}

// ImperialPints returns the volume in imperial pints
func (v Volume) ImperialPints() float64 {
	return float64(v / ImperialPint)
}

// ImperialGills returns the volume in imperial gills
func (v Volume) ImperialGills() float64 {
	return float64(v / ImperialGill)
}

// ImperialCups returns the volume in imperial cups
func (v Volume) ImperialCups() float64 {
	return float64(v / ImperialCup)
}

// ImperialFluidOunces returns the volume in imperial fluid ounces
func (v Volume) ImperialFluidOunces() float64 {
	return float64(v / ImperialFluidOunce)
}

// ImperialFluidDrams returns the volume in imperial fluid drams
func (v Volume) ImperialFluidDrams() float64 {
	return float64(v / ImperialFluidDram)
}

// ImperialPecks returns the volume in imperial pecks
func (v Volume) ImperialPecks() float64 {
	return float64(v / ImperialPeck)
}

// ImperialBushels returns the volume in imperial bushels
func (v Volume) ImperialBushels() float64 {
	return float64(v / ImperialBushel)
}

// ImperialTableSpoons returns the volume in metric/imperial tablespoons
func (v Volume) ImperialTableSpoons() float64 {
	return float64(v / ImperialTableSpoon)
}

// ImperialTeaSpoons returns the volume in metric/imperial teaspoons
func (v Volume) ImperialTeaSpoons() float64 {
	return float64(v / ImperialTeaSpoon)
}
