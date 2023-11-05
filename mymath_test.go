package mymath

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	if Sqrt(4) != math.Sqrt(4) {
		t.Errorf("unexpected Error")
	}
}

func TestCeil(t *testing.T) {
	if Ceil(4) != math.Ceil(4) {
		t.Errorf("unexpected Error")
	}
}

func TestFloor(t *testing.T) {
	if Floor(4) != math.Floor(4) {
		t.Errorf("unexpected Error")
	}
}

func TestPow(t *testing.T) {
	if Pow(4, 2) != math.Pow(4, 2) {
		t.Errorf("unexpected Error")
	}
}

func TestMax(t *testing.T) {
	if Max(4, 2) != math.Max(4, 2) {
		t.Errorf("unexpected Error")
	}
}

func TestMin(t *testing.T) {
	if Min(4, 2) != math.Min(4, 2) {
		t.Errorf("unexpected Error")
	}
}
