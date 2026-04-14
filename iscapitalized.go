package main

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if i == 0 || s[i-1] == ' ' {
			if s[i] >= 'a' && s[i] <= 'z' {
				return false
			}
		}
	} 
	return true
}
