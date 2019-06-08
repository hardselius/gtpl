package tempconv

import (
	"math"
	"testing"
)

type TestCase struct {
	F Fahrenheit
	C Celsius
	K Kelvin
}

var testCases = []TestCase{
	{68, 20, 293.15},
	{32, 0, 273.15},
	{-40, -40, 233.15},
}

func TestTempConv(t *testing.T) {
	eps := 0.0000001 // acceptable floating point error
	for _, test := range testCases {
		if math.Abs(float64(CToF(test.C)-test.F)) > eps {
			t.Errorf("CToF(%s): got %s, want %s", test.C, CToF(test.C), test.F)
		}
		if math.Abs(float64(FToC(test.F)-test.C)) > eps {
			t.Errorf("FToC(%s): got %s, want %s", test.F, FToC(test.F), test.C)
		}
		if math.Abs(float64(KToC(test.K)-test.C)) > eps {
			t.Errorf("KToC(%s): got %s, want %s", test.K, KToC(test.K), test.C)
		}
	}
}
