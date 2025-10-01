package main

func NotDecimal(dec string) string {
	// empty string
	if dec == "" {

		return "\n"
	}
	//valid numbers
	//loop through dec

	for i := 0; i < len(dec); i++ {

		ch := dec[i]

		if !(ch >= '0' && ch <= '9' || ch == '.' || ch == '-') {

			return dec + "\n"
		}
	}

	// find the dot

	dot := -1

	// loop through dec

	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			dot = i
			break
		}
	}
	// no dot
	if dot == -1 {
		return dec + "\n"
	}

	//remove zeros after dot
	allZeros := true
	//loop through dec from dot + 1 == dec[i]
	for i := dot + 1; i < len(dec); i++ {
		if dec[i] != 0 {
			allZeros = false
			break
		}
	}
	if allZeros == true {
		return dec + "\n"
	}
	//remove the dot manually
	result := ""
	for i := 0; i < len(dec); i++ {

		if i == dot {
			continue
		}
		result += string(dec[i])
	}
	//remove leading zeros
	sign := ""
	// remove - sign temporarily
	if result[0] == '-' {
		sign = "-"
		result = result[1:]
	}
	//skiping the leading zero
	i := 0
	for i < len(result) && result[i] == '0' {
		i++
	}
	result = result[i:]
	if result == "" {
		result = "0"
	}
	return sign + result + "\n"
}
