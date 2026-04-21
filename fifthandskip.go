package main

func FifthAndSkip(s string) string {
	if s == "" {
		return "\n"
	}
	if len(s) < 5 {
		return "Invalid Input\n"
	}
	result := ""
	count := 0

	for i := 0; i < len(s); i++{
		if s[i] == ' ' {
			continue
		}

		if count == 6{
			count = 0
			continue
		}
		count++

		result += string(s[i])
		if count == 5 {
			result += " "
		}
	}
	return result
}
