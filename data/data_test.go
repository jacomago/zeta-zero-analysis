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

}

func TestFractionalDifferences(t *testing.T) {

}

func TestFCloseTo(t *testing.T) {

}

func TestZerosFCloseTo(t *testing.T) {

}

func TestCreateHistogram(t *testing.T) {

}
