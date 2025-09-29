package piscine

func FifthAndSkip(str string) string {
	// case: empty string
	if len(str) == 0 {
		return "\n"
	}

	// case: shorter than 5 characters
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	count := 0 // counts characters added in current group
	skips := 0 // counts how many we’ve skipped

	for i := 0; i < len(str); i++ {
		ch := str[i]

		// ignore spaces when counting toward 5
		if ch == ' ' {
			continue
		}

		// if we've added 5 characters, insert a space and reset count
		if count == 5 {
			result += " "
			count = 0
			skips = 0
		}

		// if we've already added 5 and need to skip the 6th
		if skips == 0 && count == 0 && len(result) > 0 {
			// skip this one
			skips++
			continue
		}

		// add the character to result
		result += string(ch)
		count++
	}

	return result + "\n"
}
