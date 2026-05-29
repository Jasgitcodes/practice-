package main

import (
	"github.com/01-edu/z01"
)

func PrintAlphabet() {
	for letter := 'a'; letter <= 'z'; letter++ {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func PrintReverse() {

	for letter := 'z'; letter >= 'a'; letter-- {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}

func PrintDigits() {

	for digits := '0'; digits <= '9'; digits++ {
		z01.PrintRune(digits)
	}
	z01.PrintRune('\n')
}

func PrintDigitsRecursive(digits int) int {

	for i := 0; i < digits; i++ {

		if digits == 0 {
			return 0
		}

		result := 0
		if i > 1 {
			result = PrintDigitsRecursive(digits - 1)
		}

		return result
	}

	return 0
}

func main() {

	PrintDigitsRecursive(5)
}
