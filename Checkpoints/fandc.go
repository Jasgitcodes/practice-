package piscine

func FishChip(n int) string {

	if n < 0 {
		return "error: number is negative"
	}

	if n%2 == 0 && n%3 == 0 {

		return "fish and chips "

	} else if n%2 == 0 {

		return "Fish"

	} else if n%3 == 0 {

		return "Chips"

	} else {

		return "error: number is not divisible"
	}

}

func lasword(s string) string {

	i := len(s) - 1
	end := 0

	for i > 0 && s[i] == ' ' {

		i--

	}

	end = i

	for i > 0 && s[i] != ' ' {
		i--
	}

	return s[i+1 : end+1]
}
