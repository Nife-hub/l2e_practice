package main

func ConcatAlternate(a []int, b []int) []int {
	var slice []int

	i, j := 0, 0

	aTurn := len(a) >= len(b)

	for i < len(a) || j < len(b) {
		if aTurn{
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

// aTurn := len(a) >= len(b)

// 		for i < len(a) || j < len(b){
// 			if aTurn {
// 				if i < len(a) {
// 					slice = append(slice, a[i])
// 					i++
// 				}
				
// 			} else {
// 				if j < len(b) {
// 					slice = append(slice, b[j])
// 					j++
// 				}
// 			}
// 			aTurn = !aTurn
// 		}


// i := 0  
// 	j := 0

// 	for i > j{
// 		slice = append(slice, a[i])
// 		i++
// 	}
// 	for j > i{
// 		slice = append(slice, b[i])
// 		j++
// 	}

// 	for i < len(a) && j < len(b){
// 		slice = append(slice, a[i])
// 		slice = append(slice, b[i])
// 		i++
// 		j++
// 	}