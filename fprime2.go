package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi3(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
		n = n*10 + int(s[i]-'0')
	}
	return n
}

func PrintNbr2(n int) {
	if n > 9 {
		PrintNbr2(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		return
	}

	n := Atoi3(arg[0])
	first := true
	i := 2

	for n > 1{
		if n%i == 0{
			if !first {
				z01.PrintRune('*')
			}
			PrintNbr2(i)
			first = false
			n /= i
		} else {
			i++
		}
	}
	z01.PrintRune('\n')
}
