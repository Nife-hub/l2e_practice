package main

func CheckNumber(str string) bool {
	for _, v := range str {
		if v >= '0' && v <= '9' {
			return true
		}
	}
	return false
}