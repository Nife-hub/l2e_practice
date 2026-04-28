package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// 	// "os"
// )

func Atoi(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return -1
		}
		n = n*10 + int(ch-'0')
	}
	return n
}

func PrintNbr(n int) {
	if n > 9 {
		PrintNbr(n / 10)
	}
	// z01.PrintRune(rune(n%10 + '0'))
}

// func main() {
// 	arg := os.Args[1:]

// 	if len(arg) != 1{
// 		return
// 	}

// 	n :=Atoi(arg[0])

// 	first := true
// 	i := 2

// 	for n > 1{
// 		if n%i == 0{
// 			if !first {
// 				z01.PrintRune('*')
// 			}
// 			PrintNbr(i)
// 			n /= i
// 			first = false
// 		} else {
// 			i++
// 		}
// 	}
// }