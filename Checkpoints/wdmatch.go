package piscine

import (
	"github.com/01-edu/z01"
)

func Wdmatch(s1, s2 string) {
	i := 0

	for j := 0; j < len(s2) && i < len(s1); j++ {
		if s1[i] == s2[j] {
			i++
		}
	}

	if i == len(s1) {
		for _, r := range s1 {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	if len(os.Args) != 3 {
// 		return
// 	}
// 	Wdmatch(os.Args[1], os.Args[2])
// }
