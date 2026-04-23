package main

func ZipString(s string) string {
	result := ""
	seen := ""

	for _, ch := range s {

		// check if character is already seen
		alreadySeen := false     // initialise to false 
		for _, v := range seen {    // range through seen 
			if v == ch {
				alreadySeen = true     // if it is seen, then it is true
				break					// stop
			}
		}
		if alreadySeen {
			continue
		}

		// count occurrences
		count := 0
		for _, v := range s {
			if v == ch {
				count++
			}
		}

		// convert count to string manually (since no strconv)
		countStr := ""
		for count > 0 {
			digit := count % 10
			countStr = string('0'+rune(digit)) + countStr
			count /= 10
		}

		// build result
		result += countStr + string(ch)

		// mark as seen
		seen += string(ch)
	}

	return result
}