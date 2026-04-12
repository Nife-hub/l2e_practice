package main

func RepeatAlpha(s string) string{
	result := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			repeat := int(ch - 'a') + 1
			for i := 0; i < repeat; i++ {
				result = result + string(ch)
			}
		} 

		if ch >= 'A' && ch <= 'Z' {
			repeat := int(ch - 'A') + 1
			for i := 0; i < repeat; i++ {
				result = result + string(ch)
			}
		} else {
			result = result + string(ch)
		}
	}
	return result
}