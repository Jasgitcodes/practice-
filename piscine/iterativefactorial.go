package piscine

func IterativeFactorial(nb int) int {

	if nb < 0 || nb > 20 {
		return 0
	}

	result := 1

	for i := 1; i <= nb; i++ {
		result = result * i
	}

	return result
}

func RecursiveFactorial(nb int) int {

	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 1 {
		return 1
	}

	result := 1

	if nb > 1 {
		result = nb * RecursiveFactorial(nb-1)

	}
	return result

	// return 0
}
