package main

// import (
// 	"os"
// 	"github.com/01-edu/z01"
// )

func IsLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func ToLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func ToUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func ReverseStrCap(s string) string {
	r := []byte(s)

	for i := 0; i < len(r); i++ {
		r[i] = ToLower(r[i])
	}

	for i := 0; i < len(r); i++ {
		if IsLetter(r[i]) {
			if i == len(r)-1 || r[i+1] == ' ' {
				r[i] = ToUpper(r[i])
			}
		}
	}

	return string(r)
}

// func main() {
// 	args := os.Args[1:]

// 	if len(args) == 0 {
// 		return
// 	}

// 	for _, arg := range args {
// 		res := ReverseStrCap(arg)

// 		for _, r := range res {
// 			z01.PrintRune(r)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }