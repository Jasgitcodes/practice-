package piscine

func RepeatAlpha(s string) string {

	result := ""
	repeatCount := 0

	for _, char := range s {

		lowerCase := (char >= 'a' && char <= 'z')
		upperCase := (char >= 'A' && char <= 'Z')
		if lowerCase || upperCase {
			if lowerCase {
				repeatCount = int(char - 'a' + 1)
			} else {

				repeatCount = int(char - 'A' + 1)
			}
			for i := 0; i < repeatCount; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}
