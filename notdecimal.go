package main

func NotDecimal(dec string) string {
	// If the input is empty, just return a newline
	if dec == "" {
		return "\n"
	}

	// This will store the position (index) of the decimal point '.'
	// We start with -1 to mean "no dot found yet"
	dot := -1

	// 🔍 First loop: validate the string and find the dot
	for i := 0; i < len(dec); i++ {

		// If current character is a dot
		if dec[i] == '.' {

			// If we already found a dot before, it's invalid (more than one dot)
			if dot != -1 {
				return dec + "\n"
			}

			// Store the position of the dot
			dot = i

		} else if dec[i] < '0' || dec[i] > '9' {
			// If it's not a digit (0–9), it's invalid input
			return dec + "\n"
		}
	}

	// ❌ If no dot was found at all, return the original string
	if dot == -1 {
		return dec + "\n"
	}

	// 🔍 Check if all digits after the dot are zeros
	allZero := true

	for i := dot + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			// Found a non-zero digit → not all zeros
			allZero = false
			break
		}
	}

	// ❌ If:
	// - all digits after dot are zero (e.g. "12.00")
	// - OR dot is the last character (e.g. "12.")
	// Then return original string
	if allZero || dot == len(dec)-1 {
		return dec + "\n"
	}

	// 🔧 Otherwise, we remove the dot
	result := ""

	for i := 0; i < len(dec); i++ {
		// Add every character except the dot
		if dec[i] != '.' {
			result += string(dec[i])
		}
	}

	// Return the new string + newline
	return result + "\n"
}