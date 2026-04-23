package main

import (
	"github.com/01-edu/z01"
	"os"
)

func Atoi(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
		n = n*10 + int(s[i]-'0')
	}
	return n
}

func PrintNbr(n int) {
	if n > 9 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	args := os.Args[1:]

	// invalid number of args → print nothing
	if len(args) != 1 {
		return
	}

	n := Atoi(args[0])

	// invalid number or no prime factors → print nothing
	if n < 2 {
		return
	}

	first := true
	i := 2

	for n > 1 {        
		if n%i == 0 {
			if !first {
				z01.PrintRune('*')
			}
			PrintNbr(i)
			n /= i
			first = false
		} else {
			i++
		}
	}
	z01.PrintRune('\n')
}