package main

func FirstWord(s string) string{
	// result := ""
	// for i := 0; i < len(s); i++ {          // incase there's a space at the beginning 
	// 	if i != 0 && s[i] == ' ' {
	// 		break
	// 	}
	// 	result += string(s[i])
	// }
	//  for _, ch := range s {
	// 	if ch == ' ' {
	// 		break
	// 	}
	// 	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
	// 		result += string(ch)
	// 	}	
	// }
	// result = result + "\n"
	// return result


	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	start := i
	for i < len(s) && s[i] != ' ' {
		i++
	}
	return s[start:i] + "\n"
}




	//
	// for i := 0; i < len(s); i++ {
	// 	ch := s[i]
	// 	if i > 0 && s[i] == ' ' && s[i-1] >= 'a' && s[i-1] <= 'z' || s[i-1] >= 'A' && s[i-1] <= 'Z' {
	// 		break
	// 	}
	// 	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
	// 		result += string(ch)
	// 	}	
	// }