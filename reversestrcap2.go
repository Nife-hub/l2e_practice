package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	if len(os.Args) < 2{
		return
	}
	args := os.Args[1:]
	for _, arg := range args{
		res := ReverseStrCap2(arg)
		for _, result := range res{
			z01.PrintRune(result)
		}
	}
	z01.PrintRune('\n')
}

func ReverseStrCap2(s string) string {
	r := []byte(s)

	for i := 0; i < len(r); i++{
		r[i] = ToLower2(r[i])
	}

	for i := 0; i < len(r); i++{
		if IsLetter2(r[i]) {
			if i == len(r)-1 || r[i+1] == ' '{
				r[i] = ToUpper2(r[i])
			}
		}

	}
	return string(r)
}

func ToUpper2(c byte) byte {
	if c >= 'a' && c <= 'z' {
		c = c - 32
	}
	return c
}
func ToLower2(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c = c + 32
	}
	return c
}
func IsLetter2(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

