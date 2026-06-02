package piscine

func IsCapitalized(s string) bool {

	if len(s) == 0 {
		return false
	}

	for i, char := range s {

		if i == 0 || s[i-1] == ' ' {
			if char >= 'a' && char <= 'z' {
				return false
			}
		}
	}

	return true
}
