package piscine

func ThirdTimeIsaCharm(s string) string {

	if len(s) < 3 {
		return "\n"
	}

	var result string
	for i, r := range s {
		if (i+1)%3 == 0 {

			result += string(r)

		}
	}

	return result
}
