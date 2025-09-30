/*
Instructions
Write a function called NotDecimal() that takes as an argument a string in form of a float number
with the decimal point and returns a string converted to int without the decimal point
(you will have to multiply it by 10^n to remove the .).

If the number doesn't have a decimal point or there is only a zero after the . return the number followed by a newline \n.
If the argument is empty return a newline \n.
If the argument is not a number return it followed by a newline \n.
Expected function
*/
package main

import "fmt"

func NotDecimal(dec string) string {

	//empty argument

	if dec == "" {

		return "\n"
	}

	//check if its a valid number

	for i := 0; i < len(dec); i++ {

		ch := dec[i]

		if !(ch >= '0' && ch <= '9' || ch == '.' || ch == '-') {

			return dec + "\n"
		}
	}
	//look for the dot

	dot := -1
	//loop thru dec

	for i := 0; i < len(dec); i++ {

		if dec[i] == '.' {

			dot = i
			break
		}
	}
	if dot == -1 {

		return dec + "\n"
	}
	//all zeros

	allZeros := true
	for i := dot + 1; i < len(dec); i++ {

		if dec[i] != 0 {

			allZeros = false
			break
		}
	}
	if allZeros {

		return dec + "\n"
	}
	// 6. Remove the dot (build result manually)

	result := ""
	for i := 0; i < len(dec); i++ {

		if i == dot {

			continue
		}
		result += string(dec[i])
	}
	return result + "\n"
}
func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
