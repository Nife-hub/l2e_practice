package main

// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )

func Inter(s1, s2 string) string {
	printed := ""  
	for i := 0; i < len(s1); i++{
		ch := s1[i]
		found := false
		for k := 0; k < len(printed); k++ {
			if printed[k] == ch {
				found = true
				break
			}
		}
		if found {
			continue
		}

		for j := 0; j < len(s2); j++ {
			if s2[j] == ch {
				printed += string(ch)
				break
			}
		}
	}
	return printed
}

// func main(){
// 	arg := os.Args[1:]
// 	if len(arg) != 2 {
// 		return
// 	}

// 	result := Inter(arg[0], arg[1])
// 	for _, ch := range result {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }