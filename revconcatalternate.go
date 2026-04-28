package main

func RevConcatAlternate(slice1,slice2 []int) []int {
	var result []int

	i := len(slice1)-1
	j := len(slice2)-1

	for i > j {
		result = append(result, slice1[i])
		i--
	}
	for j > i {
		result = append(result, slice2[j])
		j--
	}
	for i >= 0 && j >= 0 {
		result = append(result, slice1[i])
		result = append(result, slice2[j])
		i--
		j--
	}
	return result
}