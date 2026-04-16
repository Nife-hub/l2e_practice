package main

import (
	"fmt"
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

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func AddPrimeSum(n int) int {
	if n < 0 {
		return 0
	}
	sum := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	n := Atoi(os.Args[1])
	if n < 0 {
		fmt.Println(0)
		return
	}

	fmt.Println(AddPrimeSum(n))
}