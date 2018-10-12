package solution

func Solution(A []int) int {
	// write your code in Go 1.4
	ct := 1
	status := true

	for status {
		for i, v := range A {
			if v == ct {
				ct++
				break
			}

			if i == (len(A) - 1) {
				status = false
				break
			}
		}
	}

	return ct
}
