package main

func ConcatAlternate(a []int, b []int) []int {
	var slice []int

	i, j := 0, 0

	aTurn := len(a) <= len(b)

	for i < len(a) || i < len(b){
		if aTurn {
			if i < len(a) {
				slice = append(slice, a[i])
				i++
			}
			
		} else {
			if j < len(b) {
				slice = append(slice, b[j])
				j++
			}
		}
		aTurn = !aTurn
	}
	return slice
}
