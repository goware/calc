package calc_test

import (
	"testing"

	"github.com/goware/calc"
)

func TestMinMax(t *testing.T) {
	x, y := 10, 50
	n := calc.Min(x, y)
	if n != x {
		t.Fatalf("min assert failed")
	}

	x, y = 10, 50
	n = calc.Max(x, y)
	if n != y {
		t.Fatalf("max assert failed")
	}

	a, b := 10.10, 50.50
	c := calc.Max(a, b)
	if c != b {
		t.Fatalf("max assert failed")
	}

}
