package piscine

import (
	"github.com/01-edu/z01"
)

func contains(s string, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func Inter(s1, s2 string) {
	printed := ""

	for _, c := range s1 {
		if contains(s2, c) && !contains(printed, c) {
			z01.PrintRune(c)
			printed += string(c)
		}
	}
}

// func main() {

// 	args := os.Args[1:]

// 	if len(args) < 2 || len(args) > 2 {
// 		return
// 	}

// 	if len(args) == 2 {
// 		Inter(args[0], args[1])
// 	}
// 	z01.PrintRune('\n')
// }
