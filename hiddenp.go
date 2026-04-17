package main

// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )

func HiddenP(s1, s2 string) int {
	if len(s1) == 0 {
		return 1
	}
	i, j := 0, 0

	for j < len(s2) {
		if s2[j] == s1[i] {
			i++
			if i == len(s1) {
				return 1
			}
		}
		j++
	}
	return 0
}

// func main(){
// 	arg := os.Args[1:]
// 	if len(arg) != 2 {
// 		return 
// 	}
// 	result := HiddenP(arg[0], arg[1])
// 	z01.PrintRune(rune(result + '0'))
// 	z01.PrintRune('\n')
// }
