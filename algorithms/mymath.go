package algorithms

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

func SumLocations(x int) int {
	ret := 0
	for x > 10 {
		ret += x % 10
		x /= 10
	}
	return ret
}

func SumOne(a int) int {
	ret := 0
	x := a
	for x > 0 {
		if x&1 == 1 {
			ret++
		}
		x >>= 1
	}
	return ret
}
