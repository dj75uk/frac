package frac_test

import (
	"frac"
	"testing"
)

type twoInputTestData struct {
	x        frac.Fraction
	y        frac.Fraction
	expected frac.Fraction
}

func getMulTestData() []twoInputTestData {
	return []twoInputTestData{
		{x: frac.Fraction{N: 0, D: 0}, y: frac.Fraction{N: 0, D: 0}, expected: frac.Fraction{N: 0, D: 0}},
	}
}

func TestMul(t *testing.T) {
	for _, data := range getMulTestData() {
		actual := frac.Mul(data.x, data.y)
		if actual != data.expected {
			t.Errorf("x: %v, y: %v, expected: %v, actual: %v", data.x, data.y, data.expected, actual)
		}
	}
}
