package main

func PrintIf(s string) string {
	v := len(s)
	 if v != 0 && v <= 2 {
		return("Invalid Input\n")
	}
	return "G\n"
}