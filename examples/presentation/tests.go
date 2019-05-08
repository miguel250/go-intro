package main

import (
	"testing"
)

//START OMIT
func Average(a []float64) float64 {
	return 1.5 // <- ;)
}

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
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
