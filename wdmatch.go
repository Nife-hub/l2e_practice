package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	s1 := args[0]
	s2 := args[1]

	i := 0

	for j := 0; j < len(s2) && i < len(s1); j++ {
		if s2[j] == s1[i] {
			i++
		}
	}

	// if we matched everything in s1
	if i == len(s1) {
		for _, c := range s1 {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}