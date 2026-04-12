package main

func HashCode(s string) string{
	length := 0
	for range s {
		length++
	}

	result := ""
	for _, ch := range s{
		ascii := int(ch)
		hash := (ascii + length) % 127

		if hash < 33 {
			hash += 33
		}
		result = result + string(hash) 
	}
	return result
}