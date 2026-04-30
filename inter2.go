package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Inter2(s1, s2 string) string {
	printed := ""

	for _, a := range s1{
		alreadyseen := false
		for _, b := range printed{
			if b == a {
				alreadyseen = true
				break
			}
		}
		if alreadyseen {
			continue
		}

		for _, v := range s2{
			if v == a {
				printed = printed + string(a)
				break
			}
		}
	}
	return printed
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}

	result := Inter2(arg[0], arg[1])
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}











// printed := ""
// 	for i := 0; i < len(s1); i++ {
// 		ch := s1[i]
// 		found := false
// 		for b := 0; b < len(printed); b++ {
// 			if printed[b] == s1[i] {
// 				found = true
// 				break
// 			}
// 		}

// 		if found {
// 			continue
// 		}

// 		for c := 0; c < len(s2); c++ {
// 			if s2[c] == ch {
// 				printed = printed + string(ch)
// 				break
// 			}
// 		}
// 	}
// 	return printed