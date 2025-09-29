package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}

	// indices starting from the last elements
	i := len(slice1) - 1
	j := len(slice2) - 1

	// figure out who goes first
	turn := 1 // 1 = slice1, 2 = slice2
	if len(slice2) > len(slice1) {
		turn = 2
	}

	// loop until both slices are finished
	for i >= 0 || j >= 0 {
		if turn == 1 && i >= 0 {
			result = append(result, slice1[i])
			i--
			turn = 2 // switch to slice2
		} else if turn == 2 && j >= 0 {
			result = append(result, slice2[j])
			j--
			turn = 1 // switch to slice1
		} else if i >= 0 { // if one slice is empty, just finish the other
			result = append(result, slice1[i])
			i--
		} else if j >= 0 {
			result = append(result, slice2[j])
			j--
		}
	}

	return result
}
