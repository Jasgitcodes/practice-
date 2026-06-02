package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return s
	}

	// c := []rune(s)
	c := []rune(s)

	for i := 0; i < len(c); i++ {
		if (c[i] < 'A' || c[i] > 'Z') && (c[i] < 'a' || c[i] > 'z') {
			return s
		}

		if i > 0 && c[i] >= 'A' && c[i] <= 'Z' && c[i-1] >= 'A' && c[i-1] <= 'Z' {
			return s
		}
	}

	if c[len(c)-1] >= 'A' && c[len(c)-1] <= 'Z' {
		return s
	}

	result := ""

	for i := 0; i < len(c); i++ {
		if c[i] >= 'A' && c[i] <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(c[i] + 32)
		} else {
			result += string(c[i])
		}
	}
	return result
}
