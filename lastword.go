package main

func LastWord(s string) string {
	result := ""
	word := ""

	for _, ch := range s {
		if ch != ' ' {
			word = word + string(ch)
		} else {
			if word != "" {
				result = word
			}
			word = ""
		}
		if word != "" {
			result = word
		}
	}
	return result + "\n"
}