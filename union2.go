package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Union2(s1, s2 string) string {
	freq := make(map[rune]bool)
	result := ""

	for _, ch := range s1{
		if !freq[ch] {
			freq[ch] = true
			result += string(ch)
		}
	}

	for _, ch := range s2{
		if !freq[ch] {
			freq[ch] = true
			result += string(ch)
		}
	}
	return result
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		z01.PrintRune('\n')
		return
	}

	res := Union2(arg[0], arg[1])
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}