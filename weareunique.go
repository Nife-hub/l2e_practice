package main

func WeAreUnique(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return -1
	}

	freq1 := make(map[rune]bool)
	freq2 := make(map[rune]bool)

	for _, ch := range s1 {
		freq1[ch] = true
	}
	for _, ch := range s2 {
		freq2[ch] = true
	}

	count := 0
	for ch := range freq1 {
		if !freq2[ch] {
			count++
		}
	}
	for ch := range freq2 {
		if !freq1[ch] {
			count++
		}
	}
	return count
}