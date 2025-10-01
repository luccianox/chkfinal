package main

func NotDecimal(dec string) string {
	// 1. Empty input
	if dec == "" {
		return "\n"
	}

	// 2. Check validity: only digits, '.', or '-'
	for i := 0; i < len(dec); i++ {
		ch := dec[i]
		if !(ch >= '0' && ch <= '9' || ch == '.' || ch == '-') {
			return dec + "\n"
		}
	}

	// 3. Find the decimal point
	dot := -1
	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			dot = i
			break
		}
	}

	// 4. If no dot → return as is
	if dot == -1 {
		return dec + "\n"
	}

	// 5. Check if everything after the dot is zero
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

	// 6. Remove the dot
	result := ""
	for i := 0; i < len(dec); i++ {
		if i == dot {
			continue
		}
		result += string(dec[i])
	}

	// 7. Remove leading zeros (but keep sign if negative)
	sign := ""
	if result[0] == '-' {
		sign = "-"
		result = result[1:]
	}
	i := 0
	for i < len(result) && result[i] == '0' {
		i++
	}
	result = result[i:]

	// If result becomes empty, it's just "0"
	if result == "" {
		result = "0"
	}

	return sign + result + "\n"
}
