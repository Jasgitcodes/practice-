package piscine

// func hiddenp(s1, s2 string) int {
// 	i := 0

// 	for j := range s2 {
// 		if i < len(s1) && s1[i] == s2[j] {
// 			i++
// 		}
// 	}

// 	if i == len(s1) {
// 		return 1
// 	}
// 	return 0
// }

func hiddenp(s1, s2 string) int {
	i := 0

	for ch := range s2 {
		if i < len(s1) && s1[i] == s2[ch] {
			i++
		}
	}

	if i == len(s1) {
		return 1
	}

	return 0
}

// func main() {

// 	input := os.Args[1:]

// 	if len(input) != 2 || len(input) > 2 {
// 		return
// 	}

// 	text1 := input[0]
// 	text2 := input[1]
// 	// Your code here

// 	if hiddenp(text1, text2) == 1 {
// 		z01.PrintRune('1')
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// }
