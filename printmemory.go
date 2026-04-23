package main

import "fmt"

func PrintMemory(arr [10]byte) {
	hexline := ""
	asciiLine := ""

	for i := 0; i < 10; i++ {
		b := arr[i]

		high := b / 16
		low := b % 16

		if high < 10 {
			hexline += string('0' + high)
		} else {
			hexline += string('a' + (high - 10))
		}

		if low < 10 {
			hexline += string('0' + low)
		} else {
			hexline += string('a' + (low - 10))
		}

		hexline += " "

		if b >= 32 && b <= 126 {
			asciiLine += string(b)
		} else {
			asciiLine += "."
		}
	}
	fmt.Println(hexline + " " + asciiLine) 
}