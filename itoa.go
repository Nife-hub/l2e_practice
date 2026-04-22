package main

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	negative := false 
		if n < 0 {
			negative = true
		}

	result := ""

	for n != 0{
		digit := n % 10
		
		if digit < 0{
			digit = -digit 
		}

		result = result + string(rune(digit + '0'))
		n /= 10
	}

	if negative {
		result = "-" + result
	}
	return result
} 

	
