package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintProgramName() {
	name := os.Args[0]

	start := 0

	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' || name[i] == '\\' {
			start = i + 1
			break
		}
	}

	for _, r := range name[start:] {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')

}
