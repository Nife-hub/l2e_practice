package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	target, err := strconv.Atoi(arg2)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	content := arg1[1 : len(arg1)-1]
	parts := strings.Split(content, ",")

	var numbers []int

	for _, part := range parts {
		part = strings.TrimSpace(part)

		num, err := strconv.Atoi(part)
		if err != nil {
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






























// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("Invalid input.")
// 		return 
// 	}

// 	arg1 := os.Args[1]
// 	arg2 := os.Args[2]
// 	// main := os.Args[1:]

// 	for i := 0; i < len(arg1); i++{
// 		if i == len(arg1)-1 && arg1[i] != ']' {
// 			fmt.Println("Invalid input.")
// 		}

// 		if i == 0 && arg1[i] != '[' {
// 			fmt.Println("Invalid input.")
// 		}

// 		ch := arg1[i]
// 		if i != 0 && i != len(arg1) {
// 			if !(ch >= '0' && ch <= '9') {
// 				fmt.Sprintf("Invalid number: %v", ch)
// 			}
// 		}
// 	}

// 	if len(os.Args[2]) != 1 {
// 		fmt.Println("Invalid target sum.")
// 	}

// 	var pairs [][]int
// 	for i := 0; i < len(arg1); i++{
// 		for j := i + 1; j <= len(arg1); j++{
// 			if arg1[i] + arg1[j] == arg2 {
// 				pairs = append(pairs, [i, j])
// 			} else {
// 				fmt.Println("No pairs found.")
// 			}
// 		}
// 	}
// 	fmt.Sprintf("Pairs with sum %v: [[%v %v] [%v %v]]", arg2, i, j)
// }