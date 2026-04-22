package main

func LastWord(s string) string {
	// result := ""
	// word := ""

	// for _, ch := range s {
	// 	if ch != ' ' {
	// 		word = word + string(ch)
	// 	} else {
	// 		if word != "" {
	// 			result = word
	// 		}
	// 		word = ""
	// 	}
	// 	if word != "" {
	// 		result = word
	// 	}
	// }
	// return result + "\n"

	i := len(s)-1
	for i >= 0 && s[i] == ' '{
		i--
	}
	end := i
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return s[i+1:end+1] + "\n"
}