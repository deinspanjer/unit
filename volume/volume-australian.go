package volume

// ...
const (
	AustralianTableSpoon = Millilitre * 20
)

// AustralianTableSpoons returns the volume in Australian tablespoons
func (v Volume) AustralianTableSpoons() float64 {
	return float64(v / AustralianTableSpoon)
}
