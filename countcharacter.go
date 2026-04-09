package main

func CountCharacter(str string, c rune) int {
	count := 0
	for v := range str {
		if v == c {
			count++
		}
	}
	return count
}