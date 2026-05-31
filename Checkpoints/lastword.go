package piscine

func LastWord(s string) string {

	runes := []rune(s)
	end := len(runes) - 1

	for end >= 0 && runes[end] == ' ' {
		end--
	}

	if end < 0 {
		return ""
	}

	start := end

	for start >= 0 && runes[start] != ' ' {
		start--
	}
	result := string(runes[start+1:end+1]) + ""
	return result
}
