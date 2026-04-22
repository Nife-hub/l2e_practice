package main

func RepeatAlpha(s string) string{
    result := ""
    for _, ch := range s{
		repeat := 1
		if ch >= 'a' && ch <= 'z' {
			repeat = int(ch - 'a' + 1)
		} else if ch >= 'A' && ch <= 'Z' {
			repeat = int(ch - 'A' + 1)
		}
		for i := 0; i < repeat; i++{
			result += string(ch)
		}
	}
	return result
}
   





// 	result := ""
// 	for _, ch := range s{
// 		if ch >= 'a' && ch <= 'z' {
// 			repeat := int(ch - 'a') + 1
// 			for i := 1; i < repeat; i++{
// 				result = result + string(ch) 
// 			}
// 		}

// 		if ch >= 'A' && ch <= 'Z' {
// 			repeat := int(ch - 'A') + 1
// 			for i := 1; i < repeat; i++{
// 				result = result + string(ch) 
// 			}
// 		} else {
// 			result += string(ch)
// 		}
// 	} 
// 	return result
// }






















