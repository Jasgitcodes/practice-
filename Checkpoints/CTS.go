package piscine

// func CTS(s string) string {

// 	if s == "" {
// 		return s
// 	}

// 	result := ""
// 	b := []byte(s)

// 	for i := 0; i < len(b); i++ {

// 		if (b[i] < 'A' || b[i] > 'Z') && (b[i] < 'a' || b[i] > 'z') {
// 			return s
// 		}

// 		if i > 0 && b[i] >= 'A' && b[i] <= 'Z' && b[i-1] >= 'A' && b[i-1] <= 'Z' {
// 			return s
// 		}
// 	}

// 	if b[len(b)-1] >= 'A' && b[len(b)-1] <= 'Z' {
// 		return s
// 	}

// 	for i := 0; i < len(b); i++ {
// 		if b[i] >= 'A' && b[i] <= 'Z' {
// 			if i != 0 {
// 				result += "_"
// 			}

// 			result += string(b[i] + 32)
// 		} else {
// 			result += string(b[i])
// 		}
// 	}

// 	return result
// }

func CTS(s string) string {
	if s == "" {
		return ""
	}

	if !isValidCamelCase(s) {
		return s
	}

	result := ""
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(char)
		} else {
			result += string(char)
		}
	}

	return result
}

func isValidCamelCase(s string) bool {
	for i, char := range s {
		if char >= '0' && char <= '9' {
			return false
		}

		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return false
		}

		if char >= 'A' && char <= 'Z' {
			if i == len(s)-1 {
				return false
			}

			if i > 0 {
				prevChar := rune(s[i-1])
				if prevChar >= 'A' && prevChar <= 'Z' {
					return false
				}
			}
		}
	}

	return true
}
