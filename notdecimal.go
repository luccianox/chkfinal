package main

func NotDecimal(dec string) string {
	// case 1: empty
	if dec == "" {
		return "\n"
	}

	// check if it's a valid number (digits, '-', '.')
	for i := 0; i < len(dec); i++ {
		ch := dec[i]
		if !(ch >= '0' && ch <= '9' || ch == '.' || ch == '-') {
			return dec + "\n"
		}
	}

	// find the dot (if any)
	dot := -1
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

	// check if only zeros after dot
	allZeros := true
	for i := dot + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			allZeros = false
			break
		}
	}
	if allZeros {
		return dec + "\n"
	}

	// remove the dot manually
	result := ""
	for i := 0; i < len(dec); i++ {
		if i == dot {
			continue
		}
		result += string(dec[i])
	}

	return result + "\n"
}
