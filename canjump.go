package main

import "fmt"

func CanJump(input []uint) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) == 1 {
		return true
	}
	pos := 0

	for pos < len(input) {

		if pos == len(input)-1 {

			return true
		}
		steps := int(input[pos])

		if steps == 0 {

			return false
		}
		pos += steps
		if pos >= len(input) {
			return false
		}
	}
	return false
}
func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
