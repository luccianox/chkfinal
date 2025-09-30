/*
cameltosnakecase
Instructions
Write a function that converts a string from camelCase to snake_case.

If the string is empty, return an empty string.
If the string is not camelCase, return the string unchanged.
If the string is camelCase, return the snake_case version of the string.
For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

lowerCamelCase
UpperCamelCase
Rules for writing in camelCase:

The word does not end on a capitalized letter (CamelCasE).
No two capitalized letters shall follow directly each other (CamelCAse).
Numbers or punctuation are not allowed in the word anywhere (camelCase1).
Expected function
*/

func CamelToSnakeCase(s string) string {
	result := ""
	//empty string

	if len(s) == 0 {

		return ""
	}

	//only letter allowed

	for _, ch := range s {

		if !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') {

			return s
		}
	}
	//no two caps to follow each other

	for i := 1; i < len(s); i++ {

		if (s[i] >= 'A' && s[i] <= 'Z') && (s[i-1] >= 'A' && s[i-1] <= 'Z') {
			return s
		}
	}

	//last letter doesn't end in caps

	last := len(s) - 1

	if s[last] >= 'A' && s[last] <= 'Z' {

		return s
	}
	//convert to camel case
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {

			if i > 0 {

				result += "_"
			}
			result += string(ch)
		} else {
			result += string(ch)
		}

	}
	return result
}
