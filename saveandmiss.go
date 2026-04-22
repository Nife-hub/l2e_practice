package main

func SaveAndMiss(s string, n int) string {
	if n <= 0 {
		return s
	}
	res := ""
	save := true
	for i := 0; i < len(s); i += n{
		end := i + n
		if end > len(s) {
			end = len(s)
		}

		if save {
			res += s[i:end]
		}
		save = !save
	}
	return res
}

