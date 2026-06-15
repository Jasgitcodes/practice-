package piscine

func HashCode(dec string) string {
	size := len(dec)
	asci := 0
	result := ""
	for _, char := range dec {
		asci = int(char)
		hashed := (asci + size) % 127

		if hashed < 33 {
			hashed += 33
		}

		result += string(rune(hashed))
	}

	return result
}
