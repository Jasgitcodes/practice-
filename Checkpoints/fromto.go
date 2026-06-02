package piscine

import "strconv"

func FromTo(from, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid input\n"
	}

	step := 1
	if from > to {
		step = -1
	}

	result := ""

	for i := from; ; i += step {
		if i < 10 {
			result += "0"
		}

		result += strconv.Itoa(i)

		if i == to {
			break
		}

		result += ", "
	}

	return result
}
