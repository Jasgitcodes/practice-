package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintParams() {
	params := os.Args[1:]

	for i := 0; i < len(params)-1; i++ {
		// fmt.Println(params[i])

		s := params[i]

		for _, r := range s {
			z01.PrintRune(r)
		}

		z01.PrintRune('\n')
	}
}

func PrintRevParams() {
	params := os.Args[1:]

	for i := len(params) - 1; i >= 0; i-- {
		s := params[i]

		for _, r := range s {
			z01.PrintRune(r)
		}

		z01.PrintRune('\n')
	}
}

func SortParams() {

	s := os.Args[1:]
	n := len(s)

	for i := 0; i < n-1; i++ {

		for _, r := range s[i] {

			result := ""
			if r >= 32 && r <= 126 {

				result += string(r)

				// z01.PrintRune('\n')
			}
		}

	}

}
