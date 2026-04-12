package main

func FirstWord(s string) string{
	result := ""
	// for i := 0; i < len(s); i++ {          // incase there's a space at the beginning 
	// 	if i != 0 && s[i] == ' ' {
	// 		break
	// 	}
	// 	result += string(s[i])
	// }
	for _, ch := range s {
		if ch == ' ' {
			break
		}
		result += string(ch)
	}
	result = result + "\n"
	return result
}