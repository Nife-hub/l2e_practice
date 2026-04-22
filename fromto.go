package main

func FromTo(a, b int) string{
	if a < 0 || a > 99 || b < 0 || b > 99 {
		return "Invalid\n"
	}

		result := ""
		if a > b{
			for i := a; i >= b; i--{
				result += FormatNumber(i)
				if i != b {
					result += ", "
				}
			}
		} else {
			for i := a; i <= b; i++{
				result += FormatNumber(i)
				if i != b {
					result += ", "
				}
			}
		}
		return result + "\n"
	}

	func FormatNumber(n int)string{
		if n < 10 {
			return "0" + string(rune(n + '0'))
		}
		tens := n/10
		ones := n%10
		
		return string(rune(tens + '0')) + string(rune(ones + '0'))
	}