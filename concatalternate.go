package main

func ConcatAlternate(a []int, b []int) []int {
	var result []int

	i, j := 0, 0

	aTurn := len(a) >= len(b)

	for i < len(a) || j < len(b) {
		if aTurn {
			if i < len(a) {
				result = append(result, a[i])
			}
		} else {
			if j < len(b) {
				result = append(result, b[j])
			}
		}
		
		aTurn = !aTurn
	}
	return result
}
