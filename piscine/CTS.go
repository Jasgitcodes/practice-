package piscine

func CTS(s string) string {

	if s == "" {
		return s
	}

	result := ""
	b := []byte(s)

	for i := 0; i < len(b); i++ {

		if (b[i] < 'A' || b[i] > 'Z') && (b[i] < 'a' || b[i] > 'z') {
			return s
		}

		if i > 0 && b[i] >= 'A' && b[i] <= 'Z' && b[i-1] >= 'A' && b[i-1] <= 'Z' {
			return s
		}
	}

	if b[len(b)-1] >= 'A' && b[len(b)-1] <= 'Z' {
		return s
	}

	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			if i != 0 {
				result += "_"
			}

			result += string(b[i] + 32)
		} else {
			result += string(b[i])
		}
	}

	return result
}

// func main() {

// 	case1 := cTS("CamelCase")
// 	case2 := cTS("camelCase")
// 	case3 := cTS("camelCase1")
// 	case4 := cTS("camelCAse")

// 	fmt.Println(case1)
// 	fmt.Println(case2)
// 	fmt.Println(case3)
// 	fmt.Println(case4)
// }
