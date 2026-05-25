package main

import "github.com/01-edu/z01"

func isNegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}

func main() {
	isNegative(2)
}
