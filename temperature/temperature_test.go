package temperature

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestTemperature(t *testing.T) {
	AssertFloatEqual(t, 1, Kelvin.Kelvin())

	AssertFloatEqual(t, 373.15, FromCelsius(100).Kelvin())
	AssertFloatEqual(t, 310.9277777777778, FromFahrenheit(100).Kelvin())
	AssertFloatEqual(t, 100.0, FromKelvin(100).Kelvin())
	AssertFloatEqual(t, 576.180303030303, FromNewton(100).Kelvin())
	AssertFloatEqual(t, 306.4833333333333, FromDelisle(100).Kelvin())
	AssertFloatEqual(t, 398.15, FromReaumur(100).Kelvin())
	AssertFloatEqual(t, 449.3404761904762, FromRomer(100).Kelvin())
	AssertFloatEqual(t, 55.555555555555515, FromRankine(100).Kelvin())

	AssertFloatEqual(t, -173.15, FromKelvin(100).Celsius())
	AssertFloatEqual(t, -279.67, FromKelvin(100).Fahrenheit())
	AssertFloatEqual(t, -57.1395, FromKelvin(100).Newton())
	AssertFloatEqual(t, 409.725, FromKelvin(100).Delisle())
	AssertFloatEqual(t, -83.40375, FromKelvin(100).Romer())
	AssertFloatEqual(t, -138.52, FromKelvin(100).Reaumur())
	AssertFloatEqual(t, 180, FromKelvin(100).Rankine())
}
