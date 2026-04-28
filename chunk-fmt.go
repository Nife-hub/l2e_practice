package main

import "fmt"

func Chunk2(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	var chunk [][]int

	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}
		chunk = append(chunk, slice[i:end])
	}
	fmt.Println(chunk)
}



