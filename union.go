package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args[1:]

// 	// If not exactly 2 arguments → print newline only
// 	if len(args) != 2 {
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	s1 := args[0]
// 	s2 := args[1]

// 	seen := ""

// 	// helper: check if char already printed
// 	isSeen := func(c byte) bool {
// 		for i := 0; i < len(seen); i++ {
// 			if seen[i] == c {
// 				return true
// 			}
// 		}
// 		return false
// 	}

// 	// process s1 first
// 	for i := 0; i < len(s1); i++ {
// 		c := s1[i]
// 		if !isSeen(c) {
// 			seen += string(c)
// 			z01.PrintRune(rune(c))
// 		}
// 	}

// 	// then process s2
// 	for i := 0; i < len(s2); i++ {
// 		c := s2[i]
// 		if !isSeen(c) {
// 			seen += string(c)
// 			z01.PrintRune(rune(c))
// 		}
// 	}

// 	z01.PrintRune('\n')
// }