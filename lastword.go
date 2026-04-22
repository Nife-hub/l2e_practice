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

	i := len(s)-1         // start at the last index of the string
	for i >= 0 && s[i] == ' ' {  // loop backwards while the character is a space till it finds a non-space character
		i--
	}
	end := i   // store the index in the variable "end"

	for i >= 0 && s[i] != ' ' {    // keep moving left while characters are not spaces till it finds the next space or reaches the beginning
		i--
	}

	return s[i+1:end+1] + "\n" // i+1 because s[i] is a space and s[i+1] is the start word after the space
							   // end+1 because slices exclude the end so we add 1 to include it 
}