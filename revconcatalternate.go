package main

func RevConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	i := len(slice1) - 1
	j := len(slice2) - 1
	// b := ""
	for i > j {
		// b = "x"
		result = append(result, slice1[i])
		i--
	}
	for j > i {
		result = append(result, slice2[j])
		j--
	}
	// if b == "x" {
		for i >= 0 && j >= 0 {
			result = append(result, slice1[i])
			result = append(result, slice2[j])
			i--
			j--
		}
	// } else {
	// 	for i >= 0 && j >= 0 {
	// 	result = append(result, slice1[j])
	// 	result = append(result, slice2[i])
	// 	i--
	// 	j--
	// }
		return result
	}

	

