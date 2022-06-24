package frac

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(x int64, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return Abs(x)
}

func LCM(x int64, y int64) int64 {
	if x == y || x == -y {
		return x
	}
	return Abs(x / GCD(x, y) * y)
}
