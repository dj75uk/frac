package frac_test

import (
	"frac"
	"testing"
)

func TestAbsNegative(t *testing.T) {
	const Input int64 = -23423423423
	const Expected int64 = -1 * Input
	actual := frac.Abs(Input)
	if actual != Expected {
		t.Errorf("expected: %d, actual: %d", Expected, actual)
	}
}

func TestAbsZero(t *testing.T) {
	const Input int64 = 0
	const Expected int64 = 0
	actual := frac.Abs(Input)
	if actual != Expected {
		t.Errorf("expected: %d, actual: %d", Expected, actual)
	}
}

func TestAbsPositive(t *testing.T) {
	const Input int64 = 6475685674
	const Expected int64 = Input
	actual := frac.Abs(Input)
	if actual != Expected {
		t.Errorf("expected: %d, actual: %d", Expected, actual)
	}
}

type twoInt64InputTestData struct {
	x        int64
	y        int64
	expected int64
}

func getGcdTestData() []twoInt64InputTestData {
	return []twoInt64InputTestData{
		{x: 0, y: 0, expected: 0},
		{x: 0, y: 0, expected: 0},
	}
}

func getLcmTestData() []twoInt64InputTestData {
	return []twoInt64InputTestData{
		{x: 0, y: 0, expected: 0},
		{x: 0, y: 0, expected: 0},
	}
}

func TestGcd(t *testing.T) {
	for _, data := range getGcdTestData() {
		actual := frac.GCD(data.x, data.y)
		if actual != data.expected {
			t.Errorf("x: %d, y: %d, expected: %d, actual: %d", data.x, data.y, data.expected, actual)
		}
	}
}

func TestLcm(t *testing.T) {
	for _, data := range getLcmTestData() {
		actual := frac.GCD(data.x, data.y)
		if actual != data.expected {
			t.Errorf("x: %d, y: %d, expected: %d, actual: %d", data.x, data.y, data.expected, actual)
		}
	}
}
