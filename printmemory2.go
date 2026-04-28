package main

import (
	"fmt"

	"github.com/01-edu/z01"
	// "github.com/01-edu/z01"
)

func PrintMemory2(arr [10]byte) {
	hex := ""
	ascii := ""

	for i := 0; i < len(arr); i++ {
		ch := arr[i]
		high := ch / 16
		low := ch % 16

		if high < 10 {
			hex += string(high + '0')
		} else {
			hex += string(high + 'a' - 10)
		}

		if low < 10 {
			hex += string(low + '0')
		} else {
			hex += string(low + 'a' - 10)
		}

		if i == 3 || i == 7 || i == 9{
			hex += "\n"
		} else {
			hex += " "
		}
	}

	for _, i := range arr {
		if i >= 32 && i <= 126 {
			ascii += string(i)
		} else {
			ascii += "."
		}
	}
	fmt.Print(hex)
	fmt.Print(ascii + "\n") 
}

func main() {
	PrintMemory3([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory3(arr [10]byte) {
	base := "0123456789abcdef"
	for i := 0; i < len(arr); i++{
		z01.PrintRune(rune(base[arr[i]/16]))
		z01.PrintRune(rune(base[arr[i]%16]))

		if i == 3 || i == 7 || i == 9{
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}


// 0 1 2 3 4 5 6 7 8 9 a(10) b(11) c(12) d(13) e(14) f(15)
// to convert a byte to hexadecimal, split the byte into two nibbles
// n / 16
// n % 16

// hex := ""
// ascii := ""

// for i, ch := range arr{
// 	high := ch / 16
// 	low := ch % 16

// 	if high < 10 {
// 		hex += string(high + '0')
// 	} else {
// 		hex += string(high + 'a' - 10)
// 	}

// 	if low < 10 {
// 		hex += string(low + '0')
// 	} else {
// 		hex += string(low + 'a' - 10)
// 	}

// 	if i == 3 || i == 7 || i == 9{
// 		hex += "\n"
// 	} else {
// 		hex += " "
// 	}
// }

// for _, i := range arr{
// 	if i >= 32 && i <= 126 {
// 		ascii += string(i)
// 	} else {
// 		ascii += "."
// 	}
// }

// fmt.Print(hex)
// fmt.Print(ascii)

// hex := ""
// ascii := ""

// for i, ch := range arr{
// 	high := ch / 16
// 	low := ch % 16

// 	if high < 10 {
// 		hex += string('0' + high)
// 	} else {
// 		hex += string('a' + high - 10)
// 	}

// 	if low < 10 {
// 		hex += string('0' + low)
// 	} else {
// 		hex += string('a' + low - 10)
// 	}

// 	if i == 3 || i == 7 || i == 9 {
// 		hex += "\n"
// 	} else {
// 		hex += " "
// 	}
// }
// for _, i := range arr{
// 	if i >= 32 && i <= 126{
// 		ascii += string(i)
// 	} else {
// 		ascii += "."
// 	}
// }
// fmt.Print(hex)
// fmt.Print(ascii)
