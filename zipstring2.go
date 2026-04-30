package main

func ZipString2(s string) string {
	result := ""
	seen := ""

	for _, ch := range s {
		alreadySeen := false
		for _, b := range seen {
			if b == ch {
				alreadySeen = true
				break
			}
		}

		if alreadySeen {
			continue
		}

		count := 0
		for _, v := range s {
			if v == ch {
				count++
			}
		}

		countstr := ""
		for count > 0 {
			digit := count % 10
			countstr = string(digit + '0') + countstr
			count /= 10
		}

		result += countstr + string(ch)
		seen += string(ch)
	}
	return result
}
