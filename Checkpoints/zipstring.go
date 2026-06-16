package piscine

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""
	count := 1

	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
		} else {
			res += itoa(count) + string(s[i-1])
			count = 1
		}
	}

	return res + "$"
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	digits := "0123456789"
	res := ""

	for n > 0 {
		res = string(digits[n%10]) + res
		n /= 10
	}

	return res
}
