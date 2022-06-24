package frac

import "fmt"

type Fraction struct {
	N int64
	D int64
}

func Mul(a Fraction, b Fraction) Fraction {
	return Fraction{N: a.N * b.N, D: a.D * b.D}
}

func (f Fraction) Float32() float32 {
	return float32(f.N) / float32(f.D)
}

func (f Fraction) Float64() float64 {
	return float64(f.N) / float64(f.D)
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.N, f.D)
}

func (f Fraction) Abs() Fraction {
	return Fraction{N: Abs(f.N), D: Abs(f.D)}
}

func (f Fraction) IsProper() bool {
	return f.N <= f.D
}

func (f Fraction) IsImproper() bool {
	return f.N > f.D
}

func (f Fraction) IsUnity() bool {
	return f.N == f.D
}
