/*slice
Instructions
The function receives a slice of strings and one or more integers, and returns a slice of strings.
 The returned slice is part of the received one but cut from the position indicated in the first int,
 until the position indicated by the second int.

In case there only exists one int, the resulting slice begins in the position indicated by the int
 and ends at the end of the received slice.

The integers can be negative.

Expected function*/
func Slice(a []string, nbrs ...int) []string {
	n := len(a)
	if len(nbrs) == 0 || len(nbrs) > 2 {

		return nil
	}

	//one index number
	if len(nbrs) == 1 {

		start := nbrs[0]
		if start < 0 {

			start = n + start
		}
		if start < 0 || start >= n {

			return nil
		}
		return a[start:]
	}
	// deal with 2 indices
	start, end := nbrs[0], nbrs[1]

	if start < 0 {
		start = n + start
	}
	if end < 0 {

		end = n + start
	}
	if start < 0 || end < 0 || start >= n || end > n || start >= end {

		return nil
	}
	return a[start:end]
}