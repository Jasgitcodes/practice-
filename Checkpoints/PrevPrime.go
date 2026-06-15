package piscine

func PrimeFindPrev(nb int) int {
	if nb <= 2 {
		return 0
	}

	for nb > 2 {
		isPrime := true

		for i := 2; i*i < nb; i++ {
			if nb%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			return nb
		}

		nb--
	}

	return 0
}
