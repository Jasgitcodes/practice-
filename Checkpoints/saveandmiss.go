package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	res := ""
	save := true

	for i := 0; i < len(arg); i += num {
		if save {
			for j := i; j < i+num && j < len(arg); j++ {
				res += string(arg[j])
			}
		}
		save = !save
	}

	return res
}

// ```
// Or
// ```

func SaveAndMiss1(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	var result string

	for i := 0; i < len(arg); i += num * 2 {
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}

		for j := i; j < end; j++ {
			result += string(arg[j])
		}
	}

	return result
}
