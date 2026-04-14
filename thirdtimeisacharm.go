package main

func ThirdTimeIsACharm(s string) string {
	if len(s) == 0 {
		return "\n"
	}
	if len(s) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i < len(s); i += 3 {
		result += string(s[i])
	}
	return result
}