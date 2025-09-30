/*fifthandskip
Instructions
Write a function FifthAndSkip() that takes a string and returns another string.
The function separates every five characters of the string with a space and removes the sixth one.

If there are spaces in the middle of a word, ignore them and get the first character
after the spaces until you reach a length of 5.
If the string is less than 5 characters return Invalid Input followed by a newline \n.
If the string is empty return a newline \n.*/

func FifthAndSkip(str string) string {

	if str == "" {

		return "\n"
	}
	if len(str) < 5 {

		return "Invalid input\n"
	}
	result := ""
	count := 0
	skips := 0

	for i := 0; i < len(str); i++ {
		ch := str[i]
		// ignore spaces when counting toward 5
		if ch == ' ' {
			contunue
		}
		// if we've added 5 characters, insert a space and reset count

		if count == 5 {
			result += " "
			count = 0
			skips = 0
		}
		// if we've already added 5 and need to skip the 6th
		if skips == 0 && count == 0 && len(result) > 0 {

			skip++
			continue
		}
		// add the character to result
		result += string(ch)
		count++
	}
	return result + "\n"
}
