package binaryGap

import "strconv"

func convertToBinary(input int) string {
	x := ""
	for input > 0 {
		a := input % 2
		x = strconv.FormatInt(int64(a), 10) + x
		input = input / 2
	}
	return x
}

func FindGap(input int) int {
	maxLong := 0
	gap := 0

	strBinary := convertToBinary(input)
	n := len(strBinary)

	for i := 0; i < n; i++ {
		if string(strBinary[i]) == "1" {
			if maxLong < gap {
				maxLong = gap
			}
			gap = 0
		}

		if string(strBinary[i]) == "0" {
			gap++
		}

	}

	return maxLong
}
