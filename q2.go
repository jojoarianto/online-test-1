package main

import (
	"fmt"
)

func main() {
	fmt.Println(h(1, "fruits"))
	fmt.Println(h(2, "fruits"))
	fmt.Println(h(5, "fruits"))
	fmt.Println(h(pow(2, 10), "fruits"))
	fmt.Println(h(pow(2, 11), "fruits"))
}

/**
 * str in : friuts
 * str out : stuirf
 */
func f(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return str
	} else {
		return f(g(str)) + string(str[0])
	}
}

/**
 * string in : fruits
 * string out : ruits
 */
func g(str string) string {
	i := 0
	new_str := ""

	for i < (len(str) - 1) {
		new_str = new_str + string(str[i+1])
		i = i + 1
	}

	return new_str
}

/**
 * recursive function
 */
func h(n int, str string) string {
	for n != 1 {
		if n%2 == 0 { // genap
			n = n / 2
		} else {
			n = 3*n + 1
		}

		str = f(str)
	}
	return str
}

/**
 * int in x = 2
 * int in y = 4
 * int out = 16
 * 2^4
 */
func pow(x int, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pow(x, y-1)
	}
}
