package main

import (
	"fmt"

	piscine "piscine/Checkpoints"
)

func main() {
	// fmt.Println(piscine.DigitLen(-150, 10))
	// fmt.Println(piscine.DigitLength(150, 10))
	// fmt.Println(piscine.LastWord("hello world"))
	// fmt.Println(piscine.FishChip(6))
	// fmt.Println(piscine.CTS("CamelCase"))
	// fmt.Println(piscine.CTS("camelCase"))
	// fmt.Println(piscine.CTS("camelCaseIsTheBestCaseEver"))
	// fmt.Println(piscine.CTS("CamelCase1"))
	// fmt.Println(piscine.CTS("CamelCAse"))
	// fmt.Println(piscine.HashCode("123456789"))
	// fmt.Println(piscine.GCD(42, 10))

	// fmt.Print(piscine.FindPrevPrime(36))
	// fmt.Print(piscine.PrimeFindPrev(36))
	// fmt.Print(piscine.FromTo(1, 10))
	// fmt.Println()
	// fmt.Print(piscine.FromTo(10, 1))
	// fmt.Println()

	// fmt.Println(piscine.IsCapitalized("Hello! How are you?"))
	// fmt.Println(piscine.IsCapitalized("Hello How Are You"))

	// fmt.Println(piscine.ThirdTimeIsaCharm("123456789"))
	// fmt.Println()
	// fmt.Println(piscine.Atoi("1251"))
	// result1 := piscine.Itoa(1251)
	// result := piscine.Atoi("1251")
	// fmt.Printf("tyep of : %T | value : %v\n", result1, result1)
	// fmt.Printf("tyep of : %T | value : %v\n", result, result)

	// piscine.PrintMemory([10]byte{'H', 'E', 'l', 'L', 'o', 16, 21, '*'})

	// piscine.Display()

	// func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(piscine.CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(piscine.CanJump(input2))

	input3 := []uint{0}
	fmt.Println(piscine.CanJump(input3))
	// }
}
