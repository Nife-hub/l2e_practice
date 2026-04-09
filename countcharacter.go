package main

func CountChar(s string, c rune) int {
	count := 0
	for _,  v := range s {
		if v == c {
			count++
		}
	}
	return count
}