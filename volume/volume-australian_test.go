package volume

import (
	. "github.com/deinspanjer/units/testing"
	"testing"
)

func TestVolume_AustralianTableSpoons(t *testing.T) {
	AssertFloatEqual(t, 50.00000000000001, Volume(1*Litre).AustralianTableSpoons())
}
