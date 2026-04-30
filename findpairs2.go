package main

import (
	"fmt"
	"os"
)

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
		if len(s) == 1 {
			return 0, false
		}
	}

	num := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		num = num*10 + int(s[i]-'0')
	}

	return num * sign, true
}

func TrimSpace(s string) string {
	start := 0
	end := len(s) - 1

	for start <= end && s[start] == ' ' {
		start++
	}

	for end >= start && s[end] == ' ' {
		end--
	}

	return s[start : end+1]
}

func Split(s string) []string {
	var result []string
	word := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			result = append(result, word)
			word = ""
		} else {
			word += string(s[i])
		}
	}

	result = append(result, word)
	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	if len(arg1) < 2 || arg1[0] != '[' || arg1[len(arg1)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	target, ok := Atoi(arg2)
	if !ok {
		fmt.Println("Invalid target sum.")
		return
	}

	content := arg1[1 : len(arg1)-1]
	parts := Split(content)

	var numbers []int

	for _, part := range parts {
		part = TrimSpace(part)

		num, ok := Atoi(part)
		if !ok {
			fmt.Printf("Invalid number: %s\n", part)
			return
		}

		numbers = append(numbers, num)
	}

	var pairs [][]int

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
}