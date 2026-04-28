package main

import (
	// "fmt"
)

func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	if len(nbrs) == 0 {
		return a
	}

	start := nbrs[0]

	if start < 0 {
		start = length + start
	}

	end := length

	if len(nbrs) > 1 {
		end = nbrs[1]
		if end < 0 {
			end = length + end
		}
	}

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

