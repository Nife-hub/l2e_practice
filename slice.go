package main

import (
	"fmt"
)


func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	// No indices provided
	if len(nbrs) == 0 {
		return a
	}

	start := nbrs[0]

	// Handle negative start
	if start < 0 {
		start = length + start
	}

	end := length

	// Handle optional end
	if len(nbrs) > 1 {
		end = nbrs[1]
		if end < 0 {
			end = length + end
		}
	}

	// Prevent out-of-range panics
	if start < 0 {
		start = 0
	}
	if start > length {
		start = length
	}
	if end < start {
		end = start
	}
	if end > length {
		end = length
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
