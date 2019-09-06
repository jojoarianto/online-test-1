package intToBinary

func Convert(input int) string {
	binary := ""

	for input > 0 {
		sisa := input % 2
		input = input / 2

		if sisa == 1 {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
	}
	return binary
}
