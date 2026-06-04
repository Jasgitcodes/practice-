package piscine

import "github.com/01-edu/z01"

func PrintRevComb() {

	var x, y, z rune

	for x = 9; x >= 2; x-- {
		for y = x - 1; y >= 1; y-- {
			for z = y - 1; z >= 0; z-- {

				z01.PrintRune(x + '0')
				z01.PrintRune(y + '0')
				z01.PrintRune(z + '0')

				if x != 2 || y != 1 || z != 0 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
