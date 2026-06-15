package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdefghijklmnopqrstuvwxyz"

	// for i := 0; i < len(arr); i++ {

	// 	char := rune(arr[i])

	// 	z01.PrintRune(rune(hex[char/16]))
	// 	z01.PrintRune(rune(hex[char%16]))

	// 	z01.PrintRune(' ')

	// 	if (i+1)%4 == 0 {
	// 		z01.PrintRune('\n')
	// 	} else {
	// 		z01.PrintRune(' ')
	// 	}
	// }

	for i, char := range arr {

		z01.PrintRune(rune(hex[char/16]))
		z01.PrintRune(rune(hex[char%16]))

		z01.PrintRune(' ')

		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
	for _, char := range arr {
		if char >= 32 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}

func PrintMemory1(arr [10]byte) {
	// hex := "0123456789abcdefghijklmnopqrstuvwxyz"
	for i, char := range arr {

		fmt.Printf("%02x ", char)

		if (i+1)%4 == 0 {
			fmt.Println()
		}

	}
	fmt.Println()

	for _, char := range arr {
		// result :=""
		if char >= 32 && char <= 126 {
			fmt.Print(string(char))
		} else {
			fmt.Print(".")
		}
	}
}
