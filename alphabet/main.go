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
