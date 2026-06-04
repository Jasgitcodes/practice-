package piscine

func DigitLength(n, base int) int {

	if base < 2 || base > 36 {
		return -1
	}

	if n < 0 {
		n = -n
	}

	count := 0
	for n >= 0 {

		if n == 0 {
			break
		}

		count++
		n = n / base
	}

	return count
}
