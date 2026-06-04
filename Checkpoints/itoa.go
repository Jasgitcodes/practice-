package piscine

func Itoa(n int) string {

	result := ""
	for n > 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10

	}

	return result
}

func Atoi(s string) int {

	result := 0
	for _, char := range s {

		digit := int(char - '0')
		result = result*10 + digit
	}

	return result
}
