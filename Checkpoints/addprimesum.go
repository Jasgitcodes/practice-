package piscine

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Atoi1(s string) int {
	n := 0
	for _, number := range s {

		if number < '0' || number > '9' {
			// z01.PrintRune('\n')
			return 0
		}
		n = n*10 + int(number-'0')
	}

	return n
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func AddPrimeSum(c int) int {
	sum := 0

	for i := 2; i <= c; i++ {
		if IsPrime(i) {
			sum += i

			fmt.Printf("%v + ", i)

		}
	}
	fmt.Println()

	return sum
}

func PrintInt(number int) {
	if number == 0 {
		z01.PrintRune('0')
		return
	}

	digit := []rune{}

	for number > 0 {
		convert := rune(number%10) + '0'
		digit = append([]rune{convert}, digit...)
		number = number / 10
	}

	for _, d := range digit {
		z01.PrintRune(d)
	}
}

func Display() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('0')
		return
	}

	n := Atoi1(args[0])

	if n <= 0 {
		PrintInt(0)
		z01.PrintRune('\n')
		return
	}

	result := AddPrimeSum(n)

	PrintInt(result)
	z01.PrintRune('\n')
}
