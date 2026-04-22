package main

func HashCode(s string) string{
	length := len(s)

	result := ""
	for _, ch := range s{
		ascii := int(ch)
		hash := (ascii + length) % 127

		if hash < 33 {
			hash += 33
		}
		result = result + string(rune(hash)) 
	}
	return result
}


func HashCode2(dec string) string {
	length := len(dec)
		var hashed string

	for _, ch := range dec {
		hashVal := (int(ch) + length) % 127
	// Make sure it’s printable
		if hashVal < 33 {
			hashVal += 33
		}
		hashed += string(rune(hashVal))
	}
	return hashed
}