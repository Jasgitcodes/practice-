package piscine

import (
	"github.com/01-edu/z01"
)

func Fprime(n int) {
	first := true

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			if !first {
				printStr("*")
			}

			printN(i)
			first = false
			n /= i
		}
	}

	if n > 1 {
		if !first {
			printStr("*")
		}
		printN(n)
	}
}

func Atoiii(s string) int {
	n := 0

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}

		n = n*10 + int(ch-'0')
	}

	return n
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func printN(n int) {
	if n >= 10 {
		printN(n / 10)
	}

	z01.PrintRune(rune(n%10 + '0'))
}

// func main() {
// 	input := os.Args[1:]

// 	n := Atoiii(input[0])

// 	if n <= 1 || len(input) != 1 {
// 		return
// 	}

// 	Fprime(n)
// 	printStr("\n")
// }
