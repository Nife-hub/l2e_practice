package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2{
		return
	}

	s1 := arg[0]
	s2 := arg[1]

	i := 0

	for j := 0; j < len(s2) && i < len(s1); j++{
		if s2[j] == s1[i] {
			i++
		}
	}

	if i == len(s1){              //if at the end of the first string
		for _, c := range s1{
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}



















// args := os.Args[1:]

// 	if len(args) != 2 {
// 		return
// 	}

// 	s1 := args[0]
// 	s2 := args[1]

// 	i := 0

// 	for j := 0; j < len(s2) && i < len(s1); j++ {
// 		if s2[j] == s1[i] {
// 			i++
// 		}
// 	}

// 	if i == len(s1) {
// 		for _, c := range s1 {
// 			z01.PrintRune(c)
// 		}
// 		z01.PrintRune('\n')
// 	}