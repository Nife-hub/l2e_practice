package main

func CheckNumber(str string) bool {
	for v := range str {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}