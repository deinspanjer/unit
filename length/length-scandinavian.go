package length

// ...
const (
	ScandinavianMile = Metre * 1e4
)

// ScandinavianMiles returns the length in in scandinavian miles (1 scandmile = 10 km)
func (l Length) ScandinavianMiles() float64 {
	return float64(l / ScandinavianMile)
}
