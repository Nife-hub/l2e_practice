package main

import (
	"os"

	"github.com/01-edu/z01"
)

func HiddenP2(a, b string) int {
	if len(a) == 0 {
		return 1
	}

	i, j := 0, 0

	for j < len(b) {
		if b[j] == a[i] {
			i++
			if i == len(a) {
				return 1
			}
		}
		j++
	}
	return 0
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		return
	}

	result := HiddenP2(arg[0], arg[1])
	z01.PrintRune(rune(result + '0'))
	z01.PrintRune('\n')
}