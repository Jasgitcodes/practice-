package piscine

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}

	for nb >= 2 {

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
