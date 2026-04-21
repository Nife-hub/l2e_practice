package main

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {      // check if the last letter is uppercase
			return s
		}

	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {        // check if two uppercase are next to each other
			return s
		}
		if !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') {   // check for anything other than a letter
			return s
		}
	}

	res := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= 'A' && ch <= 'Z' {
			if i != 0 {
				res += "_"
			} 
			res += string(ch)
		} else {
			res += string(ch)
		}
	}
	return res
}


