package piscine

func ConcatSlice(slice1, slice2 []int) []int {

	var result []int

	for i, _ := range slice1 {
		result = append(result, slice1[i])
	}

	for i, _ := range slice2 {
		result = append(result, slice2[i])
	}

	return result
}

// func main() {
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
// }
