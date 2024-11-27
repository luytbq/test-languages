package shared

import "math"

func IsPrime(number uint64) bool {
	if number < 2 {
		return false
	}

	if number%2 == 0 {
		return false
	}

	fNumber := float64(number)
	var i uint64
	for i = 3; i <= uint64(math.Sqrt(fNumber)); i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}
