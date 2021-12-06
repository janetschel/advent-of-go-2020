package maths

func Abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

func Gcd(first int, second int) int {
	var div int

	for i := 1; i <= first && i <= second; i++ {
		if first%i == 0 && second%i == 0 {
			div = i
		}
	}

	return div
}

func MaxInt() int {
	return int(^uint(0) >> 1)
}
