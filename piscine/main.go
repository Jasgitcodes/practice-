package main

import (
	"fmt"
	"piscine"
)

func main() {
	// fmt.Println(piscine.DigitLen(150, 10))
	// fmt.Println(piscine.LastWord("hello world"))
	// fmt.Println(piscine.FishChip(6))
	// fmt.Println(piscine.CTS("HelloWorld"))
	// piscine.IsNegative(25)
	// piscine.PrintComb()
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
