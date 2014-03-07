package data

import (
	"testing"
)

func TestFloatsEqual(t *testing.T) {
	const inA, inB, out = 1.2, 1.2, true
	if x := floatsEqual(inA, inB); x != out {
		t.Errorf("floatsEqual(%v, %v) = %v, want %v", inA, inB, x, out)
	}

	const inA2, inB2, out2 = 1.2, 1.3, false
	if x := floatsEqual(inA2, inB2); x != out2 {
		t.Errorf("floatsEqual(%v, %v) = %v, want %v", inA2, inB2, x, out2)
	}
}

func TestMax(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6}
	const out = 6.0
	if x := Max(in); !floatsEqual(x, out) {
		t.Errorf("max(%v) = %v, want %v", in, x, out)
	}
}

func TestAvg(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6}
	out := 3.5
	if x := Avg(in); !floatsEqual(x, out) {
		t.Errorf("Avg(%v) = %v, want %v", in, x, out)
	}
}

func TestImportZeros(t *testing.T) {
	const in string = "test.dat"
	out := Zeros{3.14: 0.14, 2.71: 0.71}
	x := ImportZeros(in)
	for k, v := range x {
		if !floatsEqual(out[k], v) {
			t.Errorf("ImportZeros(%v) = %v, want %v", in, x, out)
		}
	}
}

func TestClosestZero(t *testing.T) {
	in := Zeros{3.14: 0.14, 2.71: 0.71, 4.56: 0.56}
	const out = 0.42
	in2 := 0.14
	if x := closestZero(in, 0.14); !floatsEqual(x, out) {
		t.Errorf("closestZero(%v, %v) = %v, want %v", in, in2, x, out)
	}
}

func TestFractionalDifferences(t *testing.T) {
	in := Zeros{3.14: 0.14, 2.71: 0.71, 4.56: 0.56}
	out := []float64{0.42, 0.15, 0.15}
	x := FractionalDifferences(in)
	for i, v := range x {
		if !floatsEqual(out[i], v) {
			t.Errorf("fractionalDifferences(%v) = %v, want %v", in, x, out)
		}
	}
}

func TestfCloseTo(t *testing.T) {
	in := Zeros{3.14: 0.14, 2.71: 0.71, 4.56: 0.56}
	inZero := 0.14
	inEpsilon := 0.43
	out := 1
	if x := fCloseTo(in, inZero, inEpsilon); x != out {
		t.Errorf("fCloseTo(%v, %v, %v) = %v, want %v", in, inZero, inEpsilon, x, out)
	}
}

func TestzerosFCloseTo(t *testing.T) {
	in := Zeros{3.14: 0.14, 2.71: 0.71, 4.56: 0.56}
	inP := func(n int) float64 { return 1 / float64(n) }
	out := []int{0, 1, 1}
	r := ZerosFCloseTo(in, inP)
	for i, x := range r {
		if x != out[i] {
			t.Errorf("zerosFCloseTo(%v, %v) = %v, want %v", in, inP, x, out)
		}
	}
}
