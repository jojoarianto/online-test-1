package findRecurringChar

func Find(in string) string {
	var (
		index = 0
		n     = len(in)
		box   = map[string]int{}
	)

	for index < n {
		char := string(in[index])
		box[char] = box[char] + 1

		if box[char] == 2 {
			// force return when found
			return char
		}
		index++
	}
	return ""
}
