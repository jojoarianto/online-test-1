package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanln(&t)

	var result []int

	for i := 0; i < t; i++ {
		var a int
		fmt.Scanln(&a)

		var b int
		fmt.Scanln(&b)

		var k int
		fmt.Scanln(&k)

		// call function f
		result = append(result, f(a, b, k))
	}

	// print result
	for i := 0; i < t; i++ {
		str := fmt.Sprintf("Case %v: %v", (i + 1), result[i])
		fmt.Println(str)
	}
}

func f(a int, b int, k int) int {
	ct := 0
	for a <= b {
		if a%k == 0 {
			ct = ct + 1
		}
		a = a + 1
	}

	return ct
}
