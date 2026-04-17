package main

import "github.com/01-edu/z01"

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
	}
	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	for _, chunk := range chunks {
	z01.PrintRune('[')

		for i, num := range chunk {
			z01.PrintRune(rune(num + '0'))

			if i != len(chunk)-1 {
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune(']')
		z01.PrintRune('\n')
	} 
}
