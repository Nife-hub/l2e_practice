package main

func NotDecimal2(dec string) string {
	 if dec == ""{
		return "\n"
	}

	for i := 0; i < len(dec); i++{
		if i != 0 {
			if dec[i] == '0' && dec[i-1] == '-' {
				return dec
			}
		}
	}

	hasdecimal := false
	for _, ch := range dec{
		if ch == '.'{
			hasdecimal = true
			break
		}
	}

	if !hasdecimal{
		return dec + "\n"
	}

	first := ""
	second := ""
	result := ""
	for i:= 0; i < len(dec); i++{
		if dec[i] == '.'{
			first = dec[0:i]
			second = dec[i+1:]
		}
		if first == "0" {
    		return second + "\n"
		}
	}
	result = first + second
	result += "\n"
	return result

	// if dec == "" {
	// 	return "\n"
	// }

	// hasdecimal := false
	// for _, c := range dec {
	// 	if c == '.' {
	// 		hasdecimal = true
	// 		break
	// 	}
	// 	if !hasdecimal {
	// 		return dec + "\n"
	// 	}
	// }

	// for i := 0; i < len(dec); i++{
	// 	if i != 0 {
	// 		if dec[i] == 0 && dec[i-1] == '-'{
	// 			return dec
	// 		}
	// 	}
	// }

	// first := ""
	// second := ""
	// result := ""
	// for i := 0; i < len(dec); i++{
	// 	if dec[i] == '.'{
	// 		first = dec[0:i]
	// 		second = dec[i+1:]
	// 	} 
	// }
	// result = first + second
	// result += "\n"
	// return result 
}













