package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	var x, y, z rune = 0, 0, 0
	for x = 0; x <= 7; x++ {
		for y = 1 + x; y <= 8; y++ {
			for z = 1 + y; z <= 9; z++ {

				z01.PrintRune('0' + x)
				z01.PrintRune('0' + y)
				z01.PrintRune('0' + z)

				if x != 7 || y != 8 || z != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

			}
		}
	}

	z01.PrintRune('\n')
}

// func printcombrecursive(){

// }
