package main

func FromTo(a int, b int) string {
	if a < 0 || a > 99 || b < 0 || b > 99 {
		return "Invalid\n"
	}

	result := ""

	if a <= b {
		for i := a; i <= b; i++ {
			result += formatNumber(i)
			if i != b {
				result += ", "
			}
		}
	} else {
		for i := a; i >= b; i-- {
			result += formatNumber(i)
			if i != b {
				result += ", "
			}
		}
	}

	return result + "\n"
}

// helper function to convert int (0–99) to string
func formatNumber(n int) string {
	if n < 10 {
		return "0" + string(rune('0'+n))
	}
	tens := n / 10
	ones := n % 10
	return string(rune('0'+tens)) + string(rune('0'+ones))
}