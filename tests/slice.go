/*slice
Instructions
The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int.

In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at the end of the received slice.

The integers can be negative.*/

package piscine

func Slice(a []string, nbrs ...int) []string {
	n := len(a)

	// Not enough args → return nil
	if len(nbrs) == 0 || len(nbrs) > 2 {
		return nil
	}

	// Convert first index
	start := nbrs[0]
	if start < 0 {
		start = n + start
	}
	if start < 0 || start > n {
		return nil
	}

	// If only one argument → slice to the end
	if len(nbrs) == 1 {
		return a[start:]
	}

	// Convert second index
	end := nbrs[1]
	if end < 0 {
		end = n + end
	}
	if end < 0 || end > n {
		return nil
	}

	// If start >= end → empty slice
	if start >= end {
		return nil
	}

	return a[start:end]
}
