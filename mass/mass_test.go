package mass

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestMass(t *testing.T) {

	// SI
	AssertFloatEqual(t, 1e3, (1 * Zeptogram).Yoctograms())
	AssertFloatEqual(t, 1e3, (1 * Attogram).Zeptograms())
	AssertFloatEqual(t, 1e3, (1 * Femtogram).Attograms())
	AssertFloatEqual(t, 1e3, (1 * Picogram).Femtograms())
	AssertFloatEqual(t, 1e3, (1 * Nanogram).Picograms())
	AssertFloatEqual(t, 1e3, (1 * Microgram).Nanograms())
	AssertFloatEqual(t, 1e3, (1 * Milligram).Micrograms())

	AssertFloatEqual(t, 1e3, (1 * Gram).Milligrams())
	AssertFloatEqual(t, 1e2, (1 * Gram).Centigrams())
	AssertFloatEqual(t, 1e1, (1 * Gram).Decigrams())
	AssertFloatEqual(t, 1e0, (1 * Gram).Grams())
	AssertFloatEqual(t, 1e-1, (1 * Gram).Decagrams())
	AssertFloatEqual(t, 1e-2, (1 * Gram).Hectograms())
	AssertFloatEqual(t, 1e-3, (1 * Gram).Kilograms())

	AssertFloatEqual(t, 1e-3, (1 * Kilogram).Megagrams())
	AssertFloatEqual(t, 1e-3, (1 * Megagram).Gigagrams())
	AssertFloatEqual(t, 1e-3, (1 * Gigagram).Teragrams())
	AssertFloatEqual(t, 1e-3, (1 * Teragram).Petagrams())
	AssertFloatEqual(t, 1e-3, (1 * Petagram).Exagrams())
	AssertFloatEqual(t, 1e-3, (1 * Exagram).Zettagrams())
	AssertFloatEqual(t, 1e-3, (1 * Zettagram).Yottagrams())

	// non-SI
	AssertFloatEqual(t, 1e-3, (1 * Kilogram).Tonnes())
	AssertFloatEqual(t, 1e-3, (1 * Tonne).Kilotonnes())
	AssertFloatEqual(t, 1e-3, (1 * Kilotonne).Megatonnes())
	AssertFloatEqual(t, 1e-3, (1 * Megatonne).Gigatonnes())
	AssertFloatEqual(t, 1e-3, (1 * Gigatonne).Teratonnes())
	AssertFloatEqual(t, 1e-3, (1 * Teratonne).Petatonnes())
	AssertFloatEqual(t, 1e-3, (1 * Petatonne).Exatonnes())

	// avoirdupois
	AssertFloatEqual(t, 0.015432358352941428, (1 * Milligram).TroyGrains())
	AssertFloatEqual(t, 0.0022857142857142855, (1 * TroyGrain).AvoirdupoisOunces())
	AssertFloatEqual(t, 0.03657142857142857, (1 * TroyGrain).AvoirdupoisDrams())
	AssertFloatEqual(t, 0.00014285714285714284, (1 * TroyGrain).AvoirdupoisPounds())

	AssertFloatEqual(t, 0.07142857142857142, (1 * AvoirdupoisPound).UsStones())
	AssertFloatEqual(t, 0.058775510204081644, (1 * TroyPound).UkStones())

	// https://en.wikipedia.org/wiki/Quarter_(unit)#Weight
	AssertFloatEqual(t, 0.08818490487395102, (1 * Kilogram).UsQuarters())
	AssertFloatEqual(t, 0.07873652220888486, (1 * Kilogram).UkQuarters())

	// https://en.wikipedia.org/wiki/Hundredweight
	AssertFloatEqual(t, 50.80234544, (1 * LongHundredweight).Kilograms())
	AssertFloatEqual(t, 45.359237, (1 * ShortHundredweight).Kilograms())
	AssertFloatEqual(t, 1e0, (112 * AvoirdupoisPound).LongHundredweights())
	AssertFloatEqual(t, 1e0, (100 * AvoirdupoisPound).ShortHundredweights())

	AssertFloatEqual(t, 1, CentalHundredweight.ShortHundredweights())
	AssertFloatEqual(t, 1, ImperialHundredweight.LongHundredweights())

	// troy
	AssertFloatEqual(t, 0.0020833333333333333, (1 * TroyGrain).TroyOunces())
	AssertFloatEqual(t, 0.0026792288807189982, (1 * Gram).TroyPounds())
}
