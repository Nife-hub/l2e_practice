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

func RepeatAlpha2(s string) string {
    result := ""
    for _, r := range s {
        repeat := 1
        if r >= 'a' && r <= 'z' {
            repeat = int(r - 'a' + 1)
        } else if r >= 'A' && r <= 'Z' {
            repeat = int(r - 'A' + 1)
        }
        for i := 0; i < repeat; i++ {
            result += string(r)
        }
    }
    return result
}

