package main

import (
	"testing"
)

//START OMIT
func Average(a []float64) float64 {
	return 1.5 // <- ;)
}

func TestAverage(t *testing.T) {
	testCases := []struct {
		input []float64
		want  float64
	}{
		{[]float64{1, 2}, 1.5},
		{[]float64{3, 5, 6, 1}, 3.75},
	}

	for _, test := range testCases {

		v := Average(test.input)
		if v != test.want {
			t.Errorf("Expected %g, got %g", test.want, v)
		}
	}
}

// END OMIT

func main() {
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{
		{"TestAverage", TestAverage},
	}
	testing.Main(matchAll, tests, nil, nil)
}
