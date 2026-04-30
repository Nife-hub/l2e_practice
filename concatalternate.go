package main

func ConcatAlternate(a []int, b []int) []int {
	var result []int
	x := len(a)
	y := len(b)
	i := 0
	j := 0
	for x > y {
		result = append(result, a[i])
		i++
		x--
	}
	for y > x {
		result = append(result, b[j])
		j++
		y--
	}
	for i < len(a) && j < len(b) {
		result = append(result, a[i])
		result = append(result, b[j])
		i++
		j++
	}
	return result
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





// i, j := 0, 0

// 	aTurn := len(a) >= len(b)

// 	for i < len(a) || j < len(b) {
// 		if aTurn{
// 			if i < len(a) {
// 				slice = append(slice, a[i])
// 				i++
// 			}
// 		} else {
// 			if j < len(b) {
// 				slice = append(slice, b[j])
// 				j++
// 			}
// 		}
// 		aTurn = !aTurn
// 	}
// 	return slice